# Code Converter CLI

A command-line tool written in Go that converts code from one programming language to another. The tool analyzes the source code using [Zoekt](https://github.com/sourcegraph/zoekt), a fast trigram-based code search engine, to understand the structure of the codebase before performing the conversion.

## Features

- Convert an entire project from one programming language to another
- Preserve directory structure in the output
- Analyze code structure using Zoekt's search capabilities
- Support for multiple programming languages

## Installation

### Prerequisites

- Go 1.18 or higher
- Git

### Build from source

```bash
# Clone the repository
git clone https://github.com/b-eq/code-converter-cli.git
cd code-converter-cli

# Build the tool
go build -o code-converter-cli .
```

## Usage

```bash
# Basic usage
./code-converter-cli -input /path/to/source -output /path/to/destination -lang python

# Example: Convert a Go project to Python
./code-converter-cli -input ~/projects/my-go-app -output ~/projects/my-python-app -lang python
```

### Command-line Arguments

- `-input`: Source project directory (required)
- `-output`: Destination directory for the converted code (required)
- `-lang`: Target programming language (required)

### Supported Languages

The tool currently recognizes the following programming languages:

- Go
- JavaScript
- TypeScript
- Python
- Java
- C
- C++
- C#
- Ruby
- PHP
- Rust
- Swift
- Kotlin

## How It Works

1. The tool scans the input directory and builds a code index using Zoekt.
2. It analyzes the codebase to understand functions, imports, and dependencies.
3. For each source file, it performs language-specific conversion.
4. The converted files are written to the output directory, preserving the original folder structure.

## Implementation Notes

- The converter preserves the original code as comments in the converted files for reference.
- Non-code files (e.g., images, data files) are copied as-is to the output directory.
- Certain directories like `.git`, `node_modules`, and `vendor` are skipped during processing.

## Limitations

- The current implementation provides a basic framework for code conversion.
- Full conversion between programming languages is an extremely complex task that often requires manual adjustment after the automated conversion.
- Complex language features, custom libraries, and platform-specific code may not convert correctly.

## License

MIT

## Acknowledgements

- [Zoekt](https://github.com/sourcegraph/zoekt) for code search capabilities 