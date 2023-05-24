package items

import (
    "github.com/metux/freecity/util/cmd"
    "github.com/metux/freecity/core/base"
    "log"
)

func (tm * TerrainMap) PlaceAt(p point, cmdline cmd.Cmdline) bool {
    switch cmdline.Str(0) {
        case "road":      return tm.ErrectLine(LtRoad,  p)
        case "rail":      return tm.ErrectLine(LtRail,  p)
        case "powerline": return tm.ErrectLine(LtPower, p)
        case "rubble":    return tm.PlaceRubble(p)
    }
    return false
}

func (tm * TerrainMap) handlePlace(c cmd.Cmdline) bool {
    return tm.PlaceAt(
        point{c.Int(0), c.Int(1)},
        cmd.Cmdline(c[2:]))
}

func (tm * TerrainMap) handleZone(c cmd.Cmdline) bool {
    r := rect { c.Int(0), c.Int(1), c.Int(2), c.Int(3) }
    zt := base.ZoneTag(c.Chr(4))
    log.Println("Zoning", r, zt)
    return true
}

func (tm * TerrainMap) HandleCmd(c cmd.Cmdline, id string) bool {
    log.Println("Terrain cmd", c)
    switch c.Str(0) {
        case "": return true
        case "place": return tm.handlePlace(c.Skip(1))
        case "zone":  return tm.handleZone(c.Skip(1))
    }
    log.Println("terrain: unhandled command:", c)
    return false
}
