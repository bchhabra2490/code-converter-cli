# Code Converter CLI

A command-line tool written in Go that converts code from one programming language to another using OpenAI's GPT-4o model to perform intelligent code translation.

## Features

- Convert an entire project from one programming language to another
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

# Convert code
./code-converter-cli -input /path/to/source -output /path/to/destination -lang python
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

1. The tool scans the input directory and identifies code files based on their extensions.
2. For each recognized code file, it reads the source code.
3. The source code is sent to OpenAI's GPT-4o model with a prompt specifying the source and target languages.
4. The AI generates the equivalent code in the target language.
5. The tool cleans the response (removing markdown code block markers if present).
6. The converted files are written to the output directory with appropriate file extensions, preserving the original folder structure.

## Implementation Notes

- The tool uses the `github.com/sashabaranov/go-openai` package to interact with OpenAI's API.
- Non-code files (e.g., images, data files) are copied as-is to the output directory.
- Certain directories like `.git`, `node_modules`, and `vendor` are skipped during processing.
- The tool automatically handles file extension changes based on the target language.

## Limitations

- Requires an OpenAI API key and may incur costs based on your API usage.
- AI translation quality depends on the context and complexity of the code.
- Full conversion between programming languages is an extremely complex task that often requires manual adjustment.
- Complex language features, custom libraries, and platform-specific code may not convert perfectly.
- API rate limits may affect large-scale conversions.

## License

MIT

## Acknowledgements

- [OpenAI](https://openai.com/) for providing the GPT-4o model used for code translation
- [sashabaranov/go-openai](https://github.com/sashabaranov/go-openai) for the Go client library
