package models

//Code generated by go-validate tool; DO NOT EDIT.

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

var (
	ErrLen                = errors.New("length is incorrect")
	ErrValidationMin      = errors.New("Minimum error")
	ErrValidationMax      = errors.New("Maximum error")
	ErrValidationContains = errors.New("Contain error")
	ErrValidationRegexp   = errors.New("Regex error")
	ErrValidationType     = errors.New("Unsupported validation type")
)

func checkLen(field, str string, length int) (err ValidationError) {
	if len(str) != length {
		return ValidationError{
			Field: field,
			Err:   ErrLen,
		}
	}
	return ValidationError{
		Field: field,
		Err:   nil,
	}
}

func checkContains(typeField, val string, items string) (bool, error) {
	switch typeField {
	case "string":
		arr := strings.Split(items, ",")
		for _, item := range arr {
			if item == val {
				return true, nil
			}
		}
		return false, nil
	case "int":
		arr := strings.Split(items, ",")
		intVal, err := strconv.Atoi(val)
		if err != nil {
			return false, err
		}
		for _, item := range arr {
			result, err := strconv.Atoi(item)
			if err != nil {
				return false, err
			}
			if result == intVal {
				return true, nil
			}
		}
		return false, nil
	default:
		return false, ErrValidationType
	}
}

func checkMinMax(value interface{}, item int, typeOperation string) (bool, error) {
	typeVal := reflect.TypeOf(value).Kind().String()
	intValue, ok := value.(int)
	if !ok {
		return false, ErrValidationType
	}
	if typeVal == "int" {
		switch typeOperation {
		case "min":
			if intValue >= item {
				return true, nil
			}
			return false, nil
		case "max":
			if intValue <= item {
				return true, nil
			}
			return false, nil
		}
	}
	return false, ErrValidationType
}

func (model App) Validate() ([]ValidationError, error) {
	var vErrors []ValidationError

	errStruct := checkLen("Version", model.Version, 5)
	if errStruct.Err != nil {
		vErrors = append(vErrors, errStruct)
	}

	return vErrors, nil
}

func (model App1) Validate() ([]ValidationError, error) {
	var vErrors []ValidationError

	{
		ok, err := checkMinMax(model.Version.Build, 50, "max")
		if err != nil {
			return vErrors, err
		}
		if !ok {
			vError := ValidationError{Field: "Build", Err: ErrValidationMax}
			vErrors = append(vErrors, vError)
		}
	}

	{
		ok, err := checkMinMax(model.Version.Build, 18, "min")
		if err != nil {
			return vErrors, err
		}
		if !ok {
			vError := ValidationError{Field: "Build", Err: ErrValidationMin}
			vErrors = append(vErrors, vError)
		}
	}

	{
		for _, item := range model.Vqweqe.Build {
			ok, err := checkMinMax(item, 50, "max")
			if err != nil {
				return vErrors, err
			}
			if !ok {
				vError := ValidationError{Field: "Build", Err: ErrValidationMax}
				vErrors = append(vErrors, vError)
			}
		}
	}

	{
		for _, item := range model.Vqweqe.Build {
			ok, err := checkMinMax(item, 18, "min")
			if err != nil {
				return vErrors, err
			}
			if !ok {
				vError := ValidationError{Field: "Build", Err: ErrValidationMin}
				vErrors = append(vErrors, vError)
			}
		}
	}

	if vError, err := model.Vqweqe1.Validate(); len(vError) > 0 {
		vErrors = append(vErrors, vError[:]...)
	} else {
		return []ValidationError{}, err
	}

	return vErrors, nil
}

func (model App2) Validate() ([]ValidationError, error) {
	var vErrors []ValidationError

	if vError, err := model.App1.Validate(); len(vError) > 0 {
		vErrors = append(vErrors, vError[:]...)
	} else {
		return []ValidationError{}, err
	}

	{
		ok, err := checkMinMax(model.Build, 50, "max")
		if err != nil {
			return vErrors, err
		}
		if !ok {
			vError := ValidationError{Field: "Build", Err: ErrValidationMax}
			vErrors = append(vErrors, vError)
		}
	}

	{
		ok, err := checkMinMax(model.Build, 18, "min")
		if err != nil {
			return vErrors, err
		}
		if !ok {
			vError := ValidationError{Field: "Build", Err: ErrValidationMin}
			vErrors = append(vErrors, vError)
		}
	}

	return vErrors, nil
}

func (model Response) Validate() ([]ValidationError, error) {
	var vErrors []ValidationError

	{
		var typeField string
		typeField = reflect.TypeOf(model.Code).Kind().String()

		value := fmt.Sprint(model.Code)
		check, err := checkContains(typeField, value, "200,404,500")
		if err != nil {
			return vErrors, err
		}
		if !check {
			vError := ValidationError{Field: "Code", Err: ErrValidationContains}
			vErrors = append(vErrors, vError)
		}

	}

	return vErrors, nil
}

func (model User) Validate() ([]ValidationError, error) {
	var vErrors []ValidationError

	errStruct := checkLen("ID", model.ID, 36)
	if errStruct.Err != nil {
		vErrors = append(vErrors, errStruct)
	}

	{
		ok, err := checkMinMax(model.Age, 50, "max")
		if err != nil {
			return vErrors, err
		}
		if !ok {
			vError := ValidationError{Field: "Age", Err: ErrValidationMax}
			vErrors = append(vErrors, vError)
		}
	}

	{
		ok, err := checkMinMax(model.Age, 18, "min")
		if err != nil {
			return vErrors, err
		}
		if !ok {
			vError := ValidationError{Field: "Age", Err: ErrValidationMin}
			vErrors = append(vErrors, vError)
		}
	}

	re := regexp.MustCompile("^\\w+@\\w+\\.\\w+$")

	if !re.Match([]byte(model.Email)) {
		vError := ValidationError{Field: "Email", Err: ErrValidationRegexp}
		vErrors = append(vErrors, vError)
	}

	{
		var typeField string
		typeField = reflect.TypeOf(model.Role).Kind().String()

		value := fmt.Sprint(model.Role)
		check, err := checkContains(typeField, value, "admin,stuff")
		if err != nil {
			return vErrors, err
		}
		if !check {
			vError := ValidationError{Field: "Role", Err: ErrValidationContains}
			vErrors = append(vErrors, vError)
		}

	}

	for _, item := range model.Phones {
		errStruct := checkLen("Phones", item, 11)
		if errStruct.Err != nil {
			vErrors = append(vErrors, errStruct)
			break
		}
	}

	return vErrors, nil
}
