package main

import (
	"fmt"
	"strings"
)

/**
给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，
在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的句子。

分隔时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。

输入:
s = "catsanddog"
wordDict = ["cat", "cats", "and", "sand", "dog"]
输出:
[
  "cats and dog",
  "cat sand dog"
]
*/

type myString string

func (str myString) in(dict []string) bool {
	for _, v := range dict {
		if string(str) == v {
			return true
		}
	}
	return false
}

//func reCur(minLen int, maxLen int, res *[]string, curRes *[]string, wordDict *[]string, s string) {
func reCur(res *[]string, curRes *[]string, wordDict *[]string, s string) {
	if s == "" {
		str := strings.Join(*curRes, " ")
		*res = append(*res, str)
		return
	}
	//if len(s)<minLen{
	//	return
	//}

	//for i := minLen; i <= maxLen; i++ {
	for i := 1; i <= len(s); i++ {
		frontString := s[:i]
		tailString := s[i:]
		if myString(frontString).in(*wordDict) {
			temp := make([]string, len(*curRes))
			copy(temp, *curRes)
			temp = append(temp, frontString)
			//reCur(minLen, maxLen, res, &temp, wordDict, tailString)
			reCur(res, &temp, wordDict, tailString)
		}
	}
}

func wordBreak(s string, wordDict []string) []string {
	var res []string
	var curRes []string
	//minStrLen := 1 << 30
	//maxStrLen := 0
	//for _, v := range wordDict {
	//	curLen := len(v)
	//	if curLen < minStrLen {
	//		minStrLen = curLen
	//	}
	//	if curLen > maxStrLen {
	//		maxStrLen = curLen
	//	}
	//}
	//reCur(minStrLen, maxStrLen, &res, &curRes, &wordDict, s)
	reCur(&res, &curRes, &wordDict, s)
	return res
}

func main() {
	// 超时, 把注释的地方打开, 是只循环wordDict中最小长度到最大长度, 不过不知道是哪里出问题了, 一直出错误
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	//s := "catsandog"
	//wordDict := []string{"cats", "dog", "sand", "and", "cat"}
	strLi := wordBreak(s, wordDict)
	for _, v := range strLi {
		fmt.Println(v)
	}
}
