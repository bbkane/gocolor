// Code generated by scripts/color_gen.py. DO NOT EDIT.
package gocolor

// Color contains all colors listed in https://docs.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences?redirectedfrom=MSDN#text-formatting
type Color struct {

	// BgBlack - Applies non-bold/bright black to background
	BgBlack Code

	// BgBlackBright - Applies bold/bright black to background
	BgBlackBright Code

	// BgBlue - Applies non-bold/bright blue to background
	BgBlue Code

	// BgBlueBright - Applies bold/bright blue to background
	BgBlueBright Code

	// BgCyan - Applies non-bold/bright cyan to background
	BgCyan Code

	// BgCyanBright - Applies bold/bright cyan to background
	BgCyanBright Code

	// BgDefault - Applies only the background portion of the defaults (see 0)
	BgDefault Code

	// BgExtended - Applies extended color value to the background (see details below)
	BgExtended Code

	// BgGreen - Applies non-bold/bright green to background
	BgGreen Code

	// BgGreenBright - Applies bold/bright green to background
	BgGreenBright Code

	// BgMagenta - Applies non-bold/bright magenta to background
	BgMagenta Code

	// BgMagentaBright - Applies bold/bright magenta to background
	BgMagentaBright Code

	// BgRed - Applies non-bold/bright red to background
	BgRed Code

	// BgRedBright - Applies bold/bright red to background
	BgRedBright Code

	// BgWhite - Applies non-bold/bright white to background
	BgWhite Code

	// BgWhiteBright - Applies bold/bright white to background
	BgWhiteBright Code

	// BgYellow - Applies non-bold/bright yellow to background
	BgYellow Code

	// BgYellowBright - Applies bold/bright yellow to background
	BgYellowBright Code

	// Bold - Applies brightness/intensity flag to foreground color
	Bold Code

	// Default - Returns all attributes to the default state prior to modification
	Default Code

	// FgBlack - Applies non-bold/bright black to foreground
	FgBlack Code

	// FgBlackBright - Applies bold/bright black to foreground
	FgBlackBright Code

	// FgBlue - Applies non-bold/bright blue to foreground
	FgBlue Code

	// FgBlueBright - Applies bold/bright blue to foreground
	FgBlueBright Code

	// FgCyan - Applies non-bold/bright cyan to foreground
	FgCyan Code

	// FgCyanBright - Applies bold/bright cyan to foreground
	FgCyanBright Code

	// FgDefault - Applies only the foreground portion of the defaults (see 0)
	FgDefault Code

	// FgExtended - Applies extended color value to the foreground (see details below)
	FgExtended Code

	// FgGreen - Applies non-bold/bright green to foreground
	FgGreen Code

	// FgGreenBright - Applies bold/bright green to foreground
	FgGreenBright Code

	// FgMagenta - Applies non-bold/bright magenta to foreground
	FgMagenta Code

	// FgMagentaBright - Applies bold/bright magenta to foreground
	FgMagentaBright Code

	// FgRed - Applies non-bold/bright red to foreground
	FgRed Code

	// FgRedBright - Applies bold/bright red to foreground
	FgRedBright Code

	// FgWhite - Applies non-bold/bright white to foreground
	FgWhite Code

	// FgWhiteBright - Applies bold/bright white to foreground
	FgWhiteBright Code

	// FgYellow - Applies non-bold/bright yellow to foreground
	FgYellow Code

	// FgYellowBright - Applies bold/bright yellow to foreground
	FgYellowBright Code

	// Negative - Swaps foreground and background colors
	Negative Code

	// Nobold - Removes brightness/intensity flag from foreground color
	Nobold Code

	// Nounderline - Removes underline
	Nounderline Code

	// Positive - Returns foreground/background to normal
	Positive Code

	// Underline - Adds underline
	Underline Code

}

// DisableAll sets all colors to the empty string, effectively disabling color output
func (c *Color) DisableAll() {
	c.BgBlack = empty
	c.BgBlackBright = empty
	c.BgBlue = empty
	c.BgBlueBright = empty
	c.BgCyan = empty
	c.BgCyanBright = empty
	c.BgDefault = empty
	c.BgExtended = empty
	c.BgGreen = empty
	c.BgGreenBright = empty
	c.BgMagenta = empty
	c.BgMagentaBright = empty
	c.BgRed = empty
	c.BgRedBright = empty
	c.BgWhite = empty
	c.BgWhiteBright = empty
	c.BgYellow = empty
	c.BgYellowBright = empty
	c.Bold = empty
	c.Default = empty
	c.FgBlack = empty
	c.FgBlackBright = empty
	c.FgBlue = empty
	c.FgBlueBright = empty
	c.FgCyan = empty
	c.FgCyanBright = empty
	c.FgDefault = empty
	c.FgExtended = empty
	c.FgGreen = empty
	c.FgGreenBright = empty
	c.FgMagenta = empty
	c.FgMagentaBright = empty
	c.FgRed = empty
	c.FgRedBright = empty
	c.FgWhite = empty
	c.FgWhiteBright = empty
	c.FgYellow = empty
	c.FgYellowBright = empty
	c.Negative = empty
	c.Nobold = empty
	c.Nounderline = empty
	c.Positive = empty
	c.Underline = empty
}

// EnableAll sets all colors to their color codes
func (c *Color) EnableAll() {
	c.BgBlack = Code("\033[40m")
	c.BgBlackBright = Code("\033[100m")
	c.BgBlue = Code("\033[44m")
	c.BgBlueBright = Code("\033[104m")
	c.BgCyan = Code("\033[46m")
	c.BgCyanBright = Code("\033[106m")
	c.BgDefault = Code("\033[49m")
	c.BgExtended = Code("\033[48m")
	c.BgGreen = Code("\033[42m")
	c.BgGreenBright = Code("\033[102m")
	c.BgMagenta = Code("\033[45m")
	c.BgMagentaBright = Code("\033[105m")
	c.BgRed = Code("\033[41m")
	c.BgRedBright = Code("\033[101m")
	c.BgWhite = Code("\033[47m")
	c.BgWhiteBright = Code("\033[107m")
	c.BgYellow = Code("\033[43m")
	c.BgYellowBright = Code("\033[103m")
	c.Bold = Code("\033[1m")
	c.Default = Code("\033[0m")
	c.FgBlack = Code("\033[30m")
	c.FgBlackBright = Code("\033[90m")
	c.FgBlue = Code("\033[34m")
	c.FgBlueBright = Code("\033[94m")
	c.FgCyan = Code("\033[36m")
	c.FgCyanBright = Code("\033[96m")
	c.FgDefault = Code("\033[39m")
	c.FgExtended = Code("\033[38m")
	c.FgGreen = Code("\033[32m")
	c.FgGreenBright = Code("\033[92m")
	c.FgMagenta = Code("\033[35m")
	c.FgMagentaBright = Code("\033[95m")
	c.FgRed = Code("\033[31m")
	c.FgRedBright = Code("\033[91m")
	c.FgWhite = Code("\033[37m")
	c.FgWhiteBright = Code("\033[97m")
	c.FgYellow = Code("\033[33m")
	c.FgYellowBright = Code("\033[93m")
	c.Negative = Code("\033[7m")
	c.Nobold = Code("\033[22m")
	c.Nounderline = Code("\033[24m")
	c.Positive = Code("\033[27m")
	c.Underline = Code("\033[4m")
}

