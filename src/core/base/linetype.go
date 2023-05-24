package base

type LineType uint8

const (
    LineTypeNone  = LineType(0)
    LineTypePower = iota
    LineTypeRoad  = iota
    LineTypeRail  = iota
    LineTypePipe  = iota
)

const (
    LineTypeNamePower = "powerline"
    LineTypeNameRoad  = "road"
    LineTypeNameRail  = "rail"
    LineTypeNamePipe  = "pipe"
)

func (lt LineType) String() string {
    switch lt {
        case LineTypePower: return LineTypeNamePower
        case LineTypeRoad:  return LineTypeNameRoad
        case LineTypeRail:  return LineTypeNameRail
        case LineTypePipe:  return LineTypeNamePipe
    }
    return ""
}
