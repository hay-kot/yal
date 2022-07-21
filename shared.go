package yal

var Log = Default()

func Info(v ...interface{}) {
	Log.Info(v...)
}

func Infof(format string, v ...interface{}) {
	Log.Infof(format, v...)
}

func Debug(v ...interface{}) {
	Log.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	Log.Debugf(format, v...)
}

func Error(v ...interface{}) {
	Log.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	Log.Errorf(format, v...)
}

func Fatal(v ...interface{}) {
	Log.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	Log.Fatalf(format, v...)
}
