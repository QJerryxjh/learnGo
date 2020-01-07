package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// func main() {
// 	if len(os.Args) > 1 {
// 		fmt.Println(strings.Join(os.Args[1:], "--"))
// 	}
// }

// var NewLine = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool

// const (
// 	Space   = " "
// 	Newline = "\n"
// )

// func main() {
// 	flag.PrintDefaults()
// 	flag.Parse()          // Scans the arg list and sets up flags
// 	fmt.Println(*NewLine) // 必须在 parse 之后才能起效果
// 	var s string = ""
// 	for i := 0; i < flag.NArg(); i++ {
// 		if i > 0 {
// 			s += Space
// 			if *NewLine { // -n is parsed, flag becomes true
// 				s += Newline
// 			}
// 		}
// 		s += flag.Arg(i)
// 	}
// 	os.Stdout.WriteString(s)
// }

// 用buffer读取文件
var addOrder = flag.Bool("n", false, "add a order")
var n = 1

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Println("读取文件出错")
			continue
		}
		defer f.Close()
		cat(bufio.NewReader(f))
	}
}

// func cat(r *bufio.Reader) {
// 	for {
// 		buf, err := r.ReadBytes('\n')
// 		if err == io.EOF {
// 			break
// 		}
// 		if *addOrder {
// 			fmt.Printf("第%d行:", n)
// 			n++
// 		}
// 		fmt.Fprintf(os.Stdout, "%s", buf)
// 	}
// 	return
// }

func cat(r *bufio.Reader) {
	var buf [512]byte
	for {
		nr, err := r.Read(buf[:])
		switch true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0:
			fmt.Println("结尾")
			return
		case nr > 0:
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", ew.Error())
			}
		default:
			fmt.Println(nr, err)
		}
	}
}
