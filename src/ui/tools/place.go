package tools

type PlaceAt struct {
    name   string
    menuid string
    cmd [] string
}

func (t * PlaceAt) GetName() string {
    return t.name
}

func (t * PlaceAt) WorkAt(game * Game, p point) {
    game.Terrain.PlaceAt(p, t.cmd)
}

func (t * PlaceAt) GetMenuId() string {
    return t.menuid
}

func mkPlaceAtSimple(name string, menuid string, cmd string) * PlaceAt {
    return &PlaceAt{
        name: name,
        menuid: menuid,
        cmd: []string{cmd}}
}
