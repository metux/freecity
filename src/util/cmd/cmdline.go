package cmd

import (
    "strconv"
    "strings"
)

type Cmdline [] string

func (cmd Cmdline) Str (idx uint, df string) string {
    if int(idx)<len(cmd) {
        return cmd[idx]
    }
    return df
}

func (cmd Cmdline) Int (idx uint, df int) int {
    if int(idx)<len(cmd) {
        val,err := strconv.Atoi(cmd[idx])
        if err != nil {
            return val
        }
    }
    return df
}

func Split(s string) Cmdline {
    return Cmdline(strings.Split(strings.Split(s, "#")[0], " "))
}
