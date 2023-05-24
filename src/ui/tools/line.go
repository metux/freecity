package tools

import (
    "github.com/metux/freecity/core/base"
)

// build roads
type Road struct {
}

func (t * Road) GetName() string {
    return "Road"
}

func (t * Road) WorkAt(game * Game, p point) {
    game.Terrain.ErrectLine(p, base.LineTypeRoad)
}

func (t * Road) GetMenuId() string {
    return "infra.road"
}

// build rails
type Rail struct {
}

func (t * Rail) GetName() string {
    return "Road"
}

func (t * Rail) WorkAt(game * Game, p point) {
    game.Terrain.ErrectLine(p, base.LineTypeRail)
}

func (t * Rail) GetMenuId() string {
    return "infra.rail"
}

// build power lines
type Power struct {
}

func (t * Power) GetName() string {
    return "Powerline"
}

func (t * Power) WorkAt(game * Game, p point) {
    game.Terrain.ErrectLine(p, base.LineTypePower)
}

func (t * Power) GetMenuId() string {
    return "infra.powerline"
}

// build pipes
type Pipe struct {
}

func (t * Pipe) GetName() string {
    return "Pipe"
}

func (t * Pipe) WorkAt(game * Game, p point) {
    game.Terrain.ErrectLine(p, base.LineTypePipe)
}

func (t * Pipe) GetMenuId() string {
    return "infra.pipe"
}
