package rules

import (
    "log"
    "github.com/metux/freecity/core/base"
)

type GeneralRules struct {
    Startup struct {
        Date    date
        Funds   money
        Size    dim
    }
    Costs struct {
        Road                 money
        Rail                 money
        Bulldoze             money
        Powerline            money
        Zones struct {
            ResidentialLight money `yaml:"residential-light"`
            ResidentialDense money `yaml:"residential-dense"`
            IndustrialLight  money `yaml:"industrial-light"`
            IndustrialDense  money `yaml:"industrial-dense"`
            CommercialLight  money `yaml:"commercial-light"`
            CommercialDense  money `yaml:"commercial-dense"`
            SeaportLight     money `yaml:"seaport-light"`
            SeaportDense     money `yaml:"seaport-dense"`
            AirportLight     money `yaml:"airport-light"`
            AirportDense     money `yaml:"airport-dense"`
        }
    }
    Buildings BuildingRules
}

func (g * GeneralRules) LoadYaml(ruledir string) error {
    fn := ruledir + "/general.yaml"

    if err := base.YamlLoad(fn, g); err != nil {
        log.Println("failed loading general ruleset")
        return err
    }

    return g.Buildings.LoadYaml(ruledir)
}

func (g * GeneralRules) ZonePrice(zt zonetag) money {
    var zc = g.Costs.Zones
    switch (zt) {
        case base.ZoneResidentialLight:
            return zc.ResidentialLight
        case base.ZoneResidentialDense:
            return zc.ResidentialDense
        case base.ZoneIndustrialLight:
            return zc.IndustrialLight
        case base.ZoneIndustrialDense:
            return zc.IndustrialDense
        case base.ZoneCommercialLight:
            return zc.CommercialLight
        case base.ZoneCommercialDense:
            return zc.CommercialDense
        case base.ZoneAirportLight:
            return zc.AirportLight
        case base.ZoneAirportDense:
            return zc.AirportDense
        case base.ZoneSeaportLight:
            return zc.SeaportLight
        case base.ZoneSeaportDense:
            return zc.SeaportDense
    }
    return 0
}

func (g * GeneralRules) FindBuildingType(typename string) * BuildingType {
    return g.Buildings.ByName[typename]
}
