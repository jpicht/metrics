package metrics

import (
	"time"

	"github.com/jpicht/metrics/types"
)

type (
	// Sink implements a sink for metrics
	Sink interface {
		Publish(measurement string, tags map[string]string, values ...types.Field)
		PublishWithTimestamp(ts time.Time, measurement string, tags map[string]string, values ...types.Field)
	}
)
