package tree

import (
	"fmt"
	"os"
	"path/filepath"
)

func PrintTree(root string, depth int) error {
	return printTreeRecursive(root, 0, depth, "")
}

func printTreeRecursive(currentPath string, currentDepth, maxDepth int, prefix string) error {
	if currentDepth > maxDepth {
		return nil
	}

	files, err := os.ReadDir(currentPath)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %v", currentPath, err)
	}

	fmt.Println(prefix + currentPath)

	for i, file := range files {
		newPrefix := prefix
		if i == len(files)-1 {
			newPrefix = newPrefix + "    "
			if file.IsDir() {
				fmt.Println(prefix + "└── " + file.Name())
				printTreeRecursive(filepath.Join(currentPath, file.Name()), currentDepth+1, maxDepth, newPrefix+"    ")
			} else {
				fmt.Println(prefix + "└── " + file.Name())
			}
		} else {
			if file.IsDir() {
				fmt.Println(prefix + "├── " + file.Name())
				printTreeRecursive(filepath.Join(currentPath, file.Name()), currentDepth+1, maxDepth, newPrefix+"│   ")
			} else {
				fmt.Println(prefix + "├── " + file.Name())
			}
		}
	}

	return nil
}
