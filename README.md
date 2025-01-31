# Fast Keyword Scanner

![Go](https://img.shields.io/badge/Go-1.19-blue)
![License](https://img.shields.io/badge/License-MIT-green)

A simple high-performance keyword search tool using the **Aho-Corasick** algorithm to efficiently scan files for keyword matches. The program reads a list of keywords, optimizes them by removing supersets, and then searches through files in a given directory, reporting lines containing matches.

## Features

✅ **Optimized Keyword Processing**: Filters out redundant keywords to reduce unnecessary matches.  
✅ **Aho-Corasick Trie Search**: Uses an efficient string-matching algorithm for rapid text scanning.  
✅ **Recursive Directory Search**: Processes all files in a target directory.  
✅ **Fast & Lightweight**: Uses minimal resources while scanning large amounts of text.  

## Installation

### Prerequisites
- **Go 1.19+** installed on your machine.

### Clone & Build
```sh
git clone https://github.com/yourusername/fast-keyword-scanner.git
cd fast-keyword-scanner
go build -o scanner main.go
```

## Usage

Run the program with a **keyword file** and a **directory** to scan:

```sh
go run main.go <keywords.txt> <directory>
```
Or, if built:
```sh
./scanner <keywords.txt> <directory>
```

### Example
```sh
./scanner keywords.txt /path/to/logs
```
#### Output
```
[/path/to/logs/example.log] Suspicious activity detected.
[/path/to/logs/error.log] Critical error found!
```

## Keyword Optimization

To avoid redundant searches, the program removes keywords that contain smaller keywords as substrings.  
Example:
```
car
mycar
supercar
```
Will be optimized to:
```
car
```

## How It Works
1. **Reads keywords** from a file, removing duplicates and supersets.
2. **Builds an Aho-Corasick Trie** for fast multi-pattern matching.
3. **Walks through the target directory** and scans all files line by line.
4. **Prints matches** with the file path and matched line.

## Dependencies
- [goahocorasick](https://github.com/anknown/ahocorasick) - Aho-Corasick algorithm implementation in Go.

## Contributing
Pull requests are welcome! Feel free to open an issue if you find a bug or want to request a feature.

## License
MIT License. See [LICENSE](LICENSE) for details.
