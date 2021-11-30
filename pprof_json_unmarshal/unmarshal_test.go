package unmarshal

import (
	"testing"
)

func TestMarshal(t *testing.T) {
	var (
		m   map[string]interface{}
		bs  []byte
		err error
	)

	m = map[string]interface{}{
		"books": []map[string]interface{}{
			{
				"book":  "七侠五义",
				"price": 99.9,
			},
			{
				"book":  "黑鹰传奇",
				"price": 77.9,
			},
		},
	}

	if bs, err = Marshal(m); err != nil {
		t.Error(err.Error())
		return
	}
	t.Logf("bs: %s", bs)

}

func BenchmarkUnMarshal(b *testing.B) {
	var (
		m   map[string]interface{}
		mr  map[string]interface{}
		bs  []byte
		err error
	)

	m = map[string]interface{}{
		"books": []map[string]interface{}{
			{
				"book":  "七侠五义",
				"price": 99.9,
			},
			{
				"book":  "黑鹰传奇",
				"price": 77.9,
			},
		},
	}
	bs, _ = Marshal(m)

	mr = make(map[string]interface{})
	for i := 0; i < b.N; i++ {
		if err = UnMarshal(bs, &mr); err != nil {
			b.Error(err.Error())
			return
		}
		b.Logf("bs: %#v\n", mr)
	}
}
