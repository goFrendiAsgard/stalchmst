package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"text/template"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// TemplateRow is a single template
type TemplateRow struct {
	Code     string
	Template string
	Data     string
}

// TemplateModel is model to get template
type TemplateModel struct {
	db *sql.DB
}

func getDb(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return db, err
	}
	err = db.Ping()
	return db, err
}

// NewTemplateModel create new templateModel
func NewTemplateModel(connectionString string, maxConnectAttempt int) (*TemplateModel, error) {
	tm := &TemplateModel{}
	db, err := getDb(connectionString)
	attempt := 0
	for err != nil && attempt < maxConnectAttempt {
		time.Sleep(3 * time.Second)
		db, err = getDb(connectionString)
		attempt++
	}
	tm.db = db
	return tm, err
}

// Close close templateModel
func (tm *TemplateModel) Close() error {
	return tm.db.Close()
}

// GetAll getting all template from the database
func (tm *TemplateModel) GetAll() ([]TemplateRow, error) {
	templateRows := []TemplateRow{}
	results, err := tm.db.Query("SELECT code, template, data FROM templates")
	if err != nil {
		return templateRows, err
	}
	for results.Next() {
		templateRow := TemplateRow{}
		results.Scan(&templateRow.Code, &templateRow.Template, &templateRow.Data)
		templateRows = append(templateRows, templateRow)
	}
	return templateRows, nil
}

// GetByCode getting single template by code
func (tm *TemplateModel) GetByCode(code string) (TemplateRow, error) {
	templateRow := TemplateRow{}
	results, err := tm.db.Query(fmt.Sprintf("SELECT code, template, data FROM templates WHERE code='%s'", code))
	if err != nil {
		return templateRow, err
	}
	for results.Next() {
		results.Scan(&templateRow.Code, &templateRow.Template, &templateRow.Data)
		return templateRow, err
	}
	return templateRow, err
}

// Generate generating data based on template using data
func (tm *TemplateModel) Generate(templateRow TemplateRow, data string) (string, error) {
	var msgMapTemplate interface{}
	err := json.Unmarshal([]byte(data), &msgMapTemplate)
	if err != nil {
		return "", err
	}

	t, err := template.New(templateRow.Code).Parse(templateRow.Template)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, msgMapTemplate); err != nil {
		return "", err
	}

	return tpl.String(), err
}

// GenerateWithoutData generating data based on template without using data
func (tm *TemplateModel) GenerateWithoutData(templateRow TemplateRow) (string, error) {
	return tm.Generate(templateRow, templateRow.Data)
}
