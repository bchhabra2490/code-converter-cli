#!/bin/bash

# Build the code converter CLI
echo "Building code-converter-cli..."
go build -o code-converter-cli .

# Create output directory for the demo
DEMO_OUTPUT="./examples/converted"
mkdir -p "$DEMO_OUTPUT"

# Run the code converter to convert the Go example to Python
echo "Converting Go code to Python..."
./code-converter-cli -input ./examples/demo -output "$DEMO_OUTPUT/python" -lang python

# Run the code converter to convert the Go example to JavaScript
echo "Converting Go code to JavaScript..."
./code-converter-cli -input ./examples/demo -output "$DEMO_OUTPUT/javascript" -lang javascript

# # Show the converted files
# echo -e "\nConverted Python files:"
# ls -la "$DEMO_OUTPUT/python"

# echo -e "\nConverted JavaScript files:"
# ls -la "$DEMO_OUTPUT/javascript"

# echo -e "\nPython code:"
# cat "$DEMO_OUTPUT/python/main.py"

# echo -e "\nJavaScript code:"
# cat "$DEMO_OUTPUT/javascript/main.js"

# echo -e "\nDemo completed!" 