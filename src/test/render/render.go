package render

import (
    "github.com/metux/freecity/core/game"
    "github.com/metux/freecity/render/simple"
    "github.com/metux/freecity/render/theme"
    "github.com/metux/freecity/test/games"
)

func CreateRenderSimple(g1 * game.Game, th string) simple.Renderer {
    t := theme.ThemeSpec{}
    t.LoadYaml("../data/themes/" + th)

    renderer := simple.Renderer{
        Terrain: &g1.Terrain,
        Theme:   &t,
    }

    renderer.RenderTerrain()
    renderer.SavePNG(games.ImageFile)

    return renderer
}
