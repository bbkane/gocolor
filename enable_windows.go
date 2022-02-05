package gocolor

import (
	"os"

	"golang.org/x/sys/windows"
)

func enableConsole() error {
	// setting for os.Stdout also seems to affect os.Stderr
	stdoutPtr := os.Stdout.Fd()
	return enable(stdoutPtr)
}

func enable(out uintptr) error {
	// Code adapted from: https://github.com/jedib0t/go-pretty

	// See https://docs.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences?redirectedfrom=MSDN#text-formatting
	// and
	// https://docs.microsoft.com/en-us/windows/console/setconsolemode

	handle := windows.Handle(out)

	var stdoutConsoleMode uint32
	var flags uint32 = windows.ENABLE_PROCESSED_OUTPUT |
		windows.ENABLE_WRAP_AT_EOL_OUTPUT |
		windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING

	err := windows.SetConsoleMode(
		handle,
		stdoutConsoleMode|flags,
	)
	return err
}
