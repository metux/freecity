package common

import (
    "fmt"
    "github.com/metux/freecity/core/game"
)

func (me * MenuEntry) SetGameSpeed(speed int) {
    for x:=0; x<(game.MaxSpeed+1); x++ {
        id := fmt.Sprintf("speed.%d", x)
        if x == speed {
            me.SetChecked(id, true)
        } else {
            me.SetChecked(id, false)
        }
    }
}
