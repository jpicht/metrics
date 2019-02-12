package metrics

import (
	"time"

	"github.com/jpicht/metrics/types"
)

type (
	// TagCollector implements a convenience-interface to iteratively add tags
	TagCollector struct {
		sink Sink
		tags *tag
	}

	// tag is the linked list element used inside TagCollector
	tag struct {
		name  string
		value string
		next  *tag
	}
)

// NewTagCollector creates a new TagCollector object
func NewTagCollector(s Sink) *TagCollector {
	return &TagCollector{
		sink: s,
		tags: nil,
	}
}

// Publish pushes the tags plus the given info via the selected sink implementation
func (t *TagCollector) Publish(measurement string, values ...types.Field) {
	t.sink.Publish(measurement, t.collect(), values...)
}

// PublishWithTimestamp pushes the tags plus the given info via the selected sink implementation
func (t *TagCollector) PublishWithTimestamp(ts time.Time, measurement string, values ...types.Field) {
	t.sink.PublishWithTimestamp(ts, measurement, t.collect(), values...)
}

// WithTag adds a tag to the TagCollector
func (t *TagCollector) WithTag(name, value string) *TagCollector {
	return &TagCollector{
		t.sink,
		&tag{name, value, t.tags},
	}
}

// WithTags adds a number of tags to the TagCollector
func (t *TagCollector) WithTags(tags map[string]string) *TagCollector {
	x := t
	for name, value := range tags {
		x = t.WithTag(name, value)
	}
	return x
}

// WithTag creates a new TagCollector object with tag given and the default sink
func WithTag(name, value string) *TagCollector {
	return &TagCollector{
		sink: DefaultSink,
		tags: &tag{name, value, nil},
	}
}

// WithTags creates a new TagCollector object with the given tags and the default sink
func WithTags(tags map[string]string) *TagCollector {
	return emptyCollector.WithTags(tags)
}

var (
	emptyCollector = &TagCollector{
		sink: DefaultSink,
		tags: nil,
	}
)

func (t *tag) count() int {
	c := 0
	e := t
	for e != nil {
		c++
		e = e.next
	}
	return c
}

func (t *tag) collect() map[string]string {
	result := make(map[string]string, t.count())
	e := t
	for e != nil {
		if _, ok := result[e.name]; !ok {
			result[e.name] = e.value
		}
		e = e.next
	}
	return result
}

func (t *TagCollector) collect() map[string]string {
	if t.tags == nil {
		return map[string]string{}
	}
	return t.tags.collect()
}
