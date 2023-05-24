package items

import (
    "github.com/metux/freecity/core/base"
    "github.com/metux/freecity/core/rules"
)

type Building struct {
    TypeName            string               `yaml:"type"`
    BuildingType      * rules.BuildingType   `yaml:"-"`
    PowerGrid         * PowerGrid            `yaml:"-"`
    Position            point                `yaml:"position"`
    Consumption         base.Consumption     `yaml:"consumption"`
    Powered             bool                 `yaml:"-"`
    Terrain           * TerrainMap           `yaml:"-"`
    Destroyed           bool                 `yaml:"demolished"`
    Abandoned           bool                 `yaml:"abandoned"`
}

func (b * Building) PowerConsumption() int16 {
    return b.Consumption.Power
}

func (b * Building) IsPowerGenerator() bool {
    return b.Consumption.Power < 0
}

func (b * Building) IsPowerConsumer() bool {
    return b.Consumption.Power > 0
}

func (b * Building) OccupiedRect() rect {
    return b.Position.MakeRect(b.BuildingType.Size)
}

func (b * Building) PowerConnected() bool {
    return b.PowerGrid != nil
}

func (b * Building) ClearPowerGrid() {
    b.PowerGrid = nil
}

func (b * Building) SetPowered(enabled bool) {
    b.Powered = enabled
}

func (b * Building) RoutesPower() bool {
    return b.BuildingType.Routes.Power
}

func (b * Building) RoutesRail() bool {
    return b.BuildingType.Routes.Rail
}

func (b * Building) RoutesRoad() bool {
    return b.BuildingType.Routes.Road
}

func (b * Building) RoutesLine(lt LineType) bool {
    if (b != nil) {
        switch lt {
            case LtPower: return b.BuildingType.Routes.Power
            case LtRail:  return b.BuildingType.Routes.Rail
            case LtRoad:  return b.BuildingType.Routes.Road
        }
    }
    return false
}

func (b * Building) TileRange() TileSet {
    ts,_ := b.Terrain.TileRange(b.OccupiedRect(), true)
    return ts
}

func (b * Building) Destroy() {
    for _,t := range b.TileRange() {
        t.Tile.Building = nil
        t.Tile.Rubble = true
    }
    b.Destroyed = true
}

func (b * Building) ConnectTiles() {
    for _,t := range b.TileRange() {
        t.Tile.Building = b
    }
}

func NewBuilding(t * rules.BuildingType, pos point, tm * TerrainMap) (*Building) {
    b := Building{
        BuildingType:     t,
        Position:         pos,
        TypeName:         t.Ident,
        Consumption:      t.Consumption,
        Terrain:          tm,
    }
    return &b
}
