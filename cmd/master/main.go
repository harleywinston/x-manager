package cmd

import (
	master "github.com/harleywinston/x-manager/internal/master/app"
)

func SetupMaster() error {
	return master.InitApp()
}
