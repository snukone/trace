package trace

import (
	"fmt"
	"io"
	"log"
	"strings"
	"time"
)

// Tracer is the interface that describes an object capable of
// tracing events throughout code.
type Tracer interface {
	Trace(...interface{})
	GetDate(...interface{}) time.Time
}

type tracer struct {
	out      io.Writer
	datetime time.Time
}

func New(w io.Writer, options ...Option) Tracer {
	// Creating tracer
	tracer := &tracer{out: w}

	// Looping over options and calling them
	for _, option := range options {
		option(tracer)
	}

	// Write log
	log.Printf("Initalizing tracer ... Message: '%s' At: '%s'", strings.ReplaceAll(fmt.Sprint(tracer.out), "\n", ""), tracer.datetime)

	return tracer
}

func (t *tracer) GetDate(a ...interface{}) time.Time {
	return t.datetime
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

type Option func(*tracer)

func WithDatetime(datetime *time.Time) Option {
	return func(t *tracer) {
		t.datetime = *datetime
	}
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}
func (t *nilTracer) GetDate(a ...interface{}) time.Time {return time.Time{}}

// Off creates a Tracer that will ignore calls to Trace.
func Off() Tracer {
	return &nilTracer{}
}
