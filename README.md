## Установка
```
git clone https://github.com/partyzanex/test-task
cd ./test-task
// копируем файл конфигурации
cp ./config-sample.ini ./config.ini
```

Редактируем ./config ini
```ini
# настройки бд
[Database]
dialect=postgres
hostname=localhost
port=5432
username=
password=
dbname=
sslmode=disable
```

```bash
// уставнока миграции
./migration -tables -data
```

```bash
make deps && make build
// запуск в тестовом режиме
make run-dev
```

Теперь можно открыть http://localhost:8080

### Тестовые запросы

Авторизация
```bash
curl -i -X POST -d "login=test&pass=1234" http://localhost:8080/login

// output
HTTP/1.1 200 OK
Date: Mon, 20 Nov 2017 09:14:17 GMT
Content-Length: 0
Content-Type: text/plain; charset=utf-8

curl -i -X POST -d "login=test&pass=43210" http://localhost:8080/login

// output
HTTP/1.1 403 Forbidden
...
```

Смена пароля
```bash
curl -i -X PUT -d "login=test&pass=1234&newPass=qwerty" http://localhost:8080/login/pass

// output
HTTP/1.1 200 OK
...

curl -i -X POST -d "login=test&pass=1234" http://localhost:8080/login

// output
HTTP/1.1 403 Forbidden
...

curl -i -X POST -d "login=test&pass=qwerty" http://localhost:8080/login

// output
HTTP/1.1 200 OK
...
```

Отправка задания
```bash
curl -i -X POST -d "login=test&value={\"big_number\": 1154543234, \"text\": \"0123456789QWERTY абвндеёжзийклмн ABCDEFGH\"}" http://localhost:8080/task

// output
HTTP/1.1 200 OK
Date: Mon, 20 Nov 2017 09:23:42 GMT
Content-Length: 84
Content-Type: text/plain; charset=utf-8

big_number: 992940413
text: HGFEDCBA нмлкйизжёеднвба YTREWQ9876543210

curl -i -X POST -d "login=admin&value={\"big_number\":  0.5, \"text\": \"Content-Length\"}" http://localhost:8080/task

// output
HTTP/1.1 400 Bad Request
Date: Mon, 20 Nov 2017 09:26:20 GMT
Content-Length: 34
Content-Type: text/plain; charset=utf-8

big_number: 2147483647
text: htgneL-tnetnoC

curl -i -X POST -d "login=user&value={\"big_number\":  1234, \"text\": \"Content-Length\"}" http://localhost:8080/task

// output
HTTP/1.1 403 Forbidden
...

```