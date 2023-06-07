package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This need to go into a file - LearnCodeOnline.in"
	file, err := os.Create("./mylcogofile.txt")
	if err != nil {
		panic(err)
	}

	writer1(file, content)
	// len, err := io.WriteString(file, content)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(len)
	defer file.Close()

	reader("./mylcogofile.txt")

}

func writer1(file *os.File, content string) {
	_, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("")

}

func reader(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("This is the byte data that is present in it\n", databyte)

	fmt.Println("Converted to string:", string(databyte))

}
