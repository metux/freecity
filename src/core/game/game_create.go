package game

import (
    "github.com/metux/freecity/core/base"
)

func (g * Game) CreateGame() {
    g.InitGame(g.Config.RulesDir, g.Config.Ruleset)
    g.Notify.NotifyEmit(base.ActionGameCreate, NotifyGameCreate{g.Config.Ruleset})
    g.Terrain.CreateMap()
    g.Terrain.Update(base.ActionGameCreate)
}
