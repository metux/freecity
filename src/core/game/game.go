package game

import (
    "time"
    "log"
    "strconv"
    "github.com/metux/freecity/core/base"
    "github.com/metux/freecity/core/items"
    "github.com/metux/freecity/core/rules"
    "github.com/metux/freecity/core/simu"
)

const (
    MaxSpeed = 3
)

type Game struct {
    Config      Config
    Terrain     items.TerrainMap
    Notify      items.NotifyHandler
    Simu        simu.Simulator
    Ticker    * time.Ticker
    Speed       int
}

// initialize a game from already loaded data
func (g * Game) InitGame(rulesdir string, ruleset string) {
    if (g.Notify == nil) {
        g.Notify = new(ConsoleNotify)
    }

    g.Terrain.Init(rulesdir, ruleset, g)
    g.Simu.Init(&g.Terrain, g)
}

func (g * Game) NotifyEmit(a action, msg items.NotifyMsg) bool {
    if g.Notify != nil {
        return g.Notify.NotifyEmit(a, msg)
    }
    return false
}

func (g * Game) SetNotify(nh items.NotifyHandler) items.NotifyHandler {
    old := g.Notify
    g.Notify = nh
    return old
}

// FIXME: move this to simu ?
// FIXME: send game started signal ?
func (g * Game) Start() {
    if g.Ticker == nil {
        g.Ticker = time.NewTicker(time.Duration(g.Config.TickDelay) * time.Millisecond)
        go func() {
            for {
                select {
                    case <-g.Ticker.C:
                        if g.Speed > 0 {
                            g.Simu.Tick()
                        }
                }
            }
        }()
    }
}

func (g * Game) SetSpeed(x int) {
    // FIXME: need to tune timer
    g.Speed = x
    g.Notify.NotifyEmit(base.ActionAny, NotifyGameSpeed{x})
    g.Start()
}

func (g * Game) HandleCmd(cmd [] string, id string) bool {
    switch cmd[0] {
        case "speed": {
            i,_ := strconv.Atoi(cmd[1])
            g.SetSpeed(i)
            return true
        }
        default:
            log.Println("Game: unknown command: ", cmd, id)
            return false
    }
}

func (g * Game) FindBuildingType(bt string) * rules.BuildingType {
    return g.Terrain.GeneralRules.FindBuildingType(bt)
}
