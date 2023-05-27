package items

import (
    "math/rand"
    "github.com/metux/freecity/util/geo"
)

func (tm * TerrainMap) PlaceRubble(p point) bool {
    return tm.ModifyTile(p, func (tile * Tile) bool {
        tile.Rubble = true
        return true
    })
}

func (tm * TerrainMap) CleanRubble(act Action, p point) bool {
    return tm.ModifyTile(p, func(tile * Tile) bool {
        if tile.Rubble && tm.trySpendFunds(act, tm.GeneralRules.Costs.Bulldoze, "clean rubbble") {
            tile.Rubble = false
            return true
        }
        return false
    })
}

func (tm * TerrainMap) RandomRubble(base int, max int) {
    num_rubble := rand.Intn(max) + base
    for n := 0; n < num_rubble; n++ {
        tm.PlaceRubble(geo.RandPoint(tm.Size))
    }
}
