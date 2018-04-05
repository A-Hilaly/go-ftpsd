package handlers

import (
    "errors"
)

var (
    ErrorNonValidData = errors.New("Data not valid")
)

func validateNonEmpty(ss ...string) bool {
    for _, e := range ss {
        if e == "" {
            return false
        }
    }
    return true
}

func nonNull(s string) bool {return s != ""}
func Null(s string) bool {return s != ""}
