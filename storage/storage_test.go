package storage

import "testing"

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

func TestGetAll(t *testing.T) {
	testStorages := []storage{
		{map[interface{}]interface{}{
			1:   "a",
			"1": "b",
			"a": 1,
		}},
		{map[interface{}]interface{}{
			[3]float64{0.6, 0.9, 0.111}: [2]string{"a", "b"},
		}},
		{map[interface{}]interface{}{}},
	}

	getAllTests := []struct {
		lengthExpected int
		keysExpected   []interface{}
		valuesExpected []interface{}
	}{
		{
			3,
			[]interface{}{1, "1", "a"},
			[]interface{}{"a", "b", 1},
		},
		{
			1,
			[]interface{}{[3]float64{0.6, 0.9, 0.111}},
			[]interface{}{[2]string{"a", "b"}},
		},
		{
			0,
			nil,
			nil,
		},
	}

	for i, test := range getAllTests {
		getAllResult := testStorages[i].GetAll()
		if len(getAllResult) != test.lengthExpected {
			t.Errorf("Expected array length: %v, but got: %v", test.lengthExpected, len(getAllResult))
		}
		j := 0
		for key, value := range getAllResult {
			if key != test.keysExpected[j] {
				t.Errorf("Expected key: %v, but got: %v", test.keysExpected[j], key)
			}
			if value != test.valuesExpected[j] {
				t.Errorf("Expected value: %v, but got: %v", test.valuesExpected[j], value)
			}
			j++
		}
	}
}

func TestPut(t *testing.T) {
	testStorage := New()

	putTests := []struct {
		keyExpected   interface{}
		valueExpected interface{}
	}{
		{"1", 0},
		{"1", 1},
		{2, "a"},
		{"2", [2]string{"a", "b"}},
		{[1]byte{0x1}, [2]int{1, 2}},
	}

	for _, test := range putTests {
		testStorage.Put(test.keyExpected, test.valueExpected)

		if testStorage.data[test.keyExpected] != test.valueExpected {
			t.Errorf("Expected value: %v, but got: %v, and the key is %v",
				test.valueExpected,
				testStorage.data[test.keyExpected],
				test.keyExpected)
		}
	}
}
