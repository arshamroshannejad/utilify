package utilify

import "testing"

func TestStrToI64(t *testing.T) {
	i, err := StrToI64("666")
	if err != nil {
		t.Fatal(err)
	}
	if i != 666 {
		t.Fatalf("expected %d, got %d", 666, i)
	}
}

func BenchmarkStrToI64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StrToI64("666")
	}
}
