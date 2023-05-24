package items

import (
    "github.com/metux/freecity/core/base"
)

func (tm * TerrainMap) isPowerAt(p point) bool {
    return tm.CheckTileLine(p, LtPower)
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

func (tm * TerrainMap) ErrectPowerline(p point) (bool) {
    return tm.addLine(ActionBuildPowerline, LtPower, p, tm.updatePowerlineAt)
}

func (tm * TerrainMap) addLine(act base.Action, lt LineType, p point, cb func(p point)) (bool) {
    if tile := tm.tileForLine(p, act, lt); tile != nil {
        // FIXME: check terrain
        other := tile.PickLine(lt)
        if other.None() {
            tm.emit(ActionBuildPowerline, NotifyAlreadyOccupied{"lines", p})
            return false
        }

        tm.autoBulldoze(ActionBuildPowerline, p)

        if ! tm.trySpendFunds(ActionBuildPowerline, tm.GeneralRules.LinePrice(lt), lt.String()) {
            return false
        }

        tile.SetLine(lt, other)
        p.DoOnPointAndSurrounding(cb)

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
