package core

import "errors"

var (
    ErrorRuleNotAllowed = errors.New("RULES: Rule not allowed")
    ErrorSystemMaxStorageReached = errors.New("RULES: system reached max allowed storage")
    ErrorUserMaxStorageReached = errors.NEW5("RULES: user reached max allowed storage")
)
