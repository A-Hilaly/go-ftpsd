package system

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
    AddUserGroup(group, user, pass string) error
    AddUser(user, pass string) error
    UserExist(user string) (bool, error)
    DelUser(user string) error
    AddUserToGroup(user, group string) error
    RemoveUserFromGroup(user, group string) error
    ChangeUserName(user, nuser string) error
    ChangeUserPassword(user, npass string) error
    CleanUserDirectory(user string) error

    // Checks
    CheckOnline() error
    CheckHealth(ip string) error

    // General
    SetUlimit(limit int) error
    GetUlimit() (int, error)
    Df() ([]byte, error)
    VmStat() ([]byte, error)
    Reboot() error
    Shutdown() error
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
        AllowShellAccess  :   false,
        AllowSudo         :   false,
        MaxTotalStorage   :   1000,
        MaxStoragePerUser :     50,
        MaxEmailsPerUser  :    100,
        MaxNumberOfUsers  :    100,
        FTPGroup          : "sftp",
    }
}

func NewSystemConfig(ash, as bool, mts, msu, meu, mnu int, ftp string) *systemConfig {
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
    return systemStats{}, nil
}

func defaultsstats() *systemStats {
    s, _ := LoadSystemStats()
    return &s
}

func (ss *systemStats) loadFromStats(newstats systemStats) {
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
    defer ss.unlock()
    ss.loadFromStats(newstats)
    return err
}

func (ss *systemStats) Set(newstats systemStats) {
    ss.lock()
    defer ss.unlock()
    ss.Ulimit        = newstats.Ulimit
    ss.ActualStorage = newstats.ActualStorage
    ss.ActualFTPCxn  = newstats.ActualFTPCxn
    ss.NumberOfUsers = newstats.NumberOfUsers
    ss.SFTPstatus    = newstats.SFTPstatus
}

func (ss *systemStats) Make() error {
    return nil
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
    sm.stats = defaultsstats()
}

func (sm *SystemManager) GetStats() systemStats {
    return *sm.stats
}

func (sm *SystemManager) SetStats(sc systemStats) {
    sm.stats.loadFromStats(sc)
}

func (sm *SystemManager) MakeStats() error {
    return sm.stats.Make()
}

func (sm *SystemManager) UpdateStats() error {
    return ErrorNotImplemented
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
    return nil, ErrorRuleNotAllowed
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
    return false, ErrorRuleNotAllowed
}

func (sm *SystemManager) RenameGroup(old, name string) error {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return renameGroup(old, name)
    }
    return ErrorRuleNotAllowed
}

func (sm *SystemManager) AddUser(user, pass string) error {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return addUser(user, pass)
    }
    return ErrorRuleNotAllowed
}

func (sm *SystemManager) AddUserGroup(group, user, pass string) error {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return addUserGroup(group, user, pass)
    }
    return ErrorRuleNotAllowed
}

func (sm *SystemManager) UserExist(user string) (bool, error) {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return userExist(user)
    }
    return false, ErrorRuleNotAllowed
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
        return removeUserFromGroup(user, group)
    }
    return ErrorRuleNotAllowed
}

func (sm *SystemManager) ChangeUserName(user, nuser string) error {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return changeUserName(user, nuser)
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
        return cleanUserDirectory(user)
    }
    return ErrorRuleNotAllowed
}

func (sm *SystemManager) CheckOnline() error {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return isOnline()
    }
    return ErrorRuleNotAllowed
}

func (sm *SystemManager) CheckHealth(ip string) error {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return checkHealth(ip)
    }
    return ErrorRuleNotAllowed
}

func (sm *SystemManager) SetUlimit(limit int) error {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return setUlimit(limit)
    }
    return ErrorRuleNotAllowed
}

func (sm *SystemManager) GetUlimit() (int, error) {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return getUlimit()
    }
    return 0, ErrorRuleNotAllowed
}

func (sm *SystemManager) Df() ([]byte, error) {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return getDf()
    }
    return []byte{}, ErrorRuleNotAllowed
}

func (sm *SystemManager) VmStat() ([]byte, error) {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return getVmStat()
    }
    return []byte{}, ErrorRuleNotAllowed
}

func (sm *SystemManager) Reboot() error {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return doReboot()
    }
    return ErrorRuleNotAllowed
}

func (sm *SystemManager) Shutdown() error {
    if sm.config.AllowShellAccess && sm.config.AllowSudo {
        return doShutdown()
    }
    return ErrorRuleNotAllowed
}
/*
func (sm *SystemManager) () {

}
*/
