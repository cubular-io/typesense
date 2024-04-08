package typesense

import (
	"context"
	"errors"
	"fmt"
	"github.com/cubular-io/typesense/tapi"
	"github.com/cubular-io/typesense/tapi/pointer"
	"reflect"
	"strings"
)

func (c *Client) CreateCollection(ctx context.Context, name string, Struct any) (*tapi.CollectionResponse, error) {
	val := reflect.ValueOf(Struct)
	if val.Kind() == reflect.Ptr || val.Kind() != reflect.Struct {
		return nil, errors.New("input should be a struct")
	}

	typ := val.Type()

	fields, err := lexField(typ)
	if err != nil {
		return nil, err
	}

	for _, v := range fields {
		fmt.Println(v.Name, v.Type)
	}

	return nil, nil
}

func lexField(typ reflect.Type) ([]tapi.Field, error) {
	var collectionFields []tapi.Field
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.PkgPath != "" {
			continue // Skip unexported fields
		}
		fieldType := typeAllowed(field.Type)
		if fieldType == tapi.OBJECT {
			// Check if Object is Composition
			if field.Anonymous {
				composited, err := lexField(field.Type)
				if err != nil {
					return nil, err
				}
				collectionFields = append(collectionFields, composited...)
				continue
			}
		}
		tags := field.Tag.Get("json") // tags save the field_name and settings
		apiField, err := parseField(fieldType, tags)
		if err != nil {
			return nil, err
		}
		collectionFields = append(collectionFields, apiField)
	}
	return collectionFields, nil
}

func parseField(T tapi.Type, tag string) (tapi.Field, error) {
	params := strings.Split(tag, ",")
	var field tapi.Field
	var True bool = true

	if len(params) == 0 {
		return tapi.Field{}, errors.New("field name has to be provided for matching")
	}

	field.Name = params[0]
	field.Type = string(T)

	for _, key := range params[1:] {
		switch key {
		case "optional": // optional fields, can be null
			field.Optional = &True
		case "facet": // If a field is facet its also automatically indexed
			field.Facet = &True
			field.Index = &True
		case "index":
			field.Index = &True
		case "sort":
			field.Sort = &True
		case "infix":
			field.Infix = &True
		default:
			if ref, ok := strings.CutPrefix(key, "reference:"); ok {
				field.Reference = pointer.Ptr(ref)
			}
		}
	}

	return field, nil
}

func typeAllowed(t reflect.Type) tapi.Type {
	switch t.Kind() {
	case reflect.String:
		return tapi.STRING
	case reflect.Int32, reflect.Int:
		return tapi.INT32
	case reflect.Int64:
		return tapi.INT64
	case reflect.Float32, reflect.Float64:
		return tapi.FLOAT
	case reflect.Bool:
		return tapi.BOOL
	case reflect.Slice:
		elemType := typeAllowed(t.Elem())
		if elemType != "" {
			return elemType + "[]"
		}
	case reflect.Struct:
		return tapi.OBJECT
	case reflect.Pointer:
		return typeAllowed(t.Elem())
	}
	fmt.Println(t.Kind())

	return ""
}

// Check if nested field should be supported, this is the case if an Object or ObjectArray has Index = true
func nestedFields(fields []tapi.Field) *bool {
	for _, v := range fields {
		if v.Type == tapi.OBJECT || v.Type == tapi.OBJECTARRAY {
			return pointer.Ptr(true)
		}
	}
	return nil // equals false
}

func (c *Client) CreateCollectionWithTimestamp(ctx context.Context, name string, Struct any) {

}
