package willowlogger

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func Test_MiddlewareSetLogger(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("It panics if the child is nil", func(t *testing.T) {
		logger, err := NewZapLogger(StringToLogLevel("debug"))
		g.Expect(err).ToNot(HaveOccurred())

		g.Expect(func() { MiddlewareSetLogger(logger, nil)(nil, nil) }).To(Panic())
	})

	t.Run("It sets a logger on the request context", func(t *testing.T) {
		logger, err := NewZapLogger(StringToLogLevel("debug"))
		g.Expect(err).ToNot(HaveOccurred())

		var reqLogger *zap.Logger
		response := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "http://127.0.0.1:8080", bytes.NewBuffer([]byte{}))
		callback := MiddlewareSetLogger(logger, func(w http.ResponseWriter, r *http.Request) {
			reqLogger = MiddlewareLogger(r.Context())
		})

		callback(response, request)
		g.Expect(reqLogger).ToNot(BeNil())
	})

	t.Run("It sets a request id on the context", func(t *testing.T) {
		logger, err := NewZapLogger(StringToLogLevel("debug"))
		g.Expect(err).ToNot(HaveOccurred())

		requestID := ""
		response := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "http://127.0.0.1:8080", bytes.NewBuffer([]byte{}))
		callback := MiddlewareSetLogger(logger, func(w http.ResponseWriter, r *http.Request) {
			requestID = MiddlewareRequestID(r.Context())
		})

		callback(response, request)
		g.Expect(requestID).ToNot(Equal(""))
	})

	t.Run("It sets the request ID on the logger", func(t *testing.T) {
		observedZapCore, observedLogs := observer.New(zap.InfoLevel)
		observedLogger := zap.New(observedZapCore)

		var reqLogger *zap.Logger
		response := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "http://127.0.0.1:8080", bytes.NewBuffer([]byte{}))
		callback := MiddlewareSetLogger(observedLogger, func(w http.ResponseWriter, r *http.Request) {
			reqLogger = MiddlewareLogger(r.Context())
		})

		callback(response, request)

		reqLogger.Info("log message")
		g.Expect(len(observedLogs.FilterFieldKey(string(ResuestIDHeader)).All())).To(Equal(1))
	})
}

func Test_NamedMiddlwareLogger(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("It can return a named logger with a new context", func(t *testing.T) {
		logger, err := NewZapLogger(StringToLogLevel("debug"))
		g.Expect(err).ToNot(HaveOccurred())

		ctx := context.WithValue(context.Background(), LoggerCtxKey, logger)

		newCtx, newLogger := NamedMiddlewareLogger(ctx, "logger name")
		g.Expect(newLogger.Name()).To(Equal("logger name"))
		g.Expect(newCtx).ToNot(Equal(ctx))
	})
}
