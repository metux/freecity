package theme

func CreateTheme(dn string) (*ThemeSpec) {
    t := ThemeSpec{}
    t.LoadYaml(dn)
    return &t
}
