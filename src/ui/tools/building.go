package tools

import (
    "log"
    "fmt"
    "github.com/metux/freecity/core/rules"
)

type Building struct {
    BuildingType * rules.BuildingType
}

func (t * Building) GetName() string {
    return fmt.Sprintf("Building: %s (%dx%d)",
        t.BuildingType.Ident,
        t.BuildingType.Size.X,
        t.BuildingType.Size.Y)
}

func (t * Building) WorkAt(game * Game, p point) {
    log.Println(" --> placing rubble at", p)
    game.Terrain.ErrectBuilding(t.BuildingType.Ident, p)
}
