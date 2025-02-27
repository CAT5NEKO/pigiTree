package main

import (
	"fmt"
	"os"
	"pigiTree/tree"
	"pigiTree/utils"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		utils.PrintHelp()
		return
	}

	if os.Args[1] == "-h" {
		utils.PrintHelp()
		return
	}

	depth, err := strconv.Atoi(os.Args[1])
	if err != nil {
		utils.PrintError("Please provide a valid number for depth.")
		return
	}

	dir := "."
	if len(os.Args) > 2 {
		dir = os.Args[2]
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		utils.PrintError(fmt.Sprintf("Directory %s does not exist", dir))
		return
	}

	err = tree.PrintTree(dir, depth)
	if err != nil {
		utils.PrintError(fmt.Sprintf("Error generating tree: %v", err))
	}
}
