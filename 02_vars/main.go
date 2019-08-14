package main

import "fmt"

func main() {
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// unit unit8 unit16 unit32 unit64 uintptr
	// byte - alias for unit8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	var name string = "Fa"
	// name := "Fa"
	var age int = 26
	var isCool bool = true

	size := 1.3

	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", size)
}
