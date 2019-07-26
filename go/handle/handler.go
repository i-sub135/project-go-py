package handler

import (
	"net/http"

	"../model"
	"github.com/gin-gonic/gin"
)

func HandleIndex() gin.HandlerFunc {
	resp := map[string]string{}
	return func(res *gin.Context) {
		res.JSON(200, gin.H{
			"code":    200,
			"message": "Golang API",
			"result":  resp,
		})
	}
}

func GetCarfix(res *gin.Context) {
	var carfixs []model.Carfix

	mulai := res.Query("mulai")
	hal := res.Query("pages")

	db.Raw("SELECT * FROM tb_carfix LIMIT ?, ?", mulai, hal).Scan(&carfixs)
	res.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"message":  "OK",
		"rawCount": len(carfixs),
		"result":   carfixs,
	})
}
