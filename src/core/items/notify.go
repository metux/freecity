package items

import "fmt"

type NotifyMsg interface {
    String() string
}

// funds spent on some action, eg. errecting infrastructure objects like
// buildings, roads, etc.
type NotifyFundsSpent struct {
    Spent Money
    Funds Money
    Cause string
}

func (n NotifyFundsSpent) String() string {
    return fmt.Sprintf("spent %d funds for %s (total %d)", n.Spent, n.Cause, n.Funds)
}

// some action was tried on a non-existing tile, outside of map
type NotifyNoSuchTile struct {
    Cause string
    Pos Point
}

func (n NotifyNoSuchTile) String() string {
    return fmt.Sprintf("point %s out of map: %s", n.Pos, n.Cause)
}

// something could not be placed on the given point
type NotifyAlreadyOccupied struct {
    Obj string
    Pos Point
}

func (n NotifyAlreadyOccupied) String() string {
    return fmt.Sprintf("tile already occupied [%d:%d]: %s", n.Pos.X, n.Pos.Y, n.Obj)
}

// something could not be placed on the given point
type NotifyCantPlaceHere struct {
    Obj string
    Pos Point
}

func (n NotifyCantPlaceHere) String() string {
    return fmt.Sprintf("cant place here [%d:%d]: %s", n.Pos.X, n.Pos.Y, n.Obj)
}

// trying to errect a building failed due unknown type
type NotifyNoSuchBuildingType struct {
    BuildingType string
}

func (n NotifyNoSuchBuildingType) String() string {
    return fmt.Sprintf("no such building type: %s", n.BuildingType)
}

// signals that some action failed due lack of funds
type NotifyNotEnoughFunds struct {
    Needed Money
    Funds Money
    Cause string
}

func (n NotifyNotEnoughFunds) String() string {
    return fmt.Sprintf("not enough funds %d (total=%d) for %s", n.Needed, n.Funds, n.Cause)
}

// building has been errected successfully
type NotifyBuildingErrected struct {
    Building * Building
}

func (n NotifyBuildingErrected) String() string {
    return fmt.Sprintf("building errected type=%s", n.Building.TypeName)
}

// signals that a building got power or doesn't. called for every building
// on TerrainMap.CheckPower(). this even't can happen multiple times for
// each buiding, even if its power state didn't actually change
// call Terrain.CheckPower() to get a notify for each building
type NotifyBuildingPowered struct {
    Building * Building
}

func (n NotifyBuildingPowered) String() string {
    if n.Building.Powered {
        return fmt.Sprintf("building %s powered", n.Building.TypeName)
    } else {
        return fmt.Sprintf("building %s has no power", n.Building.TypeName)
    }
}

// notification interface for reporting game events back to upper layer code
// like user interface
type NotifyHandler interface {
    NotifyEmit(a Action, n NotifyMsg) bool
}
