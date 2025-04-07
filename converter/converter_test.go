package converter

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Mock the GenerateText function to avoid actual API calls during tests
var originalGenerateText func(string) (string, error)

func init() {
	// Save the original function
	originalGenerateText = GenerateText
}

// setupMockGPT sets up the mock function for GenerateText
func setupMockGPT(mockResponse string, mockError error) func() {
	// Save the original function
	originalGenerateText := GenerateText
	
	// Replace with mock
	GenerateText = func(prompt string) (string, error) {
		return mockResponse, mockError
	}

	// Return a cleanup function to restore the original
	return func() {
		GenerateText = originalGenerateText
	}
}

// TestFileExtensionConversion tests if file extensions are properly changed based on target language
func TestFileExtensionConversion(t *testing.T) {
	tests := []struct {
		name           string
		sourceLang     string
		targetLang     string
		inputFile      string
		expectedOutput string
	}{
		{
			name:           "Go to Python",
			sourceLang:     "Go",
			targetLang:     "python",
			inputFile:      "main.go",
			expectedOutput: "main.py",
		},
		{
			name:           "JavaScript to TypeScript",
			sourceLang:     "JavaScript",
			targetLang:     "typescript",
			inputFile:      "script.js",
			expectedOutput: "script.ts",
		},
		{
			name:           "Python to Java",
			sourceLang:     "Python",
			targetLang:     "java",
			inputFile:      "app.py",
			expectedOutput: "app.java",
		},
		{
			name:           "C++ to Rust",
			sourceLang:     "C++",
			targetLang:     "rust",
			inputFile:      "utils.cpp",
			expectedOutput: "utils.rs",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test the getTargetExtension function
			newExt := getTargetExtension(tt.targetLang)
			
			// Apply the extension change
			result := changeExtension(tt.inputFile, newExt)
			
			if result != tt.expectedOutput {
				t.Errorf("changeExtension() = %v, want %v", result, tt.expectedOutput)
			}
		})
	}
}

// TestLanguageDetection tests the language detection from file extensions
func TestLanguageDetection(t *testing.T) {
	tests := []struct {
		filename      string
		wantLanguage  string
		wantProcessed bool
	}{
		{"example.go", "Go", true},
		{"script.js", "JavaScript", true},
		{"app.py", "Python", true},
		{"style.css", "", false},
		{"data.json", "", false},
		{"readme.md", "", false},
		{"main.java", "Java", true},
		{"utils.cpp", "C++", true},
		{"program.cs", "C#", true},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			gotLanguage, gotProcessed := detectLanguage(tt.filename)
			if gotLanguage != tt.wantLanguage || gotProcessed != tt.wantProcessed {
				t.Errorf("detectLanguage(%q) = (%q, %v), want (%q, %v)",
					tt.filename, gotLanguage, gotProcessed, tt.wantLanguage, tt.wantProcessed)
			}
		})
	}
}

// TestConvertCode tests the code conversion with mocked GPT response
func TestConvertCode(t *testing.T) {
	// Set up mock GPT response
	mockResponse := "def hello_world():\n    print('Hello, World!')"
	cleanup := setupMockGPT(mockResponse, nil)
	defer cleanup()

	// Test data
	sourceCode := "func helloWorld() {\n\tfmt.Println(\"Hello, World!\")\n}"
	sourceLang := "Go"
	targetLang := "python"
	filePath := "hello.go"

	// Create converter
	converter := NewConverter("input", "output", targetLang)

	// Test conversion
	convertedCode, newExt, err := converter.convertCode(sourceCode, sourceLang, filePath)
	
	// Assertions
	if err != nil {
		t.Errorf("convertCode() unexpected error: %v", err)
	}
	
	if newExt != ".py" {
		t.Errorf("convertCode() extension = %v, want %v", newExt, ".py")
	}
	
	if !strings.Contains(convertedCode, mockResponse) {
		t.Errorf("convertCode() result does not contain expected content")
	}
}

