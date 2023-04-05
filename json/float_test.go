package cjson

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"testing"
)

type TargetFloat struct {
	Id    int           `json:"id"`
	Price CustomFloat64 `json:"price"`
}

func TestUnmarshalJSONFloat64(t *testing.T) {
	f := func(inJSON string, exp float64) {
		t.Helper()

		var target TargetFloat
		err := json.Unmarshal([]byte(inJSON), &target)
		if err != nil {
			t.Fatal(err)
		}

		if target.Price.Float64 != exp {
			t.Fatalf("unexpected result for CustomFloat64.UnmarshalJSON(%v); got %v; want %v", inJSON, target, exp)
		}
	}

	f(`{"id":1,"price":2.58}`, 2.58)
	f(`{"id":2,"price":"2.58"}`, 2.58)
	f(`{"id":3,"price":7.15}`, 7.15)
	f(`{"id":4,"price":"7.15"}`, 7.15)
}

func TestMarshalJSONFloat64(t *testing.T) {

	jsonString := `[{"id":1,"price":2.58},
		{"id":2,"price":"2.58"},
		{"id":3,"price":7.15},
		{"id":4,"price":"7.15"}]`

	targets := []TargetFloat{}

	err := json.Unmarshal([]byte(jsonString), &targets)
	if err != nil {
		t.Fatal(err)
	}

	jsonMarshaled, err := json.Marshal(targets)
	if err != nil {
		t.Fatal(err)
	}

	jsonStringOK := `[{"id":1,"price":2.58},{"id":2,"price":2.58},{"id":3,"price":7.15},{"id":4,"price":7.15}]`

	if !cmp.Equal(string(jsonMarshaled), jsonStringOK) {
		t.Fatalf("got %s; want %s", string(jsonMarshaled), jsonString)
	}

}
