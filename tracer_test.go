package trace

import (
	"bytes"
	"log"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return from New should not be nil")
	} else {
		tracer.Trace("Hello trace package.")
		if buf.String() != "Hello trace package.\n" {
			t.Errorf("Trace should not write '%s'.", buf.String())
		}

		tm := time.Date(2020, time.April, 34, 25, 72, 01, 0, time.UTC)
		tracerDate := New(&buf, WithDatetime(&tm))

		if tracerDate.GetDate() != tm {
			t.Errorf("Datetime is not in sync: '%s' vs '%s'.", tm.String(), tracerDate.GetDate().String())
		} else {
			log.Printf("Datetime is in sync: '%s' vs '%s'.", tm.String(), tracerDate.GetDate().String())
		}
		// log.Printf("hier:%s",*tracerDate.tracer.out)
		// if tracerDate.datetime != tm {
		// 	t.Errorf("Datetime is not in sync: '%s'.", tm.String())
		// }
	}
}
