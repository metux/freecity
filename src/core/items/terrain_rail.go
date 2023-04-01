package items

import (
    "github.com/metux/freecity/core/base"
)

func (tm * TerrainMap) isRailAt(pos Point) bool {
    if t := tm.tileAt(pos); t != nil {
        return t.HasRail()
    }
    return false
}

func (tm * TerrainMap) updateRailAt(pos Point) {
    tile := tm.tileAt(pos)
    if (tile == nil) || (tile.Rail.None()) {
        return
    }

    // FIXME: need to check for conflicts against powerlines and rails
    tile.Rail = base.LineDirectionFromVec(
        tm.isRailAt(pos.North()),
        tm.isRailAt(pos.East()),
        tm.isRailAt(pos.South()),
        tm.isRailAt(pos.West()))
}

func (tm * TerrainMap) ErrectRail(pos Point) (bool) {
    tile := tm.tileAt(pos)
    if tile == nil {
        tm.emit(ActionBuildRail, NotifyNoSuchTile{"rail", pos})
        return false
    }

    if tile.Building != nil {
        tm.emit(ActionBuildRail, NotifyAlreadyOccupied{"building "+tile.Building.TypeName, pos})
        return false
    }

    // FIXME: check terrain
    other := base.LineDirPick(tile.Power, tile.Rail)
    if other.None() {
        tm.emit(ActionBuildRail, NotifyAlreadyOccupied{"powerline/road", pos})
        return false
    }

    tm.autoBulldoze(ActionBuildRail, pos)

    if ! tm.trySpendFunds(ActionBuildRail, tm.GeneralRules.Costs.Rail, "rail") {
        return false
    }

    tile.Rail = other
    tm.updateRailAt(pos)
    tm.updateRailAt(pos.North())
    tm.updateRailAt(pos.East())
    tm.updateRailAt(pos.South())
    tm.updateRailAt(pos.West())

    tm.TouchObjects()
    return true
}

func (t * TerrainMap) ErrectRailH(pos Point, w int) {
    for i := 0; i<w; i++ {
        t.ErrectRail(pos)
        pos.X++
    }
}

func (t * TerrainMap) ErrectRailV(pos Point, w int) {
    for i := 0; i<w; i++ {
        t.ErrectRail(pos)
        pos.Y++
    }
}
