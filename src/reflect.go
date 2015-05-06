package main

import (
	"fmt"
	"reflect"
)

type Bird struct {
	Nmae           string
	LifeExceptance int
}

func (b *Bird) Fly() {
	fmt.Println("I am flying...")
}

func main() {
	sparrow := &Bird{"Sparrow", 3}
	s := reflect.ValueOf(sparrow).Elem()
	typeOFT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d %s %s = %v\n", i, typeOFT.Field(i).Name, f.Type(), f.Interface())
	}
}
