package opentracing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpanBaggage(t *testing.T) {
	assert := assert.New(t)

	span := NewSpan("web.request")
	span.SetBaggageItem("key", "value")
	assert.Equal("value", span.BaggageItem("key"))
}

func TestSpanContext(t *testing.T) {
	assert := assert.New(t)

	span := NewSpan("web.request")
	assert.NotNil(span.Context())
}

func TestSpanSetTag(t *testing.T) {
	assert := assert.New(t)

	span := NewSpan("web.request")
	span.Span.Meta = make(map[string]string)
	span.SetTag("component", "tracer")
	assert.Equal("tracer", span.Meta["component"])
}

func TestSpanOperationName(t *testing.T) {
	assert := assert.New(t)

	span := NewSpan("web.request")
	span.SetOperationName("http.request")
	assert.Equal("http.request", span.Span.Name)
}
