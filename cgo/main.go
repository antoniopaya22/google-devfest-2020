package main

// #include "HelloWorld.h"
import "C"
import "fmt"

func main() {
	helloWorld()
}

func helloWorld() {
	mensaje := C.hello_world()
	fmt.Println(C.GoString(mensaje))
}
