package tools

import (
    "fmt"
    "github.com/metux/freecity/core/rules"
)

type Building struct {
    BuildingType * rules.BuildingType
}

func (t * Building) GetName() string {
    return fmt.Sprintf("Building: %s (%dx%d)",
        t.BuildingType.Label,
        t.BuildingType.Size.X,
        t.BuildingType.Size.Y)
}

func (t * Building) WorkAt(game * Game, p point) {
    game.Terrain.ErrectBuilding(t.BuildingType.Ident, p)
}

func (t * Building) GetMenuId() string {
    return "building:"+t.BuildingType.Ident
}
