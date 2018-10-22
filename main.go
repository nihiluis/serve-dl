package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r.Static("/", "./dl")
	fmt.Println("Serving on :48893")
	r.Run(":48893") // listen and serve on 0.0.0.0:8080
}
