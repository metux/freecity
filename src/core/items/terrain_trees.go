package items

// just place, w/o billing
func (tm * TerrainMap) PlaceWood(pos Point, w byte) {
    if tile := tm.tileAt(pos); tile != nil {
        if tile.Wood < w {
            tile.Wood = w
        }
    }
    tm.TouchTerrain()
}

// clean zone w/ billing
func (tm * TerrainMap) CleanWood(act Action, pos Point) bool {
    if tile := tm.tileAt(pos); tile != nil {
        if tile.Wood > 0 {
            if tm.trySpendFunds(act, tm.GeneralRules.Costs.Bulldoze, "clean wood") {
                tile.Wood = 0
                return true
            }
        }
    }
    tm.TouchTerrain()
    return false
}
