package main

import (
    "log"
    "github.com/metux/freecity/test/games"
//    "github.com/metux/freecity/test/render"
    "github.com/metux/freecity/test/gtktest"
)

func main() {
    log.SetFlags(log.Lmicroseconds)

    games.TestGame1()

    g2 := games.LoadGame1()
//    renderer := render.CreateRenderSimple(g2, "parallel")
//    renderer.Nop()

    gtktest.Run(g2)
}
