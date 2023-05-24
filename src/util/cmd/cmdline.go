package cmd

import (
    "strconv"
    "strings"
)

type Cmdline [] string

func (cmd Cmdline) StrDef (idx uint, df string) string {
    if int(idx)<len(cmd) {
        return cmd[idx]
    }
    return df
}

func (cmd Cmdline) IntDef (idx uint, df int) int {
    if int(idx)<len(cmd) {
        val,err := strconv.Atoi(cmd[idx])
        if err == nil {
            return val
        }
    }
    return df
}

func (cmd Cmdline) Chr (idx uint) uint8 {
    s := cmd.Str(idx)
    if s == "" {
        return 0
    } else {
        return s[0]
    }
}

func (cmd Cmdline) Str (idx uint) string {
    return cmd.StrDef(idx, "")
}

func (cmd Cmdline) Int (idx uint) int {
    return cmd.IntDef(idx, 0)
}

func (cmd Cmdline) Skip(idx uint) Cmdline {
    return Cmdline(cmd[idx:])
}

func (cmd Cmdline) Head() (string, Cmdline) {
    return cmd.Str(0), cmd.Skip(1)
}

func Split(s string) Cmdline {
    return Cmdline(strings.Fields(strings.Split(s, "#")[0]))
}
