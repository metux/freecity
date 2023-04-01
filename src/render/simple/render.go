package simple

import (
    "log"
    "os"
    "github.com/metux/freecity/core/base"
    "github.com/metux/freecity/core/items"
    "github.com/metux/freecity/render/theme"
    "sync"
    "image"
    "image/draw"
    "image/png"
)

type Renderer struct {
    Terrain       * items.TerrainMap
    Theme         * theme.ThemeSpec
    Img           * image.RGBA
    imgDim          image.Point
    imgOrigin       image.Point
    imgStepX        image.Point
    imgStepY        image.Point
    imgTileCopyOp   draw.Op
}

func (r* Renderer) overlayTileImage(tRect image.Rectangle, name string) {
    img := r.Theme.GetImage(name)
    if (img != nil) {
        draw.Draw(r.Img, tRect, img, image.Point{0, 0}, draw.Over)
    }
}

func (r* Renderer) tileRect(pos base.Point) image.Rectangle {
    coords := r.imgOrigin.Add(r.imgStepX.Mul(pos.X)).Add(r.imgStepY.Mul(pos.Y))
    return image.Rect(
        coords.X,
        coords.Y,
        coords.X + r.Theme.TileSize.Width,
        coords.Y + r.Theme.TileSize.Height)
}

func (r* Renderer) renderTile(wg * sync.WaitGroup, tr items.TileRef) {
    tRect := r.tileRect(tr.Position)

    r.overlayTileImage(tRect, theme.ImgLandFlat)
    if tr.Tile.Rubble {
        r.overlayTileImage(tRect, theme.ImgRubble)
    }

    // render the zone background
    if z := tr.Tile.ZoneTag; z != base.ZoneNone {
        r.overlayTileImage(tRect, theme.ImgZonesPrefix + z.String())
    }

    if road := tr.Tile.Road; road != base.LineDirNone {
        r.overlayTileImage(tRect, theme.ImgRoadPrefix + road.Ident())
    }

    if p := tr.Tile.Power; p != base.LineDirNone {
        r.overlayTileImage(tRect, theme.ImgPowerlinePrefix + p.Ident())
    }

    r.overlayTileImage(tRect, theme.ImgGrid)
    wg.Done()
}

func (r * Renderer) renderTileBorder(rect base.Rect) {
    r.overlayTileImage(r.tileRect(rect.TopLeft()),     theme.ImgBorderTopLeft)
    r.overlayTileImage(r.tileRect(rect.TopRight()),    theme.ImgBorderTopRight)
    r.overlayTileImage(r.tileRect(rect.BottomLeft()),  theme.ImgBorderBottomLeft)
    r.overlayTileImage(r.tileRect(rect.BottomRight()), theme.ImgBorderBottomRight)

    if (rect.Width > 2) {
        for x:=rect.X+1; x<rect.X + rect.Width-1; x++ {
            r.overlayTileImage(r.tileRect(base.Point{x, rect.Y}), theme.ImgBorderTop)
            r.overlayTileImage(r.tileRect(base.Point{x, rect.Y + rect.Height - 1}), theme.ImgBorderBottom)
        }
    }

    if (rect.Height > 2) {
        for y:=rect.Y+1; y<rect.Y + rect.Height-1; y++ {
            r.overlayTileImage(r.tileRect(base.Point{rect.X, y}), theme.ImgBorderLeft)
            r.overlayTileImage(r.tileRect(base.Point{rect.X + rect.Width - 1, y}), theme.ImgBorderRight)
        }
    }
}

func (r * Renderer) renderBuilding(b * items.Building) {
    rect := b.OccupiedRect()

    r.renderTileBorder(rect)

    if img := r.Theme.GetImage("buildings/"+b.TypeName); img != nil {
        bounds := img.Bounds()

        p := image.Point{
            r.tileRect(rect.BottomLeft()).Min.X,
            r.tileRect(rect.BottomRight()).Max.Y - (bounds.Max.Y - bounds.Min.Y),
        }

        draw.Draw(r.Img,
                  image.Rect(p.X, p.Y, p.X + bounds.Max.X, p.Y + bounds.Max.Y),
                  img,
                  image.Point{},
                  draw.Over)
    }
}

func (r * Renderer) init() {
    r.Theme.LoadImages()

    ts  := r.Theme.TileSize.ImagePoint()
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

    if r.Img == nil {
        r.Img = image.NewRGBA(image.Rect(0, 0, r.imgDim.X, r.imgDim.Y))
    }
}

func (r * Renderer) RenderTerrain() {
    r.init()
    tl := r.Terrain.AllTiles()
    wg := new(sync.WaitGroup)
    wg.Add(len(tl))
    for _,tr := range tl {
        r.renderTile(wg, tr)
    }
    wg.Wait()
    log.Println("rendering done")

    for _,b := range r.Terrain.Buildings {
        r.renderBuilding(b)
    }
}

func (r * Renderer) SavePNG(fn string) {
    log.Println("storing PNG")
    f,_ := os.Create(fn)
    png.Encode(f, r.Img)
    log.Println("store done")
}

func (r * Renderer) Nop() {
}
