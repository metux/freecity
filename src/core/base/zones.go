package base

import (
    "gopkg.in/yaml.v3"
)

type ZoneTag byte

const (
    ZoneNone ZoneTag      = 0

    ZoneResidentialLight  = 'r'
    ZoneResidentialDense  = 'R'

    ZoneIndustrialLight   = 'i'
    ZoneIndustrialDense   = 'I'

    ZoneCommercialLight   = 'c'
    ZoneCommercialDense   = 'C'

    ZoneAirportLight      = 'a'
    ZoneAirportDense      = 'A'

    ZoneSeaportLight      = 's'
    ZoneSeaportDense      = 'S'
)

func (zt ZoneTag) MarshalYAML() (interface{}, error) {
    return string(zt), nil
}

func (zt * ZoneTag) UnmarshalYAML(value * yaml.Node) error {
    var tempStr string

    if err:= value.Decode(&tempStr); err != nil {
        return err
    }

    *zt = ZoneTag(tempStr[0])
    return nil
}

func (zt ZoneTag) MayUpgrade(nz ZoneTag) bool {
    switch (zt) {
        case ZoneNone:             return true
        case ZoneResidentialLight: return nz == ZoneResidentialDense
        case ZoneIndustrialLight:  return nz == ZoneIndustrialDense
        case ZoneCommercialLight:  return nz == ZoneCommercialDense
        case ZoneAirportLight:     return nz == ZoneAirportDense
        case ZoneSeaportLight:     return nz == ZoneSeaportDense
    }
    return false
}

func (zt ZoneTag) String() string {
    switch (zt) {
        case ZoneNone:              return "none"
        case ZoneResidentialLight:  return "residential-light"
        case ZoneResidentialDense:  return "residential-dense"
        case ZoneIndustrialLight:   return "industrial-light"
        case ZoneIndustrialDense:   return "industrial-dense"
        case ZoneCommercialLight:   return "commercial-light"
        case ZoneCommercialDense:   return "commercial-dense"
        case ZoneAirportLight:      return "airport-light"
        case ZoneAirportDense:      return "airport-dense"
        case ZoneSeaportLight:      return "seaport-light"
        case ZoneSeaportDense:      return "seaport-dense"
    }
    return string(zt)
}
