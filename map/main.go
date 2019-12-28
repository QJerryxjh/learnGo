package main

import "fmt"

// func main() {
// 	var m map[int]string
// 	fmt.Printf("%T", m)
// 	m1 := map[int]string{}
// 	fmt.Println(m1 == nil)
// 	m = make(map[int]string)
// 	fmt.Println(m == nil)
// 	m2 := make(map[int]string)
// 	fmt.Println(m2 == nil)
// 	m[1] = "a"
// 	m1[1] = "a"
// 	fmt.Println(m, m1)
// 	fmt.Println(len(m))
// }

// 容量
// func main() {
// 	// m := make(map[int]string, 10)
// 	// fmt.Println(m, len(m))
// 	// v, ok := m[2]
// 	// fmt.Println(v, ok)

// 	m := map[int]bool{0: false}
// 	v, ok := m[2]
// 	fmt.Println(v, ok)
// 	fmt.Println(m)
// 	delete(m, 1)
// 	fmt.Println(m)
// 	delete(m, 0)
// 	fmt.Println(m)

// }

// 输出顺序
// func main() {
// 	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "a": "b"}
// 	for key := range capitals {
// 		fmt.Println("Map item: Capital of", key, "is", capitals[key])
// 	}
// }

// // 练习
// func main() {
// 	m := map[int]string{
// 		0: "sunday",
// 		1: "monday",
// 		2: "tuesday",
// 		3: "wednesday",
// 		4: "thursday",
// 		5: "friday",
// 		6: "saturday",
// 	}

// 	for _, v := range m {
// 		switch v {
// 		case "tuesday":
// 			fmt.Println("有Tuesday")
// 		case "hollday":
// 			fmt.Println("有hollday")
// 		}
// 	}
// }

// 练习
func main() {
	m := map[string]string{
		"water": "水",
		"juice": "橙汁",
	}

	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
}
