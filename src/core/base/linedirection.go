package base

import (
    "log"
    "errors"
    "gopkg.in/yaml.v3"
    "github.com/metux/freecity/util"
)

type LineDirection uint8

const (
    LineDirNone LineDirection = 0

    LineDirNorthEast          = iota
    LineDirNorthEastSouth     = iota
    LineDirNorthEastSouthWest = iota
    LineDirNorthEastWest      = iota
    LineDirNorthSouth         = iota
    LineDirNorthSouthWest     = iota
    LineDirNorthWest          = iota

    LineDirEastSouth          = iota
    LineDirEastSouthWest      = iota
    LineDirEastWest           = iota

    LineDirSouthWest          = iota
    LineDirMax                = iota

    LineDirCrossing           = LineDirNorthEastSouthWest
)

func (d LineDirection) CheckCompat(other LineDirection) bool {
    switch d {
        case LineDirNorthSouth:
            return other == LineDirEastWest
        case LineDirEastWest:
            return other == LineDirNorthSouth
        case LineDirNone:
            return true
    }
    return false
}

func (d LineDirection) PickOther() LineDirection {
    if d.CheckCompat(LineDirEastWest) {
        return LineDirEastWest
    }
    if d.CheckCompat(LineDirNorthSouth) {
        return LineDirNorthSouth
    }
    return LineDirNone
}

func (d LineDirection) Ident() string {
    switch d {
        case LineDirNorthEast:          return "1100"
        case LineDirNorthEastSouth:     return "1110"
        case LineDirNorthEastSouthWest: return "1111"
        case LineDirNorthEastWest:      return "1101"
        case LineDirNorthSouth:         return "1010"
        case LineDirNorthSouthWest:     return "1011"
        case LineDirNorthWest:          return "1001"
        case LineDirEastSouth:          return "0110"
        case LineDirEastSouthWest:      return "0111"
        case LineDirEastWest:           return "0101"
        case LineDirSouthWest:          return "0011"
        case LineDirNone:               return "0000"
    }
    return "empty"
}

func (d * LineDirection) FromIdent(str string) error {
    switch str {
        case "1100": *d = LineDirNorthEast
        case "1110": *d = LineDirNorthEastSouth
        case "1111": *d = LineDirNorthEastSouthWest
        case "1101": *d = LineDirNorthEastWest
        case "1010": *d = LineDirNorthSouth
        case "1011": *d = LineDirNorthSouthWest
        case "1001": *d = LineDirNorthWest
        case "0110": *d = LineDirEastSouth
        case "0111": *d = LineDirEastSouthWest
        case "0101": *d = LineDirEastWest
        case "0011": *d = LineDirSouthWest
        case "0000": *d = LineDirNone
        default: return errors.New("cant parse str" + str)
    }
    return nil
}

// compute connection from neighboring tiles
func LineDirectionFromVec(north, east, south, west bool) (LineDirection) {
    switch {
        case !north && !east && !south && !west: return LineDirEastWest
        case !north && !east && !south &&  west: return LineDirEastWest
        case !north && !east &&  south && !west: return LineDirNorthSouth
        case !north && !east &&  south &&  west: return LineDirSouthWest
        case !north &&  east && !south && !west: return LineDirEastWest
        case !north &&  east && !south &&  west: return LineDirEastWest
        case !north &&  east &&  south && !west: return LineDirEastSouth
        case !north &&  east &&  south &&  west: return LineDirEastSouthWest
        case  north && !east && !south && !west: return LineDirNorthSouth
        case  north && !east && !south &&  west: return LineDirNorthWest
        case  north && !east &&  south && !west: return LineDirNorthSouth
        case  north && !east &&  south &&  west: return LineDirNorthSouthWest
        case  north &&  east && !south && !west: return LineDirNorthEast
        case  north &&  east && !south &&  west: return LineDirNorthEastWest
        case  north &&  east &&  south && !west: return LineDirNorthEastSouth
        case  north &&  east &&  south &&  west:
            return LineDirNorthEastSouthWest
        default:
            log.Println("unhandled vec")
            return LineDirEastWest
    }
    return LineDirEastWest
}

func (d LineDirection) Present() bool {
    return d != LineDirNone
}

func (d LineDirection) None() bool {
    return d == LineDirNone
}

func (d LineDirection) MarshalYAML() (interface{}, error) {
    return d.Ident(), nil
}

func (d * LineDirection) UnmarshalYAML(value * yaml.Node) (error) {
    var tmpStr string

    if err := value.Decode(&tmpStr); err != nil {
        return err
    }

    return d.FromIdent(tmpStr)
}

func LineDirPick(other_a, other_b LineDirection) LineDirection {
    if (other_a == LineDirNorthSouth && other_b == LineDirNone) ||
       (other_b == LineDirNorthSouth && other_a == LineDirNone) ||
       (other_a == LineDirNone && other_b == LineDirNone) {
        return LineDirEastWest
    }

    if (other_a == LineDirEastWest && other_b == LineDirNone) ||
       (other_b == LineDirEastWest && other_a == LineDirNone) {
        return LineDirNorthSouth
    }

    return LineDirNone
}

func (d * LineDirection) PickFromSurrounding(p util.Point, f func(p util.Point) bool) {
    *d = LineDirectionFromVec(
            f(p.North()),
            f(p.East()),
            f(p.South()),
            f(p.West()))
}
