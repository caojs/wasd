package main

import (
	"github.com/caojs/wasd/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
