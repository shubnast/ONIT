package Methods

import (
	"kursach/CORS"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Доп структура для обновлнеия таблицы
type UpdatePersonInput struct {
	Name     string `json:"marka"`
	LastName string `json:"model"`
}

//Метод для получения всех данных таблицы авто
func GetAllAuto(c *gin.Context) {
	gormDB := CORS.ConnectDB()
	var autos []CORS.Person
	gormDB.Find(&autos)
	c.JSON(http.StatusOK, autos)
}

//Метод для добавления авто в таблтицу
func AddAuto(c *gin.Context) {
	gormDB := CORS.ConnectDB()
	var input CORS.Person
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	auto := CORS.Person{Name: input.Name, LastName: input.LastName}
	gormDB.Create(&auto)
	c.JSON(http.StatusOK, gin.H{"message": "Auto added"})

}

//Метод для обновления авто в таблице
func UpdateAuto(c *gin.Context) {
	gormDB := CORS.ConnectDB()
	var auto CORS.Person
	if err := gormDB.Where("id = ?", c.Param("id")).First(&auto).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	var input UpdatePersonInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gormDB.Model(&auto).Update("name", input.Name)
	gormDB.Model(&auto).Update("last_name", input.LastName)
	c.JSON(http.StatusOK, gin.H{"message": "Person updated"})

}

//Метод для удаления авто из таблицы
func DelAuto(c *gin.Context) {
	gormDB := CORS.ConnectDB()
	var auto CORS.Person
	if err := gormDB.Where("id = ?", c.Param("id")).First(&auto).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	gormDB.Delete(&auto)

	c.JSON(http.StatusOK, gin.H{"autos": true})
}
