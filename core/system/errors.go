package system

import "errors"

var (
    ErrorNotImplemented          = errors.New("DEV: Not implmented")
    ErrorRuleNotAllowed          = errors.New("RULES: Rule not allowed")
    ErrorSystemMaxStorageReached = errors.New("RULES: system reached max allowed storage")
    ErrorUserMaxStorageReached   = errors.New("RULES: user reached max allowed storage")
    ErrorUserDosentExist         = errors.New("SYSCALL: User doesnt exist")
    ErrorUserAlreadyExist        = errors.New("SYSCALL: User already exist")
    ErrorGroupDoesntExist        = errors.New("SYSCALL: Group doesnt exist")
    ErrorGroupAlreadyExist       = errors.New("SYSCALL: Group already exist")
)
