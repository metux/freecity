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
    x2 := rect.X + rect.Width
    y2 := rect.Y + rect.Height
    for x := rect.X; x < x2; x++ {
        for y := rect.Y; y < y2; y++ {
            tm.SetZone(zt, point{x,y})
        }
    }
}
