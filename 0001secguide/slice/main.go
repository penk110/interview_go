package main

/*
	内存管理：
		切片长度校验
*/

// decode1: 未判断切片长度直接操作，可能会出现 index out of range
func decode1(data []byte) bool {
	if data[0] == 'A' {
		return true
	}

	// out of range
	if data[9999] == 'E' {
		return true
	}

	return false
}

// decode2: 检查长度
func decode2(data []byte) bool {

	if len(data) >= 3 && data[2] == 'a' {
		return true
	}

	return false
}

func main() {
	// good
	decode1([]byte("Aa"))
	// bad
	// panic: runtime error: index out of range [9999] with length 5
	//decode1([]byte("bad"))

	decode2([]byte("Aaa"))
}
