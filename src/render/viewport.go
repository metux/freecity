package render

import (
    "github.com/metux/freecity/util"
)

type Viewport struct {
    Offset   util.FPoint
    Scale    float64
    Prescale float64
}

func (vp Viewport) TranslateBack(ptr util.FPoint) util.FPoint {
    return util.FPoint{ptr.X / vp.Scale - vp.Offset.X * vp.Scale,
                       ptr.Y / vp.Scale - vp.Offset.Y * vp.Scale}
}

func (vp Viewport) TranslatePhys(ptr util.FPoint) util.FPoint {
    ptr.X += vp.Offset.X * vp.Scale
    ptr.Y += vp.Offset.Y * vp.Scale
    return util.FPoint{ptr.X * vp.Scale, ptr.Y * vp.Scale}
}

func (vp Viewport) TranslatedOffset() util.FPoint {
    return util.FPoint{vp.Offset.X * vp.Scale, vp.Offset.Y * vp.Scale}
}

func (vp Viewport) PrescalePoint(p util.Point) util.FPoint {
    return util.FPoint{
        float64(p.X) * vp.Prescale,
        float64(p.Y) * vp.Prescale}
}
