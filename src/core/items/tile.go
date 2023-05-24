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

func (t * Tile) PowerConnected() bool {
    return t.PowerGrid != nil
}

func (t * Tile) ClearPowerGrid() {
    t.PowerGrid = nil
}

func (t Tile) hasLineSelf(lt base.LineType) bool {
    switch (lt) {
        case base.LineTypePower: return t.Power.Present()
        case base.LineTypeRail:  return t.Rail.Present()
        case base.LineTypeRoad:  return t.Road.Present()
        case base.LineTypePipe:  return t.Pipe.Present()
    }
    return false
}

func (t Tile) HasLine(lt base.LineType) bool {
    return t.hasLineSelf(lt) || t.Building.RoutesLine(lt)
}

func (t * Tile) SetLine(lt base.LineType, line base.LineDirection) {
    switch (lt) {
        case base.LineTypePower: t.Power = line
        case base.LineTypeRail:  t.Rail  = line
        case base.LineTypeRoad:  t.Road  = line
        case base.LineTypePipe:  t.Pipe  = line
    }
}

func (t * Tile) PickLine(lt base.LineType, other1, other2 base.LineDirection) base.LineDirection {
    switch lt {
        case base.LineTypePower: return base.LineDirPick(t.Road,  t.Rail)
        case base.LineTypeRoad:  return base.LineDirPick(t.Rail,  t.Road)
        case base.LineTypeRail:  return base.LineDirPick(t.Power, t.Road)
        // currently no conflicts with others
        case base.LineTypePipe:  return base.LineDirPick(base.LineDirNone, base.LineDirNone)
    }
    return base.LineDirNone
}
