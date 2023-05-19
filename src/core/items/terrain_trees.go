package items

// just place, w/o billing
func (tm * TerrainMap) PlaceWood(p point, w byte) bool {
    return tm.ModifyTile(p, func (tile * Tile) bool {
        if tile.Wood < w {
            tile.Wood = w
            return true
        }
        return false
    })
}

// clean zone w/ billing
func (tm * TerrainMap) CleanWood(act Action, p point) bool {
    return tm.ModifyTile(p, func (tile * Tile) bool {
        if tile.Wood > 0 {
            if tm.trySpendFunds(act, tm.GeneralRules.Costs.Bulldoze, "clean wood") {
                tile.Wood = 0
                return true
            }
        }
        return false
    })
}
