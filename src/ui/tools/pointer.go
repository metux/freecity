package tools

import (
    "log"
)

type Pointer struct {
}

func (t * Pointer) GetName() string {
    return "Pointer"
}

func (t * Pointer) WorkAt(game * Game, p point) {
    log.Println("FIXME: pointer click not implemented yet")
}
