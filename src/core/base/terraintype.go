package base

// represents the possible tilt directions of a terrain tile
type TerrainType uint8
type Tilt uint8

// flat land
// hill land
// flat underwater
// hill underwater
// waterfall

const (
    flat          TerrainType = 0

    Tilt_east              = 1
    Tilt_west              = 2
    Tilt_north             = 3
    Tilt_northeast         = 4
    Tilt_nortwWest         = 5
    Tilt_south             = 6
    Tilt_southwest         = 7
    Tilt_southeast         = 8

    Water                  = 0x10
    Salted                 = 0x20
    Garbage                = 0x40
    Waterfall              = 0x50
)

const (
    LandFlat TerrainType = flat
    SaltWaterNorth TerrainType = (Water | Tilt_north)
)

func (t TerrainType) hasFlag(flag TerrainType) bool {
    return (t & flag != 0)
}

func (t TerrainType) IsWater() bool {
    return t.hasFlag(Water) || t.hasFlag(Waterfall)
}

func (t TerrainType) IsFlat() bool {
    return t & 0x0F == 0
}

func (t TerrainType) IsWaterfall() bool {
    return t.hasFlag(Waterfall)
}

func (t TerrainType) String() string {
    var w string

    if (t.hasFlag(Water)) {
        w = "Water"
        if (t.hasFlag(Salted)) {
            w = "Salt-Water"
        } else {
            w = "Fresh-Water"
        }
    } else {
        w = "Land"
    }

    if (t.hasFlag(Garbage)) {
        w = w + " (garbage)"
    }

    return w
}
