package core

import (
    "github.com/a-hilaly/supfile-api/core/data"
    "github.com/a-hilaly/supfile-api/core/system"
)

type CoreInterface interface {
    // Users
    System() (system.SystemInterface, error)
    Data() (data.DataInterface, error)
}

type CoreManager struct{
    sys  system.SystemInterface
    data data.DataInterface
}

func (sfc *CoreManager) System() (system.SystemInterface, error) {
    return sfc.sys, nil
}

func (sfc *CoreManager) Data() (data.DataInterface, error) {
    return sfc.data, nil
}

func ApiManager() CoreInterface {
    return &CoreManager{sys: &SystemManager{}, data: &DataManager{}}
}
