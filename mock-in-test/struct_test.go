package main

import (
	"testing"
)

type Try_interface interface{
	Add(num int)
}

func TestTry(t *testing.T){
	t.Run("try something", func(t *testing.T){
		var _ Try_interface = &Struct_a{0}
	})
}

