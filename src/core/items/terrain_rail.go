package items

import (
    "github.com/metux/freecity/core/base"
)

func (tm * TerrainMap) isRailAt(p point) bool {
    if t := tm.tileAt(p); t != nil {
        return t.HasRail()
    }
    return false
}

func (tm * TerrainMap) updateRailAt(p point) {
    tile := tm.tileAt(p)
    if (tile == nil) || (tile.Rail.None()) {
        return
    }

    // FIXME: need to check for conflicts against powerlines and rails
    tile.Rail = base.LineDirectionFromVec(
        tm.isRailAt(p.North()),
        tm.isRailAt(p.East()),
        tm.isRailAt(p.South()),
        tm.isRailAt(p.West()))
}

func (tm * TerrainMap) ErrectRail(p point) (bool) {
    tile := tm.tileAt(p)
    if tile == nil {
        tm.emit(ActionBuildRail, NotifyNoSuchTile{"rail", p})
        return false
    }

    if tile.Building != nil {
        tm.emit(ActionBuildRail, NotifyAlreadyOccupied{"building "+tile.Building.TypeName, p})
        return false
    }

    // FIXME: check terrain
    other := base.LineDirPick(tile.Power, tile.Rail)
    if other.None() {
        tm.emit(ActionBuildRail, NotifyAlreadyOccupied{"powerline/road", p})
        return false
    }

    tm.autoBulldoze(ActionBuildRail, p)

    if ! tm.trySpendFunds(ActionBuildRail, tm.GeneralRules.Costs.Rail, "rail") {
        return false
    }

    tile.Rail = other
    tm.updateRailAt(p)
    tm.updateRailAt(p.North())
    tm.updateRailAt(p.East())
    tm.updateRailAt(p.South())
    tm.updateRailAt(p.West())

    tm.TouchObjects()
    return true
}

func (t * TerrainMap) ErrectRailH(p point, w int) {
    for i := 0; i<w; i++ {
        t.ErrectRail(p)
        p.X++
    }
}

func (t * TerrainMap) ErrectRailV(p point, w int) {
    for i := 0; i<w; i++ {
        t.ErrectRail(p)
        p.Y++
    }
}
