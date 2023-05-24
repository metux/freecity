package items

import (
    "github.com/metux/freecity/core/base"
)

func (tm * TerrainMap) isPowerAt(p point) bool {
    return tm.CheckTileLine(p, base.LineTypePower)
}

func (tm * TerrainMap) updatePowerlineAt(p point) {
    tile := tm.tileAt(p)

    // cant use HasPowerLine() here
    if (tile == nil) || (tile.Power.None()) {
        return
    }

    // FIXME: need to check for conflicts against roads
    tile.Power.PickFromSurrounding(p, tm.isPowerAt)
}

func (tm * TerrainMap) tileForLine(p point, action base.Action, lt base.LineType, obj string) * Tile {
    tile := tm.tileAt(p)

    if tile == nil {
        tm.emit(action, NotifyNoSuchTile{obj, p})
        return nil
    }

    if tile.Building != nil {
        tm.emit(ActionBuildPowerline, NotifyAlreadyOccupied{"building "+tile.Building.TypeName, p})
        return nil
    }

    return tile
}

func (tm * TerrainMap) ErrectPowerline(p point) (bool) {
    if tile := tm.tileForLine(p, ActionBuildPowerline, base.LineTypePower, "powerline"); tile != nil {
        // FIXME: check terrain
        other := base.LineDirPick(tile.Road, tile.Rail)
        if other.None() {
            tm.emit(ActionBuildPowerline, NotifyAlreadyOccupied{"road/rail", p})
            return false
        }

        tm.autoBulldoze(ActionBuildPowerline, p)

        if ! tm.trySpendFunds(ActionBuildPowerline, tm.GeneralRules.Costs.Powerline, "powerline") {
            return false
        }

        tile.SetLine(base.LineTypePower, other)
        p.DoOnPointAndSurrounding(tm.updatePowerlineAt)

        tm.CalcPowerGrid()
        tm.TouchObjects()
        return true
    }
    return false
}

func (tm * TerrainMap) CheckPower(act Action) {
    for _,b := range tm.Buildings {
        tm.emit(act, NotifyBuildingPowered{b})
    }
}

func (tm * TerrainMap) CalcPowerGrid() {
    for idx,_ := range tm.Tiles {
        tm.Tiles[idx].ClearPowerGrid()
    }
    for idx,_ := range tm.Buildings {
        tm.Buildings[idx].ClearPowerGrid()
    }

    grids := make([] PowerGrid, 0)

    for idx,b := range tm.Buildings {
        if (!b.PowerConnected()) {
            grids = append(grids, CreatePowerGrid(tm, tm.Buildings[idx]))
        }
    }

    tm.PowerGrids = grids
}
