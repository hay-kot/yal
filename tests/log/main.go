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
	yal.Warn("This is a warn message")
	yal.Error("This is a error message")

	fmt.Println("\nGlobal Logger: Disabled")
	yal.Log.Level = yal.LevelDisable

	yal.Info("This is a info message")
	yal.Debug("This is a debug message")
	yal.Warn("This is a warn message")
	yal.Error("This is a error message")

	fmt.Println("\nGlobal Logger: No Color")
	yal.Log.Level = yal.LevelDebug
	yal.Log.Colors = false

	yal.Info("This is a info message")
	yal.Debug("This is a debug message")
	yal.Warn("This is a warn message")
	yal.Error("This is a error message")

	fmt.Println("\nGlobal Logger: Pretty Print")
	yal.Log.Colors = true

	yal.Info(pretty.Struct(Test{"John", 30}))
	yal.Info(pretty.Slice([]Test{{"John", 30}, {"Jane", 25}}))

}
