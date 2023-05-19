package gtk

import (
    "github.com/gotk3/gotk3/gdk"
)

func translateGdkKey(evkey * gdk.EventKey) string {
    prefix := "KEY_"
    if (evkey.State() & gdk.CONTROL_MASK == gdk.CONTROL_MASK) {
        prefix = "CTRL_"
    }
    return prefix+gdk.KeyValName(evkey.KeyVal())
}

func evMotionFPoint(ev * gdk.Event) fpoint {
    ev2 := gdk.EventMotion{ev}
    x, y := ev2.MotionVal()
    return fpoint{x, y}
}
