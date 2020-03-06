package tmpl

type StructTmpl struct {
	StructName string
	FieldName  []StructField
}

type StructField struct {
	Name     string
	DataType string
}

func StructTemplate() []byte {
	return []byte(`
	package models

	type {{.StructName}} struct{
		{{range .FieldName}}
			{{.Name}} {{.DataType}}
		{{end}}
	}
`)
}
