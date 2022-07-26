# Yet Another Logger

This is a tiny dependency free library for very simple logging for applications and scripts where logging is a small part of the process and not used for robust metric collection. This should, for example, not be used for a large web server or background process. It's specifically developed for easily adding a logger into a CLI application or script.

If you do end up using this, you should vendor, or copy it into your code base.

**Features**

- Leveled Logger
- Optional Colored logging

**Example**

```shell
2022-07-21 14:41:56: [INFO] This is a info message
2022-07-21 14:41:56: [DEBUG] This is a debug message
2022-07-21 14:41:56: [ERROR] This is a error message
2022-07-21 14:41:56: [FATAL] This is a fatal message
```


## Usage

### Singleton Logger

The library can be used as an easy drop in logger that acts as a singleton. There are top level methods for Info, Debug, Error, Warn, and Fatal logging that call the top level Log variable which is an instance of the Logger struct.

```go
package main

import (
	"fmt"

	"github.com/hay-kot/yal"
)

func main() {
	yal.Log.Level = yal.LevelDebug // Set the global log level

	fmt.Println("\nGlobal Logger: Color")

	yal.Info("This is a info message")
	yal.Debug("This is a debug message")
	yal.Error("This is a error message")
    yal.Fatal("This is a fatal message")
}

```