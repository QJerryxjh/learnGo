package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Address struct {
	Road string
	Num  int
}

type Person struct {
	Name  string
	Add   *Address
	Hobby []string
}

// 序列化json并存入文件
// func main() {
// 	add := &Address{"njl", 1888}
// 	p := Person{
// 		Name:  "qje",
// 		Add:   add,
// 		Hobby: []string{"reading", "jogging"},
// 	}
// 	js, err := json.Marshal(p)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Printf("%s\n", js)
// 	fmt.Printf("%T\n", js)
// 	file, _ := os.OpenFile("person.json", os.O_CREATE|os.O_WRONLY, 0666)
// 	defer file.Close()
// 	enc := json.NewEncoder(file)
// 	err = enc.Encode(p)
// 	if err != nil {
// 		log.Println("Error in encoding json")
// 	}
// }

// 反序列化json
// func main() {
// 	file, err := os.Open("person.json")
// 	if err != nil {
// 		fmt.Println("出错了")
// 		return
// 	}
// 	defer file.Close()
// 	re, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	var um interface{}
// 	json.Unmarshal(re, &um)
// 	fmt.Println(um)
// 	m := um.(map[string]interface{})
// 	fmt.Println(m)
// 	for k, v := range m {
// 		switch vv := v.(type) {
// 		case string:
// 			fmt.Println("string", vv, k)
// 		case int:
// 			fmt.Println("int", vv, k)
// 		case interface{}:
// 			fmt.Println("interface", vv, k)
// 		default:
// 			fmt.Println("no", k)
// 		}
// 	}
// }

// func main() {
// 	var p Person
// 	file, err := os.Open("person.json")
// 	if err != nil {
// 		fmt.Println("出错了")
// 		return
// 	}
// 	defer file.Close()
// 	re, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	json.Unmarshal(re, &p)
// 	fmt.Println(p)
// }

func main() {
	var p Person
	file, err := os.Open("person.json")
	if err != nil {
		fmt.Println("出错了")
		return
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&p)
	fmt.Println(p)
}
