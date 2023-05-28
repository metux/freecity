package common

import (
    "github.com/metux/freecity/util/cmd"
)

type MenuEntryHandle interface {
    SetChecked(ch bool) bool
    CreateEntrySeparator(me * MenuEntry)
    CreateEntrySimple(me * MenuEntry)
    CreateEntrySubmenu(me * MenuEntry)
    CreateEntryCheck(me * MenuEntry)
}

const SeparatorId   = "----"
const TypeCheck     = "check"
const TypeSeparator = "separator"
const TypeSubmenu   = "submenu"

type MenuEntry struct {
    Label         string          `yaml:"label"`
    Id            string          `yaml:"id"`
    Cmd           string          `yaml:"cmd"`
    Type          string          `yaml:"type"`
    Submenu    [] MenuEntry       `yaml:"submenu"`
    Handle        MenuEntryHandle `yaml:"-"`
    Parent      * MenuEntry       `yaml:"-"`
    CmdHandler    cmd.CmdHandler  `yaml:"-"`
}

func (me * MenuEntry) SetHandler(h cmd.CmdHandler) {
    me.CmdHandler = h
    for i,_ := range me.Submenu {
        me.Submenu[i].SetHandler(h)
    }
}

func (me * MenuEntry) FindById(id string) (*MenuEntry) {
    if me.Id == id {
        return me
    }
    for i,_ := range me.Submenu {
        e := me.Submenu[i].FindById(id)
        if e != nil {
            return e
        }
    }
    return nil
}

func (me * MenuEntry) SetChecked(id string, ch bool) bool {
    ent := me.FindById(id)
    if ent != nil {
        return ent.Handle.SetChecked(ch)
    }
    return false
}

func (me * MenuEntry) CreateEntries() {
    for idx,_ := range me.Submenu {
        ent := &me.Submenu[idx]
        ent.Parent = me
        if len(ent.Submenu) != 0 || ent.Type == TypeSubmenu {
            me.Handle.CreateEntrySubmenu(ent)
            ent.CreateEntries()
        } else if (ent.Id == SeparatorId || ent.Type == TypeSeparator) {
            me.Handle.CreateEntrySeparator(ent)
        } else if ent.Type == TypeCheck {
            me.Handle.CreateEntryCheck(ent)
        } else {
            me.Handle.CreateEntrySimple(ent)
        }
    }
}

func (me * MenuEntry) Activate() {
    me.CmdHandler.HandleCmd(cmd.Split(me.Cmd))
}
