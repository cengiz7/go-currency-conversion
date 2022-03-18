package cli

import "fmt"

func ClearScreen(fromTop bool) {
	if fromTop {
		fmt.Printf("\x1b[2J")
	} else {
		fmt.Printf("\x1bc")
	}
}

func EnterToContinue() {
	fmt.Printf("\n\nPress enter to continue..")
	fmt.Scanf("%s")
}