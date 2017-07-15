package cpIntoPtr

import (
	"github.com/v2pro/wombat/cp/cpAnything"
	"github.com/v2pro/wombat/gen"
	"reflect"
)

func init() {
	cpAnything.F.AddDependency(F)
}

// F the function definition
var F = &gen.FuncTemplate{
	FuncTemplateName:"cpIntoPtr",
	Dependencies: []*gen.FuncTemplate{cpAnything.F},
	TemplateParams: map[string]string{
		"DT": "the dst type to copy into",
		"ST": "the src type to copy from",
	},
	FuncName: `cp_into_{{ .DT|symbol }}_from_{{ .ST|symbol }}`,
	Source: `
{{ $cp := gen "cpAnything" "DT" (.DT|elem) "ST" .ST }}
// generated from cpIntoPtr
func {{ .funcName }}(
	err *error,
	dst {{ .DT|name }},
	src {{ .ST|name }}) {
	// end of signature
	if dst == nil {
		return
	}
	defDst := *dst
	if defDst == nil {
		{{ if .DT|elem|isMap }}
			defDst = {{ .DT|elem|name }}{}
		{{ else }}
			defDst = new({{ .DT|elem|elem|name }})
		{{ end }}
		{{ $cp }}(err, defDst, src)
		*dst = defDst
		return
	}
	{{ $cp }}(err, *dst, src)
}
`,
	GenMap: map[string]interface{}{
		"isMap": genIsMap,
	},
}

func genIsMap(typ reflect.Type) bool {
	return typ.Kind() == reflect.Map
}
