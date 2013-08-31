/*
	export GOPATH=$HOME/test
	export PKG_CONFIG_PATH=$GOPATH/src/github.com/goinggo/keyboard/pkgconfig
	export DYLD_LIBRARY_PATH=$GOPATH/src/github.com/goinggo/keyboard/DyLib

	go get -d github.com/goinggo/keyboard
	cd $GOPATH/src/github.com/goinggo/keyboard/DyLib
	make
	cd ../pkgconfig
	open -a TextEdit GoingGoKeyboard.pc
	  Change $GOPATH to literal path such as /Users/bill/test
	pkg-config --cflags --libs GoingGoKeyboard
	  -I/Users/bill/test/src/github.com/goinggo/keyboard/DyLib  -L/Users/bill/test/src/github.com/goinggo/keyboard/DyLib -lkeyboard
	cd ..
	go build
	go install
	cd $GOPATH/bin
	./keyboard

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
