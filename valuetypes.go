package metrics

import "github.com/jpicht/metrics/types"

// Int value
func Int(name string, i int) types.Field { return types.NewField(name, i) }

// Int8 value
func Int8(name string, i int8) types.Field { return types.NewField(name, i) }

// Int16 value
func Int16(name string, i int16) types.Field { return types.NewField(name, i) }

// Int32 value
func Int32(name string, i int32) types.Field { return types.NewField(name, i) }

// Int64 value
func Int64(name string, i int64) types.Field { return types.NewField(name, i) }

// UInt value
func UInt(name string, i uint) types.Field { return types.NewField(name, i) }

// UInt8 value
func UInt8(name string, i uint8) types.Field { return types.NewField(name, i) }

// UInt16 value
func UInt16(name string, i uint16) types.Field { return types.NewField(name, i) }

// UInt32 value
func UInt32(name string, i uint32) types.Field { return types.NewField(name, i) }

// UInt64 value
func UInt64(name string, i int64) types.Field { return types.NewField(name, i) }

// Float32 value
func Float32(name string, f float32) types.Field { return types.NewField(name, f) }

// Float64 value
func Float64(name string, f float64) types.Field { return types.NewField(name, f) }

// String value
func String(name string, s string) types.Field { return types.NewField(name, s) }
