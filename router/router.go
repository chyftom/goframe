package router

import (
	db "awesomeProject/base"
	"awesomeProject/model"
	"github.com/gin-gonic/gin"
)


func GetPing(r *gin.Engine){

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	 //
	 //model.GetBook()
	 //model.GetBook2()
		db.GetProduct()
		//fmt.Println(c)

	})
}

func GetPing2(r *gin.Engine) {

	r.GET("/ping2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
		model.GetUserByID(1)

	})
}

