package items

import "log"

func (tm * TerrainMap) handleErrect(cmd[] string, id string) bool {
    log.Println("handleErrect", cmd)
    return false
}

func (tm * TerrainMap) HandleCmd(cmd [] string, id string) bool {
    switch cmd[0] {
        case "errect": return tm.handleErrect(cmd[1:], id)
    }
    log.Println("terrain: unhandled command:", cmd)
    return false
}
