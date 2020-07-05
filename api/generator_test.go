package main

import (
	"testing"
)

/*
	GIVEN generate endpoint

	WHEN user access the endpoint
	AND user provide template
	AND user provide data
	THEN generated document should be returned
*/
func TestGenerate(t *testing.T) {
	template := "{{.Name}}:{{range $skill := .Skills}} {{$skill}}{{end}}"
	data := "{\"Name\": \"Tono\", \"Skills\": [\"cooking\", \"farming\", \"fishing\"]}"
	expected := "Tono: cooking farming fishing"
	generated, err := Generate(template, data)
	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}
	if generated != expected {
		t.Errorf("Unexpected value: %s\nExpecting: %s", generated, expected)
	}
}
