package tmpl

import (
	"bytes"
	"text/template"

	"github.com/snwfdhmp/decode"
)

func Replace(templatePath, valuesPath string) ([]byte, error) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return []byte{}, err
	}

	data := NewValuesFiles()
	if err := decode.YAML(valuesPath, data); err != nil {
		return []byte{}, err
	}

	buff := bytes.NewBuffer([]byte{})
	if err := t.Execute(buff, data); err != nil {
		return []byte{}, err
	}

	return buff.Bytes(), nil
}

type ValuesFile struct {
	Values map[string]interface{}
}

func NewValuesFiles() *ValuesFile {
	return &ValuesFile{
		Values: make(map[string]interface{}),
	}
}
