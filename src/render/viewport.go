package render

import (
    "github.com/metux/freecity/util/geo"
)

type fpoint = geo.FPoint
type point = geo.Point

type Viewport struct {
    Offset    fpoint
    Scale     float64
    Prescale  float64
}

func (vp Viewport) TranslateBack(ptr fpoint) fpoint {
    return fpoint{ptr.X / vp.Scale - vp.Offset.X * vp.Scale,
                  ptr.Y / vp.Scale - vp.Offset.Y * vp.Scale}
}

func (vp Viewport) TranslatePhys(ptr fpoint) fpoint {
    ptr.X += vp.Offset.X * vp.Scale
    ptr.Y += vp.Offset.Y * vp.Scale
    return ptr.Mul(vp.Scale)
}

func (vp Viewport) TranslatedOffset() fpoint {
    return vp.Offset.Mul(vp.Scale)
}

func (vp Viewport) PrescalePoint(p point) fpoint {
    return fpoint{
        float64(p.X) * vp.Prescale,
        float64(p.Y) * vp.Prescale}
}
