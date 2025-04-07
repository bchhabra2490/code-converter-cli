package converter

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// CodeAnalyzer provides basic code analysis capabilities
type CodeAnalyzer struct {
	inputDir string
	indexed  bool
}

// NewCodeAnalyzer creates a new code analyzer for the given directory
func NewCodeAnalyzer(inputDir string) *CodeAnalyzer {
	return &CodeAnalyzer{
		inputDir: inputDir,
		indexed:  false,
	}
}

// Initialize prepares the analyzer
func (a *CodeAnalyzer) Initialize() error {
	fmt.Println("Analyzing source code using simple pattern matching...")
	a.indexed = true
	return nil
}

// FindFunctionDefinitions finds all function definitions in the codebase
func (a *CodeAnalyzer) FindFunctionDefinitions() ([]string, error) {
	if !a.indexed {
		return nil, fmt.Errorf("analyzer not initialized, call Initialize() first")
	}

	var functions []string
	err := filepath.Walk(a.inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if info.IsDir() || shouldIgnoreDir(filepath.Base(path)) {
			return nil
		}
		
		// Only process code files
		lang, isCode := detectLanguage(path)
		if !isCode {
			return nil
		}
		
		// Read the file
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		
		// Do a simple pattern match based on language
		foundFunctions := findFunctionsInCode(string(content), lang, path)
		functions = append(functions, foundFunctions...)
		
		return nil
	})
	
	return functions, err
}

// FindImportsAndDependencies identifies imports and dependencies in the code
func (a *CodeAnalyzer) FindImportsAndDependencies() (map[string][]string, error) {
	if !a.indexed {
		return nil, fmt.Errorf("analyzer not initialized, call Initialize() first")
	}

	// Map of file to its imports
	dependencies := make(map[string][]string)
	
	// Walk the directory tree
	err := filepath.Walk(a.inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if info.IsDir() || shouldIgnoreDir(filepath.Base(path)) {
			return nil
		}
		
		// Only process code files
		lang, isCode := detectLanguage(path)
		if !isCode {
			return nil
		}
		
		// Read the file
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		
		// Extract imports based on language
		imports := findImportsInCode(string(content), lang)
		if len(imports) > 0 {
			relPath, err := filepath.Rel(a.inputDir, path)
			if err != nil {
				relPath = path
			}
			dependencies[relPath] = imports
		}
		
		return nil
	})
	
	return dependencies, err
}

// Helper functions for basic code analysis

// findFunctionsInCode does a very simple pattern match for functions
// Note: This is a naive implementation and would be replaced with proper parsing
func findFunctionsInCode(content, language, path string) []string {
	var functions []string
	lines := strings.Split(content, "\n")
	
	for i, line := range lines {
		line = strings.TrimSpace(line)
		
		// Very simplistic detection based on language
		switch strings.ToLower(language) {
		case "go":
			if strings.HasPrefix(line, "func ") {
				functions = append(functions, fmt.Sprintf("%s:%d:%s", path, i+1, line))
			}
		case "javascript", "typescript":
			if strings.HasPrefix(line, "function ") || strings.Contains(line, " function(") {
				functions = append(functions, fmt.Sprintf("%s:%d:%s", path, i+1, line))
			}
		case "python":
			if strings.HasPrefix(line, "def ") {
				functions = append(functions, fmt.Sprintf("%s:%d:%s", path, i+1, line))
			}
		case "java", "kotlin":
			// This is overly simplified, but just for demonstration
			if (strings.Contains(line, "public ") || strings.Contains(line, "private ") || 
				strings.Contains(line, "protected ")) && strings.Contains(line, "(") {
				functions = append(functions, fmt.Sprintf("%s:%d:%s", path, i+1, line))
			}
		}
	}
	
	return functions
}

// findImportsInCode extracts import statements from code
func findImportsInCode(content, language string) []string {
	var imports []string
	lines := strings.Split(content, "\n")
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Very simplistic detection based on language
		switch strings.ToLower(language) {
		case "go":
			if strings.HasPrefix(line, "import ") {
				imports = append(imports, strings.Trim(strings.TrimPrefix(line, "import "), "\""))
			}
		case "javascript", "typescript":
			if strings.HasPrefix(line, "import ") {
				imports = append(imports, line)
			}
		case "python":
			if strings.HasPrefix(line, "import ") || strings.HasPrefix(line, "from ") {
				imports = append(imports, line)
			}
		case "java":
			if strings.HasPrefix(line, "import ") {
				imports = append(imports, strings.TrimSuffix(strings.TrimPrefix(line, "import "), ";"))
			}
		}
	}
	
	return imports
} 