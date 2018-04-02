package core

import (
    "sync"
)
type SystemInterface interface {
    // Config
    InitConfig()
    SetConfig(sc *systemConfig)
    GetConfig() systemConfig

    InitStats()
    SetStats(sc systemStats)
    GetStats() systemStats
    MakeStats() error
    UpdateStats() error

    // Group
    AddGroup(name string) error
    GetGroups() (*[]string, error)
    DelGroup(name string) error
    GroupExist(name string) (bool, error)
    RenameGroup(old, nname string) error

    // User
    AddUserFtp(user, pass string) error
    AddUser(user, pass string) error
    UserExist(user string) (bool, error)
    DelUser(user string) error
    AddUserToGroup(user, group string) error
    RemoveUserFromGroup(user, group string) error

    ChangeUserName(user, nuser string) error
    ChangeUserPassword(user, npass string) error
    CleanUserDirectory(user string) error
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

func LoadSystemStats() (systemStats, error) {

}

func defaultsstats() systemStats {
    s, err := LoadSystemStats()
    return s
}

func (ss *systemStats) loadFromStats(newstats sytemConfig) {
    ss.lock()
    defer ss.unlock()
    ss.Ulimit        = newstats.Ulimit
    ss.ActualStorage = newstats.ActualStorage
    ss.ActualFTPCxn  = newstats.ActualFTPCxn
    ss.NumberOfUsers = newstats.NumberOfUsers
    ss.SFTPstatus    = newstats.SFTPstatus
}

func (ss *systemStats) Load() error {
    newstats, err := LoadSystemStats()
    ss.lock()
    defer ss.Unlock()
    ss.loadFromStats(newstats)
}

func (ss *systemStats) Set(conf systemConfig) error {
    ss.lock()
    defer ss.unlock()
    ss.Ulimit        = newstats.Ulimit
    ss.ActualStorage = newstats.ActualStorage
    ss.ActualFTPCxn  = newstats.ActualFTPCxn
    ss.NumberOfUsers = newstats.NumberOfUsers
    ss.SFTPstatus    = newstats.SFTPstatus
}

func (ss *systemStats) Make() error {

}


type SystemManager struct {
    mutex  sync.Mutex
    id     string
    config *systemConfig
    stats  *systemStats
}

func NewManager(id string) SystemInterface {
    return &SystemManager{
        id     : id,
        config : defaultConfig(),
        stats  : defaultsstats(),
    }
}

func NewManagerWithConfig(id string, sc *systemConfig) SystemInterface {
    return &SystemManager{
        id     : id,
        config : sc,
        stats  : defaultsstats(),
    }
}

func (sm *SystemManager) InitConfig() {
    sm.mutex.Lock()
    defer sm.mutex.Unlock()
    sm.config = defaultConfig()
}

func (sm *SystemManager) SetConfig(cfg *systemConfig) {
    sm.mutex.Lock()
    defer sm.mutex.Unlock()
    sm.config = cfg
}

func (sm *SystemManager) GetConfig() systemConfig {
    return *sm.config
}

func (sm *SystemManager) InitStats() {
    sm.mutex.Lock()
    defer sm.mutex.Unlock()
    sm.config = defaultsstats()
}

func (sm *SystemManager) GetStats() systemStats {
    return *sm.stats
}

func (sm *SystemManager) SetStats(sc systemStats) {
    sm.loadFromStats(sc)
}

func (sm *SystemManager) MakeStats() error {
    return sm.stats.Make()
}

func (sm *SystemManager) AddGroup(name string) error {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return addGroup(name)
    }
    return ErrorRuleNotAllowed
}

func (sm *SystemManager) GetGroups() (*[]string, error) {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return getGroups()
    }
    return ErrorRuleNotAllowed
}

func (sm *SystemManager) DelGroup(name string) error {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return delGroup(name)
    }
    return ErrorRuleNotAllowed
}

func (sm *SystemManager) GroupExist(name string) (bool, error) {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return groupExist(name)
    }
    return ErrorRuleNotAllowed
}

func (sm *SystemManager) RenameGroup(old, name string) error {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return renameGroup(old, name)
    }
    return ErrorRuleNotAllowed
}

func (sm *SystemManager) AddUser(user, pass string) error {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return addUser(name, pass)
    }
    return ErrorRuleNotAllowed
}

func (sm *SystemManager) AddUserGroup(group, user, pass, string) error {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return addUserSFTP(group, user, pass)
    }
    return ErrorRuleNotAllowed
}

func (sm *SystemManager) UserExist(user string) (bool, error) {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return userExist(user)
    }
    return ErrorRuleNotAllowed
}

func (sm *SystemManager) DelUser(user string) error {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return delUser(user)
    }
    return ErrorRuleNotAllowed
}

func (sm *SystemManager) AddUserToGroup(user, group string) error {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return addUserToGroup(user, group)
    }
    return ErrorRuleNotAllowed
}

func (sm *SystemManager) RemoveUserFromGroup(user, group string) error {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return removeUserFromGroup(name, group)
    }
    return ErrorRuleNotAllowed
}

func (sm *SystemManager) ChangeUserName(user, nuser string) error {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return changeuserName(user, nuser)
    }
    return ErrorRuleNotAllowed
}

func (sm *SystemManager) ChangeUserPassword(user, npass string) error {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return changeUserPassword(user, npass)
    }
    return ErrorRuleNotAllowed
}

func (sm *SystemManager) CleanUserDirectory(user string) error {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return CleanUserDirectory(name)
    }
    return ErrorRuleNotAllowed
}

/*
func (sm *SystemManager) () {

}
*/
