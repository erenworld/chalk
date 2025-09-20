package chalk

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/shiena/ansicolor"
)

const escape = "\x1b"

// Output defines the standard output of the print functions. Any io.Writer
// can be used.
var Output io.Writer = ansicolor.NewAnsiColorWriter(os.Stdout)

type Color struct {
	params []Attribute
}

// Attribute defines a single SGR Code
type Attribute int

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

func BlackString(format string, a ...interface{}) string {
	return New(FgBlack).SprintfFunc()(format, a...)
}
func RedString(format string, a ...interface{}) string {
	return New(FgRed).SprintfFunc()(format, a...)
}
func GreenString(format string, a ...interface{}) string {
	return New(FgGreen).SprintfFunc()(format, a...)
}
func BlueString(format string, a ...interface{}) string {
	return New(FgBlue).SprintfFunc()(format, a...)
}
func MagentaString(format string, a ...interface{}) string {
	return New(FgMagenta).SprintfFunc()(format, a...)
}
func CyanString(format string, a ...interface{}) string {
	return New(FgCyan).SprintfFunc()(format, a...)
}
func WhiteString(format string, a ...interface{}) string {
	return New(FgWhite).SprintfFunc()(format, a...)
}
func YellowString(format string, a ...interface{}) string {
	return New(FgYellow).SprintfFunc()(format, a...)
}




// Black is an convenient helper function to print with black foreground. A
// newline is appended to format by default.
func Black(format string, a ...interface{}) { printColor(format, FgBlack, a...) }

// Red is an convenient helper function to print with red foreground. A
// newline is appended to format by default.
func Red(format string, a ...interface{}) { printColor(format, FgRed, a...) }

// Green is an convenient helper function to print with green foreground. A
// newline is appended to format by default.
func Green(format string, a ...interface{}) { printColor(format, FgGreen, a...) }

// Yellow is an convenient helper function to print with yello foreground.
// A newline is appended to format by default.
func Yellow(format string, a ...interface{}) { printColor(format, FgYellow, a...) }

// Blue is an convenient helper function to print with blue foreground. A
// newline is appended to format by default.
func Blue(format string, a ...interface{}) { printColor(format, FgBlue, a...) }

// Magenta is an convenient helper function to print with magenta foreground.
// A newline is appended to format by default.
func Magenta(format string, a ...interface{}) { printColor(format, FgMagenta, a...) }

// Cyan is an convenient helper function to print with cyan foreground. A
// newline is appended to format by default.
func Cyan(format string, a ...interface{}) { printColor(format, FgCyan, a...) }

// White is an convenient helper function to print with white foreground. A
// newline is appended to format by default.
func White(format string, a ...interface{}) { printColor(format, FgWhite, a...) }

func New(value ...Attribute) *Color {
	c := &Color{params: make([]Attribute, 0)}
	c.Add(value...)
	return c
}

func (c *Color) Bold() *Color {
	c.Add(Bold)
	return c
}

func (c *Color) Add(value ...Attribute) *Color {
	c.params = append(c.params, value...)
	return c
}

func printColor(format string, p Attribute, a ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	c := &Color{params: []Attribute{p}}
	c.Printf(format, a...)
}

func (c *Color) prepend(value Attribute) {
	c.params = append(c.params, 0)
	copy(c.params[1:], c.params[0:])
	c.params[0] = value
}


func (c *Color) Printf(format string, a ...interface{}) (n int, err error) {
	c.Set()
	defer Unset()

	return fmt.Fprintf(Output, format, a...)
}

func (c *Color) Print(a ...interface{}) (n int, err error) {
	c.Set()
	defer Unset()

	return fmt.Fprint(Output, a...)
}

func (c *Color) Println(a ...interface{}) (n int, err error) {
	c.Set()
	defer Unset()

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

// Set sets the SGR sequence.
func (c *Color) Set() *Color {
	fmt.Fprintf(Output, "%s[%sm", escape, c.sequence())
	return c
}

func Set(p ...Attribute) *Color {
	c := New(p...)
	c.Set()
	return c
}

func Unset() {
	fmt.Fprintf(Output, "%s[%dm", escape, Reset)

}

func (c *Color) wrap(s string) string {
	return c.format() + s + c.unformat()
}

func (c *Color) format() string {
	return fmt.Sprintf("%s[%sm", escape, c.sequence()) 
}

func (c *Color) unformat() string {
	return fmt.Sprintf("%s[%dm", escape, Reset)
}



