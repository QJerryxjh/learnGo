package main

import (
	"fmt"
	"io"
	"os"
)

// 拷贝文件
func main() {
	copyfile("../test.js", "test.js")
	fmt.Println("拷贝成功")
}

func copyfile(from, to string) (wtitten int64, err error) {
	inputFrom, fromerr := os.Open(from)
	if fromerr != nil {
		fmt.Println("打开源文件出错")
		return
	}
	defer inputFrom.Close()
	outputTo, outputerr := os.OpenFile(to, os.O_WRONLY|os.O_CREATE, 0644)
	if outputerr != nil {
		fmt.Println("打开目标文件出错")
		return
	}
	defer outputTo.Close()
	return io.Copy(outputTo, inputFrom)
}
