package tools

import (
    "github.com/metux/freecity/core/game"
)

func ChooseTool(cmd [] string, g * game.Game) Tool {
    switch cmd[0] {
        case "rubble":   return &Rubble{}
        case "pointer":  return &Pointer{}
        case "building": return &Building{BuildingType: g.FindBuildingType(cmd[1])}
    }
    return &Pointer{}
}
