/*
	-- To build this code set this variable and verify the pkg-config tool can find the pkg-config file
	export PKG_CONFIG_PATH=$GOPATH/src/github.com/goinggo/keyboard/pkgconfig
	pkg-config --cflags --libs GoingGoKeyboard

	-- To run this code set this variable so the dynamic library can be found at runtime
	export DYLD_LIBRARY_PATH=$GOPATH/src/github.com/goinggo/keyboard/DyLib

	go get github.com/goinggo/keyboard
*/
package main

/*
#cgo pkg-config: GoingGoKeyboard
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
