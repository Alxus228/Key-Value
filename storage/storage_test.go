package storage

import (
	"testing"
)

func TestGet(t *testing.T) {
	testStorage := New()
	testStorage.data = map[interface{}]interface{}{
		"a": 1,
		1:   "a",
		"1": "b",
	}

	getTests := []struct {
		key      interface{}
		expected interface{}
	}{
		{"a", 1},
		{1, "a"},
		{"1", "b"},
	}

	for _, test := range getTests {
		val, err := testStorage.Get(test.key)
		if err != nil {
			t.Errorf("Expected: %v, but error occured: %v", test.expected, err)
		}
		if val != test.expected {
			t.Errorf("Expected: %v, but got: %v", test.expected, val)
		}
	}
}
