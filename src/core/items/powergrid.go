package items

import (
    "fmt"
)

type PowerGrid struct {
    Terrain        *  TerrainMap
    Production        int16
    Consumption       int16
    Producers   [] *  Building
    Consumers   [] *  Building
}

func (pg PowerGrid) String() string {
    return fmt.Sprintf("[consumers %d producers %d consumption %d production %d]",
        len(pg.Consumers), len(pg.Producers), pg.Consumption, pg.Production)
}

// add a building to powergrid (but dont trace into tiles)
// returns false if already was connected , true otherwise
func (pg * PowerGrid) addBuilding(b * Building) bool {
    if b == nil || b.PowerGrid != nil {
        return false
    }

    b.PowerGrid = pg

    if (b.IsPowerGenerator()) {
        pg.Producers = append(pg.Producers, b)
        pg.Production -= b.PowerConsumption()
        return true
    }

    if (b.IsPowerConsumer()) {
        pg.Consumers = append(pg.Consumers, b)
        pg.Consumption = b.PowerConsumption()
        return true
    }

    return true
}

func (this * PowerGrid) addTile(t * Tile) bool {
    if t == nil || t.PowerConnected() {
        return false
    }

    t.PowerGrid = this
    return true
}

func (this * PowerGrid) markTile(tr TileRef) {
    if (this.addTile(tr.Tile)) {
        this.markBuilding(tr.Tile.Building)
        if tr.Tile.Power.Present() {
            // FIXME: powering all neighbours, but we should respect direction
            // maybe make this a ruleset decision ?
            for _,tr2 := range tr.Surrounding() {
                this.markTile(tr2)
            }
        }
    }
}

func (this * PowerGrid) markBuilding(b * Building) {
    if (this.addBuilding(b)) {
        if b.BuildingType.Routes.Power {
            tiles, _ := this.Terrain.TileRange(b.OccupiedRect(), true)
            for _,tr1 := range tiles {
                if (this.addTile(tr1.Tile)) {
                    // mark surrounding as powered
                    for _,tr2 := range tr1.Surrounding() {
                        // mark the tile
                        this.markTile(tr2)
                    }
                }
            }
        }
    }
}

// FIXME: add priorities, which ones to cut off first
func (pg * PowerGrid) powerBuildings() {
    prod := pg.Production
    for _,b := range pg.Producers {
        b.SetPowered(true)
    }
    for _,b := range pg.Consumers {
        c := b.PowerConsumption()
        if (c > 0) {
            if (prod > c) {
                b.SetPowered(true)
                prod = prod - c
            } else {
                b.SetPowered(false)
            }
        }
    }
}

func CreatePowerGrid(terrain * TerrainMap, b * Building) PowerGrid {
    grid := PowerGrid { Terrain : terrain }
    grid.markBuilding(b)
    grid.powerBuildings()
    return grid
}
