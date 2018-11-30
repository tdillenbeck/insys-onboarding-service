package color

import (
	"fmt"
	"strconv"
)

type Color uint8

// Foreground text colors
const (
	FgBlack Color = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity text colors
const (
	FgHiBlack Color = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

func SprintFunc(color Color) func(...interface{}) string {
	return func(a ...interface{}) string {
		return wrap(color, fmt.Sprint(a...))
	}
}

var enabled = true

func wrap(color Color, s string) string {
	if !enabled {
		return s
	}

	return "\x1b[" + strconv.Itoa(int(color)) + "m" + s + "\x1b[0m"
}

func DisableColor() {
	enabled = false
}
