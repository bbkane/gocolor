# gocolor

A tiny cross-platform cross-platform terminal color library that supports enabling and disabling colors.

Uses colors from https://docs.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences?redirectedfrom=MSDN#text-formatting

## Usage

See [tests](./gocolor_test.go). Also see [mattn/go-isatty](https://github.com/mattn/go-isatty) to test if your program is working with a tty before you enable color.

![TestWithReflection.png](TestWithReflection.png)

## Inspirations

- https://github.com/TwiN/go-color (API inspiration)
- https://github.com/bbkane/go-color (my - now deprecated - fork of TwiN/go-color)
- https://github.com/fatih/color (color naming)
- https://github.com/jedib0t/go-pretty (Windows code)

