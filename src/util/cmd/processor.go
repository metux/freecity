package cmd

import (
    "log"
    "fmt"
)

type ScriptProcessor struct {
    Handler CmdHandler
    Vars    map[string] string
}

func (sp * ScriptProcessor) passCmd(c Cmdline, id string) bool {
    c2 := make(Cmdline,0)
    for _,ent := range c {
        for k,v := range sp.Vars {
            if ent == "$"+k {
                c2 = append(c2, v)
            } else {
                c2 = append(c2, ent)
            }
        }
    }
    return sp.Handler.HandleCmd(c2)
}

func (sp * ScriptProcessor) HandleCmd(c Cmdline, id string) bool {
    if sp.Vars == nil {
        log.Println("initializing ScriptProcessor")
        sp.Vars = make(map[string] string)
    }
    switch c.Str(0) {
        case "":    return true
        case "for": return sp.handleFor(c.Skip(0), id)
        default:
            return sp.Handler.HandleCmd(c)
    }
}

func (sp * ScriptProcessor) handleFor(c Cmdline, id string) bool {
    varname := c.Str(0)
    start := c.Int(1)
    end := c.Int(2)

    for x := start; x<end; x++ {
        sp.Vars[varname] = fmt.Sprintf("%d", x)
    }

    return true
}
