package chalk

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	Cyan.Print("Prints text in cyan.")
	Blue.Print("Prints text in blue.")

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
}