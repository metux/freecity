package base

import (
    "fmt"
    "image"
    "gopkg.in/yaml.v3"
)

type Dim struct {
    Width  int `yaml:"width,omitempty"  default: "1"`
    Height int `yaml:"height,omitempty" default: "1"`
}

func (d Dim) String() string {
    return fmt.Sprintf("%dx%d", d.Width, d.Height)
}

func (p Dim) ImagePoint(scale float64) image.Point {
    return image.Point{
        int(float64(p.Width) * scale),
        int(float64(p.Height) * scale)}
}

func (d Dim) HasPoint(p Point) bool {
    return (p.X >= 0) && (p.X < d.Width) && (p.Y >= 0) && (p.Y < d.Height)
}

func (d Dim) MarshalYAML() (interface{}, error) {
    return fmt.Sprintf("%d;%d", d.Width, d.Height), nil
}

func (d Dim) ToRect() Rect {
    return Rect{0, 0, d.Width, d.Height}
}

func (d Dim) ToPoint() Point {
    return Point{d.Width, d.Height}
}

func (d *Dim) UnmarshalYAML(value *yaml.Node) error {
    var tmpStr string

    if err := value.Decode(&tmpStr); err != nil {
        return err
    }

    if _, err := fmt.Sscanf(tmpStr, "%d;%d", &d.Width, &d.Height); err != nil {
        return err
    }
    return nil
}
