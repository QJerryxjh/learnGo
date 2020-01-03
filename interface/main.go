package main

import (
	"fmt"
)

// import "fmt"

// type Shaper interface {
// 	area() float32
// }

// // type Square struct {
// // 	side float32
// // }

// // func (s *Square) area() float32 {
// // 	return s.side * s.side
// // }

// // func main() {
// // 	// sq1 := Square{10.1}
// // 	// var shaper Shaper
// // 	// shaper = &sq1 // 传递给接口的必须是指针
// // 	// fmt.Println(shaper)
// // 	// fmt.Printf("%T\n", shaper)
// // 	// fmt.Println(shaper.area())

// // 	// sq1 := new(Square)
// // 	// sq1.side = 10.1
// // 	// var shaper Shaper
// // 	// shaper = sq1 // 传递给接口的必须是指针
// // 	// fmt.Println(shaper)
// // 	// fmt.Printf("%T\n", shaper)
// // 	// fmt.Println(shaper.area())

// // 	sq1 := new(Square)
// // 	sq1.side = 10.1
// // 	shaper := Shaper(sq1)
// // 	fmt.Println(shaper)
// // 	fmt.Printf("%T\n", shaper)
// // 	fmt.Println(shaper.area())

// // 	// sq1 := new(Square)
// // 	// sq1.side = 10.1
// // 	// shaper := sq1
// // 	// fmt.Println(shaper)
// // 	// fmt.Printf("%T\n", shaper)
// // 	// fmt.Println(shaper.area())
// // }

// type Rect struct {
// 	width  float32
// 	height float32
// }

// func (r *Rect) area() float32 {
// 	return r.width * r.height
// }

// var shapers = new([5]Shaper)

// func main() {
// 	sli := []float32{1.1, 2.2, 3.3, 4.4, 5.6}
// 	for i, v := range sli {
// 		rect := &Rect{v, v}
// 		sap := Shaper(rect)
// 		shapers[i] = sap
// 	}
// 	fmt.Println(shapers)
// 	for _, shaper := range shapers {
// 		fmt.Println(shaper.area())
// 	}
// }

// 练习2
// package main

// import "fmt"

// type Shaper interface {
// 	Area() float32
// }

// type Square struct {
// 	side float32
// }

// func (sq *Square) Area() float32 {
// 	return sq.side * sq.side
// }

// type Rectangle struct {
// 	length, width float32
// }

// func (r Rectangle) Area() float32 {
// 	return r.length * r.width
// }

// func main() {

// 	r := Rectangle{5, 3} // Area() of Rectangle needs a value
// 	q := &Square{5}      // Area() of Square needs a pointer
// 	// shapes := []Shaper{Shaper(r), Shaper(q)}
// 	// or shorter
// 	shapes := []Shaper{r, q}
// 	fmt.Printf("%T\n", shapes[0])
// 	fmt.Printf("%T\n", r)
// 	fmt.Println("Looping through shapes for area ...")
// 	for n, _ := range shapes {
// 		fmt.Println("Shape details: ", shapes[n])
// 		fmt.Println("Area of this shape is: ", shapes[n].Area())
// 	}
// }

// 练习3
// package main

// import "fmt"

// type stockPosition struct {
// 	tick  string
// 	price float32
// }

// func (s *stockPosition) getValue() float32 {
// 	return s.price
// }

// type car struct {
// 	price float32
// }

// func (c *car) getValue() float32 {
// 	return c.price
// }

// type valuable interface {
// 	getValue() float32
// }

// func showValue(v valuable) float32 {
// 	return v.getValue()
// }

// func main() {
// 	// var o valuable = stockPosition{"basketball", 1120.2}
// 	var o valuable
// 	sp := &stockPosition{"basketball", 1120.2}
// 	o = sp
// 	fmt.Println(showValue(o))
// 	var c valuable = &car{22.2}
// 	showValue(c)
// }

// 练习4

// type Simpler interface {
// 	Get() int
// 	Set(int)
// }

// type Simple struct {
// 	count int
// }

// func (s *Simple) Get() int {
// 	return s.count
// }

// func (s *Simple) Set(newCount int) {
// 	s.count = newCount
// }

// func actionAll(simpler Simpler) {
// 	fmt.Println(simpler.Get())
// 	simpler.Set(10)
// 	fmt.Println(simpler.Get())
// }

// func main() {
// 	simple := Simple{20}
// 	var simpler Simpler
// 	simpler = &simple
// 	actionAll(simpler)
// }

//

// // 练习5
// type Shaper interface {
// 	Area() int
// }

// type Square struct {
// 	side int
// }

// func (s *Square) Area() int {
// 	return s.side * s.side
// }

// func getArea(s Shaper) {
// 	fmt.Println(s.Area())
// }

// func main() {
// 	s := new(Square)
// 	s.side = 10
// 	getArea(s)
// }

// type any interface {
// }

// type vector struct {
// 	a []any
// }

// func (v *vector) at(i int) any {
// 	return v.a[i]
// }

// func (v *vector) Set(i int, value any) {
// 	v.a[i] = value
// }

// func main() {
// 	vec := new(vector)
// 	vec.a = make([]any, 10)
// 	vec.Set(0, 10)
// 	vec.Set(1, "string")
// 	vec.Set(2, true)
// 	vec.Set(3, false)
// 	vec.Set(4, []int{1, 2, 3})
// 	vec.Set(5, [3]int{1, 2, 3})
// 	fmt.Println(vec)
// 	fmt.Println(vec.at(2))
// }

type Node struct {
	left  *Node
	data  interface{}
	right *Node
}

func newNode(left *Node, data interface{}, right *Node) *Node {
	return &Node{left, data, right}
}

func (n *Node) setData(data interface{}) {
	n.data = data
}

func main() {
	n1 := newNode(nil, "n1", nil)
	n1.setData("new n1")
	fmt.Println(n1)
	n2 := newNode(nil, "n2", nil)
	n2.setData("new n2")
	fmt.Println(n2)
	n1.left = n2
	fmt.Println(n1)
}
