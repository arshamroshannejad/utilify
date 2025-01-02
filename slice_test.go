package utilify

import "testing"

func TestStrInSlice(t *testing.T) {
	multiTests := []struct {
		name     string
		mapKey   string
		strMap   map[string][]string
		msg      string
		expected bool
	}{
		{
			name:   "Test for check string value exists in slice...",
			mapKey: "one",
			strMap: map[string][]string{
				"one": {"one", "two", "three"},
			},
			msg:      "\"one\" exists in slice.",
			expected: true,
		},
		{
			name:   "Test for check string value does not exist in slice...",
			mapKey: "ana",
			strMap: map[string][]string{
				"ana": {"jack", "jill", "john"},
			},
			msg:      "\"ana\" does not exist in slice.",
			expected: false,
		},
	}
	for _, test := range multiTests {
		t.Run(test.name, func(t *testing.T) {
			result := StrInSlice(test.mapKey, test.strMap[test.mapKey])
			if result != test.expected {
				t.Errorf("expected %t, got %t", test.expected, result)
			}
		})
	}
}

func BenchmarkStrInSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StrInSlice("one", []string{"one", "two", "three"})
	}
}

func TestI64InSlice(t *testing.T) {
	multiTests := []struct {
		name     string
		mapKey   int64
		intMap   map[int64][]int64
		msg      string
		expected bool
	}{
		{
			name:   "Test for check integer value exists in slice...",
			mapKey: 1,
			intMap: map[int64][]int64{
				1: {1, 2, 3},
			},
			msg:      "1 exists in slice.",
			expected: true,
		},
		{
			name:   "Test for check integer value does not exist in slice...",
			mapKey: 4,
			intMap: map[int64][]int64{
				4: {5, 6, 7},
			},
			msg:      "4 does not exist in slice.",
			expected: false,
		},
	}
	for _, test := range multiTests {
		t.Run(test.name, func(t *testing.T) {
			result := I64InSlice(test.mapKey, test.intMap[test.mapKey])
			if result != test.expected {
				t.Errorf("expected %t, got %t", test.expected, result)
			}
		})
	}
}

func BenchmarkI64InSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		I64InSlice(1, []int64{1, 2, 3})
	}
}
