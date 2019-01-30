# Install Go tools

Fast setup go tools.

You can install the necessary tools for development in just 2 steps.

## Synopsis

    $ go get -u github.com/Code-Hex/go-install-tools
    $ go-install-tools

## Install list

It follows the following source code

https://github.com/Microsoft/vscode-go/blob/master/src/goInstallTools.ts

### Auto-completion

- `gocode` from `github.com/stamblerre/gocode`
- `gopkgs` from `github.com/uudashr/gopkgs/cmd/gopkgs`

### Go to symbol

- `go-outline` from `github.com/ramya-rao-a/go-outline`
- `go-symbols` from `github.com/acroca/go-symbols`

### Find all references and Go to implementation of symbols

- `guru` from `golang.org/x/tools/cmd/guru`

### Rename symbols

- `gorename` from `golang.org/x/tools/cmd/gorename`

### Modify tags on structs

- `gomodifytags` from `github.com/fatih/gomodifytags`

### The Go playground

- `goplay` from `github.com/haya14busa/goplay/cmd/goplay`

### Stubs for interfaces

- `impl` from `github.com/josharian/impl`

### Show errors as you type

- `gotype-live` from `github.com/tylerb/gotype-live`

### Go to definition

- `godef` from `github.com/rogpeppe/godef`
- `gogetdoc` from `github.com/zmb3/gogetdoc`

### Formatter

- `goimports` from `golang.org/x/tools/cmd/goimports`
- `goreturns` from `github.com/sqs/goreturns`
- `goformat` from `winterdrache.de/goformat/goformat`

### Linter

- `golint` from `golang.org/x/lint/golint`
- `gometalinter` from `github.com/alecthomas/gometalinter`
- `staticcheck` from `honnef.co/go/tools/...`
- `golangci-lint` from `github.com/golangci/golangci-lint/cmd/golangci-lint`
- `revive` from `github.com/mgechev/revive`

### Test
- `gotests` from `github.com/cweill/gotests/...`

### Language Server

- `go-langserver` from `github.com/sourcegraph/go-langserver`

### Debug

- `dlv` from `github.com/go-delve/delve/cmd/dlv`

### Fill structs with defaults

- `fillstruct` from `github.com/davidrjenni/reftools/cmd/fillstruct`

# Contribution

Very easy 3 steps :D

1. Fork [https://github.com/Code-Hex/go-install-tools/fork](https://github.com/Code-Hex/go-install-tools/fork)
2. Commit your changes
3. Create a new Pull Request

I'm waiting for a lot of PR.