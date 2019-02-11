package influx

import (
	"bytes"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/jpicht/metrics"
	"github.com/jpicht/metrics/types"
)

var (
	// HTTPClient is the http.Client instance used to post the data to the influx server
	HTTPClient = &http.Client{
		Timeout: 5 * time.Second,
	}
)

type (
	// Sink implements sink.Sink for InfluxDB
	Sink struct {
		target  string
		Encoder *Encoder
	}
)

func init() {
	target := os.Getenv("INFLUX_SERVER")
	database := os.Getenv("INFLUX_DATABASE")

	if target == "" {
		return
	}

	u, err := url.Parse(target)
	if err != nil {
		return
	}

	metrics.DefaultSink = NewSink(u, database)
}

// NewSink creates a new Sink object
func NewSink(u *url.URL, database string) *Sink {
	if database == "" {
		database = os.Args[0]
	}

	u.Path = "write"
	u.RawQuery = "db=" + url.QueryEscape(database)

	return &Sink{
		target:  u.String(),
		Encoder: NewEncoder(),
	}
}

// Publish pushes one metric towards the target
func (s *Sink) Publish(name string, tags map[string]string, values ...types.Field) {
	s.push(name, tags, types.Fields(values).Map(), time.Now())
}

// PublishWithTimestamp pushes one metric towards the target
func (s *Sink) PublishWithTimestamp(ts time.Time, name string, tags map[string]string, values ...types.Field) {
	s.push(name, tags, types.Fields(values).Map(), ts)
}

func (s *Sink) push(name string, tags map[string]string, values map[string]interface{}, ts time.Time) {
	t := make(map[string]string, len(tags)+len(metrics.DefaultTags))
	for k, v := range metrics.DefaultTags {
		t[k] = v
	}
	for k, v := range tags {
		t[k] = v
	}
	line := s.Encoder.Encode(name, t, values, ts)
	HTTPClient.Post(s.target, "text/plain", bytes.NewBufferString(line))
}
