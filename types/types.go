package types

type (
	// Field encapsulates a metric name value pair
	Field struct {
		name  string
		value interface{}
	}
	Fields []Field
)

func NewField(name string, value interface{}) Field {
	return Field{name, value}
}

func (f *Field) Name() string {
	return f.name
}

func (f *Field) Value() interface{} {
	return f.value
}

func (l Fields) Map() map[string]interface{} {
	m := make(map[string]interface{}, len(l))
	for _, f := range l {
		m[f.name] = f.value
	}
	return m
}
