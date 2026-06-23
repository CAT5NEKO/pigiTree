# pigiTree

ﾋﾟｷﾞﾓﾝｺﾞ ʕ◔ϖ◔ʔ — A cross-platform `tree` command clone with depth control for Windows.

On Windows, the built-in `tree` command doesn't support `-L` (max depth), so this tool fills that.

## Installation

```shell
go build -o pigiTree.exe .
```

Then place `pigiTree.exe` in a directory listed in your `PATH` (e.g. `%USERPROFILE%\go\bin`).

## Usage

```shell
pigiTree [options] [path]

Options:
  -L int    max display depth of the directory tree (default -1 = unlimit)
  -a        list all files, including hidden files
  -d        list directories only
  -f        print full path prefix for each file
  -noreport omit summary report
  -h        show this help ʕ◔ϖ◔ʔ
```
