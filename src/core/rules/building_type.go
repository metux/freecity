package rules

import (
    "gopkg.in/yaml.v3"
)

type BuildingType struct {
    // if negative, the building is a power supplier
    Name                string      `yaml:"name"`
    Size                dim    `yaml:"size"`
    Costs struct {
        Build       money   `yaml:"build" default: "0"`
        Maint       money   `yaml:"maint"`
        Destroy     money   `yaml:"destroy"`
    }
    Consumption         consumption
    Placable            bool        `yaml:"placable"`
    RoutePower          bool        `yaml:"route_power"`
    Require             string      `yaml:"require" default:"undefined"`
    Routes struct {
        Power       bool    `yaml:"power" default: "true"`
        Water       bool    `yaml:"water" default: "true"`
        Road        bool    `yaml:"road"  default: "true"`
        Rail        bool    `yaml:"rail"  default: "true"`
    }
}

func (this * BuildingType) String() string {
    d, err := yaml.Marshal(this)
    if err != nil {
        return "(yaml encode failed)"
    }
    return string(d)
}
