package gtk

import (
    "github.com/gotk3/gotk3/glib"
    "github.com/gotk3/gotk3/gtk"
    "github.com/metux/freecity/util"
)

const (
    padding = uint(0)
    labelWidth = 50
    labelHeight = 20
    margin = 10
)

type StatusBarWindow struct {
    widgetStatusBar * gtk.Box
    widgetTool      * gtk.Label
    widgetClock     * gtk.Label
    widgetMessage   * gtk.Label

    message           string
    tool              string
    date              util.Date
}

// FIXME: create our own widget

func (sb * StatusBarWindow) Update() {
    glib.TimeoutAdd(100, func() {
        if sb.widgetMessage != nil {
            sb.widgetMessage.SetText(sb.message)
        }
        if sb.widgetTool != nil {
            sb.widgetTool.SetText(sb.tool)
        }
        if sb.widgetClock != nil {
            sb.widgetClock.SetText(sb.date.String())
        }
    })
}

func (sb * StatusBarWindow) SetDate(d util.Date) {
    sb.date = d
    sb.Update()
}

func (sb * StatusBarWindow) SetMessage(s string) {
    sb.message = s
    sb.Update()
}

func (sb * StatusBarWindow) SetToolName(n string) {
    sb.tool = n
    sb.Update()
}

func (sb * StatusBarWindow) pack(w gtk.IWidget, end bool) {
    if end {
        sb.widgetStatusBar.PackEnd(w, false, false, padding)
    } else {
        sb.widgetStatusBar.PackStart(w, false, false, padding)
    }
}

func (sb * StatusBarWindow) addSep(end bool) {
    sep1,_ := gtk.SeparatorNew(gtk.ORIENTATION_VERTICAL)
    sep1.SetMarginStart(margin)
    sep1.SetMarginEnd(margin)
    sb.pack(sep1, end)
}

func (sb * StatusBarWindow) addLabel(end bool) * gtk.Label {
    l,_ := gtk.LabelNew("")
    sb.pack(l, end)
    return l
}

func (sb * StatusBarWindow) Init(parent * gtk.Box) {
    sb.widgetStatusBar,_ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
    parent.PackEnd(sb.widgetStatusBar, false, false, 0)

    // left side
    sb.widgetMessage = sb.addLabel(false)

    // right side
    sb.addSep(true)

    sb.widgetTool = sb.addLabel(true)
    sb.widgetTool.SetSizeRequest(labelWidth, labelHeight)

    sb.addSep(true)

    sb.widgetClock = sb.addLabel(true)
    sb.widgetClock.SetSizeRequest(labelWidth, labelHeight)

    sb.Update()
}
