package game

import (
    "time"
    "log"
    "github.com/metux/freecity/util/cmd"
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

func (g * Game) HandleCmd(c cmd.Cmdline, id string) bool {
    log.Println("Game cmd:", c)
    switch c.Str(0) {
        case "": return true
        case "speed": {
            g.SetSpeed(c.Int(1))
            return true
        }
        case "terrain":
            return g.Terrain.HandleCmd(c[1:], id)
        default:
            log.Println("Game: unknown command: ", c, id)
            return false
    }
}

func (g * Game) FindBuildingType(bt string) * rules.BuildingType {
    return g.Terrain.GeneralRules.FindBuildingType(bt)
}
