package main

import "fmt"

func main() {
	/*
		EXECUTABLE:
			to create a binary executable you go to the bin folder and ./gobasics
	*/

	//VARIABLES
	/*
		Main Types:
			- string
			- int
			- int8 int16 int32 int64
			- uint uint16 uint32 uint64 (unsigned)
			- byte - alias for uint8
			- rune alias for int32
			- float32 float64
			- complex64 complex128
			// complex creates a complex number from its real and imaginary components,
			 and the built-in real and imag functions extract those components:
			 var x complex128 = complex(1, 2) // 1+2i

	*/
	// Declaration
	// implicit
	name := "brad"
	name, email := "brad", "brad@gmail.com"
	// explicit
	// var name = "Brad"
	// var name string = "brad"
	// var age int32 = 37
	// for explicit declaration

	//Constants
	const isCool bool = true

	//TYPE OF THE VALUE
	fmt.Printf("%T \n", name)
	fmt.Printf("%T \n", email)
	fmt.Printf("%T \n", isCool)

	/*
		if {

		}
		else {

		}

		switch x {
		case 1:
		}

		for i:=1; i<=10; i++ {

		}

	*/

}
