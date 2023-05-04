package render

import (
    "github.com/metux/freecity/core/base"
)

type Viewport struct {
    Offset   base.FPoint
    Scale    float64
    Prescale float64
}

func (vp Viewport) TranslateBack(ptr base.FPoint) base.FPoint {
    return base.FPoint{ptr.X / vp.Scale - vp.Offset.X * vp.Scale,
                       ptr.Y / vp.Scale - vp.Offset.Y * vp.Scale}
}

func (vp Viewport) TranslatePhys(ptr base.FPoint) base.FPoint {
    ptr.X += vp.Offset.X * vp.Scale
    ptr.Y += vp.Offset.Y * vp.Scale
    return base.FPoint{ptr.X * vp.Scale, ptr.Y * vp.Scale}
}

func (vp Viewport) TranslatedOffset() base.FPoint {
    return base.FPoint{vp.Offset.X * vp.Scale, vp.Offset.Y * vp.Scale}
}

func (vp Viewport) PrescalePoint(p base.Point) base.FPoint {
    return base.FPoint{
        float64(p.X) * vp.Prescale,
        float64(p.Y) * vp.Prescale}
}
