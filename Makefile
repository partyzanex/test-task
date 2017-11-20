deps:
	go get github.com/astaxie/beego
	go get github.com/gorilla/mux
	go get github.com/jinzhu/gorm
	go get github.com/lib/pq
build:
	CGO_ENABLED=0 GOOS=linux go build -o main ./main.go
	CGO_ENABLED=0 GOOS=linux go build -o migration ./migrate/main.go	
run-dev:
	TASK_ADDR=localhost:8080 TASK_DEBUG=true ./main
run:
	./main 
clean:
	./migration -clean
	rm ./main ./migration