package util

import "fmt"

type Rect struct {
    X      int `yaml:"x"`
    Y      int `yaml:"y"`
    Width  int `yaml:"width"  default: "1"`
    Height int `yaml:"height" default: "1"`
}

func (r Rect) String() string {
    return fmt.Sprintf("[%d:%d:%d:%d]", r.X, r.Y, r.Width, r.Height)
}

func (r Rect) TopLeft() Point {
    return Point{r.X, r.Y}
}

func (r Rect) BottomLeft() Point {
    return Point{r.X, r.Y + r.Height - 1}
}

func (r Rect) TopRight() Point {
    return Point{r.X + r.Width - 1, r.Y}
}

func (r Rect) BottomRight() Point {
    return Point{r.X + r.Width - 1, r.Y + r.Height - 1}
}

func (r Rect) Valid() bool {
    return r.Width > 0 && r.Height > 0 && r.X > -1 && r.Y > -1
}
