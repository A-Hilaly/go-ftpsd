package system

import "errors"

var (
    ErrorUserDosentExist   = errors.New("SYSCALL: User doesnt exist")
    ErrorUserAlreadyExist  = errors.New("SYSCALL: User already exist")
    ErrorGroupDoesntExist  = errors.New("SYSCALL: Group doesnt exist")
    ErrorGroupAlreadyExist = errors.New("SYSCALL: Group already exist")
)
