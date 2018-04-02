package system

import (
    "strings"
    "strconv"
)

func bytesToString(b []byte) string {
    return string(b[:])
}

func splitString(s, sep string) []string {
    return strings.Split(s, sep)
}

func splitBytesStringer(b []byte, sep string) []string {
    return splitString(bytesToString(b), sep)
}

func intToString(i int) string {
    return strconv.Itoa(i)
}

func stringToInt(s string) int {
    return strconv.Atoi(s)
}
