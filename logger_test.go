package willowlogger

import (
	"testing"

	. "github.com/onsi/gomega"
	"go.uber.org/zap/zapcore"
)

func Test_StringToLogLevel(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("It accepts debug", func(t *testing.T) {
		g.Expect(StringToLogLevel("debug")).To(Equal(DEBUG))
	})

	t.Run("It accepts info", func(t *testing.T) {
		g.Expect(StringToLogLevel("info")).To(Equal(INFO))
	})

	t.Run("It returns unkown for all other values", func(t *testing.T) {
		g.Expect(StringToLogLevel("nope")).To(Equal(UNKNOWN))
	})
}

func Test_NewZapLogger(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("It can create a logger with debug log level", func(t *testing.T) {
		logger, err := NewZapLogger(StringToLogLevel("debug"))
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(logger.Level()).To(Equal(zapcore.DebugLevel))
	})

	t.Run("It can create a logger with info log level", func(t *testing.T) {
		logger, err := NewZapLogger(StringToLogLevel("info"))
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(logger.Level()).To(Equal(zapcore.InfoLevel))
	})

	t.Run("It returns an error for an invalid log level", func(t *testing.T) {
		logger, err := NewZapLogger(StringToLogLevel("bad"))
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("unknown log level received: unknown"))
		g.Expect(logger).To(BeNil())
	})
}

func Test_BaseLogger(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("It returns a new logger with no configurations other than the core", func(t *testing.T) {
		logger, err := NewZapLogger(StringToLogLevel("debug"))
		g.Expect(err).ToNot(HaveOccurred())

		namedLogger := logger.Named("named-logger")
		g.Expect(namedLogger.Name()).To(Equal("named-logger"))

		coreLogger := BaseLogger(namedLogger)
		g.Expect(coreLogger.Name()).To(Equal(""))

		g.Expect(namedLogger.Core()).To(Equal(coreLogger.Core()))
	})
}
