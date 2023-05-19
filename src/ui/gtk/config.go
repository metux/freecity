package gtk

import (
    "github.com/metux/freecity/util"
    "github.com/metux/freecity/ui/common"
)

type Config struct {
    WindowTitle   string              `yaml:"windowtitle"`
    WindowSize    point               `yaml:"windowsize"`
    DataPrefix    string              `yaml:"-"`
    Theme         string              `yaml:"theme"`
    Prescale      float64             `yaml:"prescale"`
    Scale         float64             `yaml:"scale"`
    ZoomStep      float64             `yaml:"zoomstep"`
    MoveStep      float64             `yaml:"movestep"`
    MoveInvert    bool                `yaml:"moveinvert"`
    MainMenu      common.MenuEntry    `yaml:"mainmenu"`
    KeyMap        map[string] string  `yaml:"keymap"`
}

func LoadUIYaml(prefix string) * Config {
    c := Config{
//        Prescale:       0.5,
        Prescale:       1,
        Scale:          0.5,
        Theme:          "parallel",
        ZoomStep:       0.005,
        MoveStep:       10,
        DataPrefix:     prefix,
        WindowSize:     point { 1024, 768 },
    }

    fn := prefix + "/ui/gtk.yaml"

    if err := util.YamlLoad(fn, &c); err != nil {
        return &c
    }

    return &c
}
