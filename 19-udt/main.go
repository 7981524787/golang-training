package main

import "fmt"

func main() {

	r1 := Rect{L: 23.23, B: 23.90}
	fmt.Println(r1)
	fmt.Println("Rect Length:", r1.L, "\nRect Bredth:", r1.B)
	// a1 := Area(&r1)
	// fmt.Println("Area of r1:", a1)
	// p1 := Perimeter(r1)
	// fmt.Println("Perimeter of r1:", p1)
	// fmt.Println(r1.A)
	a2 := r1.Area()
	fmt.Println("Area of r1:", a2)
	fmt.Println(r1.A)
	r2 := new(Rect) // r2 is a pointer
	r2.L = 123.23
	r2.B = 123.34

	r3 := New(123.123, 213.23)
	a4, p4 := r3.Area(), r3.Perimeter()
	fmt.Println("Area of r3:", a4)
	fmt.Println("perimeter of r3:", p4)
}

type Rect struct {
	L    float32
	B    float32
	A, P float32
}

func New(l, b float32) *Rect {
	return &Rect{L: l, B: b}
}

// func (r *Rect) New(l, b float32) *Rect {
// 	return &Rect{L: l, B: b}
// }

func NewRectDefaults() *Rect {
	return &Rect{L: 1, B: 1}
}

func Area(r *Rect) float32 { // function
	(*r).A = (*r).L * r.B
	return r.A
}

func Perimeter(r *Rect) float32 { // function
	r.P = 2 * (r.L + r.B)
	return r.P
}

func (r *Rect) Area() float32 { // method
	r.A = r.L * r.B
	return r.A
}
func (r *Rect) Perimeter() float32 { // function
	r.P = 2 * (r.L + r.B)
	return r.P
}

// type Square struct {
// 	S float32
// }

// func NewSquare(s float32) *Square {
// 	return &Square{S: s}
// }
