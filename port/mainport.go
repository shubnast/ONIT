package port

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Метод для скана UPD
func Quest2(con *gin.Context) {

	widescanresults := WideScan("localhost")

	con.JSON(http.StatusOK, widescanresults)
}

//Метод для скана TCP
func Quest3(con *gin.Context) {
	scaner := WideScan1("localhost")
	con.JSON(http.StatusOK, scaner)
}
