package main

import (
	_ "github.com/mrjxtr-dev/mr-shush/cmd/cli/generate"
	"github.com/mrjxtr-dev/mr-shush/cmd/cli/root"
)

func main() {
	root.Execute()
}
