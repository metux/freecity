package games

import (
    "github.com/metux/freecity/util/geo"
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

func TestGame1() * game.Game {
    g1 := game.DefaultConfig.CreateGame()
    terrain := &g1.Terrain

//    placeRubble(g1)

    terrain.ErrectPowerline(point{0,0})
    terrain.ErrectBuilding("powerplant/coal",        point{10, 12})
//    terrain.ErrectBuilding("powerplant/nuclear",     point{10, 20})
//    terrain.ErrectBuilding("residential/skyscraper", point{20, 20})
    terrain.ErrectLineV(power, point{13, 12}, 18)
    terrain.ErrectLineH(power, point{12, 22}, 18)
    terrain.ErrectLineV(road, point{17, 20}, 10)
    terrain.ErrectLineV(road, point{9, 19}, 12)
    terrain.ErrectLineV(road, point{40, 19}, 12)
    terrain.ErrectLineH(road, point{10, 24}, 30)
    terrain.ErrectLineH(road, point{10, 19}, 30)
    terrain.ErrectLineH(road, point{10, 30}, 30)
    terrain.ErrectLineV(power, point{12, 27}, 18)
    terrain.ZoneRect(base.ZoneIndustrialLight,  rect{0,  5, 30,  1})
    terrain.ZoneRect(base.ZoneIndustrialDense,  rect{3,  2,  2,  2})
    terrain.ZoneRect(base.ZoneResidentialLight, rect{1,  1,  1,  1})

    terrain.ZoneRect(base.ZoneCommercialDense,  rect{5,  5,  2,  2})
    terrain.ZoneRect(base.ZoneCommercialLight,  rect{8,  8,  2,  2})

    terrain.ZoneRect(base.ZoneAirportDense,     rect{28, 5, 10, 10})
    terrain.CheckPower(base.ActionGameCreate)

//    terrain.ErrectBuilding("residential/skyscraper", point{35, 35})
    terrain.DemolishBuildingAt(point{35, 35})

    g1.SaveGame(GameFile1)
    return g1
}
