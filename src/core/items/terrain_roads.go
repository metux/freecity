package items

import (
    "github.com/metux/freecity/core/base"
)

func (tm * TerrainMap) isRoadAt(pos Point) bool {
    if t := tm.tileAt(pos); t != nil {
        return t.HasRoad()
    }
    return false
}

// update the directions of neighboring roads
func (tm * TerrainMap) updateRoadAt(pos Point) {
    if tile := tm.tileAt(pos); tile != nil && tile.Road.Present() {
        // FIXME: need to check for conflicts against powerlines and rails
        tile.Road = base.LineDirectionFromVec(
            tm.isRoadAt(pos.North()),
            tm.isRoadAt(pos.East()),
            tm.isRoadAt(pos.South()),
            tm.isRoadAt(pos.West()))
    }
}

func (tm * TerrainMap) ErrectRoad(pos Point) bool {
    act := Action(base.ActionBuildRoad)
    tile := tm.tileAt(pos)
    if tile == nil {
        tm.emit(act, NotifyNoSuchTile{"road", pos})
        return false
    }

    if tile.Building != nil {
        tm.emit(act, NotifyAlreadyOccupied{"building "+tile.Building.TypeName, pos})
        return false
    }

    // FIXME: check terrain
    other := base.LineDirPick(tile.Power, tile.Rail)
    if other.None() {
        tm.emit(act, NotifyAlreadyOccupied{"powerline/rail", pos})
        return false
    }

    tm.autoBulldoze(act, pos)

    if ! tm.trySpendFunds(act, tm.GeneralRules.Costs.Road, "road") {
        return false
    }

    tile.Road = other
    tm.updateRoadAt(pos)
    tm.updateRoadAt(pos.North())
    tm.updateRoadAt(pos.East())
    tm.updateRoadAt(pos.South())
    tm.updateRoadAt(pos.West())

    tm.TouchObjects()
    return true
}

func (t * TerrainMap) ErrectRoadH(pos Point, w int) {
    for i := 0; i<w; i++ {
        t.ErrectRoad(pos)
        pos.X++
    }
}

func (t * TerrainMap) ErrectRoadV(pos Point, w int) {
    for i := 0; i<w; i++ {
        t.ErrectRoad(pos)
        pos.Y++
    }
}
