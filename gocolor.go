package gocolor

// Prepare is a convenience method that calls EnableConsole
// and returns an enabled color if there's no error and startEnabled is True,
// and a disabled color and the error otherwise.
func Prepare(startEnabled bool) (Color, error) {
	err := EnableConsole()
	if err != nil || !startEnabled {
		return Color{}, err
	}
	//nolint:exhaustruct  // This is too big to reasonably initialize
	col := Color{}
	col.EnableAll()
	return col, nil
}

// EnableConsole enables color printing on Windows and is a no-op
// on non-Windows platforms (they enable color printing by default).
func EnableConsole() error {
	return enableConsole()
}

// Add colors a message and resets the console color afterwards
func (c *Color) Add(color Code, message string) string {
	return string(color) + message + string(c.Default)
}

// ColorFunc colors a message and resets the console color afterwards.
// Use Func() to generate a ColorFunc with a specific color
type ColorFunc = func(message string) string

// Func generates a ColorFunc from a color code. It uses the address of the code instead
// of the value because the code's value might change (for example to "" when Disable() is called)
func (c *Color) Func(codePtr *Code) ColorFunc {
	return func(message string) string {
		return c.Add(*codePtr, message)
	}
}

// Code is a color code string
type Code string

// empty represents the absence of a color code. Used for Disable
const empty Code = Code("")
