package items

import (
    "github.com/metux/freecity/util/cmd"
    "github.com/metux/freecity/core/base"
    "log"
)

func (tm * TerrainMap) PlaceAt(p point, c cmd.Cmdline) bool {
    switch c.Str(0) {
        case "road":      return tm.ErrectLine(LtRoad,  p)
        case "rail":      return tm.ErrectLine(LtRail,  p)
        case "powerline": return tm.ErrectLine(LtPower, p)
        case "rubble":    return tm.PlaceRubble(p)
        case "building":  return tm.ErrectBuilding(c.Str(1), p)
    }
    return false
}

func (tm * TerrainMap) handleVLine(c cmd.Cmdline) bool {
    p := point{c.Int(0), c.Int(1)}
    n := c.Int(2)
    switch c.Str(3) {
        case "road":  tm.ErrectLineV(LtRoad, p, n)
        case "rail":  tm.ErrectLineV(LtRail, p, n)
        case "power": tm.ErrectLineV(LtPower, p, n)
        case "pipe":  tm.ErrectLineV(LtPipe, p, n)
        default: return false
    }
    return true
}

func (tm * TerrainMap) handleHLine(c cmd.Cmdline) bool {
    p := point{c.Int(0), c.Int(1)}
    n := c.Int(2)
    switch c.Str(3) {
        case "road":  tm.ErrectLineH(LtRoad, p, n)
        case "rail":  tm.ErrectLineH(LtRail, p, n)
        case "power": tm.ErrectLineH(LtPower, p, n)
        case "pipe":  tm.ErrectLineH(LtPipe, p, n)
        default: return false
    }
    return true
}

func (tm * TerrainMap) handlePlace(c cmd.Cmdline) bool {
    return tm.PlaceAt(point{c.Int(0), c.Int(1)},c.Skip(2))
}

func (tm * TerrainMap) handleZone(c cmd.Cmdline) bool {
    tm.ZoneRect(base.ZoneTag(c.Chr(4)), rect { c.Int(0), c.Int(1), c.Int(2), c.Int(3) })
    return true
}

func (tm * TerrainMap) handleRandRubble(c cmd.Cmdline) bool {
    tm.RandomRubble(c.Int(0), c.Int(1))
    return true
}

func (tm * TerrainMap) HandleCmd(c cmd.Cmdline, id string) bool {
    switch c.Str(0) {
        case "":            return true
        case "place":       return tm.handlePlace(c.Skip(1))
        case "zone":        return tm.handleZone(c.Skip(1))
        case "hline":       return tm.handleHLine(c.Skip(1))
        case "vline":       return tm.handleVLine(c.Skip(1))
        case "randrubble":  return tm.handleRandRubble(c.Skip(1))
    }
    log.Println("terrain: unhandled command:", c)
    return false
}
