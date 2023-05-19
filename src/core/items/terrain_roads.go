package items

import (
    "github.com/metux/freecity/core/base"
)

func (tm * TerrainMap) isRoadAt(p point) bool {
    if t := tm.tileAt(p); t != nil {
        return t.HasRoad()
    }
    return false
}

// update the directions of neighboring roads
func (tm * TerrainMap) updateRoadAt(p point) {
    if tile := tm.tileAt(p); tile != nil && tile.Road.Present() {
        // FIXME: need to check for conflicts against powerlines and rails
        tile.Road = base.LineDirectionFromVec(
            tm.isRoadAt(p.North()),
            tm.isRoadAt(p.East()),
            tm.isRoadAt(p.South()),
            tm.isRoadAt(p.West()))
    }
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
    tm.updateRoadAt(p)
    tm.updateRoadAt(p.North())
    tm.updateRoadAt(p.East())
    tm.updateRoadAt(p.South())
    tm.updateRoadAt(p.West())

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
