package fmt

import (
	"fmt"
	"strconv"
	"testing"
)

/*
经验总结
	1.不要用fmt用于高频的地方，容易产生性能问题，比如缓存的key生成：https://git.umlife.net/snippets/50，建议改成字符串拼接，如果有数字的存在，用strconv转化再拼接或者用bytes.Buffer：https://golang.org/pkg/bytes/#Buffer
	2.time这种获取时间的也会成为消耗性能的地方，特别是time转字符串
	3.线上性能profile，每个线上的程序都应该提供线上profile功能，参考链接：http://blog.golang.org/profiling-go-programs（用最下面的http部分）
	4.GC优化：http://www.cockroachlabs.com/blog/how-to-optimize-garbage-collection-in-go/
	5.用好 sync.Pool 来进行临时对象的复用
*/

var num = 658972981636108288
var numstr = "658972981636108288"

func BenchmarkFmt(b *testing.B) {
	for x := 0; x < b.N; x++ {
		_ = fmt.Sprintf("a_%d", num)
	}
}

func BenchmarkStrconvFmt(b *testing.B) {
	for x := 0; x < b.N; x++ {
		_ = "a_" + strconv.Itoa(num)
	}
}

func BenchmarkFmtString(b *testing.B) {
	for x := 0; x < b.N; x++ {
		_ = "a_" + numstr
	}
}
