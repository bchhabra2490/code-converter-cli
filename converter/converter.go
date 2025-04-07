package converter

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Converter handles the code conversion process
type Converter struct {
	inputDir   string
	outputDir  string
	targetLang string
}

// NewConverter creates a new Converter instance
func NewConverter(inputDir, outputDir, targetLang string) *Converter {
	return &Converter{
		inputDir:   inputDir,
		outputDir:  outputDir,
		targetLang: targetLang,
	}
}

// Convert performs the full conversion process
func (c *Converter) Convert() error {
	return c.processDirectory(c.inputDir, c.outputDir)
}

// processDirectory recursively processes all files in a directory
func (c *Converter) processDirectory(inputPath, outputPath string) error {
	// Create the output directory if it doesn't exist
	if err := os.MkdirAll(outputPath, 0755); err != nil {
		return fmt.Errorf("failed to create output directory %s: %w", outputPath, err)
	}
	
	entries, err := os.ReadDir(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %w", inputPath, err)
	}
	
	for _, entry := range entries {
		inPath := filepath.Join(inputPath, entry.Name())
		outPath := filepath.Join(outputPath, entry.Name())
		
		if entry.IsDir() {
			// Skip common directories to ignore
			if shouldIgnoreDir(entry.Name()) {
				fmt.Printf("Skipping directory: %s\n", inPath)
				continue
			}
			
			// Process subdirectory recursively
			if err := c.processDirectory(inPath, outPath); err != nil {
				return err
			}
		} else {
			// Process file
			if err := c.processFile(inPath, outPath, entry); err != nil {
				return err
			}
		}
	}
	
	return nil
}

// processFile converts a single file from source to target language
func (c *Converter) processFile(inputPath, outputPath string, info fs.DirEntry) error {
	// Check if this file should be processed based on extension
	srcLang, shouldProcess := detectLanguage(inputPath)
	if !shouldProcess {
		// Just copy the file if we're not converting it
		return copyFile(inputPath, outputPath)
	}
	
	fmt.Printf("Converting %s from %s to %s\n", inputPath, srcLang, c.targetLang)
	
	// Read the source file
	content, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", inputPath, err)
	}
	
	// Convert the code
	convertedCode, newExt, err := c.convertCode(string(content), srcLang, inputPath)
	if err != nil {
		return fmt.Errorf("failed to convert %s: %w", inputPath, err)
	}
	
	// Update the output path with the new file extension if needed
	if newExt != "" {
		outputPath = changeExtension(outputPath, newExt)
	}
	
	// Write the converted code to the output file
	if err := os.WriteFile(outputPath, []byte(convertedCode), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", outputPath, err)
	}
	
	return nil
}

// convertCode translates code from one language to another
func (c *Converter) convertCode(sourceCode, sourceLang, filePath string) (string, string, error) {
	// Get the appropriate file extension for the target language
	newExt := getTargetExtension(c.targetLang)

	convertedCode, err := convertUsingLLM(sourceCode, sourceLang, c.targetLang)
	if err != nil {
		return "", "", fmt.Errorf("failed to convert %s: %w", filePath, err)
	}
	
	return convertedCode, newExt, nil
}

// ConvertFile converts a single file from source to target language
func ConvertFile(filePath, outputDir, targetLang string) error {
    _, fileName := filepath.Split(filePath)
    outputPath := filepath.Join(outputDir, fileName)
    
    // Create output directory if it doesn't exist
    if err := os.MkdirAll(outputDir, 0755); err != nil {
        return fmt.Errorf("failed to create output directory %s: %w", outputDir, err)
    }
    
    // Get source language
    srcLang, shouldProcess := detectLanguage(filePath)
    if !shouldProcess {
        // Just copy the file if we're not converting it
        return copyFile(filePath, outputPath)
    }
    
    fmt.Printf("Converting %s from %s to %s\n", filePath, srcLang, targetLang)
    
    // Read the source file
    content, err := os.ReadFile(filePath)
    if err != nil {
        return fmt.Errorf("failed to read file %s: %w", filePath, err)
    }
    
    // Create a temporary converter just for this file
    c := NewConverter(filePath, outputDir, targetLang)
    
    // Convert the code
    convertedCode, newExt, err := c.convertCode(string(content), srcLang, filePath)
    if err != nil {
        return fmt.Errorf("failed to convert %s: %w", filePath, err)
    }
    
    // Update the output path with the new file extension if needed
    if newExt != "" {
        outputPath = changeExtension(outputPath, newExt)
    }
    
    // Write the converted code to the output file
    if err := os.WriteFile(outputPath, []byte(convertedCode), 0644); err != nil {
        return fmt.Errorf("failed to write file %s: %w", outputPath, err)
    }
    
    return nil
}

func convertUsingLLM(sourceCode, sourceLang, targetLang string) (string, error) {
	prompt := fmt.Sprintf("Convert the following %s code to %s:\n\n%s. Just return the converted code, no other text.", sourceLang, targetLang, sourceCode)

	convertedCode, err := GenerateText(prompt)
	if err != nil {
		return "", fmt.Errorf("failed to convert %s: %w", sourceCode, err)
	}
	

	return convertedCode, nil
}

// Helper functions

// shouldIgnoreDir returns true for directories that should not be processed
func shouldIgnoreDir(dirName string) bool {
	dirsToIgnore := map[string]bool{
		".git":         true,
		"node_modules": true,
		"vendor":       true,
		"dist":         true,
		"build":        true,
		".idea":        true,
		".vscode":      true,
	}
	return dirsToIgnore[dirName]
}

// detectLanguage returns the detected language of a file and whether it should be processed
func detectLanguage(filePath string) (string, bool) {
	ext := strings.ToLower(filepath.Ext(filePath))
	
	// Map of file extensions to language names
	langMap := map[string]string{
		".go":   "Go",
		".js":   "JavaScript",
		".ts":   "TypeScript",
		".py":   "Python",
		".java": "Java",
		".c":    "C",
		".cpp":  "C++",
		".cs":   "C#",
		".rb":   "Ruby",
		".php":  "PHP",
		".rs":   "Rust",
		".swift":"Swift",
		".kt":   "Kotlin",
	}
	
	lang, ok := langMap[ext]
	return lang, ok
}

// getTargetExtension returns the file extension for the target language
func getTargetExtension(targetLang string) string {
	// Map of language names to file extensions
	extMap := map[string]string{
		"go":         ".go",
		"golang":     ".go",
		"javascript": ".js",
		"typescript": ".ts",
		"python":     ".py",
		"java":       ".java",
		"c":          ".c",
		"c++":        ".cpp",
		"csharp":     ".cs",
		"c#":         ".cs",
		"ruby":       ".rb",
		"php":        ".php",
		"rust":       ".rs",
		"swift":      ".swift",
		"kotlin":     ".kt",
	}
	
	// Normalize target language to lowercase
	normalizedLang := strings.ToLower(targetLang)
	
	if ext, ok := extMap[normalizedLang]; ok {
		return ext
	}
	return ""
}

// changeExtension replaces the file extension in a path
func changeExtension(filePath, newExt string) string {
	ext := filepath.Ext(filePath)
	return filePath[:len(filePath)-len(ext)] + newExt
}

// copyFile copies a file from src to dst
func copyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()
	
	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	
	_, err = io.Copy(destination, source)
	return err
} 
