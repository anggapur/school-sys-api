package logger

import (
	"context"
	"io"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Option provides a way to customize a logger.
type Option func(opt *option)

type option struct {
	ctx           context.Context
	getLoggerFunc func() Logger
	attrs         []Attribute
	level         zapcore.Level
	output        zapcore.WriteSyncer
	config        zapcore.EncoderConfig
	encoderFunc   func(zapcore.EncoderConfig) zapcore.Encoder
	zapOption     []zap.Option
	trackingLibs  []Logger
}

func (opt *option) getLogger() Logger {
	if opt.getLoggerFunc == nil {
		return nil
	}
	return opt.getLoggerFunc()
}

var defaultOption = []Option{
	withDefaultContext(),
	withDefaultOutput(),
	withDefaultEncoder(),
	withDefaultMessage(),
	withDefaultLevel(),
	withDefaultTime(),
	withDefaultCaller(),
}

var contextKey = struct{}{}

// FromContext allow new logger to inherit configuration and attributes from
// existing logger in the context. However the configurations cannot be changed.
// If the context is nil or contains no logger, the new logger output
// will be set to io.Discard.
func FromContext(ctx context.Context) Option {
	return func(opt *option) {
		opt.getLoggerFunc = func() Logger {
			var log Logger
			defer func() {
				if log == nil {
					WithOutput(io.Discard)(opt)
				}
			}()

			if ctx == nil {
				return nil
			}
			opt.ctx = ctx

			log, ok := ctx.Value(contextKey).(Logger)
			if !ok {
				return nil
			}

			return log
		}
	}
}

// FromLogger allow new logger to inherit configuration and attributes from
// existing logger. However the configurations cannot be changed.
// If the logger is nil, the new logger output will be set to io.Discard.
func FromLogger(log Logger) Option {
	return func(opt *option) {
		opt.getLoggerFunc = func() Logger {
			defer func() {
				if log == nil {
					WithOutput(io.Discard)(opt)
				}
			}()

			if log == nil {
				return nil
			}

			return log
		}
	}
}

// WithContext is an option to specify the context for inserting the new logger.
// Default is context.Background.
func WithContext(ctx context.Context) Option {
	return func(opt *option) {
		opt.ctx = ctx
	}
}

func withDefaultContext() Option {
	return func(opt *option) {
		if opt.ctx == nil {
			opt.ctx = context.Background()
		}
	}
}

// WithAttributes is an option to add attribute to the new logger.
func WithAttributes(attrs ...Attribute) Option {
	return func(opt *option) {
		opt.attrs = append(opt.attrs, attrs...)
	}
}

// WithOutput is an option to specify the output for the new logger.
// Default is os.Stderr.
func WithOutput(w io.Writer) Option {
	return func(opt *option) {
		opt.output = zapcore.AddSync(w)
	}
}

func withDefaultOutput() Option {
	return func(opt *option) {
		if opt.output == nil {
			opt.output = zapcore.AddSync(os.Stderr)
		}
	}
}

// WithEncoder is an option to specify the encoder for the new logger.
// Default is EncoderJSON.
func WithEncoder(e Encoder) Option {
	return func(opt *option) {
		opt.encoderFunc = e.zapcoreEncoder()
	}
}

func withDefaultEncoder() Option {
	return func(opt *option) {
		if opt.encoderFunc == nil {
			opt.encoderFunc = EncoderJSON.zapcoreEncoder()
		}
	}
}

// Encoder represents the encoder for a logger.
type Encoder string

const (
	// EncodeJSON encodes the logs using JSON format.
	EncoderJSON = Encoder("json")
	// EncodeConsole encodes the logs in a plain-text format but leaves
	// the structured fields as JSON.
	EncoderConsole = Encoder("console")
)

func (e Encoder) zapcoreEncoder() func(zapcore.EncoderConfig) zapcore.Encoder {
	switch e {
	case EncoderJSON:
		return zapcore.NewJSONEncoder
	case EncoderConsole:
		return zapcore.NewConsoleEncoder
	default:
		panic("logat: invalid encoder")
	}
}

func withDefaultMessage() Option {
	return func(opt *option) {
		opt.config.MessageKey = "message"
	}
}

// WithLevel is an option to specify the severity level for the new logger.
// Default is LevelInfo.
func WithLevel(lvl Level) Option {
	return func(opt *option) {
		opt.level = lvl.zapcoreLevel()
	}
}

func WithTrackingLib(lib Logger) Option {
	return func(opt *option) {
		opt.trackingLibs = append(opt.trackingLibs, lib)
	}
}

func withDefaultLevel() Option {
	return func(opt *option) {
		opt.config.LevelKey = "level"
		opt.config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
}

// Level represents the severity level for a logger.
type Level string

const (
	LevelDebug = Level("debug")
	LevelInfo  = Level("info")
	LevelWarn  = Level("warn")
	LevelError = Level("error")
	LevelFatal = Level("fatal")
)

func (lvl Level) zapcoreLevel() zapcore.Level {
	switch lvl {
	case LevelDebug:
		return zapcore.DebugLevel
	case LevelInfo:
		return zapcore.InfoLevel
	case LevelWarn:
		return zapcore.WarnLevel
	case LevelError:
		return zapcore.ErrorLevel
	case LevelFatal:
		return zapcore.FatalLevel
	default:
		panic("logat: invalid level")
	}
}

// WithTime is an option to specify time format for the new logger.
// Default is TimeRFC3339.
func WithTime(t Time) Option {
	return func(opt *option) {
		opt.config.EncodeTime = t.zapcoreTimeEncoder()
	}
}

func withDefaultTime() Option {
	return func(opt *option) {
		opt.config.TimeKey = "time"
		if opt.config.EncodeTime == nil {
			opt.config.EncodeTime = TimeRFC3339.zapcoreTimeEncoder()
		}
	}
}

// Time represents the time format for a logger.
type Time string

const (
	TimeEpoch       = Time("epoch")
	TimeEpochMilli  = Time("epoch_milli")
	TimeEpochNano   = Time("epoch_nano")
	TimeRFC3339     = Time(time.RFC3339)
	TimeRFC3339Nano = Time(time.RFC3339Nano)
	TimeISO8601     = Time("2006-01-02T15:04:05.000Z0700")
)

func (t Time) zapcoreTimeEncoder() zapcore.TimeEncoder {
	switch t {
	case TimeEpoch:
		return zapcore.EpochTimeEncoder
	case TimeEpochMilli:
		return zapcore.EpochMillisTimeEncoder
	case TimeEpochNano:
		return zapcore.EpochNanosTimeEncoder
	case TimeRFC3339:
		return zapcore.RFC3339TimeEncoder
	case TimeRFC3339Nano:
		return zapcore.RFC3339NanoTimeEncoder
	case TimeISO8601:
		return zapcore.ISO8601TimeEncoder
	default:
		panic("logat: invalid time")
	}
}

func withDefaultCaller() Option {
	return func(opt *option) {
		opt.config.CallerKey = "caller"
		opt.config.EncodeCaller = zapcore.ShortCallerEncoder
		opt.zapOption = append(opt.zapOption, zap.AddCaller(), zap.AddCallerSkip(1))
	}
}
