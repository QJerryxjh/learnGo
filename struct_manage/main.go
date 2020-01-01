package main

import (
	"fmt"
	"os"
)

// 函数版学生管理系统
// type student struct {
// 	id   int
// 	name string
// }

// var stu = make(map[int]student, 10)

// func main() {
// 	for {
// 		fmt.Println("进行操作：")
// 		fmt.Println("\t0:显示所有数据")
// 		fmt.Println("\t1:增加数据")
// 		fmt.Println("\t2:删除数据")
// 		fmt.Println("\t3:退出程序")
// 		fmt.Println("请选择:")
// 		var operation int
// 		var id int
// 		var name string
// 		fmt.Scanf("%v", &operation)
// 		switch operation {
// 		case 0:
// 			fmt.Println("所有数据")
// 			fmt.Printf("%#v\n", stu)
// 		case 1:
// 			fmt.Print("输入学号：")
// 			fmt.Scanln(&id)
// 			fmt.Print("输入姓名：")
// 			fmt.Scanln(&name)
// 			_, ok := stu[id]
// 			fmt.Println(ok)
// 			if ok == true {
// 				fmt.Println("已经存在学号", id)
// 			} else {
// 				stuItem := student{
// 					name: name,
// 					id:   id,
// 				}
// 				stu[id] = stuItem
// 				fmt.Println("添加成功！")
// 			}
// 		case 2:
// 			fmt.Print("输入学号：")
// 			fmt.Scanln(&id)
// 			_, ok := stu[id]
// 			if ok == false {
// 				fmt.Println("不存在学号", id)
// 			} else {
// 				delete(stu, id)
// 				fmt.Println("删除成功")
// 			}
// 		case 3:
// 			os.Exit(1)
// 		default:
// 			fmt.Println("给👴爪巴")
// 		}
// 	}
// }

// 结构体版学生管理系统

type student struct {
	id   int
	name string
}

type stuMan struct {
	stu map[int]student
}

var stm = stuMan{
	stu: make(map[int]student, 10),
}

func (s *stuMan) addItem(id int, name string) {
	stu := student{
		id:   id,
		name: name,
	}
	s.stu[id] = stu
}

func (s *stuMan) removeItem(id int) {
	delete(s.stu, id)
}

func (s *stuMan) reviseName(id int, name string) {
	studentItem, ok := s.stu[id]
	if ok == false {
		println("不存在此值")
	} else {
		studentItem.name = name
		s.stu[id] = studentItem
	}
}

func (s *stuMan) selectStu(id int) student {
	return s.stu[id]
}

func main() {
	for {
		fmt.Println("--------选择操作：----------")
		fmt.Println("\t0:显示所有数据")
		fmt.Println("\t1:增加数据")
		fmt.Println("\t2:修改数据")
		fmt.Println("\t3:删除数据")
		fmt.Println("\t4:查询单个数据")
		fmt.Println("\t5:退出程序")
		fmt.Println("请输入序号:")
		var operation int
		var id int
		var name string
		fmt.Scanf("%v", &operation)
		switch operation {
		case 0:
			fmt.Println("所有数据信息如下")
			for _, v := range stm.stu {
				fmt.Printf("学号：%v ——> 姓名：%v\n", v.id, v.name)
			}
		case 1:
			fmt.Print("输入新学生的学号：")
			fmt.Scanln(&id)
			fmt.Print("输入新学生的姓名：")
			fmt.Scanln(&name)
			_, ok := stm.stu[id]
			if ok == true {
				fmt.Println("该学号已经被占用：", id)
			} else {
				stuItem := student{
					name: name,
					id:   id,
				}
				stm.stu[id] = stuItem
				fmt.Println("新学生添加成功！")
			}
		case 2:
			fmt.Print("输入要修改信息的学生学号：")
			fmt.Scanln(&id)
			fmt.Print("输入要修改信息的学生姓名：")
			fmt.Scanln(&name)
			_, ok := stm.stu[id]
			if ok == false {
				fmt.Println("没有该学号的学生", id)
			} else {
				stm.reviseName(id, name)
				fmt.Println("修改学生信息成功")
			}
		case 3:
			fmt.Print("输入要删除学生的学号：")
			fmt.Scanln(&id)
			_, ok := stm.stu[id]
			if ok == false {
				fmt.Println("不存在该学号的学生", id)
			} else {
				stm.removeItem(id)
				fmt.Println("删除学生信息成功")
			}
		case 4:
			fmt.Print("输入要查询的学生学号：")
			fmt.Scanln(&id)
			_, ok := stm.stu[id]
			if ok == false {
				fmt.Println("不存在该学号", id)
			} else {
				ret := stm.selectStu(id)
				fmt.Println(ret)
			}
		case 5:
			fmt.Println("大爷，下次再来哟")
			os.Exit(1)
		default:
			fmt.Println("给👴爪巴")
		}
	}
}
