package system

import "github.com/a-hilaly/supfile-api/core/system/syscall"

func AddGroup(name string) error {
    if exist, err := GroupExist(name); exist == false {
        return err
    }
    cmd := syscall.New(addgroup, name)
    return cmd.Run()
}

func GetGroups() (*[]string, error) {
    cmd := syscall.New(groups)
    out, err := cmd.Output()
    if err != nil {
        return nil, err
    }
    grps := bytesToString(out)
    ad := splitString(grps, " ")
    return &ad, nil
}

func GroupExist(name string) (bool, error) {
    grp, err := GetGroups()
    if err != nil {
        return false, err
    }
    for _, elem := range *grp {
        if elem == name {
            return true, nil
        }
    }
    return false, ErrorGroupDoesntExist
}

func DelGroup(name string) error {
    if exist, err := GroupExist(name); exist == false {
        return err
    }
    cmd := syscall.New(groupdel, name)
    return cmd.Run()
}
