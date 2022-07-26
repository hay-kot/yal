package yal

type LogLevel int8

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelDisable
)

var Level = LogLevel(0)

func (l LogLevel) String() string {
	switch l {
	case 0:
		return "DEBUG"
	case 1:
		return "INFO"
	case 2:
		return "WARN"
	case 3:
		return "ERROR"
	case 4:
		return "FATAL"
	}
	return "DISABLE"
}

func (l LogLevel) ColorStart() string {
	switch l {
	case 0:
		return "\x1b[36m"
	case 1:
		return "\x1b[32m"
	case 2:
		return "\x1b[33m"
	case 3:
		return "\x1b[31m"
	case 4:
		return "\x1b[35m"
	}
	return ""
}

func (l LogLevel) ColorEnd() string {
	return "\x1b[0m"
}
