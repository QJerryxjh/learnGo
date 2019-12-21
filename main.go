package main

import "fmt"

// // package main

// // import "fmt"

// // // package main

// // // import "fmt"

// // // // package main

// // // // import (
// // // // 	"fmt"
// // // // 	"os"
// // // // 	"runtime"
// // // // )

// // // // // const n = 1
// // // // // const a, b, c = 1, 2, 3
// // // // // const (
// // // // // 	e, f, g = 4, 5, 6
// // // // // 	h, i, j = 7, 8, 9
// // // // // )

// // // // // const (
// // // // // 	a = iota + 2
// // // // // 	d
// // // // // 	g
// // // // // 	b = iota
// // // // // 	c
// // // // // )

// // // // const a = iota
// // // // const b = iota
// // // // const c = iota

// // // // func test() {
// // // // }

// // // // func main() {
// // // // 	// var e = 1
// // // // 	// d := 2
// // // // 	// fmt.Print(d, e)

// // // // 	// // test();
// // // // 	// fmt.Println(a, b, c)

// // // // 	a := 1
// // // // 	a, b, c := 2, 3, 4

// // // // 	fmt.Println(runtime.GOOS, a, b, c)
// // // // 	fmt.Println(os.Getenv("PATH"))
// // // // 	fmt.Println("\n\n\n\n\tss")
// // // // 	fmt.Println(`"\n\n\n\n\tss"`)

// // // // 	goo := runtime.GOOS
// // // // 	fmt.Printf("whahslda s%s", goo)
// // // // 	fmt.Println(&goo)
// // // // 	fmt.Print("1", 2)

// // // // }

// // // // package main

// // // // var a = "G"

// // // // func main() {
// // // // 	n()
// // // // 	m()
// // // // 	n()
// // // // }

// // // // func n() { print(a) }

// // // // func m() {
// // // // 	a := "O"
// // // // 	print(a)
// // // // }

// // // // package main

// // // // var a = "G"

// // // // func main() {
// // // // 	n()
// // // // 	m()
// // // // 	n()
// // // // }

// // // // func n() {
// // // // 	print(a)
// // // // }

// // // // func m() {
// // // // 	a = "O"
// // // // 	print(a)
// // // // }

// // // // package main

// // // // var a string

// // // // func main() {
// // // // 	a = "G"
// // // // 	print(a)
// // // // 	f1()
// // // // }

// // // // func f1() {
// // // // 	a := "O"
// // // // 	print(a)
// // // // 	f2()
// // // // }

// // // // func f2() {
// // // // 	print(a)
// // // // }

// // // // package main

// // // // import (
// // // // 	"fmt"
// // // // 	"reflect"
// // // // 	"strconv"
// // // // 	"strings"
// // // // 	"time"
// // // // )

// // // // func main() {
// // // // var a int
// // // // var b float32
// // // // a = 15
// // // // b = a + a
// // // // print(b + 6)
// // // // var n int16 = 34
// // // // var m int32

// // // // m = int32(n)
// // // // print(m)

// // // // str := "strin"
// // // // var b = 0
// // // // var c string
// // // // fmt.Println(b, c)
// // // // fmt.Println(time.Now().Nanosecond())
// // // // for i := 0; i < 10; i++ {
// // // // 	fmt.Printf("%d\n", rand.Intn(10))
// // // // }
// // // // if b > 1 {
// // // // 	print(true)
// // // // }

