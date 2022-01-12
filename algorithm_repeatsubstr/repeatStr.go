package main

import (
	"fmt"
)

// 求出字符串中的无重复字符、最长的子字符串的长度
func RepeatSubStr(s string) int {
	//lastOccurred := make(map[byte]int)
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	//for i, ch := range []byte(s) {
	for i, ch := range []rune(s) {
		// i ok lastI
		// map 中 已存在该字符记录
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}


func main() {
	fmt.Println(RepeatSubStr("acbsdfcdscc"))
	fmt.Println(RepeatSubStr("ac"))
	fmt.Println(RepeatSubStr(""))
	fmt.Println(RepeatSubStr("aaa"))
	fmt.Println(RepeatSubStr("张苏i啊猝死了出 阿城可是你才bhjjb开cj打车难阿斯nbkj"))

}

/*
	1.uint8类型，或者叫 byte 型，代表了ASCII码的一个字符
	2.rune类型，代表一个 UTF-8字符 === 相当于 go语言中的 char

英文字母和中文汉字在不同字符集编码下的字节数
英文字母：
	·字节数 : 1;编码：GB2312
	字节数 : 1;编码：GBK				***
	字节数 : 1;编码：GB18030
	字节数 : 1;编码：ISO-8859-1
	字节数 : 1;编码  UTF-8			***
	字节数 : 4;编码：UTF-16
	字节数 : 2;编码：UTF-16BE
	字节数 : 2;编码：UTF-16LE

中文汉字：
	字节数 : 2;编码：GB2312
	字节数 : 2;编码：GBK				***
	字节数 : 2;编码：GB18030
	字节数 : 1;编码：ISO-8859-1
	字节数 : 3;编码：UTF-8			***
	字节数 : 4;编码：UTF-16			***
	字节数 : 2;编码：UTF-16BE
	字节数 : 2;编码：UTF-16LE
*/
