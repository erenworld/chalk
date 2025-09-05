package chalk

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"testing"
)

func TestColor(t *testing.T) {
	go func() {
		scanner := bufio.NewScanner(os.Stdout)
		for scanner.Scan() {
			fmt.Println(scanner.Text()) // Println will add back the final '\n'
		}
	}()

	time.Sleep(time.Millisecond * 300)

	c := Red.Bold()
	c.Println("Afyon")


}