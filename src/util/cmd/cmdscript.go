package cmd

import (
    "log"
    "strings"
    "github.com/metux/freecity/util"
)

func RunScriptStr(h util.CmdHandler, abort bool, id string, script string) bool {
    ret := true
    for _,l := range strings.Split(script, "\n") {
        ret = ret && RunScriptCmd(h, id, l)
        if abort && (!ret) {
            log.Println("last cmd failed, aborting script")
            return false
        }
    }
    return ret
}

func RunScriptCmd(h util.CmdHandler, id string, c0 string) bool {
    return h.HandleCmd(Split(c0), id)
}
