package core

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

type SystemManager struct {}
