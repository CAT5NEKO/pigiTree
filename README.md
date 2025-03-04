# pigiTree
Generate a directory tree up to the specified depth.

![temp](https://github.com/user-attachments/assets/c3ddc0dd-cea7-4b2c-b328-4fcde52b99ec)

When using Tree on Windows, it was not possible to create a directory tree to a specified depth by default (like `tree -L 2
`), so I implemented a similar function in Go.

## Installation

```shell
go build -o pigiTree
```
## Usage

```shell
pigiTree <depth> [path]
# example
pigiTree 2 "/Pigimon/go"

```
