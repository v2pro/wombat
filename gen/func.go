package gen

import (
	"reflect"
	"strings"
)

func panicOnError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func funcGetName(typ reflect.Type) string {
	typeName := typ.String()
	typeName = strings.Replace(typeName, ".", "__", -1)
	return typeName
}

func funcSymbol(typ reflect.Type) string {
	switch typ.Kind() {
	case reflect.Map:
		return "map_" + funcSymbol(typ.Key()) + "_to_" + funcSymbol(typ.Elem())
	case reflect.Slice:
		return "slice_" + funcSymbol(typ.Elem())
	default:
		typeName := funcGetName(typ)
		typeName = strings.Replace(typeName, "*", "ptr_", -1)
		return typeName
	}
}

func funcIsPtr(typ reflect.Type) bool {
	return typ.Kind() == reflect.Ptr
}

func funcElem(typ reflect.Type) reflect.Type {
	return typ.Elem()
}

func funcFieldOf(typ reflect.Type, fieldName string) reflect.StructField {
	field, found := typ.FieldByName(fieldName)
	if !found {
		panic(fieldName + " not found in " + typ.String())
	}
	return field
}

func funcFields(typ reflect.Type) []reflect.StructField {
	fields := make([]reflect.StructField, typ.NumField())
	for i := 0; i < len(fields); i++ {
		fields[i] = typ.Field(i)
	}
	return fields
}
