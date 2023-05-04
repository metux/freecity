package games

import (
    "github.com/metux/freecity/core/game"
    "github.com/metux/freecity/core/base"
    "math/rand"
)

type P = base.Point
type R = base.Rect

func placeRubble(g * game.Game) {
    // generate some random rubble
    num_rubble := rand.Intn(10) + 10
    for n := 0; n < num_rubble; n++ {
        g.Terrain.PlaceRubble(base.RandPoint(g.Terrain.Size))
    }
    g.Terrain.PlaceRubble(P{20,20})
}

func TestGame1() * game.Game {
    g1 := game.DefaultConfig.CreateGame()
    terrain := &g1.Terrain

//    placeRubble(g1)

    terrain.ErrectPowerline(P{0,0})
    terrain.ErrectBuilding("powerplant/coal",        P{10, 12})
//    terrain.ErrectBuilding("powerplant/nuclear",     P{10, 20})
//    terrain.ErrectBuilding("residential/skyscraper", P{20, 20})
    terrain.ErrectPowerlineV(P{13, 12}, 18)
    terrain.ErrectPowerlineH(P{12, 22}, 18)
    terrain.ErrectRoadV(P{17, 20}, 10)
    terrain.ErrectRoadV(P{9, 19}, 12)
    terrain.ErrectRoadV(P{40, 19}, 12)
    terrain.ErrectRoadH(P{10, 24}, 30)
    terrain.ErrectRoadH(P{10, 19}, 30)
    terrain.ErrectRoadH(P{10, 30}, 30)
    terrain.ErrectPowerlineV(P{12, 27}, 18)
    terrain.ZoneRect(base.ZoneIndustrialLight,  R{0,  5, 30,  1})
    terrain.ZoneRect(base.ZoneIndustrialDense,  R{3,  2,  2,  2})
    terrain.ZoneRect(base.ZoneResidentialLight, R{1,  1,  1,  1})

    terrain.ZoneRect(base.ZoneCommercialDense,  R{5,  5,  2,  2})
    terrain.ZoneRect(base.ZoneCommercialLight,  R{8,  8,  2,  2})

    terrain.ZoneRect(base.ZoneAirportDense,     R{28, 5, 10, 10})
    terrain.CheckPower(base.ActionGameCreate)

//    terrain.ErrectBuilding("residential/skyscraper", P{35, 35})
    terrain.DemolishBuildingAt(P{35, 35})

    g1.SaveGame(GameFile1)
    return g1
}
