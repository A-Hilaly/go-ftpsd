package core

import (
    "sync"

    //"github.com/a-hilaly/supfile-api/core/data"
)

type RTConfiguration struct {
    mutex              sync.Mutex
    AllowFtpProtocol   bool
    AllowCreateUserDb  bool
    AllowCreateUserSys bool

    MaxTotalStorage    int
    MaxStoragePerUser  int
    MaxEmailsPerUser   int

    AllowShellAccess   bool
    AllowSudo          bool
    Ulimit             bool
    FTPGroup           string
}

func (rtc *RTConfiguration) Lock() {
    rtc.mutex.Lock()
}

func (rtc *RTConfiguration) Unlock() {
    rtc.mutex.Unluck()
}

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

const (
    rtcDefault = &RTConfiguration{
        AllowFtpProtocol   : true,
        AllowCreateUserDb  : true,
        AllowCreateUserSys : true,

        MaxTotalStorage    : 1000,
        MaxStoragePerUser  : 50,
        MaxEmailsPerUser   : 100,

        AllowShellAccess   : true,
        AllowSudo          : true,
        Ulimit             : true,
        FTPGroup           : true,
    }
)

var (
    RealTime = &LiveStats{}
    RTConfig = &RTConfiguration{}
)

func ConfigureDefault() {
    RTConfig = rtcDefault
}

func ReplaceRTConfig(config *RTConfiguration) {
    RTConfig = config
}
