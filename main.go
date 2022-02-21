package main

import (
	"fmt"
	"os"

	"github.com/nomadicattic/nomadic/cmd"
)

var version string = "v0.0.0"
var hash string = "n/a"

func main() {
	err := cmd.Execute(version, hash)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
