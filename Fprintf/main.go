package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// func main() {
// 	fmt.Fprintf(os.Stdout, "%s\n", "hello world")
// 	buf := bufio.NewWriter(os.Stdout)
// 	fmt.Fprintf(buf, "%s\n", "hello world")
// 	buf.Flush()
// }

func main() {
	inputFile, inputError := os.Open("test.txt")
	if inputError != nil {
		fmt.Println("出错了")
		return
	}
	defer inputFile.Close()
	outputFile, outputError := os.OpenFile("testDes.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if outputError != nil {
		fmt.Println("打开目标问价出错")
		return
	}
	defer outputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	for {
		line, _, err := inputReader.ReadLine()
		if err == io.EOF {
			fmt.Println("到底了")
			break
		}
		_, err = outputWriter.WriteString(string(line[2:5]) + "\n")
		if err != nil {
			fmt.Println("写入文件出错")
			return
		}
	}
	outputWriter.Flush()
}
