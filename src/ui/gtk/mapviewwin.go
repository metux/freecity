package gtk

import (
    "log"
    "github.com/gotk3/gotk3/gtk"
    "github.com/gotk3/gotk3/gdk"
    "github.com/gotk3/gotk3/cairo"
    "github.com/metux/freecity/core/game"
    "github.com/metux/freecity/render/theme"
    render_cairo "github.com/metux/freecity/render/cairo"
)

type MapViewWindow struct {
    Config      * Config
    Renderer    * render_cairo.Renderer
    DrawingArea * gtk.DrawingArea
    Game        * game.Game
}

func (mv * MapViewWindow) Init(g * game.Game, parent * gtk.Box, cf * Config, statusmsg func(s string)) {
    mv.Config = cf
    mv.Game = g
    mv.DrawingArea, _ = gtk.DrawingAreaNew()

    mv.Renderer = render_cairo.CreateRenderer(g, theme.CreateTheme(cf.DataPrefix + "/themes/" + cf.Theme), statusmsg)
    mv.Renderer.Viewport.Prescale = cf.Prescale
    mv.Renderer.SetScale(cf.Scale)

    // Event handlers
    mv.DrawingArea.AddEvents(int(gdk.POINTER_MOTION_MASK | gdk.BUTTON_PRESS_MASK))
    mv.DrawingArea.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
        mv.Renderer.Render(cr)
    })
    mv.DrawingArea.Connect("motion-notify-event", func(da *gtk.DrawingArea, ev *gdk.Event) {
        x1, y1, x2, y2 := mv.Renderer.UpdateCursor(evMotionFPoint(ev))
        mv.DrawingArea.QueueDrawArea(x1, y1, x2, y2)
    })
    mv.DrawingArea.Connect("button-press-event", func(da *gtk.DrawingArea, ev *gdk.Event) {
        p1 := evMotionFPoint(ev)
        mv.clickAt(p1.X, p1.Y)
    })
    parent.PackStart(mv.DrawingArea, true, true, 0)
}

// FIXME
func (mv * MapViewWindow) clickAt(x, y float64) {
    p := mv.Renderer.PointerPos(fpoint{x, y})
    log.Println("placing rubble at", p)
    mv.Game.Terrain.PlaceRubble(p)
    mv.DrawingArea.QueueDraw()
}

func (mv * MapViewWindow) handleMove(cmd [] string, id string) bool {
    switch cmd[1] {
        case "left":  mv.Renderer.Move(fpoint{ mv.Config.MoveStep, 0}); break
        case "right": mv.Renderer.Move(fpoint{-mv.Config.MoveStep, 0}); break
        case "up":    mv.Renderer.Move(fpoint{0,  mv.Config.MoveStep}); break
        case "down":  mv.Renderer.Move(fpoint{0, -mv.Config.MoveStep}); break
        default:
            log.Println("MapViewWindow: unhandled move command:", cmd, id)
            return false
    }
    mv.DrawingArea.QueueDraw()
    return true
}

func (mv * MapViewWindow) handleZoom(cmd [] string, id string) bool {
    switch cmd[1] {
        case "up":   mv.Renderer.ZoomStep(mv.Config.ZoomStep);  break
        case "down": mv.Renderer.ZoomStep(-mv.Config.ZoomStep); break
        default:
            log.Println("MapViewWindow: unhandled zoom command:", cmd, id)
            return false
    }
    mv.DrawingArea.QueueDraw()
    return true
}

func (mv * MapViewWindow) HandleCmd(cmd [] string, id string) bool {
    switch cmd[0] {
        case "move": return mv.handleMove(cmd, id)
        case "zoom": return mv.handleZoom(cmd, id)
        default:
            log.Println("MapViewWindow: unhandled command:", cmd, id)
            return false
    }
    return false
}
