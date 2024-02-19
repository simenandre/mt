package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Set the directory to scan; using current directory as default.
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	if err := scanDirectoryForTodos(dir); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// scanDirectoryForTodos scans the given directory for Markdown files and prints all todo items found.
func scanDirectoryForTodos(dir string) error {
	return filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil // Skip directories
		}
		if strings.HasSuffix(path, ".md") {
			return extractAndPrintTodosFromFile(path)
		}
		return nil
	})
}

// extractAndPrintTodosFromFile reads a file and prints lines that contain todo items.
func extractAndPrintTodosFromFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "- [ ]") {
			fmt.Printf("%s: %s\n", filepath.Base(filePath), line)
		}
	}

	return scanner.Err()
}
