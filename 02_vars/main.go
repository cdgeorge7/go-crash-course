package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	//var name = "Chris"
	var age int32 = 31
	const isCool = true

	// Shorthand
	// 	name := "Biscuits"
	size := 1.3
	name, email := "Biscuits", "biscuits@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}
