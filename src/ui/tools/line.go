package tools

// build roads
type Road struct {
}

func (t * Road) GetName() string {
    return "Road"
}

func (t * Road) WorkAt(game * Game, p point) {
//    game.Terrain.PlaceRubble(p)
}

func (t * Road) GetMenuId() string {
    return "road"
}

// build rails
type Rail struct {
}

func (t * Rail) GetName() string {
    return "Road"
}

func (t * Rail) WorkAt(game * Game, p point) {
//    game.Terrain.PlaceRubble(p)
}

func (t * Rail) GetMenuId() string {
    return "road"
}

// build power lines
type Power struct {
}

func (t * Power) GetName() string {
    return "Powerline"
}

func (t * Power) WorkAt(game * Game, p point) {
//    game.Terrain.PlaceRubble(p)
}

func (t * Power) GetMenuId() string {
    return "powerline"
}

// build pipes
type Pipe struct {
}

func (t * Pipe) GetName() string {
    return "Pipe"
}

func (t * Pipe) WorkAt(game * Game, p point) {
//    game.Terrain.PlaceRubble(p)
}

func (t * Pipe) GetMenuId() string {
    return "pipe"
}
