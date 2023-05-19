package items

// FIXME: costs and constraints (tile type, ...)
func (tm * TerrainMap) SetZone(zt ZoneTag, p point) bool {
    tile := tm.tileAt(p)
    if tile == nil {
        tm.emit(ActionZoning, NotifyNoSuchTile{"zone "+string(zt), p})
        return false
    }

    if ! tile.ZoneTag.MayUpgrade(zt) {
        tm.emit(ActionZoning, NotifyAlreadyOccupied{"zone "+string(zt), p})
        return false
    }

    if ! tm.trySpendFunds(ActionZoning, tm.GeneralRules.ZonePrice(zt), "zonetype "+zt.String()) {
        return false
    }

    // fixme: compute correct zone tag / check whether already zoned
    tile.ZoneTag = zt
    tm.TouchObjects()
    return true
}

func (tm * TerrainMap) ZoneRect(zt ZoneTag, rect rect) {
    rect.DoPoints(func (p point) { tm.SetZone(zt, p) })
}
