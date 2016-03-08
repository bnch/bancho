package cmd

import (
	"fmt"
	"os"
	"runtime"
)

// Braindead helps the poor little windows users that don't know what's
// a terminal, and have opened bancho by doubleclicking on the executable
func Braindead() {
	exit := 0
	c := recover()
	if c != nil {
		fmt.Println("An error happened while attempting to execute bnch.")
		fmt.Println(c)
		exit = 1
	}
	if runtime.GOOS == "windows" {
		fmt.Print("Press a key to continue . . .")
		os.Stdin.Read(make([]byte, 1))
		fmt.Println()
	}
	if exit != 0 {
		os.Exit(exit)
	}
}
