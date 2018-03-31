package data

import (
    "errors"
)

var (
    ErrorUnkownMethod           = errors.New("DATA: Unkown method")
    ErrorAuthentificationFailed = errors.New("DATA: Authentification failed")
    ErrorNotUnique              = errors.New("DATA: Not unique select by")
    ErrorUserDosentExist        = errors.New("DATA: User doesn't exist")
    ErrorUserAlreadyExist       = errors.New("DATA: User already exist")
    ErrorAuthMethodNotSet       = errors.New("DATA: Authentification method is not given")
)