// TestConvertFile tests the individual file conversion
func TestConvertFile(t *testing.T) {
	// Create temp directories
	tempInput, err := os.MkdirTemp("", "converter-test-input")
	if err != nil {
		t.Fatalf("Failed to create temp input dir: %v", err)
	}
	defer os.RemoveAll(tempInput)
	
	tempOutput, err := os.MkdirTemp("", "converter-test-output")
	if err != nil {
		t.Fatalf("Failed to create temp output dir: %v", err)
	}
	defer os.RemoveAll(tempOutput)
	
	// Create test file
	inputFilePath := filepath.Join(tempInput, "test.go")
	testCode := "package main\n\nfunc main() {\n\tfmt.Println(\"Hello, World!\")\n}"
	if err := os.WriteFile(inputFilePath, []byte(testCode), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	
	// Set up mock GPT response
	mockResponse := "def main():\n    print('Hello, World!')"
	cleanup := setupMockGPT(mockResponse, nil)
	defer cleanup()
	
	// Convert the file
	err = ConvertFile(inputFilePath, tempOutput, "python")
	if err != nil {
		t.Fatalf("ConvertFile() error = %v", err)
	}
	
	// Check output file
	expectedOutputPath := filepath.Join(tempOutput, "test.py")
	if _, err := os.Stat(expectedOutputPath); os.IsNotExist(err) {
		t.Errorf("Expected output file %s does not exist", expectedOutputPath)
	}
	
	outputContent, err := os.ReadFile(expectedOutputPath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}
	
	if !strings.Contains(string(outputContent), mockResponse) {
		t.Errorf("Output file does not contain expected content")
	}
}

// TestDirectoryConversion tests the conversion of a directory structure
func TestDirectoryConversion(t *testing.T) {
	// Create temp directories
	tempInput, err := os.MkdirTemp("", "converter-test-input-dir")
	if err != nil {
		t.Fatalf("Failed to create temp input dir: %v", err)
	}
	defer os.RemoveAll(tempInput)
	
	tempOutput, err := os.MkdirTemp("", "converter-test-output-dir")
	if err != nil {
		t.Fatalf("Failed to create temp output dir: %v", err)
	}
	defer os.RemoveAll(tempOutput)
	
	// Create subdirectory
	subDir := filepath.Join(tempInput, "subdir")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("Failed to create subdirectory: %v", err)
	}
	
	// Create test files
	file1 := filepath.Join(tempInput, "main.go")
	file2 := filepath.Join(subDir, "utils.go")
	
	if err := os.WriteFile(file1, []byte("package main"), 0644); err != nil {
		t.Fatalf("Failed to create test file 1: %v", err)
	}
	
	if err := os.WriteFile(file2, []byte("package utils"), 0644); err != nil {
		t.Fatalf("Failed to create test file 2: %v", err)
	}
	
	// Set up mock GPT response
	mockResponse := "# Converted code"
	cleanup := setupMockGPT(mockResponse, nil)
	defer cleanup()
	
	// Create converter and convert directory
	converter := NewConverter(tempInput, tempOutput, "python")
	err = converter.Convert()
	
	if err != nil {
		t.Fatalf("Convert() error = %v", err)
	}
	
	// Check output files
	expectedFile1 := filepath.Join(tempOutput, "main.py")
	expectedFile2 := filepath.Join(tempOutput, "subdir", "utils.py")
	
	if _, err := os.Stat(expectedFile1); os.IsNotExist(err) {
		t.Errorf("Expected output file %s does not exist", expectedFile1)
	}
	
	if _, err := os.Stat(expectedFile2); os.IsNotExist(err) {
		t.Errorf("Expected output file %s does not exist", expectedFile2)
	}
}

// TestIgnoreDirectories tests that certain directories are ignored
func TestIgnoreDirectories(t *testing.T) {
	tests := []struct {
		dirName string
		want    bool
	}{
		{".git", true},
		{"node_modules", true},
		{"vendor", true},
		{"dist", true},
		{"build", true},
		{".idea", true},
		{".vscode", true},
		{"src", false},
		{"lib", false},
		{"app", false},
	}

	for _, tt := range tests {
		t.Run(tt.dirName, func(t *testing.T) {
			got := shouldIgnoreDir(tt.dirName)
			if got != tt.want {
				t.Errorf("shouldIgnoreDir(%q) = %v, want %v", tt.dirName, got, tt.want)
			}
		})
	}
} 