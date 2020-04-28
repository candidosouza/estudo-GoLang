package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2

	// forma reduzida para criar var
	area := PI * math.Pow(raio, 2)
	fmt.Println("A área da circumferência é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "xpto!"
	fmt.Println(g, h, i)

	fmt.Println(reflect.TypeOf(g))
	fmt.Println(reflect.TypeOf(h))
	fmt.Println(reflect.TypeOf(i))
}
