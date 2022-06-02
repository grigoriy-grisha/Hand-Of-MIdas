package structs

import "reflect"

var (
	DefaultTagName = "structs"
)

type Field struct {
	value      reflect.Value
	field      reflect.StructField
	defaultTag string
}

func (f *Field) Name() string {
	return f.field.Name
}

type Struct struct {
	raw     interface{}
	value   reflect.Value
	TagName string
}

func (s *Struct) Names() []string {
	fields := getFields(s.value, s.TagName)

	names := make([]string, len(fields))

	for i, field := range fields {
		names[i] = field.Name()
	}

	return names
}

func getFields(v reflect.Value, tagName string) []*Field {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	var fields []*Field

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if tag := field.Tag.Get(tagName); tag == "-" {
			continue
		}

		f := &Field{
			field: field,
			value: v.FieldByName(field.Name),
		}

		fields = append(fields, f)

	}

	return fields
}

func New(s interface{}) *Struct {
	return &Struct{
		raw:     s,
		value:   structVal(s),
		TagName: DefaultTagName,
	}
}

func structVal(s interface{}) reflect.Value {
	v := reflect.ValueOf(s)

	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("not struct")
	}

	return v
}
