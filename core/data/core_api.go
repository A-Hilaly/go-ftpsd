package data

import (
    "sync"
)

type DataInterface interface {
    // Data Interface config
    InitConfig()
    SetConfig(cfg *dataConfig)
    GetConfig() dataConfig

    // Account data
    GetUserAccountID(un, value string) (string, error)
    GetUser(accountid string) (*User, error)

    // Login and subscription
    CreateUser(user, email, type_, authtype, pass string) (*User, error)
    UserExist(un, value string) (bool, error)
    BasicAuthUser(un, value, pass string) (*User, bool, error)
    TokenizedAuthUser(email, authtype, pass string) (*User, bool, error)

    // Deletion and management
    DropUser(accountid string) error
    ActivateUser(un, value string) error
    DeactivateUser(un, value string) error

    // User updates
    // General updates
    UpdateUser(accountid string, field string, value interface{}) error
    UpdateUserMap(accountid string, values map[string]interface{}) error

    // Specifique updates
    ChangeUserPassword(accountid, npass string) error
    ChangeUserAuthMethod(accountid, newauth, pass string) error
    ChangeUserName(accountid, nuser string) error
    ChangeUserEmail(accountid, nemail string) error
    ChangeUserFLname(accountid, nfirst, nlast string) error

    // User control
    UpdateUserQuota(accountid string, nquota int) error
}

type dataConfig struct {
    mutex       sync.Mutex

    AllowAccess bool
    AllowRW     bool
    AllowRead   bool
    AllowWrite  bool
    AllowDrop   bool
    AllowUpdate bool
}

func (dc *dataConfig) lock() {dc.mutex.Lock()}

func (dc *dataConfig) unlock() {dc.mutex.Unlock()}

func defaultConfig() *dataConfig {
    return &dataConfig{
        AllowAccess : true,
        AllowRW     : true,
        AllowRead   : true,
        AllowWrite  : true,
        AllowDrop   : true,
        AllowUpdate : true,
    }
}

func NewDataConfig(aa, arw, ar, aw, ad, au bool) *dataConfig{
    return &dataConfig{
        AllowAccess : aa,
        AllowRW     : arw,
        AllowRead   : ar,
        AllowWrite  : aw,
        AllowDrop   : ad,
        AllowUpdate : au,
    }
}

type DataManager struct {
    mutex  sync.Mutex
    id     string
    config *dataConfig
}

func NewManager(id string) DataInterface {
    return &DataManager{id : id, config: defaultConfig()}
}

func NewManagerWithConfig(id string, dc *dataConfig) DataInterface {
    return &DataManager{id : id, config: dc}
}

// Interface Implementation

func (dm DataManager) InitConfig() {
    dm.mutex.Lock()
    defer dm.mutex.Unlock()
    dm.config = defaultConfig()
}

func (dm DataManager) SetConfig(cfg *dataConfig) {
    dm.mutex.Lock()
    defer dm.mutex.Unlock()
    dm.config = cfg
}

func (dm DataManager) GetConfig() dataConfig {
    return *dm.config
}

func (dm DataManager) GetUserAccountID(un, value string) (string, error){
    if dm.config.AllowAccess && dm.config.AllowRW {
        user, err := selectUserBy(un, value)
        if err != nil {
            return "", err
        }
        return user.AccountId, nil
    }
    return "", ErrorRuleNotAllowed
}
func (dm DataManager) GetUser(accountid string) (*User, error) {
    if dm.config.AllowAccess && dm.config.AllowRW {
        return selectUserBy("account_id", accountid)
    }
    return nil, ErrorRuleNotAllowed
}

func (dm DataManager) CreateUser(user, email, type_, authtype, pass string) (*User, error) {
    if dm.config.AllowAccess && dm.config.AllowRW {
        if authtype == "simple" {
            return newUser(user, email, pass, type_, authtype, "")
        } else {
            return newUser(user, email, "", type_, authtype, pass)
        }
    }
    return nil, ErrorRuleNotAllowed
}

func (dm DataManager) BasicAuthUser(un, value, pass string) (*User, bool, error) {
    if dm.config.AllowAccess && dm.config.AllowRW {
        return defaultAuthentification(un, value, pass)
    }
    return nil, false, ErrorRuleNotAllowed
}

func (dm DataManager) TokenizedAuthUser(email, authtype, pass string) (*User, bool, error) {
    if dm.config.AllowAccess && dm.config.AllowRW {
        return tokenizedAuthentification(authtype, email, pass)
    }
    return nil, false, ErrorRuleNotAllowed
}

func (dm DataManager) UserExist(un, value string) (bool, error) {
    if dm.config.AllowAccess && dm.config.AllowRW {
        return userExistBy(un, value)
    }
    return false, ErrorRuleNotAllowed
}

func (dm DataManager) DropUser(accountid string) error {
    if dm.config.AllowAccess && dm.config.AllowRW {
        return dropUserBy("account_id", accountid)
    }
    return ErrorRuleNotAllowed
}

func (dm DataManager) ActivateUser(un, value string) error {
    if dm.config.AllowAccess && dm.config.AllowRW {
        return updateUserBy(un, value, "state", "active")
    }
    return ErrorRuleNotAllowed
}

func (dm DataManager) DeactivateUser(un, value string) error {
    if dm.config.AllowAccess && dm.config.AllowRW {
        return updateUserBy(un, value, "state", "inactive")
    }
    return ErrorRuleNotAllowed
}

func (dm DataManager) UpdateUser(accountid, field string, value interface{}) error {
    if dm.config.AllowAccess && dm.config.AllowRW {
        return updateUserBy("account_id", accountid, field, value)
    }
    return ErrorRuleNotAllowed
}

func (dm DataManager) UpdateUserMap(accountid string, values map[string]interface{}) error{
    if dm.config.AllowAccess && dm.config.AllowRW {
        return updateUserMapBy("account_id", accountid, values)
    }
    return ErrorRuleNotAllowed
}

func (dm DataManager) ChangeUserPassword(accountid, npass string) error {
    if dm.config.AllowAccess && dm.config.AllowRW {
        return updateUserBy("account_id", accountid, "password", npass)
    }
    return ErrorRuleNotAllowed
}

func (dm DataManager) ChangeUserAuthMethod(accountid, newauth, pass string) error {
    return ErrorNotImplemented
}

func (dm DataManager) ChangeUserName(accountid, nuser string) error {
    if dm.config.AllowAccess && dm.config.AllowRW {
        return updateUserBy("account_id", accountid, "username", nuser)
    }
    return ErrorRuleNotAllowed
}

func (dm DataManager) ChangeUserEmail(accountid, nemail string) error {
    if dm.config.AllowAccess && dm.config.AllowRW {
        return updateUserBy("account_id", accountid, "email", nemail)
    }
    return ErrorRuleNotAllowed
}

func (dm DataManager) ChangeUserFLname(accountid, nfirst, nlast string) error {
    if dm.config.AllowAccess && dm.config.AllowRW {
        return updateUserMapBy("account_id", accountid, map[string]interface{}{"firstname": nfirst, "lastname": nlast})
    }
    return ErrorRuleNotAllowed
}

func (dm DataManager) UpdateUserQuota(accountid string, nquota int) error {
    if dm.config.AllowAccess && dm.config.AllowRW {
        return updateUserBy("account_id", accountid, "max_storage", nquota)
    }
    return ErrorRuleNotAllowed
}

/*
func (dm DataManager) () {

}
*/
