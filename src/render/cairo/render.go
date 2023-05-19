package cairo

import (
    "time"
    "fmt"
    "log"
    "github.com/metux/freecity/core/base"
    "github.com/metux/freecity/core/items"
    "github.com/metux/freecity/core/game"
    "github.com/metux/freecity/render/theme"
    "github.com/metux/freecity/render"
    "github.com/gotk3/gotk3/cairo"
    "sync"
)

type FPoint = base.FPoint

type Renderer struct {
    Terrain       * items.TerrainMap
    Theme         * theme.ThemeSpec
    Context       * cairo.Context

    revTerrain      int64
    revObjects      int64
    cursorTile      base.Point

    imgSize         FPoint
    imageCache      map[string] * cairo.Surface
    imageCacheLock  sync.RWMutex
    loaded          bool
    basemap       * cairo.Surface
    fullmap       * cairo.Surface

    tileOrigin      FPoint
    tileSize        FPoint
    tileStepX       FPoint
    tileStepY       FPoint
    tileRStepX      FPoint
    tileRStepY      FPoint

    Viewport        render.Viewport

    statusmsg func(s string)
}

const defaultScale = 0.9
//const defaultScale = 1
const defaultPrescale = 0.75

//const basemapFormat = cairo.FORMAT_RGB24 // cairo.FORMAT_ARGB32
const basemapFormat = cairo.FORMAT_ARGB32
const fullmapFormat = basemapFormat
const spriteFormat = cairo.FORMAT_ARGB32

func (r * Renderer) init() {
    if r.loaded {
        return
    }

    r.Theme.LoadImages()
    r.imageCache = make(map[string] * cairo.Surface)

    r.tileSize = r.Viewport.PrescalePoint(r.Theme.TileSize.ToPoint())

    hts := r.tileSize.Mul(0.5)
    mapSize := FPoint{float64(r.Terrain.Size.X), float64(r.Terrain.Size.Y)}

    switch r.Theme.Projection {
        case theme.ProjFlat:
            r.imgSize    = FPoint{mapSize.X * r.tileSize.X, mapSize.Y * r.tileSize.Y}
            r.tileOrigin = FPoint{0,                        0}
            r.tileStepX  = FPoint{r.tileSize.X,             0}
            r.tileStepY  = FPoint{0,                        r.tileSize.Y}
            r.tileStepX  = FPoint{1/r.tileSize.X,           0}
            r.tileStepY  = FPoint{0,                        1/r.tileSize.Y}
        break
        case theme.ProjParallel:
            r.imgSize    = FPoint{mapSize.X * r.tileSize.X, mapSize.Y * r.tileSize.Y}
            r.tileOrigin = FPoint{r.imgSize.X / 2 - hts.X,  0}
            r.tileStepX  = FPoint{ hts.X,                   hts.Y}
            r.tileStepY  = FPoint{-hts.X,                   hts.Y}
            r.tileRStepX = FPoint{hts.X,                    -hts.Y}
            r.tileRStepY = FPoint{hts.X,                    hts.Y}
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

    scaledSz := r.Viewport.PrescalePoint(base.Point{img.GetWidth(), img.GetHeight()})
    scaledImg := cairo.CreateImageSurface(spriteFormat, int(scaledSz.X), int(scaledSz.Y))

    scaledCtx := cairo.Create(scaledImg)
    defer scaledCtx.Close()

    scaledCtx.Scale(r.Viewport.Prescale, r.Viewport.Prescale)

    scaledCtx.SetSourceSurface(img, 0, 0)
    scaledCtx.Paint()

    r.imageCache[name] = scaledImg
    return scaledImg
}

func (r* Renderer) overlayTileImage(cr * cairo.Context, p base.FPoint, name string) {
    if name != "" {
        if i2 := r.getImage(name); i2 != nil {
            cr.SetSourceSurface(i2, p.X, p.Y)
            cr.Paint()
        }
    }
}

func (r* Renderer) overlayTile(cr * cairo.Context, pos base.Point, name string) {
    r.overlayTileImage(cr, r.tilePos(pos), name)
}

func (r * Renderer) PointerPos(ptr FPoint) base.Point {
    // translate back from physical to virtual coordinates
    virt := r.Viewport.TranslateBack(ptr).Sub(r.tileOrigin)

    // normalize to tile size (--> 1x1)
    norm := virt.Compress(r.tileSize)

    n2, f2 := norm.Raster()
    f3 := FPoint{1-f2.X, 1-f2.Y}

    // corner triangle correction
    trans := base.Point{ n2.X + n2.Y, -n2.X + n2.Y}
    if f2.X < 0.5 && f2.Y < 0.5 && f2.X + f2.Y < 0.5 {
        trans.X -= 1
    }
    if f3.X < 0.5 && f2.Y < 0.5 && f3.X + f2.Y < 0.5 {
        trans.Y -= 1
    }
    if f2.X < 0.5 && f3.Y < 0.5 && f2.X + f3.Y < 0.5 {
        trans.Y += 1
    }
    if f2.X > 0.5 && f3.Y < 0.5 && f3.X + f3.Y < 0.5 {
        trans.X += 1
    }

    return trans
}

// FIXME: honor prescale
func (r* Renderer) UpdateCursor(pos base.FPoint) (int, int, int, int) {
    newCursorTile := r.PointerPos(pos)

    if newCursorTile == r.cursorTile {
        return 0, 0, 0, 0
    }

    old_pos := r.tilePos(r.cursorTile)
    new_pos := r.tilePos(newCursorTile)

    r.cursorTile = newCursorTile

    if old_pos.X < 0 || old_pos.Y < 0 {
        old_pos = new_pos
    }
    if new_pos.X < 0 || new_pos.Y < 0 {
        new_pos = old_pos
    }

    p1 := base.FPoint{base.Fmin(old_pos.X, new_pos.X), base.Fmin(old_pos.Y, new_pos.Y)}
    p2 := base.FPoint{base.Fmax(old_pos.X, new_pos.X) + r.tileSize.X, base.Fmax(old_pos.Y, new_pos.Y) + r.tileSize.Y}

    pmin := r.Viewport.TranslatePhys(p1).ToPoint()
    pmax := r.Viewport.TranslatePhys(p2).ToPoint()

    return pmin.X, pmin.Y, pmax.X, pmax.Y
}

func (r* Renderer) tilePos(pos base.Point) base.FPoint {
    pX, pY := float64(pos.X), float64(pos.Y)
    return base.FPoint{r.tileOrigin.X + r.tileStepX.X * pX + r.tileStepY.X * pY,
           r.tileOrigin.Y + r.tileStepX.Y * pX + r.tileStepY.Y * pY}
}

func (r* Renderer) renderTile(cr * cairo.Context, tr items.TileRef) {
    p := r.tilePos(tr.Position)

    r.overlayTileImage(cr, p, r.Theme.ImgForRubble(tr.Tile.Rubble))
    r.overlayTileImage(cr, p, r.Theme.ImgForZone(tr.Tile.ZoneTag))
    r.overlayTileImage(cr, p, r.Theme.ImgForRoad(tr.Tile.Road))
    r.overlayTileImage(cr, p, r.Theme.ImgForRail(tr.Tile.Rail))
    r.overlayTileImage(cr, p, r.Theme.ImgForPower(tr.Tile.Power))
}

func (r* Renderer) renderTileBase(cr * cairo.Context, tr items.TileRef) {
    p := r.tilePos(tr.Position)

    r.overlayTileImage(cr, p, theme.ImgLandFlat)
    if render.DebugRaster {
        r.overlayTileImage(cr, p, theme.ImgHarness)
    } else {
        r.overlayTileImage(cr, p, theme.ImgGrid)
    }
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
        imgX := r.tilePos(rect.BottomLeft())
        imgY := r.tilePos(rect.BottomRight())
        imgY.Y += r.tileSize.Y - float64(img.GetHeight())

        cr.SetSourceSurface(img,imgX.X,imgY.Y)
        cr.Paint()
    }
}

