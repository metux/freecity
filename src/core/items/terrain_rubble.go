package items

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
