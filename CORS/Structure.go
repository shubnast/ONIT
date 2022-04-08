package CORS

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//наша структура автомобилей
type Person struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"marka"`
	LastName string `json:"model"`
}

//подключение к базе данных
func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=hoplaylala dbname=ONIT port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}
	//Миграция нашей структуры в бд
	db.AutoMigrate(&Person{})
	return db
}
