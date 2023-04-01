package base

import (
    "fmt"
    "math"
)

type FPoint struct {
    X float64
    Y float64
}

func (p FPoint) String() string {
    return fmt.Sprintf("[%7.3f:%7.3f]", p.X, p.Y)
}

func (p FPoint) Sub(p2 FPoint) FPoint {
    return FPoint{p.X - p2.X, p.Y - p2.Y}
}

func (p FPoint) Add(p2 FPoint) FPoint {
    return FPoint{p.X + p2.X, p.Y + p2.Y}
}

func (p FPoint) Mul(factor float64) FPoint {
    return FPoint{p.X * factor, p.Y * factor}
}

func (p FPoint) Raster() (Point, FPoint) {
    ix, fx := math.Modf(p.X)
    iy, fy := math.Modf(p.Y)

    if p.X < 0 {
        ix--
        fx = 1 + fx
    }
    if p.Y < 0 {
        iy--
        fy =  1 + fy
    }
    return Point{int(ix), int(iy)}, FPoint{fx, fy}
}

func (p FPoint) Compress(v FPoint) FPoint {
    return FPoint{p.X / v.X, p.Y / v.Y}
}

func (p FPoint) ToPoint() Point {
    return Point{int(p.X), int(p.Y)}
}
