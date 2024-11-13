package main


type Struct_a struct{
	a int
}

func (data *Struct_a) Add(num int){
	data.a += num
}
