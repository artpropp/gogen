package main

import "fmt"

type MyType struct {
	Foo string
	Bar string
}

//go:generate gogen ../stack.gogen MyType
//go:generate gogen ../stack.gogen int

func main() {
	mts := MyTypeStack{}
	mts.Push(MyType{"foo1", "bar1"})
	mts.Push(MyType{"foo2", "bar2"})
	p := mts.Pop()
	fmt.Printf("%#v\n", p)

	is := intStack{}
	is.Push(1)
	is.Push(2)
	is.Push(3)
	fmt.Println(is.Pop())
	fmt.Println(is.Pop())
	fmt.Println(is.Pop())
	fmt.Println(is.Pop())
}