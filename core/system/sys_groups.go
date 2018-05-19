package system

import "github.com/a-hilaly/go-ftpsd/core/system/syscall"

func addGroup(name string) error {
	if exist, err := groupExist(name); exist != false {
		return err
	}
	cmd := syscall.New(addgroup, name)
	return cmd.Run()
}

func getGroups() (*[]string, error) {
	cmd := syscall.New(groups)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	grps := bytesToString(out[:len(out)-1])
	ad := splitString(grps, " ")
	return &ad, nil
}

func groupExist(name string) (bool, error) {
	grp, err := getGroups()
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

func delGroup(name string) error {
	if exist, err := groupExist(name); exist == false {
		return err
	}
	cmd := syscall.New(groupdel, name)
	return cmd.Run()
}

func renameGroup(old, name string) error {
	return ErrorNotImplemented
}
