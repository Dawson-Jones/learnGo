/*
字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。
比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，则返回原先的字符串。
你可以假设字符串中只包含大小写英文字母（a至z）。

输入："aabcccccaaa"
输出："a2b1c5a3"

输入："abbccd"
输出："abbccd"
解释："abbccd"压缩后为"a1b2c2d1"，比原字符串长度更长。
 */
package main

import (
	"fmt"
	"strconv"
)

func compressString(S string) string {
	if S==""{
		return S
	}
	res := ""
	start := 0
	alphabet := S[0]
	alphabetCount := 0
	for i, v:=range S{
		if byte(v)!=alphabet{
			alphabetCount = i-start
			res = res + string(alphabet) + strconv.Itoa(alphabetCount)
			start = i
			alphabet = byte(v)
		}
	}
	res = res + string(alphabet) + strconv.Itoa(len(S)-start)
	if len(S)<=len(res) {
		res = S
	}
	return res
}

func main() {
	str := "aabcccccaaa"
	str = "abbccd"
	res := compressString(str)
	fmt.Println(res)
	//a := 65
	//fmt.Println(string(a))
}
