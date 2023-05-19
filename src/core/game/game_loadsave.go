package game

import (
    "log"
    "github.com/metux/freecity/core/base"
    "github.com/metux/freecity/core/items"
)

type SaveGame struct {
    Ruleset      string             `yaml:"ruleset"`
    Buildings [] items.Building     `yaml:"buildings"`
    MapSize      base.Point         `yaml:"mapsize"`
    Funds        base.Money         `yaml:"funds"`
    Date         base.Date          `yaml:"date"`
    Tiles     [] items.Tile         `yaml:"tiles"`
}

func (this * Game) SaveGame(filename string) error {
    savegame := SaveGame {
        Ruleset:   this.Config.Ruleset,
        MapSize:   this.Terrain.Size,
        Tiles:     this.Terrain.Tiles,
        Buildings: make([]items.Building, len(this.Terrain.Buildings)),
        Funds:     this.Terrain.Funds,
        Date:      this.Terrain.Date,
    }

    for idx,val := range this.Terrain.Buildings {
        if val == nil {
            log.Println("WARN: idx", idx, "is NIL")
        } else {
            savegame.Buildings[idx] = *val
        }
    }

    return base.YamlStore(filename, savegame)
}

func (g *Game) LoadGame(filename string) error {
    savegame := SaveGame {}

    if err := base.YamlLoad(filename, &savegame); err != nil {
        log.Println("failed loading file", filename)
        return err
    }

    // load ruleset
    g.InitGame(g.Config.RulesDir, savegame.Ruleset)
    g.Notify.NotifyEmit(base.ActionGameLoad, NotifyGameLoad{Filename: filename})

    g.Config.Ruleset = savegame.Ruleset
    g.Terrain.Size   = savegame.MapSize
    g.Terrain.Tiles  = savegame.Tiles
    g.Terrain.Funds  = savegame.Funds
    g.Terrain.Date   = savegame.Date

    // load the buildings
    g.Terrain.Buildings = make([]*items.Building, len(savegame.Buildings))
    for idx,_ := range savegame.Buildings {
        g.Terrain.Buildings[idx] = &savegame.Buildings[idx]
    }

    g.Terrain.Update(base.ActionGameLoad)
    return nil
}
