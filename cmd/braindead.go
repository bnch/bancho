package cmd

import (
	"fmt"
	"os"
	"runtime"
)

// Braindead helps the poor little windows users that don't know what's
// a terminal, and have opened bancho by doubleclicking on the executable
func Braindead() {
	if runtime.GOOS == "windows" {
		fmt.Print("Press a key to continue . . .")
		os.Stdin.Read(make([]byte, 1))
		fmt.Println()
	}
}
