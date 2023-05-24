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
        case LtPower: return tm.ErrectPowerline(p)
        case LtRail:  return tm.ErrectRail(p)
        case LtRoad:  return tm.ErrectRoad(p)
    }
    log.Println("ErrectLine: unsupported line type", lt)
    return false
}

func (t * TerrainMap) ErrectLineH(lt LineType, p point, w int) {
    for i := 0; i<w; i++ {
        t.ErrectLine(p, lt)
        p.X++
    }
}

func (t * TerrainMap) ErrectLineV(lt LineType, p point, w int) {
    for i := 0; i<w; i++ {
        t.ErrectLine(p, lt)
        p.Y++
    }
}

func (tm * TerrainMap) tileForLine(p point, action base.Action, lt LineType) * Tile {
    tile := tm.tileAt(p)

    if tile == nil {
        tm.emit(action, NotifyNoSuchTile{lt.String(), p})
        return nil
    }

    if tile.Building != nil {
        tm.emit(ActionBuildPowerline, NotifyAlreadyOccupied{"building "+tile.Building.TypeName, p})
        return nil
    }

    return tile
}
