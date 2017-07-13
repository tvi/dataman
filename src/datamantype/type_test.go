package datamantype

import "testing"

// TODO: add negative cases
var validValues map[DatamanType][]interface{}

func init() {
	validValues = map[DatamanType][]interface{}{
		Document: []interface{}{
			map[string]interface{}{"a": 1, "b": "c"},
		},
		String: []interface{}{
			"a",
			"asdl;fkja;sldfj",
			`asd;fljasd;flkj`,
		},
		Int: []interface{}{
			1,
			100,
			float64(123),
			"1234",
		},
		Bool: []interface{}{
			true,
			false,
		},
		// TODO:
		//DateTime: []interface{}{
		//
		//},
	}
}

func TestDatamanTypeNormalization(t *testing.T) {
	for DatamanType, valueList := range validValues {
		t.Run(string(DatamanType), func(t *testing.T) {
			for i, val := range valueList {
				if _, err := DatamanType.Normalize(val); err != nil {
					t.Fatalf("%d DatamanType=%v val=%v err=%s", i, DatamanType, val, err)
				}
			}
		})
	}
}