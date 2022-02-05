package gocolor

// Prepare is a convenience method that calls EnableConsole
// and returns an enabled color if there's no error and startEnabled is True.
func Prepare(startEnabled bool) Color {
	err := EnableConsole()
	if err != nil || !startEnabled {
		return Color{}
	}
	col := Color{}
	col.Enable()
	return col
}

// EnableConsole enables color printing on Windows and is a no-op
// on non-Windows platforms (which enable color printing by default).
func EnableConsole() error {
	return enableConsole()
}

// Add colors a message and resets the color afterwards
func (c *Color) Add(color string, message string) string {
	return color + message + c.Default
}
