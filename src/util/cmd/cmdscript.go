package cmd

import (
    "strings"
)

func RunScriptStr(h CmdHandler, abort bool, id string, script string) bool {
    ret := true
    for _,l := range strings.Split(script, "\n") {
        if l != "" {
            if ! RunScriptCmd(h, id, l) {
                if abort {
                    return false
                } else {
                    ret = false
                }
            }
        }
    }
    return ret
}

func RunScriptCmd(h CmdHandler, id string, c0 string) bool {
    return h.HandleCmd(Split(c0))
}
