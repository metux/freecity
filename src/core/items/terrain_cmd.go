package items

import (
    "github.com/metux/freecity/util/cmd"
    "github.com/metux/freecity/core/base"
    "log"
)

func (tm * TerrainMap) PlaceAt(p point, cmdline cmd.Cmdline) bool {
    switch cmdline.Str(0,"") {
        case "road":      return tm.ErrectLine(LtRoad,  p)
        case "rail":      return tm.ErrectLine(LtRail,  p)
        case "powerline": return tm.ErrectLine(LtPower, p)
        case "rubble":    return tm.PlaceRubble(p)
    }
    return false
}

func (tm * TerrainMap) handlePlace(c cmd.Cmdline) bool {
    return tm.PlaceAt(
        point{c.Int(0, 0), c.Int(1, 0)},
        cmd.Cmdline(c[2:]))
}

func (tm * TerrainMap) handleZone(c cmd.Cmdline) bool {
    r := rect { c.Int(0, 0), c.Int(1, 0), c.Int(2, 0), c.Int(3, 0) }
    zt := base.ZoneTag(c.Str(4,"")[0])
    log.Println(r, zt)
    return true
}

func (tm * TerrainMap) HandleCmd(c cmd.Cmdline, id string) bool {
    switch c.Str(0, "") {
        case "place": return tm.handlePlace(c[1:])
        case "": return true
    }
    log.Println("terrain: unhandled command:", c)
    return false
}
