package main

import(
	"bytes"
	"encoding/json"
	"text/template"
)

// Generate document
func Generate(templateStr string, data string) (string, error) {
	var dataMap interface{}
	err := json.Unmarshal([]byte(data), &dataMap)
	if err != nil {
		return "", err
	}

	t, err := template.New(templateStr).Parse(templateStr)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, dataMap); err != nil {
		return "", err
	}

	return tpl.String(), err
}
