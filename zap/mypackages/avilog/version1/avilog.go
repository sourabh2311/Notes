package avigo

import (
	"fmt"
	"os"
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Always remember that it is a pointer to SugaredLogger

type AviLogger struct {
	*zap.SugaredLogger
}

// Function to return SugaredLogger (using zap) with log rotation
// Read basics here: https://avinetworks.atlassian.net/wiki/spaces/EN/pages/787875011/Logger+for+Cloud+connector+Go

func NewLogger(filename string) *AviLogger {
	atom := zap.NewAtomicLevel()
	atom.SetLevel(zap.DebugLevel) // level has been set
	out, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	out.Close()     // was there just to check
	if err != nil { // unable to open file, logging to stdout
		cfg := zap.Config{
			Encoding:         "console",
			Level:            atom,
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey: "message",

				LevelKey:    "level",
				EncodeLevel: zapcore.CapitalLevelEncoder,

				TimeKey:    "time",
				EncodeTime: zapcore.ISO8601TimeEncoder,

				CallerKey:    "caller",
				EncodeCaller: zapcore.ShortCallerEncoder,

				LineEnding: "\n",
			},
		}
		logger, nerr := cfg.Build() // error shouldn't happen but still I am handling it
		var sugar *zap.SugaredLogger
		if nerr != nil {
			sugar = zap.NewExample().Sugar()
			sugar.Error("Was unable to create logger file!")
			sugar.Error("Was unable to create desired logger, running on simple logger!")
		} else {
			sugar = logger.Sugar()
			sugar.Error("Was unable to create logger file!")
		}
		defer sugar.Sync()
		return &AviLogger{sugar}
	} else {
		w := zapcore.AddSync(&lumberjack.Logger{
			// put you desired file path here!
			Filename: filename,
			// MaxAge is the maximum number of days to retain old log files based on the
			// timestamp encoded in their filename.  Note that a day is defined as 24
			// hours and may not exactly correspond to calendar days due to daylight
			// savings, leap seconds, etc. The default is not to remove old log files
			// based on age.
			MaxAge:  28, // days
			MaxSize: 20, // megabytes
			// MaxBackups is the maximum number of old log files to retain.  The default
			// is to retain all old log files (though MaxAge may still cause them to get
			// deleted.)
			MaxBackups: 3,
			Compress:   true, // Yes we want to compress the files
		})
		core := zapcore.NewCore(
			// could have used here zap.NewDevelopmentEncoderConfig() but still not getting line numbers...
			zapcore.NewConsoleEncoder(zapcore.EncoderConfig{ // same old settings available
				MessageKey: "message",

				LevelKey:    "level",
				EncodeLevel: zapcore.CapitalLevelEncoder,

				TimeKey:    "time",
				EncodeTime: zapcore.ISO8601TimeEncoder,

				CallerKey:    "caller",
				EncodeCaller: zapcore.FullCallerEncoder,

				LineEnding: "\n",
			}),
			w,
			atom,
		)
		logger := zap.New(core, zap.AddCaller())
		sugar := logger.Sugar()
		defer sugar.Sync()
		return &AviLogger{sugar}
	}
}

// Need to add our caller at given depth
func modifyArgs(depth int, keysAndValues []interface{}) []interface{} {
	x := fmt.Sprintf("Caller at given (%v) level: ", depth)
	_, n, o, p := runtime.Caller(depth + 1)
	if p == false {
		_, n, o, p = runtime.Caller(1)
	}
	n = n + fmt.Sprintf(":%v ", o)
	// Note that we can as well think of modifying the message instead of adding a new key value pair
	// x = x + n + "              " + msg
	var nkeysAndValues []interface{} = keysAndValues
	nkeysAndValues = append(nkeysAndValues, x, n)
	return nkeysAndValues
}

func (logger *AviLogger) InfowWithDepth(depth int, msg string, keysAndValues ...interface{}) {
	logger.Infow(msg, modifyArgs(depth, keysAndValues)...)
}

// TOCHECK
func (logger *AviLogger) WarnwWithDepth(depth int, msg string, keysAndValues ...interface{}) {
	logger.Warnw(msg, modifyArgs(depth, keysAndValues)...)
}

// TOCHECK
func (logger *AviLogger) ErrorwWithDepth(depth int, msg string, keysAndValues ...interface{}) {
	logger.Errorw(msg, modifyArgs(depth, keysAndValues)...)
}

// TOCHECK
func (logger *AviLogger) FatalwWithDepth(depth int, msg string, keysAndValues ...interface{}) {
	logger.Fatalw(msg, modifyArgs(depth, keysAndValues)...)
}

// ===================--Following to be ignored--=====================
// func Infow(sugar *zap.SugaredLogger, depth int, msg string, keysAndValues ...interface{}) {
// 	x := fmt.Sprintf("Caller at given (%v) level: ", depth)
// 	_, n, o, p := runtime.Caller(depth + 1)
// 	if p == false {
// 		_, n, o, p = runtime.Caller(1)
// 	}
// 	n = n + fmt.Sprintf(":%v ", o)
// 	x = x + n + "              " + msg
// 	// var nkeysAndValues []interface{} = keysAndValues
// 	// nkeysAndValues = append(nkeysAndValues, x, n)
// 	sugar.Infow(x, keysAndValues...)
// }
