package main

import (
	"flag"
	"fmt"
	"os"

	"pigiTree/tree"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	maxDepth := flag.Int("L", -1, "max display depth of the directory tree")
	all := flag.Bool("a", false, "list all files, including hidden")
	dirsOnly := flag.Bool("d", false, "list directories only")
	fullPath := flag.Bool("f", false, "print full path prefix for each file")
	noReport := flag.Bool("noreport", false, "omit summary report")
	help := flag.Bool("h", false, "show this help")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "ﾋﾟｷﾞﾓﾝｺﾞ ʕ◔ϖ◔ʔ\n")
		fmt.Fprintf(flag.CommandLine.Output(), "pigiTree - a directory tree generator\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options] [path]\n\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *help {
		flag.Usage()
		return nil
	}

	root := "."
	if flag.NArg() > 0 {
		root = flag.Arg(0)
	}

	info, err := os.Stat(root)
	if err != nil {
		return fmt.Errorf("%s does not exist", root)
	}
	if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", root)
	}

	opts := tree.Options{
		Writer:   os.Stdout,
		MaxDepth: *maxDepth,
		All:      *all,
		DirsOnly: *dirsOnly,
		FullPath: *fullPath,
		NoReport: *noReport,
	}

	return tree.PrintTree(root, opts)
}
