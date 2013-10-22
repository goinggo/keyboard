// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	This program provides sample code that shows you how to bind a dynamic library using CGO.

	Ardan Studios
	12973 SW 112 ST, Suite 153
	Miami, FL 33186
	bill@ardanstudios.com

	// Set these variable before getting and building the code
	export GOPATH=$HOME/keyboard
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
#cgo pkg-config: --define-variable=prefix=. GoingGoKeyboard
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
