package theme

import (
    "log"
    "image"
    "sync"
    "github.com/metux/freecity/util"
    "github.com/metux/freecity/util/geo"
    "github.com/metux/freecity/core/base"
)

const (
    ImgBorderLeft        = "tiles/border/left"
    ImgBorderRight       = "tiles/border/right"
    ImgBorderTop         = "tiles/border/top"
    ImgBorderTopLeft     = "tiles/border/top-left"
    ImgBorderTopRight    = "tiles/border/top-right"
    ImgBorderBottom      = "tiles/border/bottom"
    ImgBorderBottomLeft  = "tiles/border/bottom-left"
    ImgBorderBottomRight = "tiles/border/bottom-right"
    ImgGrid              = "tiles/grid"
    ImgHarness           = "tiles/harness"
    ImgZonesPrefix       = "tiles/zones/"
    ImgPowerlinePrefix   = "tiles/powerline/"
    ImgRoadPrefix        = "tiles/road/"
    ImgRailPrefix        = "tiles/rail/"
    ImgLandFlat          = "tiles/terrain/land/flat"
    ImgRubble            = "tiles/rubble"
)

type ThemeSpec struct {
    ThemeDir            string                    `yaml:"-"`
    TileSize            geo.Point                 `yaml:"tilesize"`
    Projection          ProjType                  `yaml:"projection"`
    ImageCache          map[string] * image.RGBA  `yaml:"-"`
    ImageCacheLock      sync.RWMutex              `yaml:"-"`
    Images struct {
        loaded bool
        Powerlines [base.LineDirMax] * image.RGBA
    } `yaml:"-"`
}

func (t * ThemeSpec) ImgForRoad(dir base.LineDirection) string {
    if dir == base.LineDirNone {
        return ""
    }
    return ImgRoadPrefix+dir.Ident()
}

func (t * ThemeSpec) ImgForRail(dir base.LineDirection) string {
    if dir == base.LineDirNone {
        return ""
    }
    return ImgRailPrefix+dir.Ident()
}

func (t * ThemeSpec) ImgForPower(dir base.LineDirection) string {
    if dir == base.LineDirNone {
        return ""
    }
    return ImgRailPrefix+dir.Ident()
}

func (t * ThemeSpec) ImgForZone(zt base.ZoneTag) string {
    if zt != base.ZoneNone {
        return "tiles/zones/"+zt.String()
    }
    return ""
}

func (t * ThemeSpec) ImgForRubble(rubble bool) string {
    if rubble {
        return ImgRubble
    } else {
        return ""
    }
}

func (t * ThemeSpec) ForZone(zt base.ZoneTag) * image.RGBA {
    if zt != base.ZoneNone {
        return t.GetImage("tiles/zones/"+zt.String())
    }
    return nil
}

func (t * ThemeSpec) ImgPath(name string) string {
    return t.ThemeDir+"/"+name+".png"
}

func (t * ThemeSpec) GetImage(name string) * image.RGBA {
    t.ImageCacheLock.Lock()
    defer t.ImageCacheLock.Unlock()

    if img, okay := t.ImageCache[name]; okay {
        return img
    }

    if img := loadPNG(t.ImgPath(name)); img != nil {
        log.Println("loaded theme image", name)
        t.ImageCache[name] = img
        return img
    }
    return nil
}

func (t * ThemeSpec) LoadImages() {
    if t.Images.loaded {
        return
    }

    t.ImageCache = make(map[string]*image.RGBA)

    // load the powerline images
    for x := 0; x<base.LineDirMax; x++ {
        t.Images.Powerlines[x] = t.loadPixmap("tiles/powerline", base.LineDirection(x).Ident())
    }

    // mark everything loaded
    t.Images.loaded = true
}

func (t * ThemeSpec) ImagePowerline(d base.LineDirection) * image.RGBA {
    return t.Images.Powerlines[d]
}

func (ts * ThemeSpec) LoadYaml(themedir string) error {
    fn := themedir + "/theme.yaml"

    if err := util.YamlLoad(fn, ts); err != nil {
        log.Println("failed loading theme from", themedir)
        return err
    }

    ts.ThemeDir = themedir
    ts.LoadImages()

    return nil
}
