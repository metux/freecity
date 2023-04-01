package items

import (
    "github.com/metux/freecity/core/base"
)

func (tm * TerrainMap) isPowerAt(pos Point) bool {
    if tile := tm.tileAt(pos); tile != nil {
        return tile.HasPowerLine()
    }
    return false
}

func (tm * TerrainMap) updatePowerlineAt(pos Point) {
    tile := tm.tileAt(pos)

    // cant use HasPowerLine() here
    if (tile == nil) || (tile.Power.None()) {
        return
    }

    // FIXME: need to check for conflicts against roads
    tile.Power = base.LineDirectionFromVec(
        tm.isPowerAt(pos.North()),
        tm.isPowerAt(pos.East()),
        tm.isPowerAt(pos.South()),
        tm.isPowerAt(pos.West()))
}

func (tm * TerrainMap) ErrectPowerline(pos Point) (bool) {
    tile := tm.tileAt(pos)

    if tile == nil {
        tm.emit(ActionBuildPowerline, NotifyNoSuchTile{"powerline", pos})
        return false
    }

    if tile.Building != nil {
        tm.emit(ActionBuildPowerline, NotifyAlreadyOccupied{"building "+tile.Building.TypeName, pos})
        return false
    }

    // FIXME: check terrain
    other := base.LineDirPick(tile.Road, tile.Rail)
    if other.None() {
        tm.emit(ActionBuildPowerline, NotifyAlreadyOccupied{"road/rail", pos})
        return false
    }

    tm.autoBulldoze(ActionBuildPowerline, pos)

    if ! tm.trySpendFunds(ActionBuildPowerline, tm.GeneralRules.Costs.Powerline, "powerline") {
        return false
    }

    tile.Power = other
    tm.updatePowerlineAt(pos)
    tm.updatePowerlineAt(pos.North())
    tm.updatePowerlineAt(pos.East())
    tm.updatePowerlineAt(pos.South())
    tm.updatePowerlineAt(pos.West())

    tm.CalcPowerGrid()
    tm.TouchObjects()
    return true
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

func (t * TerrainMap) ErrectPowerlineH(pos Point, w int) {
    for i := 0; i<w; i++ {
        t.ErrectPowerline(pos)
        pos.X++
    }
}

func (t * TerrainMap) ErrectPowerlineV(pos Point, w int) {
    for i := 0; i<w; i++ {
        t.ErrectPowerline(pos)
        pos.Y++
    }
}
