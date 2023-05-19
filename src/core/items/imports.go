package items

import (
    "github.com/metux/freecity/util"
    "github.com/metux/freecity/core/base"
)

type point         = util.Point
type rect          = util.Rect
type Money         = base.Money
type LineDirection = base.LineDirection
type ZoneTag       = base.ZoneTag
type Action        = base.Action
type date          = util.Date

const ActionZoning         = base.ActionZoning
const ActionBuildRail      = base.ActionBuildRail
const ActionBuildPowerline = base.ActionBuildPowerline
