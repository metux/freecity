package game

import "fmt"

type NotifyGameLoad struct {
    Filename string
}

func (n NotifyGameLoad) String() string {
    return fmt.Sprintf("game loaded %s", n.Filename)
}

type NotifyGameSave struct {
    Filename string
}

func (n NotifyGameSave) String() string {
    return fmt.Sprintf("game saved %s", n.Filename)
}

type NotifyGameCreate struct {
    Ruleset string
}

func (n NotifyGameCreate) String() string {
    return fmt.Sprintf("game created ruleset %s", n.Ruleset)
}

type NotifyGameSpeed struct {
    Speed int
}

func (n NotifyGameSpeed) String() string {
    return fmt.Sprintf("game speed %d", n.Speed)
}
