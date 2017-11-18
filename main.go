package main

import (
	//"encoding/binary"
	//"encoding/json"
	"net/http"
	//"reflect"


	"github.com/gorilla/mux"
	"github.com/partyzanex/test-task/config"
	"time"
	"log"
	"github.com/partyzanex/test-task/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.MainPage)
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/login/pass", handlers.ChangePass).Methods("POST")

	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         config.GetEnv("ADDR", ":8080").String(),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

//func init() {
//	Work := make(map[string]int32, 0)
//	Work["admin"] = 1000000
//}
//
//
//
//
//
//type DTO struct {
//	bigNumber int64
//	text      string
//}

//func doWork(w http.ResponseWriter, r *http.Request) {
//	var value DTO
//	login := r.FormValue("login")
//	if Work[login] <= 0 {
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	json.Unmarshal([]byte(r.FormValue("value")), &value)
//
//	v := reflect.ValueOf(value)
//	for i := 0; i < v.NumField(); i++ {
//		w.Write(reverse(v.Elem().Field(i)))
//	}
//}
//
//func reverse(val reflect.Value) []byte {
//	switch val.Kind().String() {
//	case "int64":
//		fallthrough
//	case "int32":
//		result := make([]byte, 4)
//		binary.LittleEndian.PutUint32(result, uint32(2147483647-val.Interface().(int32)))
//		return result
//	case "string":
//		var result string
//		for i := len(val.Interface().(string)); i > 0; i++ {
//			result += string(val.Interface().(string)[i])
//		}
//		return []byte(result)
//	}
//	return nil
//}
