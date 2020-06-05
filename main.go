package main

import (
	"awesomeProject/model/db"
	"awesomeProject/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.GetPing(r)
	router.GetPing2(r)
	defer db.GetDB().Close()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	fmt.Println("end..........")
}


