package system

import "github.com/a-hilaly/go-ftpsd/core/system/syscall"

func setUlimit(limit int) error {
	cmd := syscall.New(ulimit, "-n", intToString(limit))
	return cmd.Run()
}

func getUlimit() (int, error) {
	cmd := syscall.New(ulimit, "-n")
	u, err := cmd.Output()
	if err != nil {
		return 0, err
	}
	s := bytesToString(u[:len(u)-1])
	y, err := stringToInt(s)
	if err != nil {
		return 0, err
	}
	return y, nil
}

func getDf() ([]byte, error) {
	cmd := syscall.New(df)
	u, err := cmd.Output()
	return u, err
}

func getVmStat() ([]byte, error) {
	cmd := syscall.New(vmstat)
	u, err := cmd.Output()
	return u, err
}

func doReboot() error {
	cmd := syscall.New(sudo, reboot)
	// Good byte
	return cmd.Run()
}

func doShutdown() error {
	cmd := syscall.New(sudo, shutdown, "-f", "now")
	// Good byte
	return cmd.Run()
}
