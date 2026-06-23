package tree

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type Options struct {
	Writer   io.Writer
	MaxDepth int
	All      bool
	DirsOnly bool
	FullPath bool
	NoReport bool
}

func PrintTree(root string, opts Options) error {
	w := opts.Writer
	if w == nil {
		w = os.Stdout
	}

	abs, err := filepath.Abs(root)
	if err != nil {
		return fmt.Errorf("resolving path %s: %w", root, err)
	}

	info, err := os.Stat(abs)
	if err != nil {
		return fmt.Errorf("accessing path %s: %w", root, err)
	}

	p := &printer{
		w:    w,
		opts: opts,
		root: abs,
	}

	if !info.IsDir() {
		fmt.Fprintln(w, root)
		return nil
	}

	fmt.Fprintln(w, root)
	if err := p.printTreeChildren(abs, 1, ""); err != nil {
		return err
	}
	p.summary()
	return nil
}

type printer struct {
	w     io.Writer
	opts  Options
	root  string
	dirs  int
	files int
}

func (p *printer) printTreeChildren(parentPath string, depth int, prefix string) error {
	if p.opts.MaxDepth >= 0 && depth > p.opts.MaxDepth {
		return nil
	}

	entries, err := os.ReadDir(parentPath)
	if err != nil {
		return fmt.Errorf("reading directory %s: %w", parentPath, err)
	}

	entries = filterAndSort(entries, p.opts.All, p.opts.DirsOnly)

	if len(entries) == 0 {
		return nil
	}

	for i, entry := range entries {
		connector, continuation := "├── ", "│   "
		if i == len(entries)-1 {
			connector, continuation = "└── ", "    "
		}

		displayName := entry.Name()
		isDir := entry.IsDir()
		fullPath := filepath.Join(parentPath, entry.Name())

		if entry.Type()&os.ModeSymlink != 0 {
			target, err := os.Readlink(fullPath)
			if err == nil {
				displayName += " -> " + target
			}
			isDir = false
		}

		if p.opts.FullPath {
			rel, _ := filepath.Rel(p.root, fullPath)
			displayName = rel
		}

		fmt.Fprintln(p.w, prefix+connector+displayName)

		if isDir {
			p.dirs++
			if err := p.printTreeChildren(fullPath, depth+1, prefix+continuation); err != nil {
				return err
			}
		} else {
			p.files++
		}
	}

	return nil
}

func (p *printer) summary() {
	if p.opts.NoReport {
		return
	}
	dirStr := "directory"
	if p.dirs != 1 {
		dirStr = "directories"
	}
	if p.opts.DirsOnly {
		fmt.Fprintf(p.w, "\n%d %s\n", p.dirs, dirStr)
		return
	}
	fileStr := "file"
	if p.files != 1 {
		fileStr = "files"
	}
	fmt.Fprintf(p.w, "\n%d %s, %d %s\n", p.dirs, dirStr, p.files, fileStr)
}

func filterAndSort(entries []os.DirEntry, all, dirsOnly bool) []os.DirEntry {
	filtered := entries[:0]
	for _, e := range entries {
		if !all && isHidden(e.Name()) {
			continue
		}
		filtered = append(filtered, e)
	}

	slices.SortFunc(filtered, func(a, b os.DirEntry) int {
		aDir, bDir := a.IsDir(), b.IsDir()
		if aDir != bDir {
			if aDir {
				return -1
			}
			return 1
		}
		return strings.Compare(a.Name(), b.Name())
	})

	if dirsOnly {
		filtered = slices.DeleteFunc(filtered, func(e os.DirEntry) bool {
			return !e.IsDir()
		})
	}

	return filtered
}

func isHidden(name string) bool {
	return len(name) > 0 && name[0] == '.'
}
