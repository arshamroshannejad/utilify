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

func TestAddStrToSlice(t *testing.T) {
	multiTests := []struct {
		name     string
		mapKey   string
		strMap   map[string][]string
		msg      string
		expected []string
	}{
		{
			name:   "Test for add string to slice...",
			mapKey: "seven",
			strMap: map[string][]string{
				"seven": {"four", "five", "six"},
			},
			msg:      "\"seven\" added to slice.",
			expected: []string{"four", "five", "six", "seven"},
		},
	}
	for _, test := range multiTests {
		t.Run(test.name, func(t *testing.T) {
			result := AddStrToSlice(test.strMap[test.mapKey], test.mapKey)
			if !StrInSlice(test.mapKey, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func BenchmarkAddStrToSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddStrToSlice([]string{"one", "two", "three"}, "four")
	}
}

func TestRmvStrInSlice(t *testing.T) {
	multiTests := []struct {
		name     string
		mapKey   string
		strMap   map[string][]string
		msg      string
		expected []string
	}{
		{
			name:   "Test for remove string from slice...",
			mapKey: "seven",
			strMap: map[string][]string{
				"seven": {"four", "five", "six", "seven"},
			},
			msg:      "\"seven\" removed from slice.",
			expected: []string{"four", "five", "six"},
		},
	}
	for _, test := range multiTests {
		t.Run(test.name, func(t *testing.T) {
			result := RmvStrInSlice(test.strMap[test.mapKey], test.mapKey)
			if StrInSlice(test.mapKey, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func BenchmarkRmvStrInSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RmvStrInSlice([]string{"one", "two", "three"}, "two")
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

func TestAddI64ToSlice(t *testing.T) {
	multiTests := []struct {
		name     string
		mapKey   int64
		intMap   map[int64][]int64
		msg      string
		expected []int64
	}{
		{
			name:   "Test for add integer to slice...",
			mapKey: 7,
			intMap: map[int64][]int64{
				7: {4, 5, 6},
			},
			msg:      "7 added to slice.",
			expected: []int64{4, 5, 6, 7},
		},
	}
	for _, test := range multiTests {
		t.Run(test.name, func(t *testing.T) {
			result := AddI64ToSlice(test.intMap[test.mapKey], test.mapKey)
			if !I64InSlice(test.mapKey, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func BenchmarkAddI64ToSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddI64ToSlice([]int64{1, 2, 3}, 4)
	}
}

func TestRmvI64InSlice(t *testing.T) {
	multiTests := []struct {
		name     string
		mapKey   int64
		intMap   map[int64][]int64
		msg      string
		expected []int64
	}{
		{
			name:   "Test for remove integer from slice...",
			mapKey: 7,
			intMap: map[int64][]int64{
				7: {4, 5, 6, 7},
			},
			msg:      "\"7\" removed from slice.",
			expected: []int64{4, 5, 6},
		},
	}
	for _, test := range multiTests {
		t.Run(test.name, func(t *testing.T) {
			result := RmvI64InSlice(test.intMap[test.mapKey], test.mapKey)
			if I64InSlice(test.mapKey, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func BenchmarkRmvI64InSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RmvI64InSlice([]int64{1, 2, 3}, 2)
	}
}
