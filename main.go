package main

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
	"sync"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type indentJsonWriter struct {
	wrapped io.Writer
	pool    sync.Pool
}

func NewIndentJsonWriter(wrapped io.Writer) *indentJsonWriter {
	return &indentJsonWriter{
		wrapped: wrapped,
		pool: sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}
}

func (w *indentJsonWriter) Write(p []byte) (n int, err error) {
	m := make(map[string]interface{})
	if err = json.NewDecoder(bytes.NewReader(p)).Decode(&m); err != nil {
		return 0, err
	}

	for k, v := range m {
		if s, ok := v.(string); ok && strings.ContainsRune(s, '\n') {
			a := strings.Split(s, "\n")
			for i := range a {
				a[i] = strings.TrimSpace(a[i])
			}
			var j int
			for i := 0; i < len(a); i++ {
				s = a[i]
				if i+1 < len(a) && strings.HasPrefix(a[i+1], "/") {
					s += a[i+1]
					i++
				}
				if strings.HasPrefix(s, "runtime.") {
					continue
				}
				a[j] = s
				j++
			}
			m[k] = a[:j]
		}
	}

	buf := w.pool.Get().(*bytes.Buffer)
	defer w.pool.Put(buf)

	buf.Reset()
	enc := json.NewEncoder(buf)
	enc.SetIndent("", "  ")
	if err = enc.Encode(m); err != nil {
		return 0, err
	}
	return w.wrapped.Write(buf.Bytes())
}

func main() {
	cfg := zap.NewProductionConfig()

	outWriter, outWriterClose, _ := zap.Open(cfg.OutputPaths...)
	defer outWriterClose()

	errWriter, errWriterClose, _ := zap.Open(cfg.ErrorOutputPaths...)
	defer errWriterClose()

	logger, _ := cfg.Build(
		zap.ErrorOutput(errWriter),
		zap.WrapCore(func(c zapcore.Core) zapcore.Core {
			nc := zapcore.NewCore(
				zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
				zapcore.AddSync(NewIndentJsonWriter(outWriter)),
				zap.NewProductionConfig().Level,
			)
			return nc
		}))

	defer logger.Sync()
	if err := run2(); err != nil {
		logger.Error("error was", zap.Error(err))
	}
}

func run1() error {
	return errors.New("cause")
}

func run2() error {
	err := run1()
	return errors.Wrap(err, "wrap")
	// return errors.WithMessage(err, "wrap")
}
