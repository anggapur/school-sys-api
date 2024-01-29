package logger

import (
	"context"
	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

// A Logger provides leveled and structured logging.
// All methods are safe for concurrent use.
type Logger interface {
	// Debug logs debugging message that needed by the developer, it's verbose and toggle off by default.
	// A good usecase might be:
	// - incoming HTTP request
	// - incoming MQTT message
	//
	// The attrs is key-value pairs that provides additional attributes.
	Debug(msg string, attrs ...Attribute)

	// Info logs informational message that needed by the user, default level.
	// A good usecase might be:
	// - application start and stop
	// - cron job running
	//
	// The attrs is key-value pairs that provides additional attributes.
	Info(msg string, attrs ...Attribute)

	// Warn logs warning message that something might go wrong.
	// A good usecase might be:
	// - resource approaching limit
	// - connection lost with auto-reconnect
	//
	// The attrs is key-value pairs that provides additional attributes.
	Warn(msg string, attrs ...Attribute)

	// Error logs error message that something go wrong.
	// You might want to use it at top level caller and should not be use while returning error.
	// A good usecase might be:
	// - HTTP request failed
	// - Kafka consumer failed
	//
	// The attrs is key-value pairs that provides additional attributes.
	Error(msg string, attrs ...Attribute)

	// Fatal logs error message that something go wrong and the application should terminate.
	// Calls os.Exit(1) after logs the message. You might want to use it at main package/function.
	// A good usecase might be:
	// - failed to init app
	// - failed to setup dependency
	//
	// The attrs is key-value pairs that provides additional attributes.
	Fatal(msg string, attrs ...Attribute)
}

type logger struct {
	logger       *zap.Logger
	trackingLibs []Logger
}

// New creates a new logger with provided options. The returned context
// contains the logger in it. Out of the box, the logger will have
// default attributes message, level, time, and caller.
func New(opts ...Option) (context.Context, Logger) {
	opts = append(opts, defaultOption...)
	var opt option
	for _, f := range opts {
		f(&opt)
	}

	log := opt.getLogger()

	if log == nil {
		core := zapcore.NewCore(opt.encoderFunc(opt.config), opt.output, opt.level)
		log = logger{logger: zap.New(core, opt.zapOption...), trackingLibs: opt.trackingLibs}
	}

	if v, ok := log.(attributeAdder); ok {
		log = v.addAttribute(opt.attrs...)
	}

	ctx := context.WithValue(opt.ctx, contextKey, log)

	return ctx, log
}

func (log logger) Debug(msg string, attrs ...Attribute) {
	for _, lib := range log.trackingLibs {
		lib.Debug(msg, attrs...)
	}
	log.logger.Debug(msg, attrs...)
}

func (log logger) Info(msg string, attrs ...Attribute) {
	for _, lib := range log.trackingLibs {
		lib.Info(msg, attrs...)
	}
	log.logger.Info(msg, attrs...)
}

func (log logger) Warn(msg string, attrs ...Attribute) {
	for _, lib := range log.trackingLibs {
		lib.Warn(msg, attrs...)
	}
	log.logger.Warn(msg, attrs...)
}

func (log logger) Error(msg string, attrs ...Attribute) {
	for _, lib := range log.trackingLibs {
		lib.Error(msg, attrs...)
	}
	log.logger.Error(msg, attrs...)
}

func (log logger) Fatal(msg string, attrs ...Attribute) {
	for _, lib := range log.trackingLibs {
		lib.Info(msg, attrs...)
	}
	log.logger.Fatal(msg, attrs...)
}
