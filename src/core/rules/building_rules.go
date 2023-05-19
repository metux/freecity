package rules

import (
    "github.com/metux/freecity/util"
    "log"
)

type BuildingRules struct {
    BuildingTypes []            BuildingType
    Placable      [] *          BuildingType
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

    placable := make([]*BuildingType, 0)
    for idx,_:= range this.BuildingTypes {
        if this.BuildingTypes[idx].Placable {
            placable = append(placable, &this.BuildingTypes[idx])
        }
    }
    this.Placable = placable

    this.Rehash()
}

func (this * BuildingRules) LoadYaml(ruledir string) error {
    fn := ruledir + "/buildings.yaml"

    var result [] BuildingType

    if err := util.YamlLoad(fn, &result); err != nil {
        log.Println("failed loading building ruleset")
        return err
    }

    this.Set(result)

    return nil
}
