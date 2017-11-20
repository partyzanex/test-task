package main

import (
	"net/http"
	"time"
	"log"

	"github.com/gorilla/mux"

	"github.com/partyzanex/test-task/handler"
	"github.com/partyzanex/test-task/config"
	"github.com/partyzanex/test-task/model"
)

// конект к БД
var _, errDb = model.NewDBConnection()

func main() {
	// усли конект к БД не произошел
	if errDb != nil {
		panic(errDb)
	}

	r := mux.NewRouter()

	// хендлер (контроллер)
	h := handler.NewHandler()

	r.HandleFunc("/", h.MainPage).Methods("GET")
	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/login/pass", h.ChangePass).Methods("PUT")
	r.HandleFunc("/task", h.Task).Methods("POST")

	// конфигурируем сервер
	srv := &http.Server{
		Handler: r,
		// хост:порт из окружения
		Addr: config.GetEnv("TASK_ADDR", ":8080").String(),
		// таймауты
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	log.Printf("Listen on %s", srv.Addr)
	// вывод ошибки
	log.Fatal(srv.ListenAndServe())
}
