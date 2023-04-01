package items

import (
    "github.com/metux/freecity/core/base"
    "github.com/metux/freecity/core/rules"
)

type Building struct {
    TypeName            string               `yaml:"type"`
    BuildingType      * rules.BuildingType   `yaml:"-"`
    PowerGrid         * PowerGrid            `yaml:"-"`
    Position            Point                `yaml:"position"`
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

func (b * Building) OccupiedRect() Rect {
    return base.RectByPointDim(b.Position, b.BuildingType.Size)
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

func (b * Building) Destroy() {
    ts,_ := b.Terrain.TileRange(base.RectByPointDim(b.Position, b.BuildingType.Size), true)
    for _,t := range ts {
        t.Tile.Building = nil
        t.Tile.Rubble = true
    }
    b.Destroyed = true
}

func (b * Building) ConnectTiles() {
    ts,_ := b.Terrain.TileRange(base.RectByPointDim(b.Position, b.BuildingType.Size), true)
    for _,t := range ts {
        t.Tile.Building = b
    }
}

func NewBuilding(t * rules.BuildingType, pos Point, tm * TerrainMap) (*Building) {
    b := Building{
        BuildingType:     t,
        Position:         pos,
        TypeName:         t.Name,
        Consumption:      t.Consumption,
        Terrain:          tm,
    }
    return &b
}
