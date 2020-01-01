package main

import (
	"fmt"
)

// func main() {
// 	type s struct {
// 		a int
// 		b int
// 		c string
// 	}

// var s1 s // 值
// fmt.Println(s1)
// var s2 *s // 指针
// fmt.Println(s2)
// s2 = new(s)
// fmt.Println(s2)
// 	s2.a = 222

// 	s1.a = 2
// 	fmt.Println(s1)
// 	s1.c = "string"
// 	fmt.Println(s1)
// 	fmt.Println(s2)

// s3 := new(s)
// 	s3.a = 33
// 	s3.b = 44
// 	s3.c = "ccc"
// fmt.Println(s3)
// 	fmt.Println(s3.b, s3.c)

// s4 := &s{2, 3, "ss"}
// fmt.Println(s4)

// s5 := s{3, 4, "cc"}
// fmt.Println(s5)

// 	s6 := s{c: "ll"}
// 	s7 := &s{c: "ll"}
// 	fmt.Println(s6)
// 	fmt.Println(s7)
// }

// func main() {
// 	type Address struct {
// 		add string
// 	}

// 	type vcard struct {
// 		name  string
// 		addId *Address
// 		birth string
// 		img   string
// 	}

// 	addId := Address{"nj"}
// 	p := vcard{name: "qjerry", addId: &addId, birth: "1997", img: "三重刘德华.jpg"}
// 	fmt.Println(p)
// 	fmt.Printf("%v", p.addId)
// }

// func main() {
// 	type rect struct {
// 		width  int
// 		height int
// 		area   func()
// 	}

// }

// 工厂函数
// type File struct {
// 	id   int
// 	name string
// }

// func newFile(id int, name string) *File {
// 	if id < 0 {
// 		return nil
// 	}

// 	return &File{id, name}
// }

// func main() {
// 	f := newFile(9, "nine")
// 	fmt.Println(f, *f)
// 	fmt.Println(unsafe.Sizeof(File{}))
// 	fmt.Println(unsafe.Sizeof(f))
// 	fmt.Println(unsafe.Sizeof(*f))
// }

// 定义map和定义struct
// type Foo map[string]string

// func main() {
// 	u := new(Foo) // 返回一个指向nil的指针
// 	fmt.Println(u)
// 	fmt.Println(*u == nil)
// 	// (*u)["a"] = "a"
// 	// fmt.Println(u)

// 	x := make(Foo)
// 	fmt.Println(x)
// 	fmt.Println(x == nil)
// }

// tag
// type TagType struct { // tags
// 	field1 bool   "An important answer"
// 	field2 string "The name of the thing"
// 	field3 int    "How much there are"
// }

// func main() {
// 	tt := TagType{true, "Barak Obama", 1}
// 	for i := 0; i < 3; i++ {
// 		refTag(tt, i)
// 	}
// }

// func refTag(tt TagType, ix int) {
// 	ttType := reflect.TypeOf(tt)
// 	fmt.Println(ttType)
// 	fmt.Println(ix)
// 	ixField := ttType.Field(ix)
// 	fmt.Printf("%v\n", ixField.Tag)
// }

// 匿名字段
// type innerS struct {
// 	in1 int
// 	in2 int
// }
// type innerS2 struct {
// 	in3 int
// 	in4 int
// }

// type outerS struct {
// 	b      int
// 	c      float32
// 	int    // anonymous field
// 	string int
// 	innerS //anonymous field
// 	innerS2
// }

// func main() {
// 	o := new(outerS)
// 	o.int = 10
// 	o.string = 10
// 	fmt.Println(o)
// 	fmt.Println(o.in1)
// 	fmt.Println(o.int)
// }

// type S struct {
// 	f float32
// 	int
// 	string
// }

// func main() {
// 	s := S{3.0, 1, "ss"}
// 	fmt.Println(s)
// }

// type employee struct {
// 	salary float32
// }

// func (self *employee) giveRaise(per float32) {

// 	self.salary *= (100 + per) / 100
// }

// func main() {
// 	em := employee{20000}
// 	em.giveRaise(10)
// 	fmt.Println(em)
// }

// func (self employee) get() float32 {
// 	self.salary = 33
// 	return self.salary
// }
// func main() {
// 	em := employee{2}
// 	fmt.Println(em)
// 	em.get()
// 	fmt.Println(em)
// }

// 继承方法
type fa struct {
	name string
	age  int
}

func (self *fa) setAge(a int) {
	self.age = a
}

func (self *fa) Age() int {
	return self.age
}

type son struct {
	age int
	fa
}

func (self *son) Age() int {
	return self.age
}

func main() {
	s := son{10, fa{"qje", 30}}
	s.setAge(40)
	fmt.Println(s.fa.Age())
	fmt.Println(s)
}

// 练习
// type Base struct{}

// func (Base) Magic() {
// 	fmt.Println("base magic")
// }

// func (self Base) MoreMagic() {
// 	self.Magic()
// 	self.Magic()
// }

// type Voodoo struct {
// 	Base
// }

// func (Voodoo) Magic() {
// 	fmt.Println("voodoo magic")
// }

// func main() {
// 	v := new(Voodoo)
// 	v.Magic()
// 	v.MoreMagic()
// }

// 练习
// type Base struct {
// 	id int
// }

// func (self *Base) Id() int {
// 	return self.id
// }

// func (self *Base) SetId(i int) {
// 	self.id = i
// }

// type Person struct {
// 	Base
// 	FirstName string
// 	LastName  string
// }

// type Employee struct {
// 	Person
// 	salary float32
// }

// func main() {
// 	em := Employee{Person{Base{3}, "qje", "xu"}, 1000000}
// 	fmt.Println(em.id)
// 	fmt.Println(em.Id())
// }

// 练习
// type T struct {
// 	a int
// 	b float32
// 	c string
// }

// func main() {
// 	t := &T{7, -2.35, "abc\tdef"}
// 	fmt.Printf("%v\n", t)
// }

// func (self T) String() {
// }
