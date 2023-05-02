package main

import (
	"log"

	cmd "github.com/harleywinston/x-manager/cmd/master"
)

func main() {
	log.Fatal(cmd.SetupMaster())
}
