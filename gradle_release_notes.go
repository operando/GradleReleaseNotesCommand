package main

import (
	"flag"

	"github.com/codeskyblue/go-sh"
)

func main() {
	var version = flag.String("v", "", "Gradle version")
	flag.Parse()
	sh.Command("open", "https://docs.gradle.org/"+*version+"/release-notes").Run()
}
