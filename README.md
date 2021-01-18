# dir

`dir` is a utility for managing directories.

It facilitates bookmarking and switching between different folders.

## History

This project existed as a series of scripts in my dotfiles for a few years. It began as a Bash script, and was later ported to fish. Since [2019](https://blog.golang.org/using-go-modules), Go has a stable module implementation (i.e., `go.mod` & friends).

The addition of native modules to golang and the fact that `fzf` is used heavily made `go` the natural choice for implementing `dir`.