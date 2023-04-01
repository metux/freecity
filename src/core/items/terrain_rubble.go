package items

func (tm * TerrainMap) PlaceRubble(pos Point) {
    if tile := tm.tileAt(pos); tile != nil {
        tile.Rubble = true
    }
    tm.TouchTerrain()
}

func (tm * TerrainMap) CleanRubble(act Action, pos Point) bool {
    if tile := tm.tileAt(pos); tile != nil {
        if tile.Rubble {
            if tm.trySpendFunds(act, tm.GeneralRules.Costs.Bulldoze, "clean rubbble") {
                tile.Rubble = false
                return true
            }
        }
    }
    tm.TouchTerrain()
    return false
}
