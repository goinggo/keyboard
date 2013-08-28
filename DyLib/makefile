# Makefile template for dynamic library

dynamic:
	rm -f libkeyboard.dylib
	rm -f ../TestApp/libkeyboard.dylib
	gcc -c -fPIC keyboard.c
	gcc -dynamiclib -lncurses -r/usr/lib -o libkeyboard.dylib keyboard.o
	rm -f keyboard.o
	cp libkeyboard.dylib ../TestApp/libkeyboard.dylib

shared:
	rm -f libkeyboard.so
	rm -f ../TestApp/libkeyboard.so
	gcc -c -fPIC keyboard.c
	gcc -shared -W1 -lncurses -r/usr/lib -soname,libkeyboard.so -o libkeyboard.so keyboard.o
	rm -f keyboard.o
	cp libkeyboard.so ../TestApp/libkeyboard.so