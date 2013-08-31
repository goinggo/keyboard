/*
	// This will build on the Mac Only. The DyLib has been pre-built.

	// Set these variable before getting and building the code
	export GOPATH=$HOME/test
	export PKG_CONFIG_PATH=$GOPATH/src/github.com/goinggo/keyboard/pkgconfig
	export DYLD_LIBRARY_PATH=$GOPATH/src/github.com/goinggo/keyboard/DyLib

	// Get, build and install the code
	go get github.com/goinggo/keyboard

	// Run the code
	cd $GOPATH/bin
	./keyboard

	// Show things are set correctly
	pkg-config --cflags --libs GoingGoKeyboard
*/
package main

/*
#cgo pkg-config: --define-variable=prefix=./ GoingGoKeyboard
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
