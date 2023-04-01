package items

type TileRef struct {
    Position      Point
    Tile       *  Tile
    TerrainMap *  TerrainMap
}

func (tr * TileRef) PowerConnected() bool {
    return tr.Tile.PowerConnected()
}

func (tr * TileRef) Surrounding() TileSet {
    tr2,_ := tr.TerrainMap.TileRange(
        Rect{tr.Position.X - 1, tr.Position.Y - 1, 3, 3},
        true)
    return tr2
}
