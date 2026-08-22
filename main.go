// Command volt-demo-cli is the single-CLI-repo demo for volt
// (github.com/khanakia/voltkit): one main package at the repo root, bare
// vX.Y.Z tags, generated workflows and install scripts.
package main

import (
	"flag"
	"fmt"
)

// Stamped by volt via .volt.yml ldflags.vars; "dev" from source builds.
var (
	version = "dev"
	commit  = "unknown"
)

func main() {
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("volt-demo-cli %s (%s)\n", version, commit)
		return
	}
	fmt.Println("hello from volt-demo-cli", version)
}
