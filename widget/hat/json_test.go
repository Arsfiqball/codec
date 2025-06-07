package hat_test

import (
	"encoding/json"
	"testing"

	"github.com/Arsfiqball/codec/widget/hat"
)

func TestJSON(t *testing.T) {
	type sampleT0 struct {
		Child hat.Unit[string] `json:"child"`
	}

	type sampleT1 struct {
		ID    string             `json:"id"`
		Attr1 hat.Unit[string]   `json:"attr1"`
		Attr2 hat.Unit[sampleT0] `json:"attr2"`
		Attr3 hat.Unit[string]   `json:"attr3"`
	}

	t.Run("Marshal", func(t *testing.T) {
		example := sampleT1{
			ID:    "123",
			Attr1: hat.Value("value"),
			Attr2: hat.Value(sampleT0{
				Child: hat.Value("child"),
			}),
		}

		b, err := hat.MarshalJSON(example)
		if err != nil {
			t.Fatal(err)
		}

		// fmt.Println(string(b))

		if string(b) != `{"id":"123","attr1":"value","attr2":{"child":"child"}}` {
			t.Fatalf("unexpected result: %s", string(b))
		}
	})

	t.Run("Native marshal compatibility", func(t *testing.T) {
		example := sampleT1{
			ID:    "123",
			Attr1: hat.Value("value"),
			Attr2: hat.Value(sampleT0{
				Child: hat.Value("child"),
			}),
		}

		b, err := json.Marshal(example)
		if err != nil {
			t.Fatal(err)
		}

		// fmt.Println(string(b))

		if string(b) != `{"id":"123","attr1":"value","attr2":{"child":"child"},"attr3":{}}` {
			t.Fatalf("unexpected result: %s", string(b))
		}
	})

	t.Run("Native unmarshal compatibility", func(t *testing.T) {
		example := sampleT1{}

		err := json.Unmarshal([]byte(`{"id":"123","attr1":"value","attr2":{"child":"child"}}`), &example)
		if err != nil {
			t.Fatal(err)
		}

		// fmt.Printf("%+v\n", example)

		if example.ID != "123" {
			t.Fatalf("unexpected result: %+v", example)
		}

		if !example.Attr1.IsPresent() || example.Attr1.IsNull() || example.Attr1.Get() != "value" {
			t.Fatalf("unexpected result: %+v", example)
		}

		if !example.Attr2.IsPresent() || example.Attr2.IsNull() {
			t.Fatalf("unexpected result: %+v", example)
		}

		if !example.Attr2.Get().Child.IsPresent() || example.Attr2.Get().Child.IsNull() || example.Attr2.Get().Child.Get() != "child" {
			t.Fatalf("unexpected result: %+v", example)
		}

		if example.Attr3.IsPresent() || !example.Attr3.IsNull() {
			t.Fatalf("unexpected result: %+v", example)
		}
	})
}
