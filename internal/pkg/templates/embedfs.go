package templates

import (
	"bytes"
	"embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"sigs.k8s.io/yaml"
)

var (
	//go:embed *.goyaml
	FS embed.FS
)

type Template struct {
	name string
}

func New(name string) *Template {
	return &Template{
		name: fmt.Sprintf("%s.goyaml", name),
	}
}

func (tpl *Template) Render(m any, dest any) error {
	funcMap := sprig.TxtFuncMap()
	funcMap["slicetostr"] = func(items []string) string {
		if len(items) == 0 {
			return "[]"
		}

		return fmt.Sprintf("[\"%s\"]", strings.Join(items, "\",\""))
	}

	t, err := template.New(tpl.name).Funcs(funcMap).ParseFS(FS, tpl.name)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	if err := t.Execute(buf, m); err != nil {
		return err
	}

	return yaml.Unmarshal(buf.Bytes(), dest)
}
