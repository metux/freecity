package items

import (
    "github.com/metux/freecity/core/base"
)

func (tm * TerrainMap) isRoadAt(p point) bool {
    return tm.CheckTileLine(p, LtRoad)
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
    return tm.addLine(base.ActionBuildRoad, LtRoad, p, tm.updateRoadAt)
}
