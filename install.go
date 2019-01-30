package main

import (
	"log"
	"os/exec"
	"runtime"
	"sync"
)

func main() {
	list := []string{
		"github.com/acroca/go-symbols",
		"github.com/alecthomas/gometalinter",
		"github.com/cweill/gotests/...",
		"github.com/davidrjenni/reftools/cmd/fillstruct",
		"github.com/fatih/gomodifytags",
		"github.com/go-delve/delve/cmd/dlv",
		"github.com/golangci/golangci-lint/cmd/golangci-lint",
		"github.com/haya14busa/goplay/cmd/goplay",
		"github.com/josharian/impl",
		"github.com/mdempsky/gocode",
		"github.com/mgechev/revive",
		"github.com/ramya-rao-a/go-outline",
		"github.com/rogpeppe/godef",
		"github.com/sourcegraph/go-langserver",
		"github.com/sqs/goreturns",
		"github.com/stamblerre/gocode",
		"github.com/tylerb/gotype-live",
		"github.com/uudashr/gopkgs/cmd/gopkgs",
		"github.com/zmb3/gogetdoc",
		"golang.org/x/lint/golint",
		"golang.org/x/tools/cmd/goimports",
		"golang.org/x/tools/cmd/gorename",
		"golang.org/x/tools/cmd/guru",
		"honnef.co/go/tools/...",
		"winterdrache.de/goformat/goformat",
	}
	var wg sync.WaitGroup
	cpus := runtime.NumCPU()
	semaphore := make(chan struct{}, cpus)
	for _, path := range list {
		wg.Add(1)
		go func() {
			defer wg.Done()
			semaphore <- struct{}{}
			err := exec.Command("go", "install", path).Run()
			if err != nil {
				log.Printf("install error %s: %s\n", path, err)
			}
			<-semaphore
		}()
	}
	wg.Wait()
}
