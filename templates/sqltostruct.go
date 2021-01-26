package templates

import (
	"os"
	"text/template"
)

var text = `
package Models

// {{.Name }} Model is genereted by gocompare tool.
type {{ .Name }} struct {
	{{range $key, $value:= .Results}}
		{{$key}} {{$value}},
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
}

// SQLStruct ...
type SQLStruct struct {
	Name    string
	Results map[string]interface{}

	Model map[string]interface{}
}

// CreateSQLStruct ...
func CreateSQLStruct(Name string, Results map[string]interface{}) *SQLStruct {
	return &SQLStruct{Name, Results, nil}
}

// Create ...
func (s *SQLStruct) Create() error {
	s.Model = make(map[string]interface{})

	for key, val := range s.Results {
		s.Model[key] = val
	}
	return nil
}

// WriteToGOFile Write to file
func (s *SQLStruct) WriteToGOFile(path string) error {
	tmp := TemplateStruct{s.Name, s.Model}
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
