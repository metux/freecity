package tools

import (
)

type Rubble struct {
}

func (t * Rubble) GetName() string {
    return "Rubble"
}

func (t * Rubble) WorkAt(game * Game, p point) {
    game.Terrain.PlaceRubble(p)
}