func (r * Renderer) createBasemap() {
    if r.basemap != nil && r.Terrain.RevTerrain == r.revTerrain {
        return
    }

    if render.DebugMode {
        defer base.TimeTrack(time.Now(), "cairo-render createBasemap()")
    }

    r.basemap = cairo.CreateImageSurface(basemapFormat, int(r.imgSize.X), int(r.imgSize.Y))

    ctx := cairo.Create(r.basemap)
    defer ctx.Close()

    tl := r.Terrain.AllTiles()
    for _,tr := range tl {
        r.renderTileBase(ctx, tr)
    }

    r.revTerrain = r.Terrain.RevTerrain
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
    if r.fullmap != nil && r.Terrain.RevObjects == r.revObjects {
        return
    }

    if render.DebugMode {
        defer base.TimeTrack(time.Now(), "cairo-render createFullmap()")
    }

    r.createBasemap()
    r.fullmap = cairo.CreateImageSurface(basemapFormat, int(r.imgSize.X), int(r.imgSize.Y))

    ctx := cairo.Create(r.fullmap)
    defer ctx.Close()

    r.renderFullmap(ctx)
    r.revObjects = r.Terrain.RevObjects
}

func (r * Renderer) Render(cr * cairo.Context) {
    r.init()
    r.createBasemap()
    r.createFullmap()
    cr.Scale(r.Viewport.Scale, r.Viewport.Scale)

    if render.DebugMode {
        defer base.TimeTrack(time.Now(), fmt.Sprintf("cairo-render (scale %f)", r.Viewport.Scale))
    }

    // take the basemap as background
    offset := r.Viewport.TranslatedOffset()
    cr.SetSourceSurface(r.fullmap, offset.X, offset.Y)
    cr.Paint()

    if r.Terrain.Size.HasPoint(r.cursorTile) {
        r.overlayTileImage(cr, r.tilePos(r.cursorTile).Add(offset), "cursor-1x1")
    }
}

func (r * Renderer) SetScale(sc float64) {
    sc = float64(int(sc*10000)) / 10000
    r.Viewport.Scale = sc
}

func (r * Renderer) ZoomStep(z float64) {
    r.SetScale(r.Viewport.Scale + z)
}

func (r * Renderer) Move(v base.FPoint) {
    r.Viewport.Offset.X += v.X
    r.Viewport.Offset.Y += v.Y
}

func CreateRenderer(g1 * game.Game, t * theme.ThemeSpec, statusmsg func(s string)) (* Renderer) {
    return &(Renderer{
        Terrain:  &g1.Terrain,
        Theme:    t,
        Viewport: render.Viewport{
            Scale:    defaultScale,
            Prescale: defaultPrescale},
        statusmsg: statusmsg,
    })
    return nil
}
