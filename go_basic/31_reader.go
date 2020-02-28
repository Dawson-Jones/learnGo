package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func ioStart() {
	r := strings.NewReader("Hello, Reader")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n: %v, err: %v, b: %v\n", n, err, b)
		//fmt.Printf("b[:n] = %s\n", b[:8])  // 比较一下
		fmt.Printf("b[:n] = %s\n", b[:n])
		if err == io.EOF { // end of file
			break
		}
	}

}

func readLocalFile() {
	// 打开 文件
	fp, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 关闭文件
	defer fp.Close()
	// 读取数据
	bs := make([]byte, 4, 4)
	// 第一次读取
	n, err := fp.Read(bs)
	fmt.Println(err)
	fmt.Println(n)
	fmt.Println(bs)
	fmt.Println(string(bs))

}

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func writeData() {
	//fp, err := os.Open("a.txt")
	fp, err := os.OpenFile("a.txt", os.O_CREATE, os.ModePerm)
	//fp, err := os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND, os.ModePerm)  // 追加模式写数据
	HandleErr(err)
	defer fp.Close()
	bs := []byte{65, 66, 67, 68, 69, 70} // A,B,C,D,E,F
	n, err := fp.Write(bs)
	HandleErr(err)
	fmt.Println(n)

	// 直接写字符串
	n, err = fp.WriteString("Hello, World")
	HandleErr(err)
	fmt.Println(n)
	// 没有关闭之前, 第二次写的不会覆盖第一次写的, 而是追加

}

func copyFile(src string, dest string) (int, error) {
	fp1, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	fp2, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer fp1.Close()
	defer fp2.Close()
	bs := make([]byte, 1024, 1024)
	total := 0
	for {
		n, err := fp1.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("copy complete...")
			break
		} else if err != nil {
			fmt.Println("something wrong")
			return total, err
		}
		total += n
		_, _ = fp2.Write(bs[:n])
	}
	return total, nil

}

func copyFile2(src string, dest string) (int64, error) {
	fp1, err := os.Open(src)
	HandleErr(err)
	fp2, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	HandleErr(err)
	defer fp1.Close()
	defer fp2.Close()
	return io.Copy(fp2, fp1)
}

func copyFile3(src, dest string) (int, error) {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return 0, err
	}
	err = ioutil.WriteFile(dest, data, os.ModePerm)
	if err != nil {
		return 0, err
	}
	HandleErr(err)
	return len(data), nil

}

func BreakpointContinuation() {
	srcFile := "C:\\Users\\95736\\Desktop\\My_project\\Go\\vim_keyboard.gif"
	//destFile := srcFile[strings.LastIndex(srcFile, "\\")+1:]
	destFile := "vim_copy.gif"
	fmt.Println("destFile: ", destFile)
	tempFile := destFile + "temp.txt"
	fmt.Println("tempFile: ", tempFile)

	fp1, err := os.Open(srcFile)
	HandleErr(err)
	fp2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	HandleErr(err)
	tempfp, err := os.OpenFile(tempFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	HandleErr(err)
	defer fp1.Close()
	defer fp2.Close()

	// 读取临时数据再seek
	tempN, err := tempfp.Seek(0, io.SeekStart) // 相当于移动光标, 从开始偏移0个开始读
	HandleErr(err)
	fmt.Println("first return of seek function is", tempN)
	bs := make([]byte, 100, 100)
	n, _ := tempfp.Read(bs)
	countStr := string(bs[:n])
	count, _ := strconv.ParseInt(countStr, 10, 64) // 第一次temp文件没有东西, 所以不能转化, 不处理错误count默认是0
	fmt.Println("count: ", count)

	// 设置读写的位置
	fp1.Seek(count, io.SeekStart)
	fp2.Seek(count, io.SeekStart)
	data := make([]byte, 1024, 1024)
	total := int(count)

	// 复制文件
	for {
		n2, err := fp1.Read(data)
		if err == io.EOF || n2 == 0 {
			fmt.Println("file is download complete")
			_ = tempfp.Close()
			_ = os.Remove(tempFile)
			break
		}
		n3, err := fp2.Write(data[:n2])
		total += n3
		// 复制的总量放到临时文件中
		_, _ = tempfp.WriteString(strconv.Itoa(total))
		//_, _ = tempfp.WriteString(string(total))  // 这个好像是不行

		//假装断电
		//if total > 5000{
		//	fmt.Println(total)
		//	panic("断电了")
		//}
	}
}

func bufIO() {
	fp, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer fp.Close()
	// 创建reader对象
	b1 := bufio.NewReader(fp)
	//p := make([]byte, 1024)
	//n1, err := b1.Read(p)
	//fmt.Println(n1)
	//fmt.Println(string(p[:n1]))

	// ReadString(), 指定停下来的字符 byte类型就是char类型
	for {
		c1, err := b1.ReadString('\n') // 读取一行
		if err == io.EOF {
			fmt.Println("read complete")
			break
		}
		fmt.Println(c1)

	}
}

func scan() {
	//var s string
	//_, _= fmt.Scanln(&s)  // 读取到空格之前
	//fmt.Println(s)

	bufObj := bufio.NewReader(os.Stdin)
	s, _ := bufObj.ReadString('\n')
	fmt.Println(s)
}

func bufWrite() {
	fp, err := os.OpenFile("cc.txt", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer fp.Close()
	w := bufio.NewWriter(fp)
	n, err := w.WriteString("hello, world")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("the nums of data writing:", n)
	// 此时并没有写入完成, 而是写入了buf的缓冲区中
	_ = w.Flush() // 刷新缓冲区

	// 当写入的数据比缓冲区大(缓冲区默认4096Byte) 一个byte能写入一个char
	for i := 0; i < 1000; i++ {
		n, err = w.WriteString(fmt.Sprintf("%d times write\n", i))
	}
	_ = w.Flush()
}

func main() {
	//ioStart()
	//readLocalFile()
	//writeData()
	// 复制文件
	//srcFile := "a.txt"
	//destFile := "b.txt"
	//n, err := copyFile(srcFile, destFile)
	//fmt.Println(n, err)
	// 断点续传
	//BreakpointContinuation()

	// 高效读写bufio
	//bufIO()

	// 读取键盘输入
	//scan()

	// buf写操作
	bufWrite()
}
