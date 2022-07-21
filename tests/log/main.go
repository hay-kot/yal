package main

import (
	"fmt"

	"github.com/hay-kot/yal"
	"github.com/hay-kot/yal/pretty"
)

type Test struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	yal.Log.Level = yal.LevelDebug

	fmt.Println("\nGlobal Logger: Color")

	yal.Info("This is a info message")
	yal.Debug("This is a debug message")
	yal.Error("This is a error message")

	fmt.Println("\nGlobal Logger: No Color")
	yal.Log.Colors = false

	yal.Info("This is a info message")
	yal.Debug("This is a debug message")
	yal.Error("This is a error message")

	fmt.Println("\nGlobal Logger: Pretty Print")
	yal.Log.Colors = true

	yal.Debug(pretty.Struct(Test{"John", 30}))
	yal.Debug(pretty.Slice([]Test{{"John", 30}, {"Jane", 25}}))

}
