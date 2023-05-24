package tools

import (
    "github.com/metux/freecity/core/game"
)

func ChooseTool(cmd [] string, g * game.Game) Tool {
    switch cmd[0] {
        case "rubble":   return &Rubble{}
        case "pointer":  return &Pointer{}
        case "building": return &Building{buildingType: g.FindBuildingType(cmd[1])}
        case "road":     return &Road{}
        case "rail":     return &Rail{}
        case "pipe":     return &Pipe{}
        case "power":    return &Power{}
    }
    return &Pointer{}
}
