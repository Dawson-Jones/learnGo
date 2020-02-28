package main

import (
	"fmt"
	"os"
)

func file() {
	fileInfo, err := os.Stat("C:/Users/95736/Desktop/My_project/Go/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fileInfo)
	// 文件名
	fmt.Println(fileInfo.Name())
	// 其他方法
	fileInfo.IsDir()
	fileInfo.Mode()
	fileInfo.ModTime()
	fileInfo.Size()
}
func createFile() {
	file, err := os.Create("./aa.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file)

}
func openFile() {
	// open 方式只读
	fp, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fp)

	// openFile
	fp, err = os.OpenFile("text.txt", os.O_RDONLY|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fp)
	// 关闭文件
	err = fp.Close()

}
func main() {
	file()
	// 创建文件
	createFile()
	// 打开文件

}
