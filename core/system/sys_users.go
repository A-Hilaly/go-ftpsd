package system

import "github.com/a-hilaly/supfile-api/core/system/syscall"

func addUser(user, pass string) error {
    if exist, err := UserExist(user); exist != false {
        return ErrorUserAlreadyExist
    }
    cmd := syscall.New(sudo, adduser, user)
    err := cmd.Run()
    if err != nil {
        return err
    }
    cmd = syscall.New(sudo, usermod, "--password", pass, user)
    err = cmd.Run()
    return err
}

func addUserSFTP(group, user, pass string) error {
    if exist, err := UserExist(user); exist != false {
        return ErrorUserAlreadyExist
    }
    cmd := syscall.New(sudo, mkdir, "-p", "/home/ftp/" + user)
    err := cmd.Run()
    if err != nil {
        return err
    }
    cmd = syscall.New(sudo, useradd,
                      "-s", "/bin/false",
                      "-d", "/home/ftp/" + user,
                      "-g", group,
                      user)
    err = cmd.Run()
    if err != nil {
        return err
    }
    cmd = syscall.New(sudo, usermod, "--password", pass, user)
    err = cmd.Run()
    if err != nil {
        return err
    }
    cmd = syscall.New(sudo, chown, "root:root", "/home/ftp/" + user)
    err = cmd.Run()
    if err != nil {
        return err
    }
    return nil
}

func userExist(user string) (bool, error) {
    cmd := syscall.New(id, user)
    err := cmd.Run()
    if err != nil {
        return false, ErrorUserDosentExist
    }
    return true, nil
}

func delUser(user string) error {
    if exist, err := UserExist(user); exist == false {
        return err
    }
    cmd := syscall.New(sudo, userdel, "-f", user)
    return cmd.Run()
}

func addUserToGroup(user, group string) error {
    if exist, err := UserExist(user); exist == false {
        return err
    }
    cmd := syscall.New(sudo, useradd, "-G", group, user)
    err := cmd.Run()
    return err
}

func removeUserFromGroup(user, group string) error {
    if exist, err := UserExist(user); exist == false {
        return err
    }
    cmd := syscall.New(sudo, deluser, user, group)
    err := cmd.Run()
    return err
}

func userHomeFolder(user) string {
    return "/home/ftp/" + user
}

func changeUserName(user, nuser string) error {
    if exist, err := UserExist(user); exist == false {
        return err
    }
    cmd := syscall.New(sudo, usermod, "-l", nuser, user)
    err := cmd.Run()
    return err
}

func changeUserHomeFolder(user, nuser string) error {
    if exist, err := UserExist(user); exist == false {
        return err
    }
    cmd := syscall.New(sudo, usermod, "-l", nuser, user)
    err := cmd.Run()
    return err
}

func changeUserPassword(user, npass string) error {
    return ErrorNotImplemented
}

func cleanUserDirectory(user) error {

}
