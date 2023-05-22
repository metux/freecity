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
    StatusBar * gtk.Statusbar
    ToolName  * gtk.Label
}

func (sb * StatusBarWindow) SetMessage(s string) {
    if (sb.StatusBar != nil) {
        sb.StatusBar.Push(3, s)
    }
}

func (sb * StatusBarWindow) SetToolName(n string) {
    if sb.ToolName != nil {
        sb.ToolName.SetText(n)
    }
}

func (sb * StatusBarWindow) Init(parent * gtk.Box) {
    sb.StatusBar,_ = gtk.StatusbarNew()
    parent.PackEnd(sb.StatusBar, false, false, 0)

    sep1,_ := gtk.SeparatorNew(gtk.ORIENTATION_VERTICAL)
    sep1.SetMarginStart(margin)
    sep1.SetMarginEnd(margin)
    sb.StatusBar.PackStart(sep1, false, false, padding)

    sb.ToolName,_ = gtk.LabelNew("FOO")
    sb.ToolName.SetSizeRequest(labelWidth, labelHeight)
    sb.StatusBar.PackStart(sb.ToolName, false, false, padding)
}
