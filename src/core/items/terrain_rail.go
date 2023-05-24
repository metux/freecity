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
    return tm.addLine(base.ActionBuildRail, base.LineTypeRail, p, tm.updateRailAt)
}
