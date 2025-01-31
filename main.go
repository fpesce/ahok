package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "sort"
    "strings"

    goahocorasick "github.com/anknown/ahocorasick"
)

// removeSupersets removes any keyword that contains a smaller keyword as a substring.
// Example: if "car" is in the set, remove "mycar" (since any occurrence of "mycar" also contains "car").
func removeSupersets(keywords []string) []string {
    // Sort keywords by length ascending for a simple approach
    sort.Slice(keywords, func(i, j int) bool {
        if len(keywords[i]) == len(keywords[j]) {
            return keywords[i] < keywords[j]
        }
        return len(keywords[i]) < len(keywords[j])
    })

    var pruned []string
outer:
    for _, kw := range keywords {
        // If kw contains any smaller keyword in pruned, skip
        for _, existing := range pruned {
            if strings.Contains(kw, existing) {
                continue outer
            }
        }
        pruned = append(pruned, kw)
    }
    return pruned
}

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: go run main.go <keywords_file> <directory>")
        return
    }

    keywordsFile := os.Args[1]
    dir := os.Args[2]

    // 1) Read keywords from file
    f, err := os.Open(keywordsFile)
    if err != nil {
        log.Fatalf("Failed to open keywords file: %v", err)
    }
    defer f.Close()

    keywordsMap := make(map[string]struct{})
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        kw := strings.TrimSpace(scanner.Text())
        if kw != "" {
            keywordsMap[kw] = struct{}{}
        }
    }
    if err := scanner.Err(); err != nil {
        log.Fatalf("Failed to read keywords: %v", err)
    }

    // Convert map to slice
    var keywords []string
    for k := range keywordsMap {
        keywords = append(keywords, k)
    }

    // 2) Optionally remove larger keywords containing smaller ones
    keywords = removeSupersets(keywords)

    // 3) Build Aho-Corasick automaton (which is type "Machine" in the snippet)
    patterns := make([][]rune, 0, len(keywords))
    for _, kw := range keywords {
        patterns = append(patterns, []rune(kw))
    }

    var automaton goahocorasick.Machine
    if err := automaton.Build(patterns); err != nil {
        log.Fatalf("Failed to build Aho-Corasick machine: %v", err)
    }

    // 4) Walk the directory and check each file line by line
    err = filepath.Walk(dir, func(path string, info os.FileInfo, walkErr error) error {
        if walkErr != nil {
            return walkErr
        }
        // Skip directories
        if info.IsDir() {
            return nil
        }

        file, err := os.Open(path)
        if err != nil {
            log.Printf("Could not open file %s: %v", path, err)
            return nil
        }
        defer file.Close()

        fileScanner := bufio.NewScanner(file)
        for fileScanner.Scan() {
            line := fileScanner.Text()
            // MultiPatternSearch returns a slice of *Term structs if matched
            matches := automaton.MultiPatternSearch([]rune(line), false)
            if len(matches) > 0 {
                fmt.Printf("[%s] %s\n", path, line)
            }
        }
        if err := fileScanner.Err(); err != nil {
            log.Printf("Error reading file %s: %v", path, err)
        }
        return nil
    })

    if err != nil {
        log.Fatalf("Error walking directory: %v", err)
    }
}
