package handlers

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	html, err := read("./html/main.html")
	statusCode := http.StatusOK
	if err != nil {
		statusCode = http.StatusInternalServerError
		html = err.Error()
	}

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	w.WriteHeader(statusCode)
	fmt.Fprint(w, html)
}

func read(file string) (string, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
