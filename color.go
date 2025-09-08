// Package color is an ANSI color package to output colorized or SGR defined
// outputs to the standard output.
package chalk

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Color struct {
	params []Parameter
}

const escape = "\x1b" 

type Parameter int

const (
	Reset Parameter = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)
const (
	FgBlack Parameter = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

const (
	BgBlack Parameter = iota + 40
	BgRed
    BgGreen
    BgYellow
    BgBlue
    BgMagenta
    BgCyan
    BgWhite
)


var (
	Black   = &Color{params: []Parameter{FgBlack}}
	Green   = &Color{params: []Parameter{FgGreen}}
	Yellow  = &Color{params: []Parameter{FgYellow}}
	Blue    = &Color{params: []Parameter{FgBlue}}
	Magenta = &Color{params: []Parameter{FgMagenta}}
	Cyan    = &Color{params: []Parameter{FgCyan}}
	White   = &Color{params: []Parameter{FgWhite}}
)

var Output io.Writer = os.Stdout

// Red is an convenient helper function to print with red foreground.
func Red(format string, a ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	c := &Color{params: []Parameter{FgRed}}
	c.Printf(format, a...)

}

// New returns a newly created color object.
func New(value ...Parameter) *Color {
	c := &Color{params: make([]Parameter, 0)}
	c.Add(value...)
	return c
} 

func (c *Color) Bold() *Color {
	c.Add(Bold)
	return c
}

// Add is used to chain SGR parameters. Use as many as paramters to combine
// and create custom color objects. Example: Add(color.FgRed, color.Underline)
func (c *Color) Add(value ...Parameter) *Color {
	c.params = append(c.params, value...)
	return c
}

func (c *Color) prepend(value Parameter) {
	c.params = append(c.params, 0)
	copy(c.params[1:], c.params[0:])
	c.params[0] = value
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (c *Color) Printf(format string, a ...interface{}) (n int, err error) {
	c.Set()
	defer Unset()

	return fmt.Fprintf(Output, format, a...)
}

// Print formats using the default formats for its operands and writes to
// standard output. Spaces are added between operands when neither is a
// string. It returns the number of bytes written and any write error
// encountered.
func (c *Color) Print(a ...interface{}) (n int, err error) {
	c.Set()
	defer Unset()

	return fmt.Fprint(Output, a...)
}

// Println formats using the default formats for its operands and writes to
// standard output. Spaces are always added between operands and a newline is
// appended. It returns the number of bytes written and any write error
// encountered.
func (c *Color) Println(a ...interface{}) (n int, err error) {
	c.Set() 		// applique "\033[31m"
	defer Unset()	// à la fin, applique "\033[0m"

	return fmt.Fprintln(Output, a...)
}

// sequence returns a formated SGR sequence to be plugged into a "\x1b[...m"
// an example output might be: "1;36" -> bold cyan
func (c *Color) sequence() string {
	format := make([]string, len(c.params))
	for i, v := range c.params {
		format[i] = strconv.Itoa(int(v))
	}

	return strings.Join(format, ";")
}

// Set sets the SGR sequence.
func (c *Color) Set() {
	fmt.Fprintf(Output, "%s[%sm", escape, c.sequence())
}

func Unset() {
	fmt.Fprintf(Output, "%s[%dm", escape, Reset)
}