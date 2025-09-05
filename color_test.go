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
	c.Println("Prints cyan text with an underline.")
	c.Printf("This prints bold cyan %s\n", "too!.")

	// Create custom color objects:
	d := New(FgWhite, BgGreen)
	d.Println("White with green bg")

	Yellow.Set()
	fmt.Println("Existing text in your codebase will be now in Yellow")
	fmt.Printf("This one %s\n", "too")
	Unset()

	New(FgMagenta, Bold).Set()
	defer Unset()

	fmt.Println("All text will be now bold red with white background.")
}