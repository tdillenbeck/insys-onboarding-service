package config

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strings"
	"unicode/utf8"

	"weavelab.xyz/wlib/color"
)

// Print calls the DefaultConfig.Print
func Print() {
	DefaultConfig.Print()
}

func PrettyPrint() {
	DefaultConfig.PrettyPrint()
}

// PrettyPrint shows all the env variables and their values in nice color and a box
func (c *Config) PrettyPrint() {
	width := terminalWidth()
	//range to get slice of keys
	var keys []string
	for k := range c.Settings {
		keys = append(keys, k)
	}
	//sort slice of keys so we can print the values alphbetical
	sort.Strings(keys)

	envs := formatEnv(keys, c.Settings)

	maxLength := maxLineLength(envs, width)

	leftPadding := (width - maxLength) / 2
	if leftPadding < 0 {
		leftPadding = 0
	}

	w := bytes.NewBuffer(nil)

	printSpace(w)
	printWeaveStamp(w, width)

	printSpace(w)

	printLineOfStars(w, leftPadding, maxLength)
	printHeader(w, leftPadding, maxLength)
	printLineOfStars(w, leftPadding, maxLength)

	printEnvVars(w, envs, leftPadding, maxLength)
	printLineOfStars(w, leftPadding, maxLength)
	printLineOfStars(w, leftPadding, maxLength)
	printSpace(w)

	// there's a better way...
	_, err := io.Copy(os.Stdout, w)
	if err != nil {
		fmt.Printf("unable to copy to std out: %s\n", err)
	}
}

// Print shows each config setting on a new line
func (c *Config) Print() {

	w := bytes.NewBuffer(nil)
	for _, v := range c.Settings {
		fmt.Fprintf(w, "%#v\n\n", v)
	}

	_, err := io.Copy(os.Stdout, w)
	if err != nil {
		fmt.Printf("unable to copy to std out: %s\n", err)
	}
}

func logo() []string {
	return []string{
		"oooooooooooooooooooooooooooooooooooooooooooooooooo",
		"oooooooooooooooooooooooooooooooooooooooooooooooooo",
		"oooooooooooooooooooo+///::///+oooooooooooooooooooo",
		"ooooooooooooooo+:-`   ``````   `-:+ooooooooooooooo",
		"oooooooooooo+-` `-:/+oooooooo+/:-` `-+oooooooooooo",
		"oooooooooo+. `-+oooooooooooooooooo+-` .+oooooooooo",
		"oooooooo+- `:+oooooooooooooooooooooo+:` -+oooooooo",
		"ooooooo+` -+oooooooooooooooooooooooooo+- `+ooooooo",
		"oooooo+` :oooooooooooooooooooooooooooooo: `+oooooo",
		"oooooo. -oooooo--+ooooo+--+ooooo+--oooooo- .oooooo",
		"ooooo/  +ooooo+  :ooooo/  /ooooo:  +ooooo+  /ooooo",
		"ooooo- .oooooo+  :ooooo/  /ooooo:  +oooooo. -ooooo",
		"ooooo. -oooooo+  :ooooo/  /ooooo:  +oooooo- .ooooo",
		"ooooo- .oooooo+  :ooooo/  /ooooo:  +oooooo. -ooooo",
		"ooooo/  +oooooo` ./////-  -/////. `oooooo+  /ooooo",
		"oooooo. -oooooo+-`      ..      `-+oooooo- .oooooo",
		"oooooo+` :oooooooo+++++oooo+++++oooooooo: `+oooooo",
		"ooooooo+` -+oooooooooooooooooooooooooo+- `+ooooooo",
		"oooooooo+- `:+oooooooooooooooooooooo+:` -+oooooooo",
		"oooooooooo+. `-+oooooooooooooooooo+-` .+oooooooooo",
		"oooooooooooo+-` `-:/+oooooooo+/:-` `-+oooooooooooo",
		"ooooooooooooooo+:-`   ``````   `-:+ooooooooooooooo",
		"oooooooooooooooooooo+///::///+oooooooooooooooooooo",
		"oooooooooooooooooooooooooooooooooooooooooooooooooo",
		"oooooooooooooooooooooooooooooooooooooooooooooooooo",
	}
}

func logoWidth() int {
	return utf8.RuneCountInString(logo()[0])
}

func printWeaveStamp(w io.Writer, width int) {
	weave := color.SprintFunc(color.FgCyan)

	wl := logo()

	l := logoWidth()
	p := (width - l) / 2
	if p < 0 {
		p = 1
	}

	padding := strings.Repeat(" ", p)

	for _, v := range wl {
		fmt.Fprintf(w, "%s%s\n", padding, weave(v))
	}
}

const (
	borderLeft  = "**** "
	borderRight = " ****"
)

func borderWidth() int {
	return utf8.RuneCountInString(borderLeft) + utf8.RuneCountInString(borderRight)
}

func printEnvVars(w io.Writer, envs []env, leftPadding int, maxLength int) {

	padding := strings.Repeat(" ", leftPadding)

	ls := borderLeft
	rs := borderRight

	d := maxLength - utf8.RuneCountInString(ls) - utf8.RuneCountInString(rs)
	if d < 0 {
		d = 80
	}

	for _, v := range envs {

		space := d - v.w
		if space < 0 {
			space = 1
		}

		spaces := strings.Repeat(" ", space)

		fmt.Fprintf(w, "%s%s%s%s%s\n", padding, ls, v.s, spaces, rs)
	}
}

type env struct {
	s string
	w int
}

func formatEnv(keys []string, settings map[string]*Setting) []env {

	key := color.SprintFunc(color.FgHiBlue)
	value := color.SprintFunc(color.FgHiGreen)

	envs := make([]env, len(keys))
	for i, k := range keys {

		a := settings[k].Env
		b := settings[k].value()

		s := fmt.Sprintf("%s=%s", key(a), value(b))

		w := utf8.RuneCountInString(a) + 1 + utf8.RuneCountInString(b)

		envs[i] = env{
			s: s,
			w: w,
		}
	}

	return envs
}

func printSpace(w io.Writer) {
	fmt.Fprint(w, "\n\n")
}

func printLineOfStars(w io.Writer, leftPadding int, width int) {
	p := strings.Repeat(" ", leftPadding)
	s := strings.Repeat("*", width)
	fmt.Fprintf(w, "%s%s\n", p, s)
}

func maxLineLength(envs []env, width int) int {
	maxLength := 0
	b := borderWidth()

	for _, v := range envs {
		w := v.w + b

		if w > maxLength {
			maxLength = w
		}

		if maxLength > width {
			break
		}
	}

	if maxLength > width {
		return width
	}

	lw := logoWidth()
	if maxLength < lw {
		return lw
	}

	return maxLength
}

func printHeader(w io.Writer, leftPadding int, maxLength int) {

	padding := strings.Repeat(" ", leftPadding)

	text := " Config Vars "

	// how many stars do we need?
	c := maxLength - utf8.RuneCountInString(text)

	// split the stars between the left half and the right
	rc := c / 2
	lc := int(math.Ceil(float64(c) / 2))

	title := color.SprintFunc(color.FgHiRed)
	t := title(text)

	ls := strings.Repeat("*", lc)
	rs := strings.Repeat("*", rc)

	fmt.Fprintf(w, "%s%s%s%s\n", padding, ls, t, rs)

}

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}
