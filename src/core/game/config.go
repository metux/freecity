package game

type Config struct {
    RulesDir       string
    SavegameDir    string
    Ruleset        string
    TickDelay      int
}

var DefaultConfig = Config{
    RulesDir:       "../data/rules",
    SavegameDir:    "..",
    Ruleset:        "default",
    TickDelay:      100,
}

func (cf Config) CreateGame() * Game {
    game := Game{Config: cf}
    game.CreateGame()
    return &game
}

func (cf Config) LoadGame(filename string) * Game {
    game := Game{Config: cf}
    game.LoadGame(filename)
    return &game
}
