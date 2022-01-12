package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

// byte to string
func byte2str(s []byte) (p string) {
	data := make([]byte, len(s))
	for i, v := range s {
		data[i] = v
	}
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&p))
	// 新字符串底层的数组: 设置长度和指定第一个元素的指针
	hdr.Len = len(s)
	hdr.Data = uintptr(unsafe.Pointer(&data[0]))

	return p
}

// string to byte
func string2byte(s string) (p []byte) {
	p = make([]byte, len(s))
	//for i, v := range []byte(s) {
	//	p[i] = v
	//}

	for i := 0; i < len(s); i++ {
		// 中间变量保存字符串中的每个元素
		c := s[i]
		p[i] = c
	}
	return p
}

// string to rune
func string2rune(s []byte) (p []rune) {
	p = make([]rune, len(s))
	for len(s) > 0 {
		// DecodeRune unpacks the first UTF-8 encoding in p and returns the rune and
		// its width in bytes. If p is empty it returns (RuneError, 0). Otherwise, if
		// the encoding is invalid, it returns (RuneError, 1). Both are impossible
		// results for correct, non-empty UTF-8.
		//
		// An encoding is invalid if it is incorrect UTF-8, encodes a rune that is
		// out of range, or is not the shortest possible UTF-8 encoding for the
		// value. No other validation is performed.

		// 看源码，大概意思是根据第一个字符判断是 每个字符占用多少个 byte
		// 循环中打印每次返回的 size 发现中文字符和英文字符的 size 不一样
		r, size := utf8.DecodeRune(s)
		//fmt.Printf("utf8.DecodeRune: %v %v\n", r, size)
		p = append(p, int32(r))
		// 取切片，为进项转换部分
		s = s[size:]
	}

	return p
}

// rune to string
func rune2string(s []rune) (p string) {
	temp := make([]byte, len(s))
	buffer := make([]byte, 3)
	for _, v := range s {
		// EncodeRune writes into p (which must be large enough) the UTF-8 encoding of the rune.
		// It returns the number of bytes written.
		// func EncodeRune(p []byte, r rune) int {}
		n := utf8.EncodeRune(buffer, v)
		temp = append(temp, buffer[:n]...)
	}
	p = string(temp)
	return p
}

func main() {
	r0 := byte2str([]byte("hello, 狗浪！"))
	fmt.Printf("byte2str: %v\n", r0)

	r1 := string2byte("hello, 狗浪！")
	fmt.Printf("string2byte: %s\n", r1)

	fmt.Printf("string2rune: %v\n",
		string2rune(
			[]byte("我要测试string转rune，rune内部会自动帮我将每个字符转成int32类型")))

	r2 := rune2string([]rune("我要测试rune转string"))
	fmt.Printf("rune2string: %v\n", r2)
}
