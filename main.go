package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/b-eq/code-converter-cli/converter"
)

func main() {
	// Define command-line flags
	inputDir := flag.String("input", "", "Input project directory (required)")
	outputDir := flag.String("output", "", "Output directory for converted code (required)")
	targetLang := flag.String("lang", "", "Target programming language (required)")
	
	// Parse flags
	flag.Parse()
	
	// Validate required flags
	if *inputDir == "" || *outputDir == "" || *targetLang == "" {
		fmt.Println("Error: input, output, and lang flags are required")
		flag.Usage()
		os.Exit(1)
	}
	
	// Convert relative paths to absolute
	absInputDir, err := filepath.Abs(*inputDir)
	if err != nil {
		fmt.Printf("Error resolving input directory path: %v\n", err)
		os.Exit(1)
	}
	
	absOutputDir, err := filepath.Abs(*outputDir)
	if err != nil {
		fmt.Printf("Error resolving output directory path: %v\n", err)
		os.Exit(1)
	}
	
	// Validate that input directory exists
	if _, err := os.Stat(absInputDir); os.IsNotExist(err) {
		fmt.Printf("Error: input directory %s does not exist\n", absInputDir)
		os.Exit(1)
	}
	
	// Create output directory if it doesn't exist
	if err := os.MkdirAll(absOutputDir, 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Printf("Converting code from %s to %s language\n", absInputDir, *targetLang)
	fmt.Printf("Output will be saved to %s\n", absOutputDir)
	
	inputPaths := strings.Split(*inputDir, ",")
	for _, path := range inputPaths {
		path = strings.TrimSpace(path)
		fileInfo, err := os.Stat(path)
		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", path, err)
			continue
		}
		
		if fileInfo.IsDir() {
			// Process directory (existing code)
			conv := converter.NewConverter(path, *outputDir, *targetLang)
			if err := conv.Convert(); err != nil {
				fmt.Printf("Error during directory conversion: %v\n", err)
			}
		} else {
			// Process individual file
			if err := converter.ConvertFile(path, *outputDir, *targetLang); err != nil {
				fmt.Printf("Error converting file %s: %v\n", path, err)
			}
		}
	}
	
	fmt.Println("Conversion completed successfully!")
} 