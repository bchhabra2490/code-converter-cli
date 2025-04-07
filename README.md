# Code Converter CLI

A command-line tool written in Go that converts code from one programming language to another using OpenAI's GPT-4o model to perform intelligent code translation.

## Features

- Convert an entire project from one programming language to another
- Process individual files or multiple files
- Preserve directory structure in the output
- AI-powered code translation using OpenAI's GPT-4o model
- Automatic handling of file extensions based on target language
- Support for multiple programming languages

## Installation

### Prerequisites

- Go 1.18 or higher
- Git
- OpenAI API key (required)

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
# Set your OpenAI API key
export OPENAI_API_KEY="your-api-key"

# Convert an entire directory
./code-converter-cli -input /path/to/source -output /path/to/destination -lang python

# Convert a single file
./code-converter-cli -input /path/to/source/file.js -output /path/to/destination -lang python

# Convert multiple specific files
./code-converter-cli -input /path/to/file1.go,/path/to/file2.go -output /path/to/destination -lang python
```

### Command-line Arguments

- `-input`: Source project directory or file path(s) (required)
  - Can specify a directory to process all files
  - Can specify a single file path
  - Can specify multiple files as a comma-separated list
- `-output`: Output directory for converted code (required)
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

1. The tool processes the input, which can be a directory, a single file, or multiple files.
2. For directories, it recursively scans and identifies code files based on their extensions.
3. For each recognized code file, it reads the source code.
4. The source code is sent to OpenAI's GPT-4o model with a prompt specifying the source and target languages.
5. The AI generates the equivalent code in the target language.
6. The tool cleans the response (removing markdown code block markers if present).
7. The converted files are written to the output directory with appropriate file extensions, preserving the original folder structure for directory inputs.

## Implementation Notes

- The tool uses the `github.com/sashabaranov/go-openai` package to interact with OpenAI's API.
- Non-code files (e.g., images, data files) are copied as-is to the output directory.
- Certain directories like `.git`, `node_modules`, and `vendor` are skipped during processing.
- The tool automatically handles file extension changes based on the target language.
- When processing individual files, the output directory structure is flattened for those files.

## Limitations

- Requires an OpenAI API key and may incur costs based on your API usage.
- AI translation quality depends on the context and complexity of the code.
- Full conversion between programming languages is an extremely complex task that often requires manual adjustment.
- Complex language features, custom libraries, and platform-specific code may not convert perfectly.
- API rate limits may affect large-scale conversions.
- When converting multiple files separately, inter-file dependencies may not be handled as well as when converting entire directories.

## License

MIT

## Acknowledgements

- [OpenAI](https://openai.com/) for providing the GPT-4o model used for code translation
- [sashabaranov/go-openai](https://github.com/sashabaranov/go-openai) for the Go client library
