package items

type TileRef struct {
    Position      point
    Tile       *  Tile
    TerrainMap *  TerrainMap
}

func (tr * TileRef) PowerConnected() bool {
    return tr.Tile.PowerConnected()
}

func (tr * TileRef) Surrounding() TileSet {
    tr2,_ := tr.TerrainMap.TileRange(
        tr.Position.Surrounding(),
        true)
    return tr2
}
