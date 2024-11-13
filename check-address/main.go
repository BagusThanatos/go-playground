package main

import "fmt"

func check_address(input string){
	new_string := input

	// Does not actually allocates?
	fmt.Printf("address input: %p, address new_string: %p", &input, &new_string)
}

func check_address_int(input int){
	new_int := input

	// Does not actually allocates?
	fmt.Printf("address input: %p, address new_int: %p", &input, &new_int)
}

func main(){
	check_address("hellowasdfasdfasdfasfd")
	check_address_int(2)
}