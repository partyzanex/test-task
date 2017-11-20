package config

import (
	"os"
	"reflect"
	"strconv"
	"strings"
)

// код вытащил из своего репозитория https://github.com/partyzanex/golang-test-task/blob/master/conf/conf.go
// адаптировал немного под задачу
// можно было через import

// интерфейс для преобразованния типов
type TypeConvert interface {
	Int() int
	String() string
	Float() float64
	Bool() bool
}

type Value struct {
	data interface{}
}

// convert data to int
func (v Value) Int() int {
	var result = 0

	switch reflect.TypeOf(v.data).String() {
	case "float64":
		result = int(v.data.(float64))
	case "bool":
		b := int(0)
		if v.data.(bool) {
			b = int(1)
		}
		result = b
	case "string":
		b, err := strconv.ParseInt(v.data.(string), 10, 0)
		if err != nil {
			return 0
		}
		result = int(b)
	}

	return result
}

// convert data to string
func (v Value) String() string {
	var result string

	switch reflect.TypeOf(v.data).String() {
	case "float64":
		result = strconv.FormatFloat(v.data.(float64), 'g', -1, 64)
	case "bool":
		b := "false"
		if v.data.(bool) {
			b = "true"
		}
		result = b
	case "string":
		result = v.data.(string)
	}

	return result
}

// convert data to float64
func (v Value) Float() float64 {
	var result float64

	switch reflect.TypeOf(v.data).String() {
	case "float64":
		result = v.data.(float64)
	case "bool":
		b := 0.
		if v.data.(bool) {
			b = 1.
		}
		result = b
	case "string":
		result, _ = strconv.ParseFloat(v.data.(string), 64)
	}

	return result
}

// convert data to boolean
func (v Value) Bool() bool {
	var result bool

	switch reflect.TypeOf(v.data).String() {
	case "float64":
		b := false
		if v.data.(float64) > 0 {
			b = true
		}
		result = b
	case "bool":
		result = v.data.(bool)
	case "string":
		b := true
		if v.data.(string) == "" ||  v.data.(string) == "1" || strings.ToLower(v.data.(string)) != "true" {
			b = false
		}
		result = b
	}

	return result
}

// получение значений переменный окружений
func GetEnv(key, fallback string) TypeConvert {
	var result string
	var ok bool
	// если нет переменной
	if result, ok = os.LookupEnv(key); !ok {
		// знаечние по умолчанию
		result = fallback
	}

	return Value{data: result}
}
