package tools

import (
    "fmt"
    "github.com/metux/freecity/core/rules"
)

type Building struct {
    buildingType * rules.BuildingType
}

func (t * Building) GetName() string {
    return fmt.Sprintf("Building: %s (%dx%d)",
        t.buildingType.Label,
        t.buildingType.Size.X,
        t.buildingType.Size.Y)
}

func (t * Building) WorkAt(game * Game, p point) {
    game.Terrain.ErrectBuilding(t.buildingType.Ident, p)
}

func (t * Building) GetMenuId() string {
    return "building:"+t.buildingType.Ident
}
