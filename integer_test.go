package utilify

import "testing"

func TestI64ToStr(t *testing.T) {
	s := I64ToStr(666)
	if s != "666" {
		t.Fatalf("expected %s, got %s", "666", s)
	}
}

func BenchmarkI64ToStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		I64ToStr(666)
	}
}
