package games

import (
    "github.com/metux/freecity/core/game"
)

func LoadGame1() * game.Game {
    g2 := game.DefaultConfig.LoadGame(GameFile1)
    g2.SaveGame(GameFile2)
    return g2
}
