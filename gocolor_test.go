package gocolor_test

import (
	"fmt"
	"reflect"
	"testing"

	"go.bbkane.com/gocolor"
)

// NOTE: these "tests" must be manually inspected to ensure they pass.
// I would love to use example tests ( https://go.dev/blog/examples ) but
// those seem to swallow stdout, which I need to see

func printTestName(t *testing.T, c *gocolor.Color) {
	fmt.Println()

	fmt.Println(
		c.Add(
			c.Underline+c.Bold+c.FgWhiteBright,
			"# "+t.Name(),
		),
	)
	fmt.Println()
}

// TestPrepare demos the Prepare convienience method to get colors
func TestPrepare(t *testing.T) {
	color, err := gocolor.Prepare(true)
	if err != nil {
		t.Fatal(err)
	}

	printTestName(t, &color)

	fmt.Println(
		color.Add(color.FgRed, "FgRed"),
		color.Add(color.FgCyanBright+color.Negative, "FgCyanBright+Negative"),
	)
	fmt.Println()
}

// TestManualCreation demos creating a Color object without using Prepare
func TestManualCreation(t *testing.T) {
	err := gocolor.EnableConsole()
	if err != nil {
		t.Fatal(err)
	}
	//nolint:exhaustruct  // This is too big to reasonably initialize
	color := gocolor.Color{}
	color.EnableAll()

	printTestName(t, &color)

	fmt.Println(
		color.Add(color.FgYellowBright, "Hello!"),
		color.Add(color.BgBlue+color.FgRed, "World!"),
	)
	fmt.Println()
}

// TestWithReflection demoes all colors - though not combined with each other :)
func TestWithReflection(t *testing.T) {
	color, err := gocolor.Prepare(true)
	if err != nil {
		t.Fatal(err)
	}

	printTestName(t, &color)

	e := reflect.ValueOf(&color).Elem()

	for i := 0; i < e.NumField(); i++ {
		name := e.Type().Field(i).Name
		val := e.Field(i).Interface().(gocolor.Code)
		fmt.Print(color.Add(val, name))
		fmt.Print(" ")
		if (i+1)%6 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func TestFunc(t *testing.T) {
	color, err := gocolor.Prepare(true)
	if err != nil {
		t.Fatal(err)
	}

	printTestName(t, &color)

	green := color.Func(&color.FgGreenBright)
	fmt.Println(green("Enabled!"))

	color.DisableAll()
	fmt.Println(green("Disabled!"))

	fmt.Println()
}
