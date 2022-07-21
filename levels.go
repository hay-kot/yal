package yal

type LogLevel int8

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelError
	LevelFatal
)

var Level = LogLevel(0)

func (l LogLevel) String() string {
	switch l {
	case 0:
		return "DEBUG"
	case 1:
		return "INFO"
	case 2:
		return "ERROR"
	case 3:
		return "FATAL"
	}
	return "UNKNOWN"
}

func (l LogLevel) ColorStart() string {
	switch l {
	case 0:
		return "\x1b[36m"
	case 1:
		return "\x1b[32m"
	case 2:
		return "\x1b[31m"
	case 3:
		return "\x1b[35m"
	}
	return ""
}

func (l LogLevel) ColorEnd() string {
	return "\x1b[0m"
}
