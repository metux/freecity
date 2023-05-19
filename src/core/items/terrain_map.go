package items

import (
    "log"
    "errors"
    "github.com/metux/freecity/core/base"
    "github.com/metux/freecity/core/rules"
)

type TerrainMap struct {
    Size                 point
    Tiles           []   Tile
    Buildings       [] * Building
    PowerGrids      []   PowerGrid
    GeneralRules         rules.GeneralRules
    Funds                Money
    Ruleset              string
    Notify               NotifyHandler
    Date                 date
    RevTerrain           int64
    RevObjects           int64
}

func (tm * TerrainMap) CreateMap() {
    tm.Size  = tm.GeneralRules.Startup.Size
    tm.Funds = tm.GeneralRules.Startup.Funds
    tm.Date  = tm.GeneralRules.Startup.Date
    tm.Tiles = make([]Tile, tm.Size.X * tm.Size.Y)
}

func (tm * TerrainMap) tilePoint(p point) (TileRef, error) {
    if (!tm.Size.HasPoint(p)) {
        return TileRef{}, errors.New("coords out of range: "+p.String())
    }
    return TileRef{
        TerrainMap: tm,
        Tile:       &tm.Tiles[p.Y * tm.Size.X + p.X],
        Position:   p,
    }, nil
}

func (tm * TerrainMap) tileAt(p point) * Tile {
    if (!tm.Size.HasPoint(p)) {
        return nil
    }
    return &tm.Tiles[p.Y * tm.Size.X + p.X]
}

func (tm * TerrainMap) AllTiles() (TileSet) {
    tiles,_ := tm.TileRange(tm.Size.SpanRect(), true)
    return tiles
}

func (tm * TerrainMap) TileRange(rect rect, ignore bool) (TileSet, error) {
    buffer := make(TileSet, 0)
    for x := rect.X; x < (rect.X + rect.Width); x++ {
        for y := rect.Y; y < (rect.Y + rect.Height); y++ {
            ref, err := tm.tilePoint(point{x, y})
            if err != nil {
                if ignore {
                    log.Println("reached the border")
                } else {
                    log.Println("error: "+err.Error())
                    return buffer, err
                }
            } else {
                buffer = append(buffer, ref)
            }
        }
    }
    return buffer, nil
}

func (tm * TerrainMap) checkFunds(value base.Money, cause string) bool {
    if tm.Funds < value {
        return false
    }
    return true
}

func (tm * TerrainMap) trySpendFunds(act base.Action, value base.Money, cause string) bool {
    if tm.Funds < value {
        tm.emit(act, NotifyNotEnoughFunds{
            Needed: value,
            Funds:  tm.Funds,
            Cause: cause,
        })
        return false
    }
    tm.Funds -= value
    tm.emit(act, NotifyFundsSpent{
        Spent: value,
        Funds: tm.Funds,
        Cause: cause,
    })
    return true
}

func (tm * TerrainMap) Init(rulesdir string, ruleset string, n NotifyHandler) {
    tm.Ruleset = ruleset
    tm.GeneralRules.LoadYaml(rulesdir+"/"+tm.Ruleset)
    tm.Notify = n
    tm.RevTerrain++
    tm.RevObjects++
}

func (tm * TerrainMap) emit(act Action, n NotifyMsg) {
    if tm.Notify == nil {
        log.Println("WARN: no notify handler for ", n)
    } else {
        tm.Notify.NotifyEmit(act, n)
    }
}

func (tm * TerrainMap) Update(act Action) {
    tm.ConnectBuildings()
    tm.CheckPower(act)
}

func (tm * TerrainMap) autoBulldoze(act Action, pos point) {
    tm.CleanRubble(act, pos)
    tm.CleanWood(act, pos)
}

func (tm * TerrainMap) TouchTerrain() {
    tm.RevTerrain++
    tm.RevObjects++
}

func (tm * TerrainMap) TouchObjects() {
    tm.RevObjects++
}

func (tm * TerrainMap) ModifyTile(p point, f func(tile * Tile) bool) bool {
    if tile := tm.tileAt(p); tile != nil {
        if f(tile) {
            tm.TouchTerrain()
            return true
        }
    }
    return false
}

func (tm * TerrainMap) CheckTile(p point, f func(tile Tile) bool) bool {
    if tile := tm.tileAt(p); tile != nil {
        return f(*tile)
    }
    return false
}
