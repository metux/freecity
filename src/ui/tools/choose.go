package tools

import (
    "github.com/metux/freecity/core/game"
)

func ChooseTool(cmd [] string, g * game.Game) Tool {
    switch cmd[0] {
        case "rubble":   return mkPlaceAtSimple("Rubble",    "rubble",          "rubble")
        case "road":     return mkPlaceAtSimple("Road",      "infra.road",      "road")
        case "rail":     return mkPlaceAtSimple("Rail",      "infra.rail",      "rail")
        case "pipe":     return mkPlaceAtSimple("Pipe",      "infra.pipe",      "pipe")
        case "power":    return mkPlaceAtSimple("Powerline", "infra.powerline", "powerline")
        case "building": return &Building{buildingType: g.FindBuildingType(cmd[1])}
        case "pointer":  return &Pointer{}
    }
    return &Pointer{}
}
