package gtk

import (
    "github.com/gotk3/gotk3/gtk"
    "github.com/gotk3/gotk3/gdk"
    "github.com/metux/freecity/core/game"
    "log"
)

type MainWindow struct {
    Win     * gtk.Window
    MapView   MapViewWindow
    Game    * game.Game
}

func (mw * MainWindow) Init(g * game.Game, datadir string) {
    err := error(nil)

    mw.Game = g

    // Initialize GTK without parsing any command line arguments.
    gtk.Init(nil)

    // Create a new toplevel window, set its title, and connect it to the
    // "destroy" signal to exit the GTK main loop when it is destroyed.
    mw.Win, err = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
    if err != nil {
        log.Fatal("Unable to create window:", err)
    }
    mw.Win.SetTitle("FreeCity")
    mw.Win.Connect("destroy", func() {
        gtk.MainQuit()
    })

    // Set the default window size.
    mw.Win.SetDefaultSize(800, 600)

    mw.MapView = MapViewWindow{}
    mw.MapView.Init(mw.Game, mw.Win, LoadUIYaml(datadir))

    // FIXME: handle mouse click to center

    // FIXME: gtk keymaps
    mw.Win.Connect("key-press-event", func(win *gtk.Window, ev *gdk.Event) {
        keyEvent := &gdk.EventKey{ev}
        switch (keyEvent.KeyVal()) {
            case gdk.KEY_Left:
                mw.MapView.MoveLeft()
            break;
            case gdk.KEY_Right:
                mw.MapView.MoveRight()
            break;
            case gdk.KEY_Up:
                mw.MapView.MoveUp()
            break
            case gdk.KEY_Down:
                mw.MapView.MoveDown()
            break;
            case gdk.KEY_plus:
                mw.MapView.ZoomUp()
            break;
            case gdk.KEY_minus:
                mw.MapView.ZoomDown()
            break;
            case gdk.KEY_q:
                if (keyEvent.State() & gdk.CONTROL_MASK) == gdk.CONTROL_MASK {
                    gtk.MainQuit()
                }
            break;
            default:
                log.Println("key event", keyEvent.KeyVal())
                return
            break;
        }
        mw.Win.QueueDraw()
    })

    // Recursively show all widgets contained in this window.
    mw.Win.ShowAll()
}
