package utils

import (
	"fmt"
	"os"
)

func PrintHelp() {
	fmt.Println("ﾋﾟｷﾞﾓﾝｺﾞ ʕ◔ϖ◔ʔ")
	fmt.Println("pigiTree Generate a directory tree up to the specified depth.")
	fmt.Println("Usage: pigiTree <depth> [path]")
	fmt.Println("If no path is specified, the current directory will be used.")
	fmt.Println("Example: pigiTree 3")
	fmt.Println("Example with path: pigiTree 3 /path/to/dir")
	os.Exit(0)
}

func PrintError(message string) {
	fmt.Println("Error:", message)
	os.Exit(1)
}
