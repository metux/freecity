package items

func (tm * TerrainMap) SetZone(zt ZoneTag, p point) bool {
    return tm.ModifyTile(p, func (tile * Tile) bool {
        if ! tile.ZoneTag.MayUpgrade(zt) {
            tm.emit(ActionZoning, NotifyAlreadyOccupied{"zone "+string(zt), p})
            return false
        }

        if ! tm.trySpendFunds(ActionZoning, tm.GeneralRules.ZonePrice(zt), "zonetype "+zt.String()) {
            return false
        }

        tile.ZoneTag = zt
        return true
    })
}

func (tm * TerrainMap) ZoneRect(zt ZoneTag, rect rect) {
    rect.DoPoints(func (p point) { tm.SetZone(zt, p) })
}
