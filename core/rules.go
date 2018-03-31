package core

import (
    "sync"
)

type LiveStats struct {
    mutex          sync.Mutex
    Storage        int
    FTPConnections int
}

func (rt *LiveStats) Update(storage, ftpc int) {
    rt.mutex.Lock()
    defer rt.mutex.Unlock()
    rt.Storage = storage
    rt.FTPConnections = ftpc
}

var (
    RealTime = &LiveStats{}

    mutex              sync.Mutex
    AllowFtpProtocol   bool = true
    AllowCreateUser    bool = true
    AllowCreateUserDb  bool = true
    AllowCreateUserSys bool = true

    MaxTotalStorage     int = 1000
    MaxStoragePerUser   int = 50
    MaxEmailsPerUser    int = 50

    AllowShellAccess   bool = true
    AllowSudo          bool = true
    Ulimit              int = 8192
    FTPGroup         string = "sftp"
)

func LockRules() {
    mutex.Lock()
}

func UnlockRules() {
    mutex.Unlock()
}
