package main

import (
	"bytes"
	"errors"
	"go/ast"
	"go/printer"
	"go/token"
	"log"
	"regexp"
	"strings"
	"text/template"
)

type NastedStruct struct {
	Parent     []string
	Name       string
	Type       string
	Validators map[string]string
	IsStruct   bool
	IsArray    bool
}

type TemplateData struct {
	Package       string
	NastedStructs map[string][]NastedStruct
}

var validateRegexp = regexp.MustCompile(`validate:"(.*)?"`)

func parseStruct(strct *ast.StructType, nastedStruct []NastedStruct, fieldName []string) (result []NastedStruct) {
	fset := token.NewFileSet()
	for _, field := range strct.Fields.List {
		if nestedStruct, isStructField := field.Type.(*ast.StructType); isStructField { //nolint
			nastedStruct = append(nastedStruct, NastedStruct{
				Parent:   fieldName,
				Name:     field.Names[0].Name,
				Type:     field.Names[0].Name,
				IsStruct: true,
			})
			parents := append(fieldName, field.Names[0].Name)
			nastedStruct = parseStruct(nestedStruct, nastedStruct, parents)
		} else {
			var typeNameBuf bytes.Buffer
			err := printer.Fprint(&typeNameBuf, fset, field.Type)
			if err != nil {
				log.Fatalf("failed printing %s", err)
			}
			tp := typeNameBuf.String()

			var name string
			if len(field.Names) != 0 {
				name = field.Names[0].Name
			} else {
				name = tp
			}

			validators := make(map[string]string)

			if field.Tag != nil {
				matches := validateRegexp.FindAllStringSubmatch(field.Tag.Value, -1)
				if len(matches) > 0 && len(matches[0]) > 1 {
					for _, expr := range strings.Split(matches[0][1], "|") {
						params := strings.Split(expr, ":")
						validators[params[0]] = params[1]
					}
				}
			}
			_, isArray := field.Type.(*ast.ArrayType)
			nastedStruct = append(nastedStruct, NastedStruct{
				Parent:     fieldName,
				Name:       name,
				Type:       tp,
				Validators: validators,
				IsStruct:   false,
				IsArray:    isArray,
			})
		}
	}

	return nastedStruct
}

func getMetaDataForTamplate(pkg string, nastedStructs map[string][]NastedStruct) TemplateData {
	return TemplateData{
		Package:       pkg,
		NastedStructs: nastedStructs,
	}
}

func parseAstFile(file *ast.File) (data TemplateData, err error) { //nolint
	nastedStructsMap := make(map[string][]NastedStruct)

	for _, decl := range file.Decls {
		decl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range decl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			spec, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}
			nastedStructs := make([]NastedStruct, 0, 1)
			nastedStructs = parseStruct(spec, nastedStructs, []string{})
			nastedStructsMap[typeSpec.Name.Name] = nastedStructs

			if err != nil {
				return data, err
			}
		}
	}

	for key, value := range nastedStructsMap {
		countOfEmptyTagsFields := 0
		for _, item := range value {
			if len(item.Validators) == 0 {
				childs := 0
				for _, it := range value {
					if len(it.Parent) > 0 && it.Parent[0] == item.Name {
						childs++
					}
				}
				if childs == 0 {
					countOfEmptyTagsFields++
				}
			}
		}
		if countOfEmptyTagsFields == len(value) {
			nastedStructsMap[key] = []NastedStruct{}
		}
	}

	data = getMetaDataForTamplate("package "+file.Name.Name, nastedStructsMap)

	return data, nil
}

func writeToTemplate(file *ast.File) (*bytes.Buffer, error) {
	data, err := parseAstFile(file)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	buf.Write([]byte(data.Package))
	tmpl, err := template.New("validation").Parse(codeTemplate)
	if err != nil {
		return nil, errors.New("cannot parse template correctly") //nolint
	}
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}

	return &buf, nil
}
