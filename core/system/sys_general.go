package system

import "github.com/a-hilaly/supfile-api/core/system/syscall"

func setUlimit(limit int) error {
    cmd := syscall.New(ulimit, "-n", intToString(limit))
    return cmd.Run()
}

func getUlimit() (int, error) {

    cmd := syscall.New(ulimit, "-n", intToString(limit))
    u, err := cmd.Output()
    if err != nil {
        return err
    }
    return stringToInt(bytesToString(u[:len(u)-1])), nil
}

func df() (*[]byte, error){
    cmd := syscall.New(df)
    u, err := cmd.Output()
    return &u, err
}

func vmStat() (*[]byte, error) {
    cmd := syscall.New(vmstat)
    u, err := cmd.Output()
    return &u, err
}

func reboot() error {
    cmd := syscall.New(sudo, reboot)
    // Good byte
    return cmd.Run()
}

func shutdown() error {
    cmd := syscall.New(sudo, shutdown, "-f", "now")
    // Good byte
    return cmd.Run()
}
