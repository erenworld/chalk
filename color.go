package chalk

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/shiena/ansicolor"
)

// NoColor defines if the output should be colorized or not.
// It's global and affects all Colors.
var NoColor bool = false

// Color defines a custom color object which is defined by SGR attributes.
type Color struct {
	params 	[]Attributes
	NoColor *bool
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
var Output io.Writer = ansicolor.NewAnsiColorWriter(os.Stdout)

func printColor(format string, p Attributes, a ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	c := &Color{params: []Attributes{p}}
	c.Printf(format, a...)
}

// Set sets the given attributes immediately. It will change the color of
// output with the given SGR attributes until color.Unset() is called.
func Set(p ...Attributes) *Color {
	c := New(p...)
	c.set()
	return c
}

// Unset resets all escape attributes and clears the output. Usualy should
// be called after Set().
func Unset() {
	if NoColor {
        return
    }
	fmt.Fprintf(Output, "%s[%dm", escape, Reset)
}

// Set sets the SGR sequence.
func (c *Color) Set() *Color {
	if NoColor {
		return c
	}
    fmt.Fprint(Output, c.format())
	return c
}

// Set sets the SGR sequence.
func (c *Color) set() *Color {
	if c.isNoColorSet() {
		return c
	}
	fmt.Fprint(Output, c.format())

	return c
}

func (c *Color) unset() {
	if c.isNoColorSet() {
		return 
	}
	Unset()
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

// Add is used to chain SGR attributes. Use as many as paramters to combine
// and create custom color objects. Example: Add(color.FgRed, color.Underline)
func (c *Color) Add(value ...Attributes) *Color {
	c.params = append(c.params, value...)
	return c
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
// Standard fmt.PrintF() method wrapped with the given color.
func (c *Color) Printf(format string, a ...interface{}) (n int, err error) {
	c.set()
	defer c.unset()

	return fmt.Fprintf(Output, format, a...)
}

// Print formats using the default formats for its operands and writes to
// standard output. Spaces are added between operands when neither is a
// string. It returns the number of bytes written and any write error
// encountered.
// Standard fmt.Print() method wrapped with the given color.
func (c *Color) Print(a ...interface{}) (n int, err error) {
	c.set()
	defer c.unset()

	return fmt.Fprint(Output, a...)
}

// Println formats using the default formats for its operands and writes to
// standard output. Spaces are always added between operands and a newline is
// appended. It returns the number of bytes written and any write error
// encountered.
// Standard fmt.Println() method wrapped with the given color.
func (c *Color) Println(a ...interface{}) (n int, err error) {
	c.set() 		// applique "\033[31m"
	defer c.unset()	// à la fin, applique "\033[0m"

	return fmt.Fprintln(Output, a...)
}

// PrintFunc returns a new function prints the passed arguments as colorized
// with color.Print().
func (c *Color) PrintFunc() func(a ...interface{}) {
	return func(a ...interface{}) { c.Print(a...) }
}

// PrintfFunc returns a new function prints the passed arguments as colorized
// with color.Printf().
func (c *Color) PrintfFunc() func(format string, a ...interface{}) {
	return func(format string, a ...interface{}) { c.Printf(format, a...) }
}

// PrintlnFunc returns a new function prints the passed arguments as colorized
// with color.Println().
func (c *Color) PrintlnFunc() func(a ...interface{}) {
	return func(a ...interface{}) { c.Println(a...) }
}

// SprintFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprint(). Useful to put into or mix into other
// string.
// Windows users should use this in conjuction with color.Output, example:
//	put := New(FgYellow).SprintFunc()
//	fmt.Ffprintf(color.Output, "This is a %s", put("warning"))
func (c *Color) SprintFunc() func(a ...interface{}) string {
	return func(a ...interface{}) string {
		return c.wrap(fmt.Sprint(a...))
	}
}

// SprintfFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprintf(). Useful to put into or mix into other
// string. Windows users should use this in conjuction with color.Output.
func (c *Color) SprintfFunc() func(format string, a ...interface{}) string {
	return func(format string, a ...interface{}) string {
		return c.wrap(fmt.Sprintf(format, a...))
	}
}

// SprintlnFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprintln(). Useful to put into or mix into other
// string. Windows users should use this in conjuction with color.Output.
func (c *Color) SprintlnFunc() func(a ...interface{}) string {
	return func(a ...interface{}) string {
		return c.wrap(fmt.Sprintln(a...))
	}
}

func (c *Color) wrap(s string) string {
	if c.isNoColorSet() {
		return s
	}
    return c.format() + s + c.unformat()
}

func (c *Color) unformat() string { 
	return fmt.Sprintf("%s[%dm", escape, Reset)
}

func (c *Color) format() string { 
	return fmt.Sprintf("%s[%sm", escape, c.sequence())
}

// DisableColor disables the color output. Useful to not change any existing
// code and still being able to output. Can be used for flags like
// "--no-color". To enable back use EnableColor() method.
func (c *Color) DisableColor() {
	t := new(bool)
	*t = true
	c.NoColor = t
}

// EnableColor enables the color input. Use it in conjuction with
// DisableColor(). Otherwise this method has no side effects.
func (c *Color) EnableColor() {
	t := new(bool)
	*t = false
	c.NoColor = t
}

func (c *Color) isNoColorSet() bool {
	// check first if we have user setted action
	if c.NoColor != nil {
		return *c.NoColor
	}

	// if not return the global option, which is disabled by default
	return NoColor
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


