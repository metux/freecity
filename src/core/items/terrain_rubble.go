package items

func (tm * TerrainMap) PlaceRubble(p point) {
    if tile := tm.tileAt(p); tile != nil {
        tile.Rubble = true
    }
    tm.TouchTerrain()
}

func (tm * TerrainMap) CleanRubble(act Action, p point) bool {
    if tile := tm.tileAt(p); tile != nil {
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
