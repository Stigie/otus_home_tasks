package main

import (
	"github.com/stretchr/testify/require"
	"go/parser"
	"go/token"
	"testing"
)

func TestCodeParsing(t *testing.T) {
	t.Run("should parse package name correctly", func(t *testing.T) {
		input, _ := parser.ParseFile(token.NewFileSet(), "./test_data/models.go", nil, parser.AllErrors)
		data, err := parseAstFile(input)

		userData := []NastedStruct{
			{
				Parent: []string{},
				Name:   "ID",
				Type:   "string",
				Validators: map[string]string{
					"len": "36",
				},
				IsStruct: false,
				IsArray:  false,
			},
			{
				Parent:     []string{},
				Name:       "Name",
				Type:       "string",
				Validators: map[string]string{},
				IsStruct:   false,
				IsArray:    false,
			},
			{
				Parent: []string{},
				Name:   "Age",
				Type:   "int",
				Validators: map[string]string{
					"max": "50",
					"min": "18",
				},
				IsStruct: false,
				IsArray:  false,
			},
			{
				Parent: []string{},
				Name:   "Email",
				Type:   "string",
				Validators: map[string]string{
					"regexp": "^\\\\w+@\\\\w+\\\\.\\\\w+$",
				},
				IsStruct: false,
				IsArray:  false,
			},
			{
				Parent: []string{},
				Name:   "Role",
				Type:   "UserRole",
				Validators: map[string]string{
					"in": "admin,stuff",
				},
				IsStruct: false,
				IsArray:  false,
			},
			{
				Parent: []string{},
				Name:   "Phones",
				Type:   "[]string",
				Validators: map[string]string{
					"len": "11",
				},
				IsStruct: false,
				IsArray:  true,
			},
		}

		appData := []NastedStruct{
			{
				Parent: []string{},
				Name:   "Version",
				Type:   "string",
				Validators: map[string]string{
					"len": "5",
				},
				IsStruct: false,
				IsArray:  false,
			},
		}

		app1Data := []NastedStruct{
			{
				Parent:     []string{},
				Name:       "Version",
				Type:       "Version",
				Validators: nil,
				IsStruct:   true,
				IsArray:    false,
			},
			{
				Parent: []string{
					"Version",
				},
				Name: "Build",
				Type: "int",
				Validators: map[string]string{
					"max": "50",
					"min": "18",
				},
				IsStruct: false,
				IsArray:  false,
			},
			{
				Parent:     []string{},
				Name:       "Vqweqe",
				Type:       "Vqweqe",
				Validators: nil,
				IsStruct:   true,
				IsArray:    false,
			},
			{
				Parent: []string{
					"Vqweqe",
				},
				Name: "Build",
				Type: "[]int",
				Validators: map[string]string{
					"max": "50",
					"min": "18",
				},
				IsStruct: false,
				IsArray:  true,
			},
			{
				Parent:     []string{},
				Name:       "Vqweqe1",
				Type:       "App",
				Validators: map[string]string{},
				IsStruct:   false,
				IsArray:    false,
			},
		}

		expectedData := map[string][]NastedStruct{
			"User": userData,
			"App1": app1Data,
			"App":  appData,
		}

		templateData := TemplateData{
			Package:       "package test_data",
			NastedStructs: expectedData,
		}

		require.Nil(t, err)
		require.Equal(t, data, templateData)
	})
}
