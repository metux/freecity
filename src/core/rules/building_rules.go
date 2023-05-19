package rules

import (
    "github.com/metux/freecity/core/base"
    "log"
)

type BuildingRules struct {
    BuildingTypes []            BuildingType
    ByName        map[string] * BuildingType
}

func (this * BuildingRules) Rehash() {
    newmap := make(map[string]*BuildingType)
    for i := range this.BuildingTypes {
        elem := &this.BuildingTypes[i]
        newmap[elem.Name] = elem

        // fix missing values
        if (elem.Size.X == 0) {
            elem.Size.X = 1
        }
        if (elem.Size.Y == 0) {
            elem.Size.Y = 1
        }
    }
    this.ByName = newmap
}

func (this * BuildingRules) Set(l [] BuildingType) {
    this.BuildingTypes = l
    this.Rehash()
}

func (this * BuildingRules) LoadYaml(ruledir string) error {
    fn := ruledir + "/buildings.yaml"

    var result [] BuildingType

    if err := base.YamlLoad(fn, &result); err != nil {
        log.Println("failed loading building ruleset")
        return err
    }

    this.Set(result)

    return nil
}
