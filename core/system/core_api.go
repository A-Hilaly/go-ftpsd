package core

import (
    "sync"
)
type SystemInterface interface {
    // Config
    Init()
    SetConfig()
    GetConfig()

    // User
    AddUser()
    UserExist()
    DelUser()
    ChangeUserGroup()
    ChangeUserName()
    ChangeUserPassword()
    RemoveUserFromGroup()
    CleanUserDirectory()

    // Group
    AddGroup()
    GroupExist()
    DelGroup()
    ChangeGroupName()
}

type systemConfig struct {
    mutex             sync.Mutex

    AllowShellAccess  bool
    AllowSudo         bool

    MaxTotalStorage   int
    MaxStoragePerUser int
    MaxEmailsPerUser  int
    MaxNumberOfUsers  int
    FTPGroup          string
}

func (sc *systemConfig) lock() {sc.mutex.Lock()}

func (sc *systemConfig) unlock() {sc.mutex.Unlock()}

func defaultConfig() *systemConfig {
    return &systemConfig{
        AllowShellAccess  :   true,
        AllowSudo         :   true,
        MaxTotalStorage   :   1000,
        MaxStoragePerUser :     50,
        MaxEmailsPerUser  :    100,
        MaxNumberOfUsers  :    100,
        FTPGroup          : "sftp",
    }
}

func NewSystemConfig(ash, as bool, mts, msu, meu, mnu int, ftp string) {
    return &systemConfig{
        AllowShellAccess  : ash,
        AllowSudo         :  as,
        MaxTotalStorage   : mts,
        MaxStoragePerUser : msu,
        MaxEmailsPerUser  : meu,
        MaxNumberOfUsers  : mnu,
        FTPGroup          : ftp,
    }
}

type systemStats struct {
    mutex          sync.Mutex

    Ulimit         int
    ActualStorage  int
    ActualFTPCxn   int
    NumberOfUsers  int
    SFTPstatus     bool
}

func (ss *systemStats) lock() {ss.mutex.Lock()}

func (ss *systemStats) unlock() {ss.mutex.Unlock()}

func (ss *systemStats) Load() error {

}

func (ss *systemStats) Set() error {

}

func (ss *systemStats) Make() error {

}


type SystemManager struct {
    id     string
    config *systemConfig
    stats  *systemStats
}
