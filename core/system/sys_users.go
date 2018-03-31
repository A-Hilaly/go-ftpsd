package system

import "github.com/a-hilaly/supfile-api/core/system/syscall"

func AddUser(group, user, pass string) error {
    if exist, err := UserExist(user); exist != false {
        return err
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

func UserExist(user string) (bool, error) {
    cmd := syscall.New(id, user)
    err := cmd.Run()
    if err != nil {
        return false, ErrorUserDosentExist
    }
    return true, nil
}

func DelUser(user string) error {
    if exist, err := UserExist(user); exist == false {
        return err
    }
    cmd := syscall.New(sudo, userdel, "-f", user)
    return cmd.Run()
}
