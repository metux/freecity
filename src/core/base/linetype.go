package base

type LineType uint8

const (
    LineTypeNone  = 0
    LineTypePower = iota
    LineTypeRoad  = iota
    LineTypeRail  = iota
    LineTypePipe  = iota
)
