package logger

import (
	"go.uber.org/zap"
)

func NewLogger(name, revision string) (*zap.Logger, error) {
	return &zap.Logger{}, nil
}

func NewDevelopmentLogger() (*zap.Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	return logger, nil
}

// func NewConfig() zap.Config {
// 	return zap.Config{
// 		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
// 		Development:       false,
// 		DisableCaller:     false,
// 		DisableStacktrace: false,
// 		Sampling: &zap.SamplingConfig{
// 			Initial:    0,
// 			Thereafter: 0,
// 			Hook:       func(zapcore.Entry, zapcore.SamplingDecision) { panic("not implemented") },
// 		},
// 		Encoding: "json",
// 		EncoderConfig: zapcore.EncoderConfig{
// 			MessageKey:       "message",
// 			LevelKey:         "severity",
// 			TimeKey:          "eventTime",
// 			NameKey:          "logger",
// 			CallerKey:        "caller",
// 			StacktraceKey:    "stacktrace",
// 			EncodeLevel:      func(zapcore.Level, zapcore.PrimitiveArrayEncoder) { panic("not implemented") },
// 			EncodeTime:       func(time.Time, zapcore.PrimitiveArrayEncoder) { panic("not implemented") },
// 			EncodeDuration:   func(time.Duration, zapcore.PrimitiveArrayEncoder) { panic("not implemented") },
// 			EncodeCaller:     func(zapcore.EntryCaller, zapcore.PrimitiveArrayEncoder) { panic("not implemented") },
// 			EncodeName:       func(string, zapcore.PrimitiveArrayEncoder) { panic("not implemented") },
// 			ConsoleSeparator: "",
// 		},
// 		OutputPaths:      []string{"stdout"},
// 		ErrorOutputPaths: []string{"stderr"},
// 		InitialFields: map[string]interface{}{
// 			"": nil,
// 		},
// 	}
// }
