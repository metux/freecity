package gtktest

import (
    "os"
    "log"
    "flag"
    "github.com/gotk3/gotk3/gtk"
    "github.com/gotk3/gotk3/glib"
    "github.com/metux/freecity/core/game"
    ui_gtk "github.com/metux/freecity/ui/gtk"
)

const dataPrefix = "../data/"

func Run(g * game.Game) {
    app,err := gtk.ApplicationNew("de.metux.freecity", glib.APPLICATION_FLAGS_NONE)
    if err != nil {
        log.Fatal("failed creating gtk application %v", err)
        panic("cant proceed")
    }
    app.Connect("activate", func() {
        mw := ui_gtk.MainWindow{}
        mw.Init(app, g, dataPrefix)
    })
    os.Exit(app.Run(flag.Args()))
}
