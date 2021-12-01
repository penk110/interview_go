package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(sum(10, 20))
}

/*

case '0', '1', '2'

go build -x main.go

-x	查看编译与链接过程

➜  build_-x git:(master) ✗ go build -x main.go
WORK=/var/folders/wp/0kxj_ktn6xqg0qrrrty0jcr80000gn/T/go-build786854944
mkdir -p $WORK/b001/
cat >$WORK/b001/_gomod_.go << 'EOF' # internal
package main
import _ "unsafe"
//go:linkname __debug_modinfo__ runtime.modinfo
var __debug_modinfo__ = "0w\xaf\f\x92t\b\x02A\xe1\xc1\a\xe6\xd6\x18\xe6path\tcommand-line-arguments\nmod\tgithub.com/penk110/interview_go\t(devel)\t\n\xf92C1\x86\x18 r\x00\x82B\x10A\x16\xd8\xf2"
EOF
cat >$WORK/b001/importcfg << 'EOF' # internal
# import config
packagefile fmt=/work/sdk/go1.16/pkg/darwin_amd64/fmt.a
packagefile runtime=/work/sdk/go1.16/pkg/darwin_amd64/runtime.a
EOF
cd /work/gopath/src/github.com/penk110/interview_go/build_-x
/work/sdk/go1.16/pkg/tool/darwin_amd64/compile -o $WORK/b001/_pkg_.a -trimpath "$WORK/b001=>" -p main -lang=go1.16 -complete -buildid nRKyrpxlSQcOOxq64fgz/nRKyrpxlSQcOOxq64fgz -goversion go1.16 -D _/work/gopath/src/github.com/penk110/interview_go/build_-x -importcfg $WORK/b001/importcfg -pack -c=4 ./main.go $WORK/b001/_gomod_.go
/work/sdk/go1.16/pkg/tool/darwin_amd64/buildid -w $WORK/b001/_pkg_.a # internal
cp $WORK/b001/_pkg_.a /work/Library/Caches/go-build/0e/0e6c5d3894c6257dda6b73f33c8e8ba6055131d3c014d793e5fc891571ada5c0-d # internal
cat >$WORK/b001/importcfg.link << 'EOF' # internal
packagefile command-line-arguments=$WORK/b001/_pkg_.a
packagefile fmt=/work/sdk/go1.16/pkg/darwin_amd64/fmt.a
packagefile runtime=/work/sdk/go1.16/pkg/darwin_amd64/runtime.a
packagefile errors=/work/sdk/go1.16/pkg/darwin_amd64/errors.a
packagefile internal/fmtsort=/work/sdk/go1.16/pkg/darwin_amd64/internal/fmtsort.a
packagefile io=/work/sdk/go1.16/pkg/darwin_amd64/io.a
packagefile math=/work/sdk/go1.16/pkg/darwin_amd64/math.a
packagefile os=/work/sdk/go1.16/pkg/darwin_amd64/os.a
packagefile reflect=/work/sdk/go1.16/pkg/darwin_amd64/reflect.a
packagefile strconv=/work/sdk/go1.16/pkg/darwin_amd64/strconv.a
packagefile sync=/work/sdk/go1.16/pkg/darwin_amd64/sync.a
packagefile unicode/utf8=/work/sdk/go1.16/pkg/darwin_amd64/unicode/utf8.a
packagefile internal/bytealg=/work/sdk/go1.16/pkg/darwin_amd64/internal/bytealg.a
packagefile internal/cpu=/work/sdk/go1.16/pkg/darwin_amd64/internal/cpu.a
packagefile runtime/internal/atomic=/work/sdk/go1.16/pkg/darwin_amd64/runtime/internal/atomic.a
packagefile runtime/internal/math=/work/sdk/go1.16/pkg/darwin_amd64/runtime/internal/math.a
packagefile runtime/internal/sys=/work/sdk/go1.16/pkg/darwin_amd64/runtime/internal/sys.a
packagefile internal/reflectlite=/work/sdk/go1.16/pkg/darwin_amd64/internal/reflectlite.a
packagefile sort=/work/sdk/go1.16/pkg/darwin_amd64/sort.a
packagefile math/bits=/work/sdk/go1.16/pkg/darwin_amd64/math/bits.a
packagefile internal/oserror=/work/sdk/go1.16/pkg/darwin_amd64/internal/oserror.a
packagefile internal/poll=/work/sdk/go1.16/pkg/darwin_amd64/internal/poll.a
packagefile internal/syscall/execenv=/work/sdk/go1.16/pkg/darwin_amd64/internal/syscall/execenv.a
packagefile internal/syscall/unix=/work/sdk/go1.16/pkg/darwin_amd64/internal/syscall/unix.a
packagefile internal/testlog=/work/sdk/go1.16/pkg/darwin_amd64/internal/testlog.a
packagefile io/fs=/work/sdk/go1.16/pkg/darwin_amd64/io/fs.a
packagefile sync/atomic=/work/sdk/go1.16/pkg/darwin_amd64/sync/atomic.a
packagefile syscall=/work/sdk/go1.16/pkg/darwin_amd64/syscall.a
packagefile time=/work/sdk/go1.16/pkg/darwin_amd64/time.a
packagefile internal/unsafeheader=/work/sdk/go1.16/pkg/darwin_amd64/internal/unsafeheader.a
packagefile unicode=/work/sdk/go1.16/pkg/darwin_amd64/unicode.a
packagefile internal/race=/work/sdk/go1.16/pkg/darwin_amd64/internal/race.a
packagefile path=/work/sdk/go1.16/pkg/darwin_amd64/path.a
EOF
mkdir -p $WORK/b001/exe/
cd .
/work/sdk/go1.16/pkg/tool/darwin_amd64/link -o $WORK/b001/exe/a.out -importcfg $WORK/b001/importcfg.link -buildmode=exe -buildid=pLtAR0UFxnLX5oOElsIf/nRKyrpxlSQcOOxq64fgz/-aOQuBC0Nl3xTILAC4oC/pLtAR0UFxnLX5oOElsIf -extld=clang $WORK/b001/_pkg_.a
/work/sdk/go1.16/pkg/tool/darwin_amd64/buildid -w $WORK/b001/exe/a.out # internal
mv $WORK/b001/exe/a.out main
rm -r $WORK/b001/

*/
