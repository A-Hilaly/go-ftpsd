package system

import "github.com/a-hilaly/supfile-api/core/system/syscall"


func isOnline() error {
    cmd := syscall.New(curl, googleip, "-m", "1")
    // Good byte
    return cmd.Run()
}
