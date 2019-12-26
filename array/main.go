package main

import "fmt"

// func main() {
// 	var id [5]int

// 	fmt.Println(id)
// 	fmt.Println(id[len(id)-1])
// 	fmt.Println(len(id))
// 	for i, v := range id {
// 		fmt.Println(i, v)
// 	}

// }

// func main() {
// 	a := [...]string{"a", "b", "c", "d"}
// 	for i, v := range a {
// 		fmt.Println(i, v)
// 	}
// }

// func main() {
// 	var arr1 [5]int
// 	var arr2 = new([5]int)
// 	fmt.Println(arr1, arr2)
// 	fmt.Printf("%T    %T\n", arr1, arr2)
// 	// arr1 = *arr2
// 	arr2 = &arr1
// 	arr1[2] = 2
// 	fmt.Println(arr1, arr2)
// }

// func main() {
// 	var arr [5]int
// 	fmt.Println(arr)
// 	set(&arr)
// 	fmt.Println(arr)
// }

// func set(a *[5]int) {
// 	fmt.Println(a)
// 	a[2] = 2
// }

// func main() {
// 	var arr [5]int
// 	var arr1 = arr
// 	arr1 = arr
// 	arr[1] = 1
// 	fmt.Println(arr, arr1)
// }

// func main() {
// 	var arr [16]int
// 	for i := 0; i <= 15; i++ {
// 		arr[i] = i
// 	}
// 	fmt.Println(arr)
// }

// func main() {
// 	var arr = [...]int{1, 2, 3}
// 	var arr1 = []int{1, 2, 3, 4}
// 	fmt.Println(arr, arr1)
// 	fmt.Printf("%T", arr)
// }

// 切片
// func main() {
// 	var sli [3]int
// 	var sli2 [3]int
// 	fmt.Println(sli == sli2)
// 	sli[0] = 1
// 	fmt.Println(sli)
// }

// func main() {
// 	arr := [5]int{1, 2, 3, 4, 5}
// 	sli := arr[:2]
// 	sli2 := arr[2:]
// 	fmt.Println(sli, sli2)

// 	sli3 := arr[3:5]
// 	sli4 := sli3[0:cap(sli3)]
// 	fmt.Println(sli4)
// 	sli4[1] = 33
// 	fmt.Println(arr)
// }

// func main() {
// 	s := make([]int, 20, 50)
// 	fmt.Println(s)
// 	fmt.Println(cap(s))
// }

// func main() {
// 	s := make([]int, 5)
// 	s1 := s[2:4]
// 	fmt.Println(len(s), cap(s))
// 	fmt.Println(len(s1), cap(s1))
// }

// func main() {
// 	s1 := []byte{'p', 'o', 'e', 'm'}
// 	s2 := s1[2:]
// 	fmt.Printf("%c", s2)
// 	s2[1] = 't'
// 	fmt.Printf("%c     %c", s1, s2)
// 	fmt.Printf("%c", s1)
// 	fmt.Printf("%c", s2)
// }

// func main() {
// 	arr := [3]int{1, 2, 3}
// 	pre, next := arr[:2], arr[2:]
// 	fmt.Println(pre, next)
// }

// func main() {
// 	items := [...]int{10, 20, 30, 40, 50}
// 	for _, item := range items {
// 		item *= 2
// 	}
// 	fmt.Printf("%v", items)
// }

// func main() {
// 	arr := [5]int{1, 2, 3, 4, 5}
// 	slice := arr[0:3]
// 	fmt.Println(slice)
// 	slice1 := slice[0:4]
// 	fmt.Println(slice1)
// }

// func main() {
// 	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	a := arr[5:7]
// 	fmt.Println(a)
// 	a2 := a[0:1]
// 	fmt.Println(a2)
// }

// func main() {
// sliF := []int{1, 2, 3}
// sliT := make([]int, 2)
// n := copy(sliT, sliF)

// fmt.Println(n)
// fmt.Println(sliT, sliF)

// sli3 := []int{1, 2, 3}
// sli4 := append(sli3, 4, 5, 6)
// fmt.Println(sli3)
// fmt.Println(sli4)

// 	arr := [5]int{1, 2, 3, 4, 5}
// 	sli := arr[1:3]
// 	fmt.Println(sli)
// 	copy(sli[0:1], []int{9})

// 	fmt.Println(sli)
// }

// func main() {
// s := "strings"
// slice := []byte(s)
// fmt.Printf("%c\n", slice)
// slice2 := make([]int, 7)

// slice2 = append(slice2, 2, 2, 2)
// copy(slice2, s)
// fmt.Printf("%c\n", slice2)
// fmt.Println(len(slice2))
// r := []int32(s)
// fmt.Printf("%c\n", r)

// str := s[2:4]
// fmt.Println(str)

// 	str := "strings"
// 	fmt.Println(str[len(str)/2:] + str[:len(str)/2])
// }

// func main() {
// str := "aabccd"
// slice := str[:]
// var prev string
// for _, v := range slice {
// 	fmt.Println(prev)
// 	if prev != string(v) {
// 		fmt.Printf("%c", v)
// 	}
// 	prev = string(v)
// }

// 	arr := [5]byte{'a', 'a', 'c', 'c', 'e'}
// 	var prev byte
// 	for _, v := range arr {
// 		if prev != v {
// 			fmt.Printf("%c", v)
// 		}
// 		prev = v
// 	}
// }

// func main() {
// 	sli := []int{1, 2, 0, 4, 3}
// 	for i := 0; i < len(sli)-1; i++ {
// 		for j := i + 1; j < len(sli); j++ {
// 			if sli[i] > sli[j] {
// 				sli[i], sli[j] = sli[j], sli[i]
// 			}
// 		}
// 	}
// 	fmt.Println(sli)
// }

// func main() {
// 	fmt.Println(mapFun(func(i int) int {
// 		return 10 * i
// 	}, []int{1, 2, 3, 4, 5}))
// }

// func mapFun(fn func(int) int, sli []int) (ret []int) {
// 	for i := 0; i < len(sli); i++ {
// 		ret = append(ret, fn(sli[i]))
// 	}
// 	return ret
// }

// func main() {
// arr := [5]int{1: 2, 2: 2}
// fmt.Println(arr)

// 	arr := [3][2]int{
// 		[2]int{1, 2},
// 		[2]int{2, 3},
// 		[2]int{3, 4},
// 	}
// 	fmt.Println(arr)
// }

// func main() {
// 	a := 2
// 	switch a {
// 	case 1:
// 		fmt.Println("1")
// 	case 2:
// 		fmt.Println("2")
// 		fallthrough
// 	case 3:
// 		fmt.Println("3")
// 	default:
// 		fmt.Println("default")
// 	}
// }

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	// sli1 := arr[:3]
	// copy(sli1, sli2)
	// fmt.Println(sli1, sli2)
	// sli2 := append(sli1[1:], 333, 444, 55, 66)
	// fmt.Println(sli1, sli2, arr)
	// sli2[0] = 111
	// fmt.Println(sli1, sli2, arr)

	sliceTo := arr[:3]
	sliceFrom := []int{33, 44}
	l := copy(sliceTo, sliceFrom)
	fmt.Println(sliceTo, arr, l)

}
