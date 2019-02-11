package defaultsink

import (
	"os"
	"path"
	"time"

	"github.com/jpicht/metrics/sink"
	"github.com/jpicht/metrics/sink/influx"
)

// DefaultSink is the default instance, configured via environment variables
var DefaultSink sink.Sink = &NullSink{}

func init() {
	DefaultTags["host"], _ = os.Hostname()
	DefaultTags["application"] = path.Base(os.Args[0])

	if influx.DefaultSink != nil {
		DefaultSink = influx.DefaultSink
	}
}

// Publish is a short-hand for DefaultSink.Publish
func Publish(name string, tags map[string]string, values map[string]interface{}) {
	DefaultSink.Publish(name, tags, values)
}

// WithTimestamp is a short-hand for DefaultSink.WithTimestamp
func WithTimestamp(ts time.Time) Sink {
	return DefaultSink.WithTimestamp(ts)
}
