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
    act := base.Action(ActionBuildRail)
    lt  := base.LineType(base.LineTypeRail)
    cb  := tm.updateRailAt

    if tile := tm.tileForLine(p, act, lt); tile != nil {
        other := tile.PickLine(lt)
        if other.None() {
            tm.emit(act, NotifyAlreadyOccupied{"lines", p})
            return false
        }

        tm.autoBulldoze(act, p)

        if ! tm.trySpendFunds(act, tm.GeneralRules.LinePrice(lt), lt.String()) {
            return false
        }

        tile.SetLine(lt, other)
        p.DoOnPointAndSurrounding(cb)

        tm.TouchObjects()
        return true
    }
    return false
}
