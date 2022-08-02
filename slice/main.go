package main

import "fmt"

type Coba struct{
	a, b int
	c float64
}

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	coba := make([]Coba, 2)
	printCoba(coba)

	coba = append(coba, Coba{1,2,3})
	printCoba(coba)

	fmt.Println("Print loop")
	for _, value := range coba {
		fmt.Printf("%v\n", value)
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func printCoba(coba []Coba){
	fmt.Printf("Coba: %v\n", coba)
}