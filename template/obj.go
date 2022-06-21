package {{.PkgName}}

{{with .Imp}}
import (
{{range .}}
"{{.}}"
{{- end }}
)
{{- end }}

{{range .Objs}}
type {{.ObjName}} struct {
	{{.ObjExtFrom}}

	{{- range .FieldList}}
		{{.FieldName}}	{{.FieldType}}	{{.FieldTag}}
	{{- end }}
}
{{- end }}