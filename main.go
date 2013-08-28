package main

/*
#cgo CFLAGS: -IDyLib
#cgo LDFLAGS: -L. -lkeyboard
#include <keyboard.h>
*/
import "C"
import (
	"fmt"
)

func main() {

	C.InitKeyboard()

	fmt.Printf("\nEnter: ")

	for {
		r := C.GetCharacter()

		fmt.Printf("%c", r)

		if r == 'q' {
			break
		}
	}

	C.CloseKeyboard()
}
