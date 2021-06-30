package main

import (
	"fmt"

	"github.com/shimmerglass/clip"
)

type A struct {
	Foo string `clip:"foo,blue"`
	Bar B      `clip:",inline"`
	Buz C      `clip:"buz"`
}

type B struct {
	F1 int `clip:"f1"`
	F2 int `clip:"f2"`
}

type C struct {
	Hello string `clip:"hello"`
	World []int  `clip:"world"`
}

func main() {
	err := clip.Print(A{
		Foo: "bar",
		Bar: B{
			F1: 0,
			F2: 50,
		},
		Buz: C{
			Hello: "hello",
			World: []int{
				1, 2, 3, 4,
			},
		},
	})
	if err != nil {
		fmt.Println(err)
	}
}
