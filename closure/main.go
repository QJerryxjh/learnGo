package main

import "fmt"

// func main() {
// 	func(a int, b int) {
// 		fmt.Println(0, a, b)
// 	}(2, 3)
// }

// func main() {
// 	fv := func() {
// 		fmt.Println("hello world")
// 	}
// 	fv()
// 	fmt.Println(reflect.TypeOf(fv))
// }

// func f() (ret int) {
// 	fmt.Println(ret)
// 	defer func() {
// 		fmt.Println(ret)
// 		ret++
// 	}()
// 	fmt.Println(ret)
// 	return 7
// }
// func main() {
// 	fmt.Println(f())
// }

func main() {
	retFn := Add2()
	fmt.Println(retFn(3))
	fmt.Println(retFn(3))
	fmt.Println(retFn(3))
}

func Add2() func(c int) int {
	o := 0
	return func(b int) int {
		o++
		return b + o
	}
}
