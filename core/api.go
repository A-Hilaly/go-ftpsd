package core

import (
	"github.com/a-hilaly/go-ftpsd/core/data"
	"github.com/a-hilaly/go-ftpsd/core/system"
)

type CoreInterface interface {
	// Users
	System() system.SystemInterface
	Data() data.DataInterface
}

type CoreManager struct {
	sys  system.SystemInterface
	data data.DataInterface
}

func (sfc *CoreManager) System() system.SystemInterface {
	return sfc.sys
}

func (sfc *CoreManager) Data() data.DataInterface {
	return sfc.data
}

func NewManager(id string) CoreInterface {
	return &CoreManager{sys: system.NewManager(id), data: data.NewManager(id)}
}
