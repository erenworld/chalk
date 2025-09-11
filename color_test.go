package chalk

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

// Testing colors is kinda different. First we test for given colors and their
// escaped formatted results. Next we create some visual tests to be tested.
// Each visual test includes the color name to be compared.
func TestColor(t *testing.T) {
	rb := new(bytes.Buffer)
	Output = rb

	testColors := []struct {
		text string
		code Attributes
	}{
		{text: "black", code: FgBlack},
		{text: "red", code: FgRed},
		{text: "green", code: FgGreen},
		{text: "yellow", code: FgYellow},
		{text: "blue", code: FgBlue},
		{text: "magent", code: FgMagenta},
		{text: "cyan", code: FgCyan},
		{text: "white", code: FgWhite},
	}

	for _, c := range testColors {
		New(c.code).Print(c.text)

		line, _ := rb.ReadString('\n')
		scannedLine := fmt.Sprintf("%q", line)
		colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", c.code, c.text)
		escapedForm := fmt.Sprintf("%q", colored)

		fmt.Printf("%s\t: %s\n", c.text, line)

		if scannedLine != escapedForm {
			t.Errorf("Expecting %s, got '%s'\n", escapedForm, scannedLine)
		}
	}

	// First Visual Test
	fmt.Println("")
	Output = os.Stdout

	New(FgRed).Printf("red\t")
	New(BgRed).Print("         ")
	New(FgRed, Bold).Println(" red")

	New(FgGreen).Printf("green\t")
	New(BgGreen).Print("         ")
	New(FgGreen, Bold).Println(" green")

	New(FgYellow).Printf("yellow\t")
	New(BgYellow).Print("         ")
	New(FgYellow, Bold).Println(" yellow")

	New(FgBlue).Printf("blue\t")
	New(BgBlue).Print("         ")
	New(FgBlue, Bold).Println(" blue")

	New(FgMagenta).Printf("magenta\t")
	New(BgMagenta).Print("         ")
	New(FgMagenta, Bold).Println(" magenta")

	New(FgCyan).Printf("cyan\t")
	New(BgCyan).Print("         ")
	New(FgCyan, Bold).Println(" cyan")

	New(FgWhite).Printf("white\t")
	New(BgWhite).Print("         ")
	New(FgWhite, Bold).Println(" white")

	fmt.Println("")

	c := New(FgRed).Add(Underline)
	c.Println("Prints red text with an underline.")

	New(FgMagenta, Bold).Set()
	defer Unset()
	fmt.Println("All text will be now bold magenta.")

	// Second Visual test
	Red("red")
	Black("black")
	Green.Printf("green\n")

	// Third visual test
	fmt.Println()
	Set(FgBlue)
	fmt.Println("is this blue ?")
	Unset()

	Set(FgMagenta)
	fmt.Println("and this magenta?")
	Unset()

	// Fourth visual test
	fmt.Println()
	blue := New(FgBlue).PrintlnFunc()
	blue("blue text with custom print func")

	red := New(FgRed).PrintfFunc()
	red("red text with a printf func: %d\n", 123)

	put := New(FgYellow).SprintFunc()
	warn := New(FgRed).SprintFunc()

	fmt.Printf("this is a %s and this is %s. Aborting\n", put("warning"), warn("error"))

	info := New(FgWhite, BgGreen).SprintFunc()
	fmt.Printf("this %s rocks!\n", info("package"))
}