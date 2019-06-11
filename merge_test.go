package dict

import (
	"bytes"
	"encoding/json"
	"testing"
)

var mergetests = []struct {
	q        string
	p        string
	expected string
}{
	{
		`{ "x": { "b":6, "d": 8}, "y": 5, "z": 9 }`,
		`{ "x": { "a":1, "b": 2, "c": 3}, "y": 7, "u": 4}`,
		`{"x": {"a": 1, "b": 6, "c": 3, "d": 8}, "y": 5, "u": 4, "z": 9 }`,
	},
	{
		`{}`,
		`{}`,
		`{}`,
	},
	{
		`{"b":2}`,
		`{"a":1}`,
		`{"a":1,"b":2}`,
	},
	{
		`{"a":{"x":3}}`,
		`{"a":{"x":2}}`,
		`{"a":{"x":3}}`,
	},
	{
		`{"a":{"y":2}}`,
		`{"a":{"x":1}}`,
		`{"a":{"x":1, "y":2}}`,
	},

	{
		`{"a":{"y":7, "z":8}}`,
		`{"a":{"x":1, "y":4}}`,
		`{"a":{"x":1, "y":7, "z":8}}`,
	},
}

func TestMerge(t *testing.T) {
	for _, d := range mergetests {
		var p map[string]interface{}
		if err := json.Unmarshal([]byte(d.p), &p); err != nil {
			t.Error(err)
			continue
		}

		var q map[string]interface{}
		if err := json.Unmarshal([]byte(d.q), &q); err != nil {
			t.Error(err)
			continue
		}

		var expected map[string]interface{}
		if err := json.Unmarshal([]byte(d.expected), &expected); err != nil {
			t.Error(err)
			continue
		}

		actual := merge(p, q)
		assert(t, expected, actual)
	}
}

func assert(t *testing.T, expected, actual map[string]interface{}) {
	expectedBuf, err := json.Marshal(expected)
	if err != nil {
		t.Error(err)
		return
	}
	actualBuf, err := json.Marshal(actual)
	if err != nil {
		t.Error(err)
		return
	}
	if bytes.Compare(expectedBuf, actualBuf) != 0 {
		t.Errorf("expected %s, got %s", string(expectedBuf), string(actualBuf))
		return
	}
}