// // // // 	type TZ int
// // // // 	var a int
// // // // 	var b TZ
// // // // 	b = 2
// // // // 	var c = "abcd"
// // // // 	fmt.Println(a, b)
// // // // 	fmt.Println(reflect.TypeOf(""))
// // // // 	fmt.Println(len(c))
// // // // 	fmt.Println("kl" + "paa")
// // // // 	fmt.Println(strings.HasPrefix("example", "aex"))
// // // // 	fmt.Println(strings.HasSuffix("example", "l"))
// // // // 	fmt.Println(strings.Contains("example", "PL"))
// // // // 	fmt.Println(strings.Index("example", "xaa"))
// // // // 	fmt.Println(strings.LastIndex("example", "e"))
// // // // 	fmt.Println(strings.Replace("eeeaaacccssss", "s", "ttt", -1))
// // // // 	fmt.Println(strings.Count("ssss", "s"))
// // // // 	fmt.Println(strings.Repeat("str", 8))
// // // // 	fmt.Println(strings.ToLower("SHDAHSD"))
// // // // 	fmt.Println(strings.ToUpper("kjlsjdlJLHALSHDL"))
// // // // 	fmt.Println(strings.TrimSpace("   asdas "))
// // // // 	fmt.Println(strings.Trim("sstings", "s"))
// // // // 	fmt.Println(reflect.TypeOf(strings.Fields("sss      strings strings")))
// // // // 	fmt.Println(strings.Fields("   "))
// // // // 	fmt.Println(strings.Split("asldkja alsjdka alks lkk k", " "))
// // // // 	fmt.Println([]string{"ss", "dd"})
// // // // 	fmt.Println(strings.NewReader("strinssg"))
// // // // 	fmt.Println(strconv.IntSize)
// // // // 	fmt.Println(strconv.Itoa(077))
// // // // 	fmt.Println(strconv.Atoi("99"))
// // // // 	fmt.Println(time.Now().Weekday())
// // // // 	fmt.Printf("asdasd %09d\n", 11)
// // // // 	fmt.Println(time.Now().Format("02 Jan 2006 15:04"))
// // // // 	var p *int
// // // // 	// var i1 = 11
// // // // 	// p = &i1
// // // // 	// fmt.Println(p, &i1, p == &i1, i1, *p, i1 == *p)
// // // // 	*p = 2
// // // // }

// // // // func main() {
// // // // 	const (
// // // // 		a = iota
// // // // 		b = iota + iota
// // // // 	)

// // // // 	fmt.Println(a, b)
// // // // 	fmt.Println(runtime.GOOS)
// // // // 	p := fmt.Sprintf("sss %d", a)
// // // // 	fmt.Println(p)
// // // // 	if c := 1; c > 2 {
// // // // 		fmt.Println(c)
// // // // 	} else {
// // // // 		fmt.Println("ssss", c)
// // // // 	}
// // // // 	fmt.Println("")

// // // // 	count, err := fmt.Println(11)
// // // // 	fmt.Println(count, reflect.TypeOf(err))
// // // // 	test := 1
// // // // 	switch test {
// // // // 	case 0:
// // // // 		fmt.Println(1)
// // // // 	case 1, 2, 3:
// // // // 		fmt.Println(2)
// // // // 		fallthrough
// // // // 	default:
// // // // 		fmt.Println(3)
// // // // 	}
// // // // }

// // // func main() {
// // // 	// for i := 1; i < 5; i++ {
// // // 	// 	fmt.Println(i)
// // // 	// }

// // // 	// str := "abcABC"

// // // 	// for i := 0; i < len(str); i++ {
// // // 	// 	fmt.Printf("%d---->%c", i, str[i])
// // // 	// 	fmt.Println(str[i])
// // // 	// }

// // // 	// for j := 1; j <= 25; j++ {
// // // 	// 	fmt.Println(strings.Repeat("G", j))
// // // 	// }

// // // 	// for k := 1; k <= 100; k++ {
// // // 	// 	switch {
// // // 	// 	case k%5 == 0 && k%3 == 0:
// // // 	// 		fmt.Println("FizzBuzz")
// // // 	// 	case k%5 == 0:
// // // 	// 		fmt.Println("Buzz")
// // // 	// 	case k%3 == 0:
// // // 	// 		fmt.Println("Fizz")
// // // 	// 	default:
// // // 	// 		fmt.Println(k)
// // // 	// 	}
// // // 	// }

// // // 	// for l := 0; l < 10; l++ {
// // // 	// 	fmt.Println(strings.Repeat("*", 20))
// // // 	// }

// // // 	// for {
// // // 	// 	fmt.Println(1)
// // // 	// }

// // // 	// string := "love games"
// // // 	// for i, v := range string {
// // // 	// 	fmt.Printf("%d ===> %c \n", i, v)
// // // 	// }

// // // 	// for i := 0; i < 5; i++ {
// // // 	// 	var v int
// // // 	// 	fmt.Printf("%d ", v)
// // // 	// 	v = 5
// // // 	// }

// // // 	// for i := 0; ; i++ {
// // // 	// 	fmt.Println("Value of i is now:", i)
// // // 	// }

// // // 	// for i := 0; i < 3; {
// // // 	// 	fmt.Println("Value of i:", i)
// // // 	// }

