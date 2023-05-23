package geo

func Fmin(x float64, y float64) float64 {
    if x < y {
        return x
    }
    return y
}

func Fmax(x float64, y float64) float64 {
    if x > y {
        return x
    }
    return y
}
