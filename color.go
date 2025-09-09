package chalk

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Color defines a custom color object which is defined by SGR attributess.
type Color struct {
	params []Attributes
}

const escape = "\x1b" 

type Attributes int

// Base paramaters
const (
	Reset Attributes = iota
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

// Foreground text colors
const (
	FgBlack Attributes = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Background text colors
const (
	BgBlack Attributes = iota + 40
	BgRed
    BgGreen
    BgYellow
    BgBlue
    BgMagenta
    BgCyan
    BgWhite
)

var (
	Green   = &Color{params: []Attributes{FgGreen}}
	Yellow  = &Color{params: []Attributes{FgYellow}}
	Blue    = &Color{params: []Attributes{FgBlue}}
	Magenta = &Color{params: []Attributes{FgMagenta}}
	Cyan    = &Color{params: []Attributes{FgCyan}}
	White   = &Color{params: []Attributes{FgWhite}}
)

// Output defines the standard output of the print functions. By default
// os.Stdout is used.
var Output io.Writer = os.Stdout

func printColor(format string, p Attributes, a ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	c := &Color{params: []Attributes{p}}
	c.Printf(format, a...)
}

// Red is an convenient helper function to print with red foreground.
func Red(format string, a ...interface{}) { printColor(format, FgRed, a...) }

// Black is an convenient helper function to print with black foreground.
func Black(format string, a ...interface{}) { printColor(format, FgBlack, a...) }

// New returns a newly created color object.
func New(value ...Attributes) *Color {
	c := &Color{params: make([]Attributes, 0)}
	c.Add(value...)
	return c
} 

func (c *Color) Bold() *Color {
	c.Add(Bold)
	return c
}

// Add is used to chain SGR attributess. Use as many as paramters to combine
// and create custom color objects. Example: Add(color.FgRed, color.Underline)
func (c *Color) Add(value ...Attributes) *Color {
	c.params = append(c.params, value...)
	return c
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
// Standard fmt.PrintF() method wrapped with the given color.
func (c *Color) Printf(format string, a ...interface{}) (n int, err error) {
	c.Set()
	defer Unset()

	return fmt.Fprintf(Output, format, a...)
}

// Print formats using the default formats for its operands and writes to
// standard output. Spaces are added between operands when neither is a
// string. It returns the number of bytes written and any write error
// encountered.
// Standard fmt.Print() method wrapped with the given color.
func (c *Color) Print(a ...interface{}) (n int, err error) {
	c.Set()
	defer Unset()

	return fmt.Fprint(Output, a...)
}

// Println formats using the default formats for its operands and writes to
// standard output. Spaces are always added between operands and a newline is
// appended. It returns the number of bytes written and any write error
// encountered.
// Standard fmt.Println() method wrapped with the given color.
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

// Set sets the given attributess immediately. It will change the color of
// output with the given SGR attributess until color.Unset() is called.
func Set(p ...Attributes) *Color {
	c := New(p...)
	c.Set()
	return c
}

// Unset resets all escape attributes and clears the output. Usualy should
// be called after Set().
func Unset() {
	fmt.Fprintf(Output, "%s[%dm", escape, Reset)
}

// Set sets the SGR sequence.
func (c *Color) Set() *Color {
	fmt.Fprintf(Output, "%s[%sm", escape, c.sequence())
	return c
}
