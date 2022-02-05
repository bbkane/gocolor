package gocolor_test

import (
	"fmt"
	"testing"

	"github.com/bbkane/gocolor"
)

// NOTE: these must be manually inspected to ensure they pass

func TestPrepare(t *testing.T) {
	color := gocolor.Prepare(true)

	fmt.Println(
		color.Add(color.FgRed, "FgRed"),
		color.Add(color.FgCyanBright+color.Negative, "FgCyanBright+Negative"),
	)
}

func TestManualCreation(t *testing.T) {
	err := gocolor.EnableConsole()
	if err != nil {
		t.Fatal(err)
	}
	color := gocolor.Color{}
	color.Enable()

	fmt.Println(
		color.Add(color.FgYellowBright, "Hello!"),
		color.Add(color.BgBlue+color.FgRed, "World!"),
	)
}
