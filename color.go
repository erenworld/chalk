package chalk

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

const escape = "\x1b"

// Output defines the standard output of the print functions. 
// Any io.Writer can be used.
var Output io.Writer = colorable.NewColorableStdout()

// This is a global variable and affects all colors.
// For more control over each color use the method DisableColor() individually.
var NoColor = !isatty.IsTerminal(os.Stdout.Fd()) || os.Getenv("TERM") == "dumb"

type Color struct {
	params  []Attribute
	noColor *bool
}


// Attribute defines a single SGR Code
type Attribute int

const (
	Reset Attribute = iota
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
	FgBlack Attribute = iota + 30
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
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background colors
const (
	BgBlack Attribute = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background Hi-Intensity text colors
const (
	BgHiBlack Attribute = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)


func New(value ...Attribute) *Color {
	c := &Color{params: make([]Attribute, 0)}
	c.Add(value...)
	return c
}

func BlackString(format string, a ...interface{}) string {
	return printString(format, FgBlack, a...)
}
func RedString(format string, a ...interface{}) string {
	return printString(format, FgRed, a...)
}
func GreenString(format string, a ...interface{}) string {
	return printString(format, FgGreen, a...)
}
func BlueString(format string, a ...interface{}) string {
	return printString(format, FgBlue, a...)
}
func MagentaString(format string, a ...interface{}) string {
	return printString(format, FgMagenta, a...)
}
func CyanString(format string, a ...interface{}) string {
	return printString(format, FgCyan, a...)
}
func WhiteString(format string, a ...interface{}) string {
	return printString(format, FgWhite, a...)
}
func YellowString(format string, a ...interface{}) string {
	return printString(format, FgYellow, a...)
}

func Black(format string, a ...interface{})   { printColor(format, FgBlack, a...) }
func Red(format string, a ...interface{})     { printColor(format, FgRed, a...) }
func Green(format string, a ...interface{})   { printColor(format, FgGreen, a...) }
func Yellow(format string, a ...interface{})  { printColor(format, FgYellow, a...) }
func Blue(format string, a ...interface{})    { printColor(format, FgBlue, a...) }
func Magenta(format string, a ...interface{}) { printColor(format, FgMagenta, a...) }
func Cyan(format string, a ...interface{})    { printColor(format, FgCyan, a...) }
func White(format string, a ...interface{})   { printColor(format, FgWhite, a...) }

func (c *Color) Bold() *Color {
	c.Add(Bold)
	return c
}

func (c *Color) Add(value ...Attribute) *Color {
	c.params = append(c.params, value...)
	return c
}

func printColor(format string, p Attribute, a ...interface{}) {
	// If no arguments (a...) are given, it treats format as the text itself.
	if len(a) == 0 {
		a = append(a, format)
		format = "%s"
	}

	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	c := &Color{params: []Attribute{p}}
	c.Printf(format, a...)
}

func printString(format string, p Attribute, a ...interface{}) string {
	if len(a) == 0 {
		a = append(a, format)
		format = "%s"
	}

	return New(p).SprintfFunc()(format, a...)
}

func (c *Color) Printf(format string, a ...interface{}) (n int, err error) {
	c.Set()
	defer c.unset()

	return fmt.Fprintf(Output, format, a...)
}

func (c *Color) Print(a ...interface{}) (n int, err error) {
	c.Set()
	defer c.unset()

	return fmt.Fprint(Output, a...)
}

func (c *Color) Println(a ...interface{}) (n int, err error) {
	c.Set()
	defer c.unset()

	return fmt.Fprintln(Output, a...)
}

func (c *Color) PrintFunc() func(a ...interface{}) {
	return func(a ...interface{}) { c.Print(a...) }
}

func (c *Color) PrintfFunc() func(format string, a ...interface{}) {
	return func(format string, a ...interface{}) { c.Printf(format, a...) }
}

func (c *Color) PrintlnFunc() func(a ...interface{}) {
	return func(a ...interface{}) { c.Println(a...) }
}

func (c *Color) SprintFunc() func(a ...interface{}) string {
	return func(a ...interface{}) string {
		return c.wrap(fmt.Sprint(a...))
	}
}

func (c *Color) SprintfFunc() func(format string, a ...interface{}) string {
	return func(format string, a ...interface{}) string {
		return c.wrap(fmt.Sprintf(format, a...))
	}
}

func (c *Color) SprintlnFunc() func(a ...interface{}) string {
	return func(a ...interface{}) string {
		return c.wrap(fmt.Sprintln(a...))
	}
}

// sequence returns a formated SGR sequence to be plugged into a "\033[...m"
// an example output might be: "1;36" -> bold cyan
func (c *Color) sequence() string {
	format := make([]string, len(c.params))
	for i, v := range c.params {
		format[i] = strconv.Itoa(int(v))
	}

	return strings.Join(format, ";")
}

func Set(p ...Attribute) *Color {
	c := New(p...)
	c.Set()
	return c
}

func Unset() {
	if NoColor {
		return
	}
	fmt.Fprintf(Output, "%s[%dm", escape, Reset)
}

func (c *Color) Set() *Color {
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

func (c *Color) wrap(s string) string {
	if c.isNoColorSet() {
		return s
	}

	return c.format() + s + c.unformat()
}

func (c *Color) format() string {
	return fmt.Sprintf("%s[%sm", escape, c.sequence())
}

func (c *Color) unformat() string {
	return fmt.Sprintf("%s[%dm", escape, Reset)
}

// DisableColor disables the color output. Useful to not change any existing
// code and still being able to output. Can be used for flags like
// "--no-color". To enable back use EnableColor() method.
func (c *Color) DisableColor() {
	c.noColor = boolPtr(true)
}

// EnableColor enables the color output. Use it in conjuction with
// DisableColor(). Otherwise this method has no side effects.

func (c *Color) EnableColor() {
	c.noColor = boolPtr(false)

}

func (c *Color) isNoColorSet() bool {
	// check first if we have user setted action
	if c.noColor != nil {
		return *c.noColor
	}

	// if not return the global option, which is disabled by default
	return NoColor
}

func boolPtr(v bool) *bool {
	return &v
}

// Equals returns a boolean value indicating whether two colors are equals.
func (c *Color) Equal(c2 *Color) bool {
	if len(c.params) != len(c2.params) {
		return false
	}

	for _, attr := range c.params {
		if !c2.hasAttr(attr) {
			return false
		}
	}

	return true
}

func (c *Color) hasAttr(a Attribute) bool {
	for _, attr := range c.params {
		if attr == a {
			return true
		}
	}

	return false
}