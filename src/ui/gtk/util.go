package gtk

import (
    "github.com/gotk3/gotk3/gdk"
    "github.com/metux/freecity/core/base"
)

func translateGdkKey(evkey * gdk.EventKey) string {
    prefix := "KEY_"
    if (evkey.State() & gdk.CONTROL_MASK == gdk.CONTROL_MASK) {
        prefix = "CTRL_"
    }
    return prefix+gdk.KeyValName(evkey.KeyVal())
}

func evMotionFPoint(ev * gdk.Event) base.FPoint {
    ev2 := gdk.EventMotion{ev}
    x, y := ev2.MotionVal()
    return base.FPoint{x, y}
}
