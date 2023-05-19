package items

import (
    "github.com/metux/freecity/core/base"
)

func (tm * TerrainMap) isRoadAt(p point) bool {
    return tm.CheckTile(p, Tile.HasRoad)
//    if t := tm.tileAt(p); t != nil {
//        return t.HasRoad()
//    }
//    return false
}

// update the directions of neighboring roads
func (tm * TerrainMap) updateRoadAt(p point) {
    tm.ModifyTile(p, func(tile * Tile) bool {
        if tile.Road.Present() {
            // FIXME: need to check for conflicts against powerlines and rails
            tile.Road.PickFromSurrounding(p, tm.isRoadAt)
            return true
        }
        return false
    })
}

func (tm * TerrainMap) ErrectRoad(p point) bool {
    act := Action(base.ActionBuildRoad)
    tile := tm.tileAt(p)
    if tile == nil {
        tm.emit(act, NotifyNoSuchTile{"road", p})
        return false
    }

    if tile.Building != nil {
        tm.emit(act, NotifyAlreadyOccupied{"building "+tile.Building.TypeName, p})
        return false
    }

    // FIXME: check terrain
    other := base.LineDirPick(tile.Power, tile.Rail)
    if other.None() {
        tm.emit(act, NotifyAlreadyOccupied{"powerline/rail", p})
        return false
    }

    tm.autoBulldoze(act, p)

    if ! tm.trySpendFunds(act, tm.GeneralRules.Costs.Road, "road") {
        return false
    }

    tile.Road = other

    p.DoOnPointAndSurrounding(tm.updateRoadAt)

    tm.TouchObjects()
    return true
}

func (t * TerrainMap) ErrectRoadH(p point, w int) {
    for i := 0; i<w; i++ {
        t.ErrectRoad(p)
        p.X++
    }
}

func (t * TerrainMap) ErrectRoadV(p point, w int) {
    for i := 0; i<w; i++ {
        t.ErrectRoad(p)
        p.Y++
    }
}
