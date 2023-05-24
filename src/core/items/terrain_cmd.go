package items

import (
    "log"
    "strconv"
)

func (tm * TerrainMap) PlaceAt(p point, cmd [] string) bool {
    switch cmd[0] {
        case "road":      return tm.ErrectLine(LtRoad,  p)
        case "rail":      return tm.ErrectLine(LtRail,  p)
        case "powerline": return tm.ErrectLine(LtPower, p)
        case "rubble":    return tm.PlaceRubble(p)
    }
    return false
}

func (tm * TerrainMap) handleErrect(cmd[] string, id string) bool {
    x,_ := strconv.Atoi(cmd[0])
    y,_ := strconv.Atoi(cmd[1])
    return tm.PlaceAt(point{x,y}, cmd[2:])
}

func (tm * TerrainMap) HandleCmd(cmd [] string, id string) bool {
    switch cmd[0] {
        case "place": return tm.handleErrect(cmd[1:], id)
    }
    log.Println("terrain: unhandled command:", cmd)
    return false
}
