package game

import (
    "log"
    "fmt"
    "github.com/metux/freecity/core/items"
)

// Simple notification handler that just prints out to log console
type ConsoleNotify struct {
    NotifyCount int
}

func (cn * ConsoleNotify) notifyfmt(act action, msg string, i ...interface{}) bool {
    log.Printf("game: [%4d] %s: %s", cn.NotifyCount, act.String(), fmt.Sprintf(msg, i...))
    cn.NotifyCount++
    return true
}

func (cn * ConsoleNotify) NotifyEmit(a action, n items.NotifyMsg) bool {
    return cn.notifyfmt(a, n.String())
}
