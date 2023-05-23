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
    widgetStatusBar * gtk.Statusbar
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

func (sb * StatusBarWindow) Init(parent * gtk.Box) {
    sb.widgetStatusBar,_ = gtk.StatusbarNew()
    parent.PackEnd(sb.widgetStatusBar, false, false, 0)

    sep1,_ := gtk.SeparatorNew(gtk.ORIENTATION_VERTICAL)
    sep1.SetMarginStart(margin)
    sep1.SetMarginEnd(margin)
    sb.widgetStatusBar.PackStart(sep1, false, false, padding)

    sb.widgetMessage,_ = gtk.LabelNew("")
    sb.widgetStatusBar.PackStart(sb.widgetMessage, false, false, padding)

    sb.widgetTool,_ = gtk.LabelNew("")
    sb.widgetTool.SetSizeRequest(labelWidth, labelHeight)
    sb.widgetStatusBar.PackStart(sb.widgetTool, false, false, padding)

    sb.Update()
}
