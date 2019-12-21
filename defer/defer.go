package main

import "fmt"

// func main() {
// 	fmt.Println(1)
// 	defer fmt.Println(2)
// 	fmt.Println(3)
// }
// 1  3  2

// func main() {
// 	for i := 0; i < 5; i++ {
// 		defer fmt.Println(i)
// 	}
// }

// func main() {
// 	i := 0
// 	defer fmt.Println(i)
// 	i++
// 	return
// }

func main() {
	fmt.Println(1)
	defer fmt.Println(test2())
	fmt.Println(2)
}

func test2() int {
	fmt.Println("test2")
	return 3
}
