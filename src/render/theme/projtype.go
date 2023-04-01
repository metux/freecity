package theme

import (
    "errors"
    "gopkg.in/yaml.v3"
)

type ProjType byte

const (
    ProjUnknown ProjType = 0
    ProjFlat = iota
    ProjParallel = iota
)

func (p ProjType) String() string {
    switch p {
        case ProjFlat:     return "flat"
        case ProjParallel: return "parallel"
    }
    return "unknown"
}

func (p * ProjType) FromString(str string) error {
    switch str {
        case "flat":     *p = ProjFlat
        case "parallel": *p = ProjParallel
        default:
            return errors.New("cant parse ProjType "+str)
    }
    return nil
}

func (p ProjType) MarshalYAML() (interface{}, error) {
    return p.String(), nil
}

func (p * ProjType) UnmarshalYAML(value * yaml.Node) error {
    var tmpStr string

    if err := value.Decode(&tmpStr); err != nil {
        return err
    }

    return p.FromString(tmpStr)
}
