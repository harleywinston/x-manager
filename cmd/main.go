package main

import (
	"log"

	"github.com/harleywinston/x-manager/cmd/master"
)

func main() {
	log.Fatal(master.SetupMaster())
}
