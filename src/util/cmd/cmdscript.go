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

func RunScriptCmd(h util.CmdHandler, id string, cmd string) bool {
    c1 := strings.Split(cmd, "#")
    if c1[0] == "" {
        log.Println("SKIP: ", cmd)
        return true
    }
    return h.HandleCmd(strings.Split(c1[0], " "), id)
}
