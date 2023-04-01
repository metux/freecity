package theme

import (
    "os"
    "log"
    "image"
    "image/png"
    "image/draw"
)

func loadPNG(fn string) * image.RGBA {
    infile, err := os.Open(fn)
    if err != nil {
        log.Println("error opening", fn, err)
        return nil
    }
    defer infile.Close()
    img, err := png.Decode(infile)
    if err != nil {
        log.Println("error decoding", fn, err)
        return nil
    }

    b := img.Bounds()
    m := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
    draw.Draw(m, m.Bounds(), img, b.Min, draw.Src)
    return m
}

func (t * ThemeSpec) loadPixmap(group string, name string) * image.RGBA {
    return loadPNG(t.ImgPath(group+"/"+name))
}
