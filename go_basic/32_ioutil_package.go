package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func readFile() {
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))

}

func fileWrite() {
	temp := "hello, world"
	err := ioutil.WriteFile("b.txt", []byte(temp), os.ModePerm) // 写入之前会清空
	if err != nil {
		log.Fatal(err)
		return
	}
}

func readDir(dirName string) {
	filesInfo, err := ioutil.ReadDir(dirName) // 读取一层
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("nums of files:", len(filesInfo))
	for i, v := range filesInfo {
		fmt.Printf("the [%d] file named [%s], isDir? [%t]\n", i, v.Name(), v.IsDir())
	}
}

func recursiveReadDir(dirName string, level int) {
	filesInfo, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
		return
	}
	s := "|--"
	for i := 0; i < level; i++ {
		s = "| " + s
	}

	for _, v := range filesInfo {
		if v.IsDir() {
			subfile := dirName + "/" + v.Name()
			recursiveReadDir(subfile, level+1)
		}
		fmt.Printf("%s %s\n", s, v.Name())
	}
}

func main() {
	// 读取文件中的所有数据
	//readFile()
	// 写文件
	fileWrite()
	// 读取目录
	//readDir("C:\\Users\\95736\\Desktop\\My_project\\intelligence_learn")
	// 递归读取目录
	//recursiveReadDir("/Users/dawson/my_project/learnGo", 0)
}
