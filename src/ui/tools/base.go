package tools

import (
    "github.com/metux/freecity/util"
    "github.com/metux/freecity/core/game"
)

type point = util.Point
type Game = game.Game

type Tool interface {
    GetName() string
    WorkAt(g * Game, p point)
}
