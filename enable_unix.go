//go:build !windows
// +build !windows

package gocolor

func enableConsole() error {
	return nil
}
