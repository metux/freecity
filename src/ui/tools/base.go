package tools

import (
    "github.com/metux/freecity/util/geo"
    "github.com/metux/freecity/core/game"
)

type point = geo.Point
type Game = game.Game

type Tool interface {
    GetName() string
    WorkAt(g * Game, p point)
}
