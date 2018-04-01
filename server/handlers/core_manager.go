package handlers

import (

    "github.com/a-hilaly/supfile-api/core"
)


var Manager core.CoreManager

func InitManager() {
    Manager = core.NewManager()
}
