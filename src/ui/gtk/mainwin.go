package gtk

import (
    "github.com/gotk3/gotk3/gtk"
    "github.com/gotk3/gotk3/gdk"
    "github.com/metux/freecity/core/base"
    "github.com/metux/freecity/util"
    "github.com/metux/freecity/core/game"
    "github.com/metux/freecity/core/items"
    "github.com/metux/freecity/ui/common"
    "github.com/metux/freecity/ui/tools"
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
    Tool        tools.Tool
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

func (mw * MainWindow) chooseTool(cmd [] string) tools.Tool {
    switch cmd[0] {
        case "rubble":   return &tools.Rubble{}
        case "pointer":  return &tools.Pointer{}
        case "building": return &tools.Building{BuildingType: mw.Game.FindBuildingType(cmd[1])}
    }
    return nil
}

func (mw * MainWindow) handleCmdTool(cmd [] string, id string) bool {
    log.Println("Tool: ", cmd, "("+id+")")
    t := mw.chooseTool(cmd)

    if t != nil {
        mw.SetTool(t)
        return true
    }
    return false
}

func (mw * MainWindow) HandleCmd(cmd [] string, id string) bool {
    switch cmd[0] {
        case "": return false
        case "mapview": return mw.MapView.HandleCmd(cmd[1:], id)
        case "quit": mw.App.Quit(); break
        case "game": return mw.Game.HandleCmd(cmd[1:], id)
        case "tool": return mw.handleCmdTool(cmd[1:], id)
        default:
            log.Println("MainWindow: unhandled command: ", cmd, id)
            return false
    }
    return true
}

func (mw * MainWindow) InitBuildMenu() {
    m := mw.Config.MainMenu.FindById("build")

    log.Println("builder menu", m)

    for idx,bt := range mw.Game.Terrain.GeneralRules.Buildings.BuildingTypes {
//    for idx,bt := range mw.Game.Terrain.GeneralRules.Buildings.Placable {
        log.Println("building --> ", idx, bt)

//        if bt.Placable {
        m.Submenu = append(m.Submenu, common.MenuEntry{
            Label: bt.Label,
            Id:    bt.Ident,
            Cmd:   "place-building "+bt.Ident,
            CmdHandler: m.CmdHandler,
        })
//        }
    }

    m.CreateEntries()
}

func (mw * MainWindow) SetTool(t tools.Tool) {
    mw.Tool = t
    mw.StatusBar.Push(3, "Tool: "+t.GetName())
}

func (mw * MainWindow) Init(app * gtk.Application, g * game.Game, datadir string) {
    mw.Config = LoadUIYaml(datadir)
    mw.App = app
    mw.Game = g
    mw.Console = g.SetNotify(mw)

    // set initial tool
    mw.SetTool(&tools.Rubble{})

    // create main window
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

    // create the buildings menu
    mw.InitBuildMenu()

    // create the map viewer
    mw.MapView.DoWorkAt = func(p point) {
        log.Println("Do work at", p, mw.Tool.GetName())
        mw.Tool.WorkAt(mw.Game, p)
    }
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
            mw.HandleCmd(util.SplitCmdline(id), "key")
        } else {
            log.Println("key not bound", key)
        }
    })

    mw.window.ShowAll()
}
