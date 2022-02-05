```go

Brush
Color
Crayon
Dye
Shade
Strobe

// Prepare is a convenience method returns an enabled Color unless
// EnableConsole() returns an error or startEnabled == false
func Prepare(startEnabled bool) Color {
    err := EnableConsole()
    if err != nil || startEnabled == false {
        return Color{}
    }
    col := Color
    col.Enable()
    return col
}

color.Enable()
color.Disable()
color.With()
```
