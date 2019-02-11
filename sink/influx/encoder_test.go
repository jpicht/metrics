package influx_test

import (
	"testing"

	"gitlab.4com.de/metrics/sink/influx"
)

func TestEncoderEscapeFieldName(t *testing.T) {
	cases := map[string]string{
		"test":   "test",
		"te,st":  "te\\,st",
		"te=st":  "te\\=st",
		"te st":  "te\\ st",
		"te\\st": "te\\\\st",
	}
	e := influx.New()

	for in, out := range cases {
		if ret := e.EscapeFieldName(in); out != ret {
			t.Errorf("Expected EscapeFieldName('%s') = '%s' (got '%s')", in, out, ret)
		}
	}
}

func TestEncoderEncodeValue(t *testing.T) {
	cases := [][]interface{}{
		{1, "1i"},
		{1.0, "1"},
		{1.1, "1.1"},
		{"a", "\"a\""},
		{" ", "\" \""},
		{"\"", "\"\\\"\""},
	}
	e := influx.New()

	for _, d := range cases {
		in := d[0]
		out := d[1]

		if ret := e.EncodeValue(in); out != ret {
			t.Errorf("Expected EscapeFieldName(%#v) = '%s' (got '%s')", in, out, ret)
		}
	}
}
