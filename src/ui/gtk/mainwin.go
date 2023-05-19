package gtk

import (
    "github.com/gotk3/gotk3/gtk"
    "github.com/gotk3/gotk3/gdk"
    "github.com/metux/freecity/core/base"
    "github.com/metux/freecity/core/cmd"
    "github.com/metux/freecity/core/game"
    "github.com/metux/freecity/core/items"
    "log"
)

type MainWindow struct {
    App       * gtk.Application
    window    * gtk.ApplicationWindow
    MapView     MapViewWindow
    Game      * game.Game
    Console     items.NotifyHandler
    Box       * gtk.Box
    Config    * Config
    StatusBar * gtk.Statusbar
}

func (mw * MainWindow) NotifyEmit(a base.Action, n items.NotifyMsg) bool {
    mw.Console.NotifyEmit(a, n)
    switch n2 := n.(type) {
        case game.NotifyGameSpeed: {
            mw.Config.MainMenu.SetGameSpeed(n2.Speed)
            return true
        }
    }
    mw.StatusBar.Push(2, n.String())
    return false
}

func (mw * MainWindow) HandleCmd(cmd [] string, id string) bool {
    switch cmd[0] {
        case "": return false
        case "mapview": return mw.MapView.HandleCmd(cmd[1:], id)
        case "quit": mw.App.Quit(); break
        case "game": return mw.Game.HandleCmd(cmd[1:], id)
        default:
            log.Println("MainWindow: unhandled command: ", cmd, id)
            return false
    }
    return true
}

func (mw * MainWindow) Init(app * gtk.Application, g * game.Game, datadir string) {
    mw.Config = LoadUIYaml(datadir)
    mw.App = app

    mw.Game = g
    mw.Console = g.SetNotify(mw)

    mw.window,_= gtk.ApplicationWindowNew(app)
    mw.window.SetTitle(mw.Config.WindowTitle)
    mw.window.Connect("destroy", func() { mw.HandleCmd([]string{"quit"}, "") })
    mw.window.SetDefaultSize(mw.Config.WindowSize.X, mw.Config.WindowSize.Y)

    mw.Box,_ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)
    mw.window.Add(mw.Box)

    mw.StatusBar,_ = gtk.StatusbarNew()

    // init menu
    mw.Config.MainMenu.SetHandler(mw)
    mw.Box.PackStart(GtkLoadMenuBar(&mw.Config.MainMenu), false, false, 0)

    mw.MapView = MapViewWindow{}
    mw.MapView.Init(mw.Game, mw.Box, mw.Config, func(s string) {
        mw.StatusBar.Push(3, s)})

    // statusbar
    mw.Box.PackEnd(mw.StatusBar, false, false, 0)
    mw.StatusBar.Push(1, "game startup")

    // FIXME: handle mouse click to center

    mw.window.Connect("key-press-event", func(win *gtk.ApplicationWindow, ev *gdk.Event) {
        key := translateGdkKey(&gdk.EventKey{ev})
        id,okay := mw.Config.KeyMap[key]
        if okay {
            mw.HandleCmd(cmd.SplitCmdline(id), "key")
        } else {
            log.Println("key not bound", key)
        }
    })

    mw.window.ShowAll()
}
