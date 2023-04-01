package cmd

import (
    "strings"
)

func SplitCmd(s string) (string, string) {
    idx := strings.Index(s, ".")
    if idx > -1 {
        return s[:idx],s[idx+1:]
    } else {
        return s,""
    }
}

func SplitCmdline(s string) [] string {
    return strings.Split(s, " ")
}
