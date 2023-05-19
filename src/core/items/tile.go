package items

import (
    "fmt"
    "github.com/metux/freecity/core/base"
)

type Tile struct {
    Terrain     base.TerrainType    `yaml:"type"`
    ZoneTag     base.ZoneTag        `yaml:"zone,omitempty"`
    Wood        byte                `yaml:"wood,omitempty"`
    Height      byte                `yaml:"height,omitempty"`
    Water       byte                `yaml:"water,omitempty"`
    PowerGrid * PowerGrid           `yaml:"-"`
    Power       LineDirection       `yaml:"power,omitempty"`
    Road        LineDirection       `yaml:"road,omitempty"`
    Rail        LineDirection       `yaml:"rail,omitempty"`
    Pipe        LineDirection       `yaml:"pipe,omitempty"`
    Rubble      bool                `yaml:"rubble,omitempty"`
    Building  * Building            `yaml:"-"`
}

func (t * Tile) String() string {
    return fmt.Sprintf("Tile: %s wood=%d power=%s road=%d pipe=%d",
        t.Terrain, t.Wood, t.Power, t.Road, t.Pipe)
}

func (t * Tile) IsWater() bool {
    return t.Terrain.IsWater()
}

func (t * Tile) IsLand() bool {
    return ! t.Terrain.IsWater()
}

func (t * Tile) IsFlat() bool {
    return t.Terrain.IsFlat()
}

func (t * Tile) IsWaterfall() bool {
    return t.Terrain.IsWaterfall()
}

func (t * Tile) HasPowerLine() bool {
    return (t.Power != base.LineDirNone) || ((t.Building != nil) && t.Building.RoutesPower())
}

func (t * Tile) PowerConnected() bool {
    return t.PowerGrid != nil
}

func (t * Tile) ClearPowerGrid() {
    t.PowerGrid = nil
}

func (t * Tile) HasRoad() bool {
    return (t.Road != base.LineDirNone) || ((t.Building != nil) && t.Building.RoutesRoad())
}

func (t * Tile) HasRail() bool {
    return (t.Rail != base.LineDirNone) || ((t.Building != nil) && t.Building.RoutesRail())
}