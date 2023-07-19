package gtk

import (
    "log"
    "github.com/gotk3/gotk3/gtk"
    "github.com/gotk3/gotk3/gdk"
    "github.com/gotk3/gotk3/cairo"
    "github.com/metux/freecity/core/game"
    "github.com/metux/freecity/render/theme"
    "github.com/metux/freecity/util/cmd"
    render_cairo "github.com/metux/freecity/render/cairo"
)

type MapViewWindow struct {
    Config      * Config
    Renderer    * render_cairo.Renderer
    DrawingArea * gtk.DrawingArea
    Game        * game.Game
    DoWorkAt    func (p point)
}

func (mv * MapViewWindow) Init(g * game.Game, parent * gtk.Box, cf * Config, statusmsg func(s string)) gtk.IWidget {
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
    return mv.DrawingArea
}

// FIXME: differentiate buttons
func (mv * MapViewWindow) clickAt(x, y float64) {
    mv.DoWorkAt(mv.Renderer.PointerPos(fpoint{x, y}))
    mv.DrawingArea.QueueDraw()
}

func (mv * MapViewWindow) handleMove(c cmd.Cmdline, id string) bool {
    switch c.Str(0) {
        case "left":  mv.Renderer.Move(fpoint{ mv.Config.MoveStep, 0}); break
        case "right": mv.Renderer.Move(fpoint{-mv.Config.MoveStep, 0}); break
        case "up":    mv.Renderer.Move(fpoint{0,  mv.Config.MoveStep}); break
        case "down":  mv.Renderer.Move(fpoint{0, -mv.Config.MoveStep}); break
        default:
            log.Println("MapViewWindow: unhandled move command:", c, id)
            return false
    }
    mv.DrawingArea.QueueDraw()
    return true
}

func (mv * MapViewWindow) handleZoom(c cmd.Cmdline, id string) bool {
    switch c.Str(0) {
        case "up":   mv.Renderer.ZoomStep(mv.Config.ZoomStep);  break
        case "down": mv.Renderer.ZoomStep(-mv.Config.ZoomStep); break
        default:
            log.Println("MapViewWindow: unhandled zoom command:", c, id)
            return false
    }
    mv.DrawingArea.QueueDraw()
    return true
}

func (mv * MapViewWindow) HandleCmd(c cmd.Cmdline, id string) bool {
    switch c.Str(0) {
        case "move": return mv.handleMove(c.Skip(1), id)
        case "zoom": return mv.handleZoom(c.Skip(1), id)
        case "repaint": {
            mv.Game.Terrain.TouchTerrain()
            mv.DrawingArea.QueueDraw()
            return true
        }
        default:
            log.Println("MapViewWindow: unhandled command:", c, id)
            return false
    }
    return false
}
