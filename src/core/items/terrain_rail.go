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
    if tile := tm.tileForLine(p, ActionBuildRail, base.LineTypeRail, "rail"); tile != nil {
        other := tile.PickLine(base.LineTypeRail)
        if other.None() {
            tm.emit(ActionBuildRail, NotifyAlreadyOccupied{"powerline/road", p})
            return false
        }

        tm.autoBulldoze(ActionBuildRail, p)

        if ! tm.trySpendFunds(ActionBuildRail, tm.GeneralRules.LinePrice(base.LineTypeRail), "rail") {
            return false
        }

        tile.SetLine(base.LineTypeRail, other)
        p.DoOnPointAndSurrounding(tm.updateRailAt)

        tm.TouchObjects()
        return true
    }
    return false
}
