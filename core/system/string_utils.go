package system

import (
    "strings"
)

func bytesToString(b []byte) string {
    return string(b[:])
}

func splitString(s, sep string) *[]string {
    return &strings.Split(s, sep)
}

func splitBytesStringer(b []byte, sep string) *[]string {
    return splitString(bytesToString(b), sep)
}
