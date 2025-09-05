package chalk

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	Cyan.Print("Prints text in cyan.")
	Blue.Print("Prints text in blue.")

	// Chain SGR parameters
	Green.Add(Bold).Println("Green with bold.")
	Red.Add(BgWhite, Underline).Printf("Red with White background and underscore: %s\n", "format too!")

	c := Cyan.Add(Underline)
	c.Println("Prints bold cyan.")
	c.Printf("Thir prints bold cyan %s\n", "too!.")

	// Create custom color objects:
	d := New(FgGreen, BgCyan, Italic)
	d.Print("Italic green with cyan background")

	Cyan.Set()
	fmt.Println("Existing text in your codebase will be now in Cyan")
	fmt.Printf("This one %s\n", "too")
	Unset()

	New(FgBlack, BgWhite, Bold).Set()
	defer Unset()

	fmt.Println("All text will be now bold red with white background.")
}