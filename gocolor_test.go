package gocolor_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bbkane/gocolor"
)

// NOTE: these must be manually inspected to ensure they pass

// TestPrepare demos the Prepare convienience method to get colors
func TestPrepare(t *testing.T) {
	color := gocolor.Prepare(true)

	fmt.Println(
		color.Add(color.FgRed, "FgRed"),
		color.Add(color.FgCyanBright+color.Negative, "FgCyanBright+Negative"),
	)
}

// TestManualCreation demos the less magic way of getting colors. This way let's you test for errors explicitly
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

// TestWithReflection demoes all colors - though not combined with each other :)
func TestWithReflection(t *testing.T) {
	color := gocolor.Prepare(true)

	e := reflect.ValueOf(&color).Elem()

	for i := 0; i < e.NumField(); i++ {
		name := e.Type().Field(i).Name
		val := e.Field(i).Interface().(string)
		fmt.Print(color.Add(val, name))
		fmt.Print(" ")
		if (i+1)%6 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()

}
