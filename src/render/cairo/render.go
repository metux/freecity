package cairo

import (
    "log"
    "time"
    "fmt"
    "github.com/metux/freecity/core/base"
    "github.com/metux/freecity/core/items"
    "github.com/metux/freecity/core/game"
    "github.com/metux/freecity/render/theme"
    "github.com/metux/freecity/render"
    "github.com/gotk3/gotk3/cairo"
    "sync"
    "image"
)

type Renderer struct {
    Terrain       * items.TerrainMap
    Theme         * theme.ThemeSpec
    Context       * cairo.Context
    Prescale        float64
    imgDim          image.Point
    imgOrigin       image.Point
    imgStepX        image.Point
    imgStepY        image.Point
    tileWidth       int
    tileHeight      int
    imageCache      map[string] * cairo.Surface
    imageCacheLock  sync.RWMutex
    Scale           float64
    loaded          bool
    basemap       * cairo.Surface
    fullmap       * cairo.Surface
    OffsetX         float64
    OffsetY         float64
}

const defaultScale = 0.9
const defaultPrescale = 0.75

//const basemapFormat = cairo.FORMAT_RGB24 // cairo.FORMAT_ARGB32
const basemapFormat = cairo.FORMAT_ARGB32
const fullmapFormat = basemapFormat
const spriteFormat = cairo.FORMAT_ARGB32

func (r * Renderer) prescaleLen(i int) int {
    return int(float64(i) * r.Prescale)
}

func (r * Renderer) init() {
    if r.loaded {
        return
    }

    r.Theme.LoadImages()
    r.imageCache = make(map[string] * cairo.Surface)

    r.tileWidth = r.prescaleLen(r.Theme.TileSize.Width)
    r.tileHeight = r.prescaleLen(r.Theme.TileSize.Height)

    ts := image.Point{r.tileWidth, r.tileHeight}

    hts := image.Point{ts.X/2,ts.Y/2}

    switch r.Theme.Projection {
        case theme.ProjFlat:
            r.imgDim    = image.Point{r.Terrain.Size.Width  * ts.X,
                                      r.Terrain.Size.Height * ts.Y}
            r.imgOrigin = image.Point{0, 0}
            r.imgStepX  = image.Point{ts.X, 0}
            r.imgStepY  = image.Point{0, ts.Y}
        break
        case theme.ProjParallel:
            r.imgDim    = image.Point{r.Terrain.Size.Width  * ts.X,
                                      r.Terrain.Size.Height * ts.Y}
            r.imgOrigin = image.Point{r.imgDim.X / 2 - hts.X, 0}
            r.imgStepX  = image.Point{hts.X, hts.Y}
            r.imgStepY  = image.Point{-hts.X, hts.Y}
        break
        default:
            panic("unsupported projection: "+r.Theme.Projection.String())
    }

    r.loaded = true
}

func (r* Renderer) getImage(name string) * cairo.Surface {
    r.imageCacheLock.Lock()
    defer r.imageCacheLock.Unlock()

    if img, okay := r.imageCache[name]; okay {
        return img
    }

    img, err := cairo.NewSurfaceFromPNG(r.Theme.ImgPath(name))
    if err != nil {
        log.Println("failed loading surface from png", name, err)
        return nil
    }

    scaledImg := cairo.CreateImageSurface(spriteFormat,
        r.prescaleLen(img.GetWidth()),
        r.prescaleLen(img.GetHeight()))

    scaledCtx := cairo.Create(scaledImg)
    defer scaledCtx.Close()

    scaledCtx.Scale(r.Prescale, r.Prescale)

    scaledCtx.SetSourceSurface(img, 0, 0)
    scaledCtx.Paint()

    r.imageCache[name] = scaledImg
    return scaledImg
}

func (r* Renderer) overlayTileImage(cr * cairo.Context, x float64, y float64, name string) {
    if name != "" {
        if i2 := r.getImage(name); i2 != nil {
            cr.SetSourceSurface(i2, x, y)
            cr.Paint()
        }
    }
}

func (r* Renderer) overlayTile(cr * cairo.Context, pos base.Point, name string) {
    px, py := r.tilePos(pos)
    r.overlayTileImage(cr, px, py, name)
}

func (r* Renderer) tilePos(pos base.Point) (float64, float64) {
    coords := r.imgOrigin.Add(r.imgStepX.Mul(pos.X)).Add(r.imgStepY.Mul(pos.Y))
    return float64(coords.X), float64(coords.Y)
}

func (r* Renderer) renderTile(cr * cairo.Context, tr items.TileRef) {
    x, y := r.tilePos(tr.Position)

    r.overlayTileImage(cr, x, y, r.Theme.ImgForRubble(tr.Tile.Rubble))
    r.overlayTileImage(cr, x, y, r.Theme.ImgForZone(tr.Tile.ZoneTag))
    r.overlayTileImage(cr, x, y, r.Theme.ImgForRoad(tr.Tile.Road))
    r.overlayTileImage(cr, x, y, r.Theme.ImgForRail(tr.Tile.Rail))
    r.overlayTileImage(cr, x, y, r.Theme.ImgForPower(tr.Tile.Power))
}

