package main

import (
	"fmt"
	"reflect"
)

// func main() {
// 	var x float64 = 20
// 	v := reflect.ValueOf(x)
// 	v2 := reflect.ValueOf(&x)
// 	fmt.Println(v)
// 	fmt.Println(v2)
// 	fmt.Println(v.CanSet())
// 	fmt.Println(v2.CanSet())
// 	v2 = v2.Elem()
// 	fmt.Println(v2.CanSet())
// 	fmt.Println(v2)
// 	fmt.Println(v.Kind() == reflect.Float64)
// 	fmt.Println(reflect.Float64)
// 	fmt.Println(v.Interface())
// }

// type NotknownType struct {
// 	s1, s2, s3 string
// }

// func (n NotknownType) String() string {
// 	return n.s1 + " - " + n.s2 + " - " + n.s3
// }

// var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

// type test interface {
// }

// func main() {
// 	value := reflect.ValueOf(secret)
// 	fmt.Println(value)
// 	typ := reflect.TypeOf(secret)
// 	fmt.Println(typ)
// 	var t test = 1
// 	fmt.Println(t)
// 	fmt.Println(t.(test))
// 	fmt.Printf("%T", t)
// 	knd := value.Kind()
// 	fmt.Println(knd)
// }

type T struct {
	A int
	B string
}

func main() {
	t := T{22, "ww"}
	setableT := reflect.ValueOf(&t).Elem()
	fmt.Println(setableT)
	fmt.Println(setableT.CanSet())
	typT := setableT.Type()
	fmt.Println(typT)
	fmt.Println(setableT.NumField())
	fmt.Println(setableT.Field(0).Type())
	fmt.Println(setableT.Field(0).Interface())
	fmt.Println(typT.Field(0).Name)
	setableT.Field(0).SetInt(100)
	fmt.Println(setableT)
	fmt.Println(t)
}
