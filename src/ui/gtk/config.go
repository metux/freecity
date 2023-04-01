package gtk

import (
    "github.com/metux/freecity/core/base"
)

type Config struct {
    DataPrefix  string  `yaml:"-"`
    Theme       string  `yaml:"theme"`
    Prescale    float64 `yaml:"prescale"`
    Scale       float64 `yaml:"scale"`
    ZoomStep    float64 `yaml:"zoomstep"`
    MoveStep    float64 `yaml:"movestep"`
    MoveInvert  bool    `yaml:"moveinvert"`
}

func LoadUIYaml(prefix string) * Config {
    c := Config{
        Prescale:       0.5,
        Scale:          0.5,
        Theme:          "parallel",
        ZoomStep:       0.005,
        MoveStep:       10,
        DataPrefix:     prefix,
    }

    fn := prefix + "/ui/gtk.yaml"

    if err := base.YamlLoad(fn, &c); err != nil {
        return &c
    }

    return &c
}