func (r* Renderer) renderTileBase(cr * cairo.Context, tr items.TileRef) {
    x, y := r.tilePos(tr.Position)

    r.overlayTileImage(cr, x, y, theme.ImgLandFlat)
    r.overlayTileImage(cr, x, y, theme.ImgGrid)
}

func (r * Renderer) renderTileBorder(cr * cairo.Context, rect base.Rect) {
    r.overlayTile(cr, rect.TopLeft(),     theme.ImgBorderTopLeft)
    r.overlayTile(cr, rect.TopRight(),    theme.ImgBorderTopRight)
    r.overlayTile(cr, rect.BottomLeft(),  theme.ImgBorderBottomLeft)
    r.overlayTile(cr, rect.BottomRight(), theme.ImgBorderBottomRight)

    if (rect.Width > 2) {
        for x:=rect.X+1; x<rect.X + rect.Width-1; x++ {
            r.overlayTile(cr, base.Point{x, rect.Y}, theme.ImgBorderTop)
            r.overlayTile(cr, base.Point{x, rect.Y + rect.Height - 1}, theme.ImgBorderBottom)
        }
    }

    if (rect.Height > 2) {
        for y:=rect.Y+1; y<rect.Y + rect.Height-1; y++ {
            r.overlayTile(cr, base.Point{rect.X, y}, theme.ImgBorderLeft)
            r.overlayTile(cr, base.Point{rect.X + rect.Width - 1, y}, theme.ImgBorderRight)
        }
    }
}

func (r * Renderer) renderBuilding(cr * cairo.Context, b * items.Building) {
    rect := b.OccupiedRect()

    r.renderTileBorder(cr, rect)

    if img := r.getImage("buildings/"+b.TypeName); img != nil {
        imgX,_ := r.tilePos(rect.BottomLeft())
        _,imgY := r.tilePos(rect.BottomRight())
        imgY   += float64(r.tileHeight - img.GetHeight())

        cr.SetSourceSurface(img,imgX,imgY)
        cr.Paint()
    }
}

func (r * Renderer) createBasemap() {
    if r.basemap != nil {
        return
    }

    if render.DebugMode {
        defer base.TimeTrack(time.Now(), "cairo-render createBasemap()")
    }

    r.basemap = cairo.CreateImageSurface(basemapFormat, r.imgDim.X, r.imgDim.Y)

    ctx := cairo.Create(r.basemap)
    defer ctx.Close()

    tl := r.Terrain.AllTiles()
    for _,tr := range tl {
        r.renderTileBase(ctx, tr)
    }
}

func (r * Renderer) renderFullmap(ctx * cairo.Context) {
    // take the basemap as background
    ctx.SetSourceSurface(r.basemap, 0, 0)
    ctx.Paint()

    tl := r.Terrain.AllTiles()
    for _,tr := range tl {
        r.renderTile(ctx, tr)
    }
    for _,b := range r.Terrain.Buildings {
        r.renderBuilding(ctx, b)
    }
}

func (r * Renderer) createFullmap() {
    if r.fullmap != nil {
        return
    }

    if render.DebugMode {
        defer base.TimeTrack(time.Now(), "cairo-render createFullmap()")
    }

    r.createBasemap()
    r.fullmap = cairo.CreateImageSurface(basemapFormat, r.imgDim.X, r.imgDim.Y)

    ctx := cairo.Create(r.fullmap)
    defer ctx.Close()

    r.renderFullmap(ctx)
}

func (r * Renderer) Render(cr * cairo.Context) {
    r.init()
    r.createBasemap()
    r.createFullmap()
    cr.Scale(r.Scale, r.Scale)

    if render.DebugMode {
        defer base.TimeTrack(time.Now(), fmt.Sprintf("cairo-render (scale %f)", r.Scale))
    }

    // take the basemap as background
    cr.SetSourceSurface(r.fullmap, r.OffsetX, r.OffsetY)
    cr.Paint()
}

func (r * Renderer) SetScale(sc float64) {
    sc = float64(int(sc*10000)) / 10000
    r.Scale = sc
}

func (r * Renderer) SetOffset(x, y int) {
    r.OffsetX = float64(x)
    r.OffsetY = float64(y)
}

func (r * Renderer) ZoomStep(x float64) {
    r.SetScale(r.Scale + x)
}

func (r * Renderer) MoveX(x float64) {
    r.OffsetX += x
}

func (r * Renderer) MoveY(y float64) {
    r.OffsetY += y
}

func CreateRenderer(g1 * game.Game, t * theme.ThemeSpec) (* Renderer) {
    return &(Renderer{
        Terrain:  &g1.Terrain,
        Theme:    t,
        Scale:    defaultScale,
        Prescale: defaultPrescale,
    })
    return nil
}
