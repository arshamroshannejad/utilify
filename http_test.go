package utilify

import (
	"context"
	"net/http"
	"testing"
)

func TestGetByCtx(t *testing.T) {
	key := "testKey"
	value := "testValue"
	req := &http.Request{}
	ctx := context.WithValue(context.Background(), key, value)
	req = req.WithContext(ctx)
	if result := GetByCtx(req, key); result != value {
		t.Fatalf("expected %s, got %s", value, result)
	}
}

func BenchmarkGetByCtx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		key := "testKey"
		value := "testValue"
		req := &http.Request{}
		ctx := context.WithValue(context.Background(), key, value)
		req = req.WithContext(ctx)
		GetByCtx(req, key)
	}
}
