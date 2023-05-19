package items

import (
    "log"
)

type TileSet [] TileRef

func (tiles TileSet) AllFlat() bool {
    for _, t := range tiles {
        if ! t.Tile.IsFlat() {
            return false
        }
    }
    return true
}

func (tiles TileSet) AllLandFlat() bool {
    for _, t := range tiles {
        if ! t.Tile.IsLand() || ! t.Tile.IsFlat() {
            return false
        }
    }
    return true
}

func (tiles TileSet) AllLand() bool {
    for _, t := range tiles {
        if ! t.Tile.IsLand() {
            return false
        }
    }
    return true
}

func (tiles TileSet) AllWater() bool {
    for _, t := range tiles {
        if ! t.Tile.IsWater() {
            return false
        }
    }
    return true
}

func (tiles TileSet) AllWaterfall() bool {
    for _, t := range tiles {
        if ! t.Tile.IsWater() || t.Tile.IsFlat() {
            return false
        }
    }
    return true
}

func (tiles TileSet) FlatLandAndWater() bool {
    var height byte = 0
    hasLand := false
    hasWater := false

    for _, t := range tiles {
        if t.Tile.IsLand() {
            if hasLand {
                if height != t.Tile.Height {
                    log.Println("already have land with different height: ", height, " vs ", t.Tile.Height)
                    return false
                }
            } else {
                height = t.Tile.Height
                hasLand = true
            }
        } else {
            hasWater = true
        }
    }

    return hasLand && hasWater
}

func (tiles TileSet) Check(check string) bool {
    switch (check) {
        case "any_land":
            return tiles.AllLand()
        case "any_water":
            return tiles.AllWater()
        case "flat_land":
            return tiles.AllLandFlat()
        case "waterfall":
            return tiles.AllWaterfall()
    }
    return true
}
