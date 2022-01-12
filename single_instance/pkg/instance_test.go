package pkg

import "testing"

/******test******/

func TestGetNormalIns(t *testing.T) {
	initInstance := GetNormalIns()
	t.Logf("init instance: %v\n", initInstance)

}

func TestDoubleLockIns(t *testing.T) {

}

func TestSyncIns(t *testing.T) {

}

/******benchmark******/

func BenchmarkGetNormalIns(b *testing.B) {
	b.Logf("benchmark normal lock\n")
	for i := 0; i < b.N; i++ {
		go func() {
			ins := GetNormalIns()
			b.Logf("id: %v, instance: %v\n", i, ins)
		}()
	}
}

func BenchmarkDoubleLockIns(b *testing.B) {
	b.Logf("benchmark double lock\n")
	for i := 0; i < b.N; i++ {
		go func() {
			ins := DoubleLockIns()
			b.Logf("id: %v, instance: %v\n", i, ins)
		}()
	}
}

func BenchmarkSynOnceIns(b *testing.B) {
	b.Logf("benchmark sync once lock\n")
	for i := 0; i < b.N; i++ {
		go func() {
			ins := SynOncecIns()
			b.Logf("id: %v, instance: %v\n", i, ins)
		}()
	}
}

/*

➜  pkg git:(master) ✗ go test -bench=. -benchmem
goos: darwin
goarch: amd64
BenchmarkGetNormalIns-12          674938              2730 ns/op             555 B/op          3 allocs/op
--- BENCH: BenchmarkGetNormalIns-12
    instance_test.go:22: benchmark normal lock
    instance_test.go:26: id: 1, instance: &{}
    instance_test.go:22: benchmark normal lock
    instance_test.go:26: id: 16, instance: &{}
    instance_test.go:26: id: 41, instance: &{}
    instance_test.go:26: id: 49, instance: &{}
    instance_test.go:26: id: 49, instance: &{}
    instance_test.go:26: id: 27, instance: &{}
    instance_test.go:26: id: 49, instance: &{}
    instance_test.go:26: id: 49, instance: &{}
        ... [output truncated]
BenchmarkDoubleLockIns-12        1000000              2509 ns/op             505 B/op          4 allocs/op
--- BENCH: BenchmarkDoubleLockIns-12
    instance_test.go:32: benchmark double lock
    instance_test.go:36: id: 1, instance: &{}
    instance_test.go:32: benchmark double lock
    instance_test.go:36: id: 100, instance: &{}
    instance_test.go:36: id: 100, instance: &{}
    instance_test.go:36: id: 100, instance: &{}
    instance_test.go:36: id: 100, instance: &{}
    instance_test.go:36: id: 100, instance: &{}
    instance_test.go:36: id: 100, instance: &{}
    instance_test.go:36: id: 100, instance: &{}
        ... [output truncated]
BenchmarkSynOnceIns-12           1000000              1896 ns/op             405 B/op          4 allocs/op
--- BENCH: BenchmarkSynOnceIns-12
    instance_test.go:42: benchmark sync once lock
    instance_test.go:46: id: 1, instance: &{}
    instance_test.go:42: benchmark sync once lock
    instance_test.go:46: id: 100, instance: &{}
    instance_test.go:46: id: 100, instance: &{}
    instance_test.go:46: id: 100, instance: &{}
    instance_test.go:46: id: 100, instance: &{}
    instance_test.go:46: id: 100, instance: &{}
    instance_test.go:46: id: 100, instance: &{}
    instance_test.go:46: id: 100, instance: &{}
        ... [output truncated]
PASS

ok single_instance/pkg 13.459s

*/
