package handlers

import (
	"github.com/a-hilaly/go-ftpsd/core"
)

var Manager = InitManager()

func InitManager() core.CoreInterface {
	return core.NewManager("default-id")
}
