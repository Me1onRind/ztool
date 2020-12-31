package main

import (
	"log"

	"github.com/Me1onRind/ztool/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
