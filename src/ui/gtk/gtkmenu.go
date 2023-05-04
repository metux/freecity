package gtk

import (
    "log"
    "github.com/gotk3/gotk3/gtk"
    "github.com/metux/freecity/ui/common"
)

type GtkMenuEntryHandle struct {
    Shell * gtk.MenuShell
    Item * gtk.MenuItem
    Check * gtk.CheckMenuItem
}

func (meh GtkMenuEntryHandle) SetChecked(ch bool) bool {
    if meh.Check != nil {
        old := meh.Check.GetActive()
        meh.Check.SetActive(ch)
        return old
    }
    log.Println("WARN: menu item not CheckMenuItem")
    return false
}

func (meh GtkMenuEntryHandle) CreateEntrySeparator(me * common.MenuEntry) {
    mi,_ := gtk.SeparatorMenuItemNew()
    me.Handle = GtkMenuEntryHandle{Item: &mi.MenuItem}
    meh.Shell.Append(mi)
}

func (meh GtkMenuEntryHandle) CreateEntrySimple(me * common.MenuEntry) {
    mi,_ := gtk.MenuItemNewWithMnemonic(me.Label)
    me.Handle = GtkMenuEntryHandle{Item: mi}
    meh.Shell.Append(mi)
    me2 := me
    mi.Connect("activate", func() { me2.Activate() })
}

func (meh GtkMenuEntryHandle) CreateEntryRadio(me * common.MenuEntry) {
    mi,_ := gtk.RadioMenuItemNewWithMnemonic(nil, me.Label)
    me.Handle = GtkMenuEntryHandle{Item: &mi.MenuItem}
    meh.Shell.Append(mi)
    me2 := me
    mi.Connect("activate", func() { me2.Activate() })
}

func (meh GtkMenuEntryHandle) CreateEntryCheck(me * common.MenuEntry) {
    mi,_ := gtk.CheckMenuItemNewWithMnemonic(me.Label)
    me.Handle = GtkMenuEntryHandle{Item: &mi.MenuItem, Check: mi}
    meh.Shell.Append(mi)
    cl_me := me
    cl_mi := mi
    mi.Connect("activate", func() {
        if cl_mi.GetActive() {
            cl_me.Activate()
        }
    })
}

func (meh GtkMenuEntryHandle) CreateEntrySubmenu(me * common.MenuEntry) {
    mi,_ := gtk.MenuItemNewWithMnemonic(me.Label)
    menu,_ := gtk.MenuNew()
    me.Handle = GtkMenuEntryHandle{Item: mi, Shell: &menu.MenuShell}
    mi.SetSubmenu(menu)
    meh.Shell.Append(mi)
}

func GtkLoadMenuBar(root * common.MenuEntry) * gtk.MenuShell {
    menubar,_ := gtk.MenuBarNew()
    root.Handle = GtkMenuEntryHandle{Shell: &menubar.MenuShell}
    root.CreateEntries()
    return &menubar.MenuShell
}
