

### 空接口转换
```go
// 空接空结构体   runtime2.go
type eface struct {
	_type *_type
	data  unsafe.Pointer
}
```
##### 查看空接口转换
```text
// 截取出部分汇编，查看对应源码是发现和动态接口类型转化类似都涉及到内存逃逸和寻址问题
runtime.convT32
```


### switch类型推断
##### 源码	compile/gc/swt.go

```go
func (s *typeSwitch) flush() {
	cc := s.clauses
	s.clauses = nil
	if len(cc) == 0 {
		return
	}

	sort.Slice(cc, func(i, j int) bool { return cc[i].hash < cc[j].hash })

	// Combine adjacent cases with the same hash.
	merged := cc[:1]
	for _, c := range cc[1:] {
		last := &merged[len(merged)-1]
		if last.hash == c.hash {
			last.body.AppendNodes(&c.body)
		} else {
			merged = append(merged, c)
		}
	}
	cc = merged

    // TODO: 降低复杂度到 o(log2n)
	binarySearch(len(cc), &s.done,
		func(i int) *Node {
			return nod(OLE, s.hashname, nodintconst(int64(cc[i-1].hash)))
		},
		func(i int, nif *Node) {
			// TODO(mdempsky): Omit hash equality check if
			// there's only one type.
			c := cc[i]
			nif.Left = nod(OEQ, s.hashname, nodintconst(int64(c.hash)))
			nif.Nbody.AppendNodes(&c.body)
		},
	)
}



```

哈希值有函数typehash计算出来 		800397251    269349216
```go
// typehash computes a hash value for type t to use in type switch statements.
func typehash(t *types.Type) uint32 {
	p := t.LongString()

	// Using MD5 is overkill, but reduces accidental collisions.
	h := md5.Sum([]byte(p))
	return binary.LittleEndian.Uint32(h[:4])
}
```

```
   0x0049 00073 (eface_escape.go:13)       MOVQ    CX, ""..autotmp_3+40(SP)
   0x004e 00078 (eface_escape.go:13)       JMP     80
   0x0050 00080 (eface_escape.go:13)       PCDATA  $0, $0
   0x0050 00080 (eface_escape.go:13)       TESTB   AL, (AX)
   0x0052 00082 (eface_escape.go:13)       MOVL    type.uint32+16(SB), AX
   0x0058 00088 (eface_escape.go:13)       MOVL    AX, ""..autotmp_5+8(SP)
   0x005c 00092 (eface_escape.go:13)       CMPL    AX, $-800397251				// TODO: 先比较 uint32
   0x0061 00097 (eface_escape.go:13)       JEQ     101
   0x0063 00099 (eface_escape.go:13)       JMP     164
   0x0065 00101 (eface_escape.go:13)       PCDATA  $0, $1
   0x0065 00101 (eface_escape.go:13)       LEAQ    type.uint32(SB), AX
   0x006c 00108 (eface_escape.go:13)       PCDATA  $0, $0
   0x006c 00108 (eface_escape.go:13)       PCDATA  $1, $0
   0x006c 00108 (eface_escape.go:13)       CMPQ    ""..autotmp_3+32(SP), AX
   0x0071 00113 (eface_escape.go:13)       JEQ     117
   0x0073 00115 (eface_escape.go:13)       JMP     160
   0x0075 00117 (eface_escape.go:13)       MOVL    $1, AX
   0x007a 00122 (eface_escape.go:13)       JMP     124
   0x007c 00124 (eface_escape.go:13)       MOVB    AL, ""..autotmp_4+3(SP)
   0x0080 00128 (eface_escape.go:13)       TESTB   AL, AL
   0x0082 00130 (eface_escape.go:13)       JNE     134
   0x0084 00132 (eface_escape.go:13)       JMP     156
   0x0086 00134 (eface_escape.go:16)       PCDATA  $0, $-1
   0x0086 00134 (eface_escape.go:16)       PCDATA  $1, $-1
   0x0086 00134 (eface_escape.go:16)       JMP     136
   0x0088 00136 (eface_escape.go:17)       PCDATA  $0, $0
   0x0088 00136 (eface_escape.go:17)       PCDATA  $1, $0
   0x0088 00136 (eface_escape.go:17)       MOVL    $3232, "".ui+4(SP)
   0x0090 00144 (eface_escape.go:13)       JMP     146
   0x0092 00146 (eface_escape.go:13)       PCDATA  $0, $-1
   0x0092 00146 (eface_escape.go:13)       PCDATA  $1, $-1
   0x0092 00146 (eface_escape.go:13)       MOVQ    48(SP), BP
   0x0097 00151 (eface_escape.go:13)       ADDQ    $56, SP
   0x009b 00155 (eface_escape.go:13)       RET
   0x009c 00156 (eface_escape.go:13)       JMP     158
   0x009e 00158 (eface_escape.go:13)       JMP     146
   0x00a0 00160 (eface_escape.go:13)       PCDATA  $0, $0
   0x00a0 00160 (eface_escape.go:13)       PCDATA  $1, $0
   0x00a0 00160 (eface_escape.go:13)       XORL    AX, AX
   0x00a2 00162 (eface_escape.go:13)       JMP     124
   0x00a4 00164 (eface_escape.go:13)       PCDATA  $1, $1
   0x00a4 00164 (eface_escape.go:13)       CMPL    AX, $-269349216	// TODO: 后比较uint16
   0x00a9 00169 (eface_escape.go:13)       JEQ     173

```



##### 结论

1.空接口转化损耗主要在于内存逃逸、创建eface、堆内设置垃圾回收逻辑等

2.使用fmt格式化输出等时涉及到空接口类型转化，从而会造成性能损耗

3.Go语言对单字节做了特殊优化

```go

```



