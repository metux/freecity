package items

import (
    "github.com/metux/freecity/util"
    "github.com/metux/freecity/util/geo"
    "github.com/metux/freecity/core/base"
)

type point         = geo.Point
type rect          = geo.Rect
type Money         = base.Money
type LineDirection = base.LineDirection
type LineType      = base.LineType
type ZoneTag       = base.ZoneTag
type Action        = base.Action
type date          = util.Date

const ActionZoning         = base.ActionZoning
const ActionBuildPowerline = base.ActionBuildPowerline

const (
    LtPower = base.LineTypePower
    LtRoad  = base.LineTypeRoad
    LtRail  = base.LineTypeRail
    LtPipe  = base.LineTypePipe
)
