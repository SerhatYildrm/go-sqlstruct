package templates

import (
	"os"
	"text/template"
)

var text = `
package Models

{{if .Package}}
import (
{{range $index, $val := .Packages}}
	"{{$val}}"
{{end}}
)
{{- end}}

// {{.Name }} Model is genereted by gocompare tool.
type {{ .Name }} struct {
	{{range $key, $value:= .Results}}
		{{$key}} {{$value}}
	{{end}}
}
	`

// SQLColumn ...
type SQLColumn struct {
	Name string
	Type interface{}
}

// TemplateStruct ...
type TemplateStruct struct {
	Name    string
	Results map[string]interface{}

	Package  bool
	Packages []string
}

// SQLStruct ...
type SQLStruct struct {
	Name    string
	Results map[string]interface{}

	Model map[string]interface{}

	Package  bool
	Packages []string
}

// CreateSQLStruct ...
func CreateSQLStruct(Name string, Results map[string]interface{}) *SQLStruct {
	var p = make([]string, 0)
	return &SQLStruct{Name, Results, nil, false, p}
}

// Create ...
func (s *SQLStruct) Create() error {
	s.Model = make(map[string]interface{})

	for key, val := range s.Results {
		s.Model[key] = val
		s.InsertPackage(val)
	}
	return nil
}

// WriteToGOFile Write to file
func (s *SQLStruct) WriteToGOFile(path string) error {
	tmp := TemplateStruct{s.Name, s.Model, s.Package, s.Packages}
	tpl := template.Must(template.New("SqlToStruct").Parse(text))

	f, err := os.Create(path + "/" + s.Name + ".go")
	defer f.Close()

	if err != nil {
		return err
	}
	_ = tpl.Execute(f, &tmp)
	f.Sync()
	return nil
}

// InsertPackage ...
func (s *SQLStruct) InsertPackage(columnType interface{}) {
	if columnType == "time.Time" {
		if !HasValue(s.Packages, "time") {
			s.Package = true
			s.Packages = append(s.Packages, "time")
		}
	}
}

// HasValue ...
func HasValue(arr []string, value string) bool {
	for _, val := range arr {
		if value == val {
			return true
		}
	}
	return false
}
