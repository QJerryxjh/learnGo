package main

import "fmt"

// scan 输入
// var (
// 	fn, ln, s string
// 	i         int
// 	f         float32
// 	input     = "55.1 / 22 / go"
// 	format    = "%f / %d / %s"
// )

// func main() {
// 	fmt.Println("请输入full name")
// 	fmt.Scanln(&fn, &ln)
// 	fmt.Println("全名是：", fn, ln)
// 	fmt.Sscanf(input, format, &f, &i, &s)
// 	fmt.Println(f, i, s)
// }

// buffered reader 读取输入
// func main() {
// 	inputReader := bufio.NewReader(os.Stdin)
// 	fmt.Println("请开始表演：")
// 	input, err := inputReader.ReadString('\t') // 读取到指定的字符，并且返回的字符串也会包含这个字符
// 	if err != nil {
// 		fmt.Println(err)
// 		fmt.Println("出错了")
// 	} else {
// 		fmt.Println(input)
// 	}
// }

// switch 版本的输入判断
// func main() {
// 	inputReader := bufio.NewReader(os.Stdin)
// 	fmt.Println("输入你的名字，开始表演吧：")
// 	input, err := inputReader.ReadString('\n')
// 	if err != nil {
// 		fmt.Println("笨比，就是你不会用")
// 		return
// 	}
// 	fmt.Println("你的名字是：", input)
// 	switch input {
// 	case "qjerry\n":
// 		fallthrough
// 	case "ww\n":
// 		fallthrough
// 	case "江浦彭于晏\n":
// 		fmt.Println("welcome", input)
// 	default:
// 		fmt.Println("爪巴")
// 	}
// }

// 练习1
// 编写一个程序，从键盘读取输入。当用户输入 'S' 的时候表示输入结束，这时程序输出 3 个数字：
// i) 输入的字符的个数，包括空格，但不包括 '\r' 和 '\n'
// ii) 输入的单词的个数
// iii) 输入的行数

var (
	nrchar, nrword, nrline int
)

// 1
// func main() {
// 	iR := bufio.NewReader(os.Stdin)
// 	fmt.Println("行吧，开始表演吧：")
// 	for {
// 		i, err := iR.ReadString('\n')
// 		if err != nil {
// 			fmt.Println("会不会用啊你")
// 			return
// 		}
// 		if i == "S\n" {
// 			fmt.Println("char number", nrchar)
// 			fmt.Println("word number", nrword)
// 			fmt.Println("line number", nrline)
// 			os.Exit(0)
// 		}
// 		numAdd(i)
// 	}

// }
// func numAdd(input string) {
// 	nrchar += len(input) - 1
// 	nrword += len(strings.Fields(input))
// 	nrline++
// }

// 2
// func main() {
// 	iR := bufio.NewReader(os.Stdin)
// 	fmt.Println("行吧，开始表演吧：")
// 	i, err := iR.ReadString('S')
// 	if err != nil {
// 		fmt.Println("会不会用啊你")
// 		return
// 	}
// 	ret := strings.Split(i, "\n")
// 	fmt.Println(ret)
// 	for _, v := range ret {
// 		numAdd(v)
// 	}
// 	// 最后一行的回车符并不算在输入内容之内，但是在计算个数的时候被减去了
// 	nrchar++
// 	fmt.Println(nrword, nrline, nrchar)
// }
// func numAdd(input string) {
// 	nrchar += len(input) - 1
// 	nrword += len(strings.Fields(input))
// 	nrline++
// }

// 练习
// 编写一个简单的逆波兰式计算器，它接受用户输入的整型数（最大值 999999）和运算符 +、-、*、/。
// 输入的格式为：number1 ENTER number2 ENTER operator ENTER --> 显示结果
func main() {
	var (
		num1, num2 int
		opt        string
	)
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	fmt.Scanln(&opt)
	fmt.Println(num1, num2, opt)
	switch opt {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)
	default:
		fmt.Println("会不会玩啊你")
	}
}
