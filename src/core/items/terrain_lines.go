package items

import (
    "log"
    "github.com/metux/freecity/core/base"
)

func (tm * TerrainMap) CheckTileLine(p point, lt LineType) bool {
    if tile := tm.tileAt(p); tile != nil {
        return tile.HasLine(lt)
    }
    return false
}

func (tm * TerrainMap) ErrectLine(p point, lt LineType) bool {
    switch lt {
        case base.LineTypePower: return tm.ErrectPowerline(p)
        case base.LineTypeRail:  return tm.ErrectRail(p)
        case base.LineTypeRoad:  return tm.ErrectRoad(p)
    }
    log.Println("ErrectLine: unsupported line type", lt)
    return false
}
