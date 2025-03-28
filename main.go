package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Define command line flags
	rootPath := flag.String("path", ".", "Root path to start the search")
	outputDir := flag.String("output", "./envs", "Directory to store copied .env files")
	flag.Parse()

	// Create the output directory if it doesn't exist
	err := os.MkdirAll(*outputDir, 0755)
	if err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		return
	}

	// Get absolute path of root to determine relative paths later
	absRootPath, err := filepath.Abs(*rootPath)
	if err != nil {
		fmt.Printf("Error getting absolute path: %v\n", err)
		return
	}

	// Directories to exclude
	excludeDirs := map[string]bool{
		"node_modules":  true,
		".git":          true,
		"vendor":        true,
		"dist":          true,
		"build":         true,
		".vscode":       true,
		".idea":         true,
		"__pycache__":   true,
		"venv":          true,
		".env":          true, // Exclude .env directories (not files)
		"bin":           true,
		"obj":           true, // Common for .NET builds
		"target":        true, // Common for Java/Maven builds
		"coverage":      true,
		"logs":          true,
		"tmp":           true,
		"temp":          true,
		"cache":         true,
		".cache":        true,
		"public/assets": true,
		"public/dist":   true,
		"public/build":  true,
	}

	// Make sure we don't scan our output directory
	excludeDirs[filepath.Base(*outputDir)] = true

	// Keep track of found .env files
	foundFiles := []string{}
	copiedFiles := []string{}

	// Walk through the directory tree
	err = filepath.Walk(absRootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", path, err)
			return nil // Continue walking despite error
		}

		// Check if this is a directory we want to exclude
		if info.IsDir() {
			dirName := filepath.Base(path)
			if excludeDirs[dirName] {
				fmt.Printf("Skipping directory: %s\n", path)
				return filepath.SkipDir
			}
		}

		// Check if file is a .env file (could be .env, .env.local, .env.development, etc.)
		if !info.IsDir() && (strings.HasPrefix(filepath.Base(path), ".env") || strings.HasSuffix(filepath.Base(path), ".env")) {
			foundFiles = append(foundFiles, path)
			fmt.Printf("Found .env file: %s\n", path)

			// Get the relative path from the root path
			relPath, err := filepath.Rel(absRootPath, path)
			if err != nil {
				fmt.Printf("Error getting relative path for %s: %v\n", path, err)
				return nil
			}

			// Create the destination path
			destPath := filepath.Join(*outputDir, relPath)

			// Create destination directory structure
			destDir := filepath.Dir(destPath)
			err = os.MkdirAll(destDir, 0755)
			if err != nil {
				fmt.Printf("Error creating directory structure for %s: %v\n", destPath, err)
				return nil
			}

			// Copy the file
			err = copyFile(path, destPath)
			if err != nil {
				fmt.Printf("Error copying %s to %s: %v\n", path, destPath, err)
			} else {
				fmt.Printf("Copied to: %s\n", destPath)
				copiedFiles = append(copiedFiles, destPath)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %s: %v\n", absRootPath, err)
		return
	}

	// Summary
	fmt.Println("\n--- Summary ---")
	fmt.Printf("Found %d .env files:\n", len(foundFiles))
	for i, file := range foundFiles {
		fmt.Printf("%d. %s\n", i+1, file)
	}

	fmt.Printf("\nSuccessfully copied %d .env files to %s:\n", len(copiedFiles), *outputDir)
	for i, file := range copiedFiles {
		fmt.Printf("%d. %s\n", i+1, file)
	}
}

// Helper function to copy a file
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	// Preserve file permissions
	sourceInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	return os.Chmod(dst, sourceInfo.Mode())
}
