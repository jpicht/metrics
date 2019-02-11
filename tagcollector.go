package metrics

import (
	"time"

	"github.com/jpicht/metrics/types"
)

type (
	// TagCollector implements a convenience-interface to iteratively add tags
	TagCollector struct {
		sink Sink
		tags map[string]string
	}
)

// NewTagCollector creates a new TagCollector object
func NewTagCollector(s Sink) *TagCollector {
	return &TagCollector{
		sink: s,
		tags: map[string]string{},
	}
}

// Publish pushes the tags plus the given info via the selected sink implementation
func (t *TagCollector) Publish(measurement string, values ...types.Field) {
	t.sink.Publish(measurement, t.tags, values...)
}

// PublishWithTimestamp pushes the tags plus the given info via the selected sink implementation
func (t *TagCollector) PublishWithTimestamp(ts time.Time, measurement string, values ...types.Field) {
	t.sink.PublishWithTimestamp(ts, measurement, t.tags, values...)
}

// WithTag adds a tag to the TagCollector
func (t *TagCollector) WithTag(name, value string) *TagCollector {
	t.tags[name] = value
	return t
}

// WithTags adds a number of tags to the TagCollector
func (t *TagCollector) WithTags(tags map[string]string) *TagCollector {
	for name, value := range tags {
		t.tags[name] = value
	}
	return t
}

// WithTag creates a new TagCollector object with tag given and the default sink
func WithTag(name, value string) *TagCollector {
	return &TagCollector{
		sink: DefaultSink,
		tags: map[string]string{name: value},
	}
}

// WithTags creates a new TagCollector object with the given tags and the default sink
func WithTags(tags map[string]string) *TagCollector {
	return &TagCollector{
		sink: DefaultSink,
		tags: tags,
	}
}
