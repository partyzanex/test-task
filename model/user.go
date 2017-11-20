package model

import (
	"fmt"
	"crypto/md5"
	"io"
)

type User struct {
	ID         int64  `gorm:"AUTO_INCREMENT;primary key"`
	Login      string `gorm:"type:varchar;unique_index;not nil"`
	Pass       string `gorm:"type:varchar;not nil"`
	WorkNumber int32
}

func (u *User) Get(login string, pass string) error {
	// если login = "", то h.Auth[login] == "" => true - лишний запрос в бд
	// проверка логина и пароля
	// отсекаем очевидно не нужные запросы
	if login == "" {
		return fmt.Errorf("invalid login")
	}

	if pass == "" {
		return fmt.Errorf("invalid pass")
	}
	// чтобы избежать любые SQL-инъекции при авторизации
	// создается md5 хеш для сроки login+pass
	hash := md5.New()
	io.WriteString(hash, fmt.Sprintf("%s%s", login, pass))
	// where
	cond := fmt.Sprintf("md5(CONCAT(login, pass)) = '%x'", hash.Sum(nil))
	return DBConn.Where(cond).First(u).Error
}

// поиск пользователя в БД по логину
// предполагается что поле login - уникальный индекс
func (u *User) FindByLogin(login string) error {
	// на всякий случай переводим login в хеш
	hash := md5.New()
	io.WriteString(hash, login)
	return DBConn.Where(fmt.Sprintf("md5(login) = '%x'", hash.Sum(nil))).First(u).Error
}

// обновление пароля
// todo: хранить пароль в зашифрованном виде
func (u *User) UpdatePass(pass string) error {
	if pass == "" {
		return fmt.Errorf("password is empty")
	}

	u.Pass = pass
	return DBConn.Model(u).Update("pass", pass).Error
}
