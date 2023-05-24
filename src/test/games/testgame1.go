package games

import (
    "github.com/metux/freecity/util/geo"
    "github.com/metux/freecity/util/cmd"
    "github.com/metux/freecity/core/game"
    "github.com/metux/freecity/core/base"
    "math/rand"
)

type point = geo.Point
type rect = geo.Rect

const (
    power = base.LineTypePower
    road = base.LineTypeRoad
)

func placeRubble(g * game.Game) {
    // generate some random rubble
    num_rubble := rand.Intn(10) + 10
    for n := 0; n < num_rubble; n++ {
        g.Terrain.PlaceRubble(geo.RandPoint(g.Terrain.Size))
    }
    g.Terrain.PlaceRubble(point{20,20})
}

const script = `
terrain place 0 0   road
terrain place 10 12 building powerplant/coal
terrain place 10 20 building powerplant/nuclear
terrain place 20 20 building residential/skyscraper
terrain place 35 35 building residential/skyscraper

terrain vline 13 12 18 power
terrain hline 12 22 18 power
terrain vline 12 27 18 power

terrain vline 17 20 10 road
terrain vline  9 19 12 road
terrain vline 40 19 12 road
terrain hline 10 24 30 road
terrain hline 10 19 30 road
terrain hline 10 30 30 road

terrain zone  0  5 30  1 i
terrain zone  3  2  2  2 I
terrain zone  1  1  1  1 r
terrain zone  5  5  2  2 C
terrain zone  8  8  2  2 c
terrain zone 28  5 10  10 A
`

func TestGame1() * game.Game {
    g1 := game.DefaultConfig.CreateGame()

    cmd.RunScriptStr(g1, false, "testgame1", script)

    terrain := &g1.Terrain

//    placeRubble(g1)

    terrain.CheckPower(base.ActionGameCreate)

//    terrain.ErrectBuilding("residential/skyscraper", point{35, 35})
    terrain.DemolishBuildingAt(point{35, 35})

    g1.SaveGame(GameFile1)
    return g1
}
