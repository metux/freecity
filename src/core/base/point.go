package base

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

func (p Point) MarshalYAML() (interface{}, error) {
    return fmt.Sprintf("%d;%d", p.X, p.Y), nil
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

func RandPoint(d Dim) Point {
    return Point{rand.Intn(d.Width), rand.Intn(d.Height)}
}
