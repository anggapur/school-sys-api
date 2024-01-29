package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Attribute is a key-value pair represents an attribute.
type Attribute = zapcore.Field

type attributeAdder interface {
	addAttribute(attrs ...Attribute) Logger
}

func (log logger) addAttribute(attrs ...Attribute) Logger {
	return logger{logger: log.logger.With(attrs...), trackingLibs: log.trackingLibs}
}

// Field returns an attribute for the given key and value.
func Field(key string, value any) Attribute {
	return zap.Any(key, value)
}

// FieldErr returns an error attribute with key=error and value=err.
func FieldErr(err error) Attribute {
	return zap.Error(err)
}

// FieldCorrelationID returns a correlation id attribute for distributed tracing
// with key=correlationID and value=id.
func FieldCorrelationID(id string) Attribute {
	return zap.String("correlationID", id)
}
