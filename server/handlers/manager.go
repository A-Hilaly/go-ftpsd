package handlers

import (

    "github.com/a-hilaly/supfile-api/core"
)


var Manager = InitManager()

func InitManager() core.CoreInterface {
    return core.NewManager("XD")
}
