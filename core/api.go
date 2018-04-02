package core

import (
    "github.com/a-hilaly/supfile-api/core/data"
    "github.com/a-hilaly/supfile-api/core/system"
)

type CoreInterface interface {
    // Users
    System() system.SystemInterface, error
    Data() data.DataInterface
}

type CoreManager struct{
    sys  system.SystemInterface
    data data.DataInterface
}

func (sfc *CoreManager) System() (system.SystemInterface) {
    return sfc.sys
}

func (sfc *CoreManager) Data() (data.DataInterface) {
    return sfc.data
}

func NewManager(id string) CoreInterface {
    return &CoreManager{sys: system.NewManager(id), data: data.NewManager(id)}
}
