package tree

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func PrintTree(root string, depth int) error {
	fmt.Println(root)

	if depth <= 0 {
		return nil
	}

	return printTreeChildren(root, 1, depth, "")
}

func printTreeChildren(parentPath string, currentDepth, maxDepth int, prefix string) error {
	if currentDepth > maxDepth {
		return nil
	}

	entries, err := os.ReadDir(parentPath)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %v", parentPath, err)
	}

	entries = filterAndSort(entries)

	for i, entry := range entries {
		connector := "├── "
		continuation := "│   "
		if i == len(entries)-1 {
			connector = "└── "
			continuation = "    "
		}

		fmt.Println(prefix + connector + entry.Name())

		if entry.IsDir() {
			err := printTreeChildren(
				filepath.Join(parentPath, entry.Name()),
				currentDepth+1,
				maxDepth,
				prefix+continuation,
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func filterAndSort(entries []os.DirEntry) []os.DirEntry {
	var filtered []os.DirEntry
	for _, e := range entries {
		name := e.Name()
		if len(name) > 0 && name[0] != '.' {
			filtered = append(filtered, e)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		iIsDir := filtered[i].IsDir()
		jIsDir := filtered[j].IsDir()
		if iIsDir != jIsDir {
			return iIsDir
		}
		return filtered[i].Name() < filtered[j].Name()
	})

	return filtered
}
