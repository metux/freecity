package items

import (
    "github.com/metux/freecity/core/base"
)

func (tm * TerrainMap) isRoadAt(p point) bool {
    return tm.CheckTileLine(p, base.LineTypeRoad)
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
    if tile := tm.tileForLine(p, act, base.LineTypeRoad, "road"); tile != nil {
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

        tile.SetLine(base.LineTypeRoad, other)
        p.DoOnPointAndSurrounding(tm.updateRoadAt)

        tm.TouchObjects()
        return true
    }
    return false
}
