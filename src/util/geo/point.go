package geo

import (
    "fmt"
    "math/rand"
    "gopkg.in/yaml.v3"
)

type Point struct {
    X int `yaml:"x,omitempty"`
    Y int `yaml:"y,omitempty"`
}

func (p Point) String() string {
    return fmt.Sprintf("[%d:%d]", p.X, p.Y)
}

func (p Point) North() Point {
    return Point{p.X, p.Y-1}
}

func (p Point) South() Point {
    return Point{p.X, p.Y+1}
}

func (p Point) East() Point {
    return Point{p.X+1, p.Y}
}

func (p Point) West() Point {
    return Point{p.X-1, p.Y}
}

func (p Point) ToFPoint() FPoint {
    return FPoint{float64(p.X), float64(p.Y)}
}

func (p Point) MarshalYAML() (interface{}, error) {
    return fmt.Sprintf("%d;%d", p.X, p.Y), nil
}

func (p Point) MakeRect(size Point) Rect {
    return Rect{ X: p.X, Y: p.Y, Width: size.X, Height: size.Y }
}

func (p Point) SpanRect() Rect {
    return Rect{ X: 0, Y: 0, Width: p.X, Height: p.Y }
}

func (p Point) HasPoint(p2 Point) bool {
    return (p2.X >= 0) && (p2.X < p.X) && (p2.Y >= 0) && (p2.Y < p.Y)
}

func (p Point) Surrounding() Rect {
    return Rect{p.X - 1, p.Y - 1, 3, 3}
}

func (p Point) DoOnPointAndSurrounding(f func(p Point)) {
    f(p)
    f(p.North())
    f(p.East())
    f(p.South())
    f(p.West())
}

func (p *Point) UnmarshalYAML(value *yaml.Node) error {
    var tmpStr string

    if err := value.Decode(&tmpStr); err != nil {
        return err
    }

    if _, err := fmt.Sscanf(tmpStr, "%d;%d", &p.X, &p.Y); err != nil {
        return err
    }
    return nil
}

func RandPoint(d Point) Point {
    return Point{rand.Intn(d.X), rand.Intn(d.Y)}
}
