package gtk

import (
    "github.com/gotk3/gotk3/gtk"
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

func (mv * MapViewWindow) Init(g * game.Game, win * gtk.Window, cf * Config) {
    mv.Config = cf
    mv.DrawingArea, _ = gtk.DrawingAreaNew()

    mv.Renderer = render_cairo.CreateRenderer(g, theme.CreateTheme(cf.DataPrefix + "/themes/" + cf.Theme))
    mv.Renderer.Prescale = cf.Prescale
    mv.Renderer.SetScale(cf.Scale)

    // Event handlers
    mv.DrawingArea.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
        mv.Renderer.Render(cr)
    })
    win.Add(mv.DrawingArea)
}

func (mv * MapViewWindow) ZoomUp() {
    mv.Renderer.ZoomStep(mv.Config.ZoomStep)
}

func (mv * MapViewWindow) ZoomDown() {
    mv.Renderer.ZoomStep(-mv.Config.ZoomStep)
}

func (mv * MapViewWindow) MoveUp() {
    mv.Renderer.MoveY(mv.Config.MoveStep)
}

func (mv * MapViewWindow) MoveDown() {
    mv.Renderer.MoveY(-mv.Config.MoveStep)
}

func (mv * MapViewWindow) MoveLeft() {
    mv.Renderer.MoveX(mv.Config.MoveStep)
}

func (mv * MapViewWindow) MoveRight() {
    mv.Renderer.MoveX(-mv.Config.MoveStep)
}
