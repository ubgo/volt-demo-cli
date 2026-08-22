// Command volt-demo-cli is the single-CLI-repo demo for volt
// (github.com/khanakia/voltkit): one main package at the repo root, bare
// vX.Y.Z tags, generated workflows and install scripts.
package main

import (
	"flag"
	"fmt"
	"os"
)

// Stamped by volt via .volt.yml ldflags.vars; "dev" from source builds.
var (
	version = "dev"
	commit  = "unknown"
)

func main() {
	// `skills` subcommand: serves this project's published SKILL.md files,
	// always matching this binary's version (voltkit/skillcmd; wiring in
	// skills_gen.go is volt-generated). Dispatched before flag parsing —
	// it owns its own args.
	if len(os.Args) > 1 && os.Args[1] == "skills" {
		cmd := newSkillsCommand()
		cmd.SetArgs(os.Args[2:])
		if err := cmd.Execute(); err != nil {
			os.Exit(1)
		}
		return
	}
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("volt-demo-cli %s (%s)\n", version, commit)
		return
	}
	fmt.Println("hello from volt-demo-cli", version)
}