// // // 	// s := ""
// // // 	// for s != "aaaaa" {
// // // 	// 	fmt.Println("Value of s:", s)
// // // 	// 	s = s + "a"
// // // 	// }

// // // 	// for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
// // // 	// 	s = i+1, j+1, s+"a" {
// // // 	// 	fmt.Println("Value of i, j, s:", i, j, s)
// // // 	// }

// // // 	// i := 5
// // // 	// for {
// // // 	// 	i = i - 1
// // // 	// 	fmt.Printf("The variable i is now: %d\n", i)
// // // 	// 	if i < 0 {
// // // 	// 		break
// // // 	// 	}
// // // 	// }

// // // 	// 	for i := 0; i <= 5; i++ {
// // // 	// 		for j := 0; j <= 5; j++ {
// // // 	// 			if j == 4 {
// // // 	// 				goto test
// // // 	// 			}
// // // 	// 			fmt.Printf("i is: %d, and j is: %d\n", i, j)
// // // 	// 		}
// // // 	// 	}
// // // 	// test:
// // // 	// 	fmt.Println("sjajlsdjl")
// // // 	// i := 0
// // // 	// for { //since there are no checks, this is an infinite loop
// // // 	// 	if i >= 3 {
// // // 	// 		break
// // // 	// 	}
// // // 	// 	//break out of this for loop when this condition is met
// // // 	// 	fmt.Println("Value of i is:", i)
// // // 	// 	i++
// // // 	// }
// // // 	// fmt.Println("A statement just after for loop.")
// // // 	// for i := 0; i < 7; i++ {
// // // 	// 	if i%2 == 0 {
// // // 	// 		continue
// // // 	// 	}
// // // 	// 	fmt.Println("Odd:", i)
// // // 	// }

// // // 	test := func() {
// // // 		fmt.Println("kl;s")
// // // 	}

// // // 	test()
// // // }
// // // func test2() {

// // // }

// // func main() {
// // 	fmt.Println("sss")

// // 	str := "string"
// // 	fmt.Println(&str)
// // 	fmt.Println(*&str)
// // }

// package main

// import (
// 	"errors"
// 	"fmt"
// 	"math"
// )

// func main() {
// 	// x1, y1 := test(2, 3, 4)
// 	// fmt.Println(x1, y1)
// 	// x2 := test2(2, 3, 4)
// 	// // fmt.Println(x2)
// 	// add1, mul1, red1 := namedReturnFn(2, 3)
// 	// add2, mul2, red2 := unNamedReturnFn(2, 3)
// 	// fmt.Println(add1, add2, mul1, mul2, red1, red2)
// 	s, err := mySqrt(-99)
// 	fmt.Println(s, err)
// }

// // func test(a int, b int, c int) (x int, y int) {
// // 	x = 2
// // 	return a + b, y
// // }

// // func test2(a int, b int, c int) int {
// // 	return a + b
// // }

// // func namedReturnFn(a int, b int) (add int, mul int, red int) {
// // 	return a + b, a * b, a - b
// // }

// // func unNamedReturnFn(a int, b int) (int, int, int) {
// // 	return a + b, a * b, a - b
// // }

// func mySqrt(f float64) (float64, error) {
// 	if f >= 0 {
// 		return math.Sqrt(f), nil
// 	}
// 	return float64(math.NaN()), errors.New("出错了")

// }

func main() {
	// a := 1
	// s := &a
	// test(s)
	// fmt.Println(a, *s)
	// fmt.Println(test(1, 2, 3, 3, 4))
	// arr := []int{2, 3, 3}
	// fmt.Println(test(9, 9, arr...))

	arr := []int{1, 2, 4}
	printSlice(arr...)
	// test(1, 2)

	// arr2 := {a: 1}
}

// func test(reply *int) {
// 	*reply = 2
// }

// func test(a int, b int, s ...int) int {
// 	sum := 0
// 	fmt.Println(len(s))
// 	for i := 0; i < len(s); i++ {
// 		sum += s[i]
// 	}
// 	fmt.Println(sum)
// 	return a + b + sum
// }

func printSlice(args ...int) {
	for _, v := range args {
		fmt.Println(v)
	}
}

// type Options struct {
// 	c int
// 	d string
// }

// func test(a, b, opt ...Options) {
// 	for _, v := range opt {
// 		fmt.Println(v)
// 	}
// }

func test2(values ...interface{}) {
	for _, value := range values {
		fmt.Println(value)
	}
}
