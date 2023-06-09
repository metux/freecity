package main

import (
    "log"
    "github.com/metux/freecity/test/games"
    "github.com/metux/freecity/test/gtktest"
)

func main() {
    log.SetFlags(log.Lmicroseconds)

    games.TestGame1()

    g2 := games.LoadGame1()
    gtktest.Run(g2)
}
