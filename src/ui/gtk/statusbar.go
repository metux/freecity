package gtk

import (
    "github.com/gotk3/gotk3/gtk"
)

const (
    padding = uint(0)
    labelWidth = 50
    labelHeight = 20
    margin = 10
)

type StatusBarWindow struct {
    widgetStatusBar * gtk.Statusbar
    widgetToolName  * gtk.Label

    message string
    toolName string
}

func (sb * StatusBarWindow) SetMessage(s string) {
    sb.message = s
    if (sb.widgetStatusBar != nil) {
        sb.widgetStatusBar.Push(3, s)
    }
}

func (sb * StatusBarWindow) SetToolName(n string) {
    sb.toolName = n
    if sb.widgetToolName != nil {
        sb.widgetToolName.SetText(n)
    }
}

func (sb * StatusBarWindow) Init(parent * gtk.Box) {
    sb.widgetStatusBar,_ = gtk.StatusbarNew()
    parent.PackEnd(sb.widgetStatusBar, false, false, 0)

    sep1,_ := gtk.SeparatorNew(gtk.ORIENTATION_VERTICAL)
    sep1.SetMarginStart(margin)
    sep1.SetMarginEnd(margin)
    sb.widgetStatusBar.PackStart(sep1, false, false, padding)

    sb.widgetToolName,_ = gtk.LabelNew("")
    sb.widgetToolName.SetSizeRequest(labelWidth, labelHeight)
    sb.widgetStatusBar.PackStart(sb.widgetToolName, false, false, padding)

    sb.SetMessage(sb.message)
    sb.SetToolName(sb.toolName)
}
