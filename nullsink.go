package metrics

import (
	"time"

	"github.com/jpicht/metrics/types"
)

// NullSink implements a metrics black hole
type NullSink struct{}

func (*NullSink) Publish(measurement string, tags map[string]string, values ...types.Field) {
}

func (n *NullSink) PublishWithTimestamp(ts time.Time, measurement string, tags map[string]string, values ...types.Field) {
}
