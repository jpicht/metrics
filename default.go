package metrics

import (
	"os"
	"path"
	"time"

	"github.com/jpicht/metrics/types"
)

var (
	// DefaultSink is the default instance, configured via environment variables
	DefaultSink Sink = &NullSink{}

	// DefaultTags to add to all measurements
	DefaultTags = map[string]string{}
)

func init() {
	DefaultTags["host"], _ = os.Hostname()
	DefaultTags["application"] = path.Base(os.Args[0])
}

// Publish is a short-hand for DefaultSink.Publish
func Publish(measurement string, tags map[string]string, values ...types.Field) {
	DefaultSink.Publish(measurement, tags, values...)
}

// PublishWithTimestamp is a short-hand for DefaultSink.PublishWithTimestamp
func PublishWithTimestamp(ts time.Time, measurement string, tags map[string]string, values ...types.Field) {
	DefaultSink.PublishWithTimestamp(ts, measurement, tags, values...)
}
