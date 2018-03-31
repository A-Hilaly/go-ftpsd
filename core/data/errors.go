package data

import (
    "errors"
)

var (
    ErrorUnkownMethod           = errors.New("Unkown method")
    ErrorAuthentificationFailed = errors.New("Authentification failed")
    ErrorNotUnique              = errors.New("Not unique select by")
    ErrorUserDosentExist        = errors.New("User doesn't exist")
    ErrorUserAlreadyExist       = errors.New("User already exist")
)
