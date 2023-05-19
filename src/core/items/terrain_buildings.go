package items

import (
    "log"
    "github.com/metux/freecity/core/base"
    "github.com/metux/freecity/core/rules"
)

func (tm * TerrainMap) mayPlaceBuildingType(bt * rules.BuildingType, pos point) bool {
    rect := pos.MakeRect(bt.Size)

    tiles, err := tm.TileRange(rect, false)
    if err != nil {
        log.Println("failed getting tile range")
        return false
    }
    for _,tr := range tiles {
        tm.autoBulldoze(base.ActionErrectBuilding, tr.Position)
        if tr.Tile.Building != nil {
            tm.emit(base.ActionErrectBuilding, NotifyCantPlaceHere{"building "+bt.Name, tr.Position})
            return false
        }
    }

    return tiles.Check(bt.Require)
}

func (tm * TerrainMap) ConnectBuildings() {
    for _,b := range tm.Buildings {
        b.Terrain = tm
        if b.BuildingType == nil {
            b.BuildingType = tm.GeneralRules.FindBuildingType(b.TypeName)
        }
        b.ConnectTiles()
    }
    // FIXME: need to fix neighboring titles --> powerlines
    tm.CalcPowerGrid()
}

func (tm * TerrainMap) ErrectBuilding (typename string, pos point) bool {
    bt := tm.GeneralRules.FindBuildingType(typename)
    if bt == nil {
        tm.emit(base.ActionErrectBuilding, NotifyNoSuchBuildingType{typename})
        return false
    }

    if ! tm.mayPlaceBuildingType(bt, pos) {
        log.Println("cant place building type", typename, "at", pos)
        return false
    }

    if ! tm.trySpendFunds(base.ActionErrectBuilding, bt.Costs.Build, "building type" + typename) {
        return false
    }

    b := NewBuilding(bt, pos, tm)
    tm.Buildings = append(tm.Buildings, b)
    b.ConnectTiles()
    tm.CalcPowerGrid()
    tm.emit(base.ActionErrectBuilding, NotifyBuildingErrected{b})
    return true
}

func (tm * TerrainMap) RemoveBuilding(b * Building) {
    buf := make([]*Building, 0)
    for _,walk := range tm.Buildings {
        if walk != nil && walk != b {
            buf = append(buf, walk)
        }
    }
    tm.Buildings = buf
}

func (tm * TerrainMap) DemolishBuildingAt(pos point) bool {
    tile := tm.tileAt(pos)
    if tile == nil || tile.Building == nil {
        return false
    }

    tm.RemoveBuilding(tile.Building)
    tile.Building.Destroy()
    tile.Building = nil

    // FIXME: remove from tm.Buildings
    return true
}
