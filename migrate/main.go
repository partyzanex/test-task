package main

import (
	"github.com/partyzanex/test-task/model"
	"flag"
	"log"
)

var (
	tables = flag.Bool("tables", false, "tables migrate")
	data = flag.Bool("data", false, "data import")
	clean = flag.Bool("clean", false, "remove data")
)


func main() {
	flag.Parse()

	db, err := model.NewDBConnection()
	if err != nil {
		panic(err)
	}

	if *tables {
		err = db.AutoMigrate(&model.User{}).Error
		if err != nil {
			panic(err)
		}
		log.Println("Tables was been created")
	}

	if *data {
		user := &model.User{
			Login: "test",
			Pass: "1234",
			WorkNumber: 100000,
		}

		err = db.Save(user).Error
		if err != nil {
			panic(err)
		}
		log.Println("Test user was been created")
	}

	if *clean {
		err = db.Delete(&model.User{}, "login = ?", "test").Error
		if err != nil {
			panic(err)
		}
		log.Println("Test user was been removed")
	}
}
