package tools

type Rubble struct {
}

func (t * Rubble) GetName() string {
    return "Rubble"
}

func (t * Rubble) WorkAt(game * Game, p point) {
    game.Terrain.PlaceAt(p, []string{"rubble"})
}

func (t * Rubble) GetMenuId() string {
    return "rubble"
}
