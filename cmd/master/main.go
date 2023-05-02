package master

import (
	"github.com/harleywinston/x-manager/internal/master/app"
)

func SetupMaster() error {
	return app.InitApp()
}
