package data

import (
    "errors"
)

var (
    ErrorNotImplemented         = errors.New("DEV: Not implemented yet")
    ErrorRuleNotAllowed         = errors.New("RULES: Rule not allowed")
    ErrorMaxActiveUsersReached  = errors.New("RULES: max active users reached")
    ErrorMaxDBStorageReached    = errors.New("RULES: max db storage reached")
    ErrorUnkownMethod           = errors.New("DATA: Unkown method")
    ErrorAuthentificationFailed = errors.New("DATA: Authentification failed")
    ErrorNotUnique              = errors.New("DATA: Not unique select by")
    ErrorUserDosentExist        = errors.New("DATA: User doesn't exist")
    ErrorUserAlreadyExist       = errors.New("DATA: User already exist")
    ErrorAuthMethodNotSet       = errors.New("DATA: Authentification method is not given")
    ErrorUsernameExist          = errors.New("DATA: Username already exist")
    ErrorUserEmailExist         = errors.New("DATA: Email already exist")
)
