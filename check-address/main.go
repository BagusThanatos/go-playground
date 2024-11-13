package main

import "fmt"


// why is all the input and new variabls on the heap
func check_address(input string){
	new_string := input

	fmt.Printf("address input: %p, address new_string: %p\n", &input, &new_string)
}

func check_address_int(input int){
	new_int := input

	fmt.Printf("address input: %p, address new_int: %p\n", &input, &new_int)
}

func check_address_int_long(input int){
	var new_int_long int64 = int64(input)

	fmt.Printf("address input: %p, address new_int: %p\n", &input, &new_int_long)
}

// but this is not on the heap, probably only due to inline?
func check_address_int_long_with_return(input int) (int64){
	var new_int_long int64 = int64(input)

	return new_int_long
}


func main(){
	hello := "hellowasdfasdfasdfasfd"
	fmt.Printf("address of hello: %p\n", &hello)
	check_address(hello)
	check_address_int(2)
	check_address_int_long(2)
	check_address_int_long_with_return(2)
}