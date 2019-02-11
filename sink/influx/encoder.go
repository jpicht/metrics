package influx

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type (
	// Encoder implements the influx line protocol encoding
	Encoder struct {
	}

	// FieldEncoder encodes a value to be stored in influx
	FieldEncoder interface {
		InfluxValue() string
	}
)

// New creates an Encoder instance
func NewEncoder() *Encoder {
	return &Encoder{}
}

// Encode formats a data point conforming the influx the line protocol
func (e *Encoder) Encode(name string, tags map[string]string, values map[string]interface{}, ts time.Time) string {
	return e.Tags(name, tags) + " " + e.Values(values) + " " + strconv.FormatInt(ts.UnixNano(), 10)
}

var nameEscapingReplacer = strings.NewReplacer(",", "\\,", "=", "\\=", " ", "\\ ", "\\", "\\\\")

// EscapeFieldName prefixes all special characters with a backslash
func (e *Encoder) EscapeFieldName(s string) string {
	if s == "" {
		return ""
	}

	return nameEscapingReplacer.Replace(s)
}

// EscapeString encodes a string for use as a value inside the influx line protocol
func (e *Encoder) EscapeString(s string) string {
	return "\"" + strings.Replace(s, "\"", "\\\"", -1) + "\""
}

// Tags encodes the tags conforming to the influx line protocol
func (e *Encoder) Tags(name string, tags map[string]string) string {
	elements := make([]string, len(tags)+1)
	elements[0] = name
	i := 1
	for k, v := range tags {
		elements[i] = e.EscapeFieldName(k) + "=" + v
		i++
	}
	return strings.Join(elements, ",")
}

// Values encodes a map of values into a line protocol string
func (e *Encoder) Values(values map[string]interface{}) string {
	elements := make([]string, len(values))
	i := 0
	for k, v := range values {
		elements[i] = e.EscapeFieldName(k) + "=" + e.EncodeValue(v)
		i++
	}
	return strings.Join(elements, ",")
}

// EncodeValue encodes a single value for use in the wire protocol
func (e *Encoder) EncodeValue(v interface{}) string {
	switch vv := v.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", vv) + "i"
	case float32:
		return strconv.FormatFloat(float64(vv), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(vv, 'f', -1, 64)
	case string:
		return e.EscapeString(vv)
	default:
	}
	return e.EscapeString(fmt.Sprintf("%v", v))
}
