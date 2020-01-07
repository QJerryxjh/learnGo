package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 读文件1
// func main() {
// 	inputFile, inputError := os.Open("../test.js")
// 	if inputError != nil {
// 		fmt.Printf("读取文件错误\n" + "这个文件存在吗\n" + "你真的有权限吗\n")
// 		return
// 	}
// 	defer inputFile.Close()
// 	inputReader := bufio.NewReader(inputFile)
// 	for {
// 		inputString, readerError := inputReader.ReadString('\n')
// 		fmt.Printf("读取的内容是：%s", inputString)
// 		if readerError == io.EOF {
// 			return
// 		}
// 	}
// }

// 读文件2
// 以二进制的形式读取文件
// func main() {
// 	inputFile, inputError := os.Open("../test.js")
// 	if inputError != nil {
// 		fmt.Println("打开文件出错")
// 		return
// 	}
//	defer inputFile.Close()
// 	inputReader := bufio.NewReader(inputFile)
// 	buf := make([]byte, 200)

// n 为读取到的字节数，把读到的字节存在 buf 内
// 	n, err := inputReader.Read(buf)
// 	if err != nil {
// 		fmt.Println("读取出错")
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(n)
// 	fmt.Println(buf)
// }

// 读文件3
// 按照列读取
// func main() {
// 	file, err := os.Open("../test.txt")
// 	if err != nil {
// 		fmt.Println("出错了")
// 		return
// 	}
// 	defer file.Close()
// 	var col1, col2, col3 []string
// 	for {
// 		var v1, v2, v3 string
// 		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
// 		if err != nil {
// 			break
// 		}
// 		col1 = append(col1, v1)
// 		col2 = append(col2, v2)
// 		col3 = append(col3, v3)
// 	}
// 	fmt.Println(col1, col2, col3)
// }

// 读文件4
// 文件 products.txt 的内容如下：
// "The ABC of Go";25.5;1500
// "Functional Programming with Go";56;280
// "Go for It";45.9;356
// "The Go Way";55;500
// 每行的第一个字段为 title，第二个字段为 price，第三个字段为 quantity。内容的格式基本与 示例 12.3c 的相同，除了分隔符改成了分号。请读取出文件的内容，创建一个结构用于存取一行的数据，然后使用结构的切片，并把数据打印出来。
// type Book struct {
// 	title    string
// 	price    float64
// 	quantity int
// }

// func main() {
// 	books := make([]Book, 0)
// 	inputFile, inputError := os.Open("products.txt")
// 	if inputError != nil {
// 		fmt.Println("打开文件出错")
// 		return
// 	}
// 	defer inputFile.Close()
// 	inputReader := bufio.NewReader(inputFile)
// 	for {
// 		line, err := inputReader.ReadString('\n')
// 		if err == io.EOF {
// 			break
// 		}
// 		line = string(line[:len(line)-1])
// 		strSl := strings.Split(line, ";")
// 		book := new(Book)
// 		book.title = strSl[0]
// 		book.price, err = strconv.ParseFloat(strSl[1], 32)
// 		if err != nil {
// 			fmt.Println("出错了文件", err)
// 		}
// 		book.quantity, err = strconv.Atoi(strSl[2])
// 		if err != nil {
// 			fmt.Println("文件出错")
// 		}
// 		books = append(books, *book)
// 	}
// 	for _, v := range books {
// 		fmt.Println(v)
// 	}
// }

// 写文件1
// func main() {
// 	inputFile := "../test.js"
// 	outFile := "../test_1.js"
// 	buf, err := ioutil.ReadFile(inputFile)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "file error: %s\n", err)
// 	}
// 	fmt.Println(string(buf))
// 	err = ioutil.WriteFile(outFile, buf, 0644)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

// 写文件2
// func main() {
// 	只写模式，没有就创建
// 	outputFile, outputError := os.OpenFile("write.txt", os.O_WRONLY|os.O_CREATE, 0666)
// 	if outputError != nil {
// 		fmt.Println("读文件出错")
// 		return
// 	}
// 	defer outputFile.Close()
// 	outWriter := bufio.NewWriter(outputFile)
// 	outputString := "hello world\n"
// 	for i := 0; i < 10; i++ {
// 		outWriter.WriteString(outputString)
// 	}
// 	fmt.Fprintf(outputFile, "ssss\n")
// 	outWriter.Flush() // 把缓冲中的数据写入下层的接口
// }

// func main() {
// 	os.Stdout.WriteString("hello world\n")
// 	f, _ := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY, 0)
// 	defer f.Close()
// 	f.WriteString("hello world in a file\n")
// }

// 练习
// 请给这个结构编写一个 save 方法，将 Title 作为文件名、Body作为文件内容，写入到文本文件中。
// 再编写一个 load 函数，接收的参数是字符串 title，该函数读取出与 title 对应的文本文件。
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() {
	outputFile, outputError := os.OpenFile(p.Title, os.O_RDWR|os.O_CREATE, 0644)
	fmt.Println(1)
	if outputError != nil {
		fmt.Println("打开文件出错")
		return
	}
	defer outputFile.Close()
	outputWrite := bufio.NewWriter(outputFile)
	n, err := outputWrite.Write(p.Body)
	fmt.Println(err)
	fmt.Println(n)
	fmt.Println(p.Body)
	outputWrite.Flush()
}

func (p *Page) load(filename string) (str string, err error) {
	inputFile, inputError := os.Open(filename)
	if inputError != nil {
		fmt.Println("打开文件出错")
		return str, inputError
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
		line, err := inputReader.ReadString('\n')
		str += line
		if err == io.EOF {
			return str, nil
		}
	}
}

func main() {
	p1 := new(Page)
	p1.Title = "p1.txt"
	p1.Body = []byte{'a', 'b', 'c'}
	p1.save()
	s, err := p1.load("p1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}
