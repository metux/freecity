package tools

import (
    "log"
)

type Rubble struct {
}

func (t * Rubble) GetName() string {
    return "Rubble"
}

func (t * Rubble) WorkAt(game * Game, p point) {
    log.Println(" --> placing rubble at", p)
    game.Terrain.PlaceRubble(p)
}
