package pretty

import "encoding/json"

// Prettify takes in a struct or list and returns the JSON representation of the struct
// pretty printed with indentation.
func Struct(v interface{}) string {
	val, _ := json.MarshalIndent(v, "", "  ")
	return string(val)
}

// Prettify takes in a struct or list and returns the JSON representation of the struct
// pretty printed with indentation.
func Slice(v interface{}) string {
	val, _ := json.MarshalIndent(v, "", "  ")
	return string(val)
}
