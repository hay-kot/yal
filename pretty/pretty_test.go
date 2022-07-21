package pretty

import "testing"

func Test_Struct(t *testing.T) {
	type Test struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	asPretty := Struct(Test{"John", 30})

	if asPretty != `{
  "name": "John",
  "age": 30
}` {
		t.Errorf("Expected %s, got %s", `{
  "name": "John",
  "age": 30
}`, asPretty)
	}
}

func Test_Slice(t *testing.T) {
	type Test struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	asPretty := Slice([]Test{{"John", 30}, {"Jane", 25}})

	if asPretty != `[
  {
    "name": "John",
    "age": 30
  },
  {
    "name": "Jane",
    "age": 25
  }
]` {
		t.Errorf("Expected \n%s, got \n%s", `[
  {
    "name": "John",
    "age": 30
  },
  {
    "name": "Jane",
    "age": 25
  }
]`, asPretty)
	}
}
