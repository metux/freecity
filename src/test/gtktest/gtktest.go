package gtktest

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/metux/freecity/core/game"
	ui_gtk "github.com/metux/freecity/ui/gtk"
)

const dataPrefix = "../data/"

func Run(g * game.Game) {
    // Initialize GTK without parsing any command line arguments.
    gtk.Init(nil)

    mw := ui_gtk.MainWindow{}
    mw.Init(g, dataPrefix)

    // Begin executing the GTK main loop.  This blocks until
    // gtk.MainQuit() is run. 
    gtk.Main()
}
