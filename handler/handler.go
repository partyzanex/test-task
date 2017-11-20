package handler

import (
	"net/http"
	"fmt"
	"io/ioutil"

	"github.com/partyzanex/test-task/model"
	"encoding/json"
	"reflect"
	//"encoding/binary"
	//"math"
	//"strconv"
	"math"
	"strconv"
	"strings"
	"bytes"
	"github.com/partyzanex/test-task/config"
)

type DTO struct {
	BigNumber int32  `json:"big_number"`
	Text      string `json:"text"`
}

// хендлер (контроллер)
// todo: для каждой категории запросов сделать отдельный тип
// todo: надо где-то в другом месте хранить кеш
type Handler struct {
	// кеш для авторизации
	Auth map[string]string
	// кеш для чего-то другого
	// видимо для каждого пользователя есть ограничение на кол-во заданий
	// но это количество не изменяется очевидным способом здесь
	Work map[string]int32
}

// прогрев кеша для админа и теста
func (h *Handler) init() {
	h.Auth["admin"] = "123456"
	h.Work["admin"] = 1000000
}

func (Handler) MainPage(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusOK
	// html в файл
	html, err := read("./html/main.html")
	if err != nil {
		statusCode = http.StatusInternalServerError
		// вывод ошибки вместо страницы
		html = err.Error()
	}

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	w.WriteHeader(statusCode)
	fmt.Fprint(w, html)
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusOK

	login := r.FormValue("login")
	pass := r.FormValue("pass")

	if h.Auth[login] == "" {
		user := &model.User{}
		err := user.Get(login, pass)

		if err != nil {
			statusCode = http.StatusForbidden
		} else {
			h.Auth[login] = pass
			h.Work[login] = user.WorkNumber
		}
	} else if h.Auth[login] != pass {
		statusCode = http.StatusForbidden
	}

	w.WriteHeader(statusCode)
}

func (h Handler) ChangePass(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusOK

	login := r.FormValue("login")
	pass := r.FormValue("pass")
	newPass := r.FormValue("newPass")

	if h.Auth[login] != "" && h.Auth[login] == pass {
		user := &model.User{}
		err := user.FindByLogin(login)
		if err != nil {
			statusCode = http.StatusInternalServerError
		}

		err = user.UpdatePass(newPass)
		if err != nil {
			statusCode = http.StatusBadRequest
		} else {
			// обновляем пароль
			h.Auth[login] = user.Pass
		}
	} else if h.Auth[login] == "" {
		statusCode = http.StatusUnauthorized
	} else {
		statusCode = http.StatusForbidden
	}

	w.WriteHeader(statusCode)
}

// весьма странная авторизация (по логину)
func (h Handler) Task(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusOK
	login := r.FormValue("login")

	if h.Work[login] <= 0 {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var value DTO
	err := json.Unmarshal([]byte(r.FormValue("value")), &value)
	if err != nil {
		// в т.ч. для вещественных чисел
		statusCode = http.StatusBadRequest
	}

	refVal := reflect.ValueOf(value)
	// буфер
	var buf bytes.Buffer
	// инвертированное значение
	var rev string
	n := refVal.NumField() - 1

	// т.к. явных требований к ответу нет, то формат будет произвольным:
	// key => reversed value
	for i := 0; i <= n; i++ {
		rev = fmt.Sprintf("%s: %s", refVal.Type().Field(i).Tag.Get("json"), h.reverse(refVal.Field(i)))
		// если строка не последняя
		if i < n {
			// то делаем перенос строки
			fmt.Fprintln(&buf, rev)
		} else {
			fmt.Fprint(&buf, rev)
		}
	}

	w.WriteHeader(statusCode)
	fmt.Fprint(w, buf.String())
}

// сейчас у нас всего 2 типа, которые жестко определны в DTO
// вероятно достаточным будет инверсия значений только 2-х этих типов
func (Handler) reverse(val reflect.Value) string {
	switch val.Type().Kind() {
	// хз, тербований нет - изначально был int32
	case reflect.Int32:
		i := int(val.Int())
		num := math.MaxInt32 - i
		return strconv.Itoa(num)
	case reflect.String:
		// вероятно самый эффективный способ инверсии строк
		s := strings.Split(val.String(), "")
		n := len(s)
		k := n / 2
		for i := 0; i < k; i++ {
			s[i], s[n-1-i] = s[n-1-i], s[i]
		}
		return strings.Join(s, "")
	}

	return ""
}

// для удобства
func read(file string) (string, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

var debug = config.GetEnv("TASK_DEBUG", "").Bool()

// конструктор
func NewHandler() *Handler {
	hdl := &Handler{
		Auth: make(map[string]string),
		Work: make(map[string]int32),
	}

	// если переменная окружение DEBUG=true|1
	if debug {
		hdl.init()
	}

	return hdl
}
