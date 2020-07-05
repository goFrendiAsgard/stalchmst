package main

import (
	"testing"
)

/*
	GIVEN list template endpoint
	WHEN user access the endpoint
	THEN list of available template should be shown
*/
func TestGetAllTemplate(t *testing.T) {
	tm, err := CreateDefaultTemplateModel()
	if err != nil {
		t.Error(err)
		return
	}
	defer tm.Close()

	templateRows, err := tm.GetAll()
	if err != nil {
		t.Error(err)
		return
	}
	if len(templateRows) == 0 {
		t.Errorf("Template rows should be more than zero")
	}
}

/*
	GIVEN template by code endpoint
	WHEN user access the endpoint
	AND user provide template code
	THEN the corresponding template should be shown
*/
func TestGetTemplateByCode(t *testing.T) {
	tm, err := CreateDefaultTemplateModel()
	if err != nil {
		t.Error(err)
		return
	}
	defer tm.Close()

	templateRow, err := tm.GetByCode("cv-markdown-example")
	if err != nil {
		t.Error(err)
		return
	}
	if templateRow.Code != "cv-markdown-example" {
		t.Errorf("Unexpected template row: %#v", templateRow)
	}
}
