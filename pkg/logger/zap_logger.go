package logger

import (
	"context"
	"os"
	"strings"

	"github.com/lugondev/send-sen/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLogger is an implementation of the Logger interface using Zap.
type ZapLogger struct {
	sugar *zap.SugaredLogger
}

// NewZapLogger creates a new Zap logger instance based on configuration.
func NewZapLogger(cfg config.Config) (Logger, error) {
	var level zapcore.Level
	switch strings.ToLower(cfg.Log.Level) {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	case "panic":
		level = zapcore.PanicLevel
	default:
		level = zapcore.InfoLevel // Default to Info
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // Human-readable time format
	encoderConfig.TimeKey = "timestamp"                   // Standardize time key
	encoderConfig.LevelKey = "level"
	encoderConfig.NameKey = "logger"
	encoderConfig.CallerKey = "caller"
	encoderConfig.MessageKey = "message"
	encoderConfig.StacktraceKey = "stacktrace"

	var encoder zapcore.Encoder
	if strings.ToLower(cfg.Log.Format) == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig) // Default to JSON
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(os.Stdout), // Log to standard output
		level,
	)

	// Add caller information
	// Use AddCallerSkip(1) to report the caller of the logger methods (Debug, Info, etc.)
	// instead of the logger wrapper itself.
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))
	sugar := logger.Sugar()

	sugar.Info("Zap logger initialized successfully.")

	return &ZapLogger{sugar: sugar}, nil
}

// Debug logs a message at DebugLevel.
func (l *ZapLogger) Debug(ctx context.Context, args ...any) {
	l.sugar.Debug(args...)
}

// Info logs a message at InfoLevel.
func (l *ZapLogger) Info(ctx context.Context, args ...any) {
	l.sugar.Info(args...)
}

// Warn logs a message at WarnLevel.
func (l *ZapLogger) Warn(ctx context.Context, args ...any) {
	l.sugar.Warn(args...)
}

// Error logs a message at ErrorLevel.
func (l *ZapLogger) Error(ctx context.Context, args ...any) {
	l.sugar.Error(args...)
}

// Fatal logs a message at FatalLevel then calls os.Exit(1).
func (l *ZapLogger) Fatal(ctx context.Context, args ...any) {
	l.sugar.Fatal(args...)
}

// Panic logs a message at PanicLevel then panics.
func (l *ZapLogger) Panic(ctx context.Context, args ...any) {
	l.sugar.Panic(args...)
}

// Debugf logs a formatted message at DebugLevel.
func (l *ZapLogger) Debugf(ctx context.Context, template string, args ...any) {
	l.sugar.Debugf(template, args...)
}

// Infof logs a formatted message at InfoLevel.
func (l *ZapLogger) Infof(ctx context.Context, template string, args ...any) {
	l.sugar.Infof(template, args...)
}

// Warnf logs a formatted message at WarnLevel.
func (l *ZapLogger) Warnf(ctx context.Context, template string, args ...any) {
	l.sugar.Warnf(template, args...)
}

// Errorf logs a formatted message at ErrorLevel.
func (l *ZapLogger) Errorf(ctx context.Context, template string, args ...any) {
	l.sugar.Errorf(template, args...)
}

// Fatalf logs a formatted message at FatalLevel then calls os.Exit(1).
func (l *ZapLogger) Fatalf(ctx context.Context, template string, args ...any) {
	l.sugar.Fatalf(template, args...)
}

// Panicf logs a formatted message at PanicLevel then panics.
func (l *ZapLogger) Panicf(ctx context.Context, template string, args ...any) {
	l.sugar.Panicf(template, args...)
}

// WithFields adds structured context to the logger.
func (l *ZapLogger) WithFields(fields map[string]any) Logger {
	zapFields := make([]any, 0, len(fields)*2)
	for k, v := range fields {
		zapFields = append(zapFields, k, v)
	}
	newSugar := l.sugar.With(zapFields...)
	return &ZapLogger{sugar: newSugar}
}

// Sync flushes any buffered log entries.
func (l *ZapLogger) Sync() error {
	return l.sugar.Sync()
}
