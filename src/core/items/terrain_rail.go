package items

import (
    "github.com/metux/freecity/core/base"
)

func (tm * TerrainMap) isRailAt(p point) bool {
    return tm.CheckTileLine(p, base.LineTypeRail)
}

// FIXME: need to check for conflicts against powerlines and rails
func (tm * TerrainMap) updateRailAt(p point) {
    tm.ModifyTile(p, func (tile * Tile) bool {
        if tile.Rail.Present() {
            tile.Rail.PickFromSurrounding(p, tm.isRailAt)
            return true
        }
        return false
    })
}

func (tm * TerrainMap) ErrectRail(p point) (bool) {
    tile := tm.tileAt(p)
    if tile == nil {
        tm.emit(ActionBuildRail, NotifyNoSuchTile{"rail", p})
        return false
    }

    // FIXME: check terrain
    other := base.LineDirPick(tile.Power, tile.Rail)
    if other.None() {
        tm.emit(ActionBuildRail, NotifyAlreadyOccupied{"powerline/road", p})
        return false
    }

    if tile.Building != nil {
        tm.emit(ActionBuildRail, NotifyAlreadyOccupied{"building "+tile.Building.TypeName, p})
        return false
    }

    tm.autoBulldoze(ActionBuildRail, p)

    if ! tm.trySpendFunds(ActionBuildRail, tm.GeneralRules.Costs.Rail, "rail") {
        return false
    }

    tile.SetLine(base.LineTypeRail, other)
    p.DoOnPointAndSurrounding(tm.updateRailAt)

    tm.TouchObjects()
    return true
}
