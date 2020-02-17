package main

import "fmt"

func main() {
	//Arrays
	var fruitArray [2]string
	//fruitArray := [2]string{apple, orange}

	//fruitArray[0] = "apple"
	// fruitSlice := []string{"Apple", "orange", "grape"}

	// fmt.Println(len(fruitSlice))
	// //range
	// fmt.Println(fruitSlice[1:2])

	emails := make(map[string]string) //key, value

	emails["bob"] = "bob@gmail.com"
	emails["caz"] = "caz@gmail.com"

	// emails := map[string]string{key1:val1, key2:val2, ...}

	delete(emails, "bob")

	ids := []int{33, 76, 343, 64, 45, 32, 2, 12}

	for i, id : = range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//0 - ID: 33

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
	//bob: bob@gmail.com
}
