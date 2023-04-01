package base

type Action byte

const (
    ActionAny            = Action(0)
    ActionGameLoad       = iota
    ActionGameCreate     = iota
    ActionGameSave       = iota
    ActionErrectBuilding = iota
    ActionBuildRoad      = iota
    ActionBuildRail      = iota
    ActionBuildPowerline = iota
    ActionClearLand      = iota
    ActionZoning         = iota
    ActionSimu           = iota
)

func (act Action) String() string {
    switch (act) {
        case ActionGameLoad:       return "game-load"
        case ActionGameCreate:     return "game-create"
        case ActionGameSave:       return "game-save"
        case ActionErrectBuilding: return "errect-builing"
        case ActionBuildRoad:      return "build-road"
        case ActionBuildRail:      return "build-rail"
        case ActionBuildPowerline: return "build-powerline"
        case ActionClearLand:      return "clear-land"
        case ActionZoning:         return "zoning"
        case ActionSimu:           return "simu"
    }
    return "(anything)"
}
