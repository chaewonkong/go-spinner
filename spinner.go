package spinner

import (
	"fmt"
	"time"
)

var (
	// Default simple stick
	Default = []string{"-", "\\", "|", "/"}

	// Arrow arrow
	Arrow = []string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"}

	// Bar bar
	Bar = []string{"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▁"}
)

// Spin print spinner with given delay
func Spin(delay time.Duration, spinText []string) {
	for {
		for _, r := range spinText {
			fmt.Printf("\r%s", r)
			time.Sleep(delay)
		}
	}
}
