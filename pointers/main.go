package main

import "fmt"

func main() {
	a := 5
	b := &a //memory address

	fmt.Println(a, b)
	fmt.Println("%T\n", b) //ponter *int type will be printed

	fmt.Println(*b) // 5
	fmt.Println(b)  // memory address

	*b = 10 //sets the pinter where b is pointing to to 10 so a is 10

}
