package main

import "fmt"

/*
	字符组合：
		["1", "2", "4", "5"]

	递归？递归次数？
	怎么感觉可以构建一棵树？
*/

func combine(ts []string, fin string, skin int) []string {
	var (
		tss []string
		fti string
		i   int
		v   string
	)
	// 第一个取出来
	fti = ts[0]
	tss = []string{fti}

	var tmp = func(in []string, fin string, skin int) []string {
		var out []string
		for i, n := range in {
			if skin == i {
				continue
			}
			out = append(out, fin+n)
		}

		return out
	}
	for i, v = range ts[1:] {
		var til = fti + v
		tss = append(tss, til)
		var tss2 = tmp(ts[1:], til, i)

		tss = append(tss, tss2...)
	}

	return tss
}

func combine2(ts []string, skin, max int, out []string) []string {
	// 第一个取出来
	if len(ts) == 0 {
		return out
	}
	if len(ts) == 1 {
		out = append(out, ts[0])
		return out
	}

	for i, v := range ts {
		if i == skin {
			continue
		}
		out = append(out, ts[skin]+v)
	}
	out = append(out, combine2(ts, 0, max, out)...)

	return out
}

func main() {
	var (
		ts []string
	)
	/*
		默认入参是有序且是无重复元素的且长度大于1
		[]string{"1", "2", "4", "5"}

		1 12 124 125 1245 1254
		1 14 142 145 1425 1452
		1 15 152 154 1524 1542

		21 24 25 214 215

	*/
	ts = []string{"1", "2", "4", "5"}

	fmt.Println("source: ", ts)

	//fmt.Println(combine(ts, ts[0], 1))
	fmt.Println(combine2(ts, 0, len(ts), []string{}))

}
