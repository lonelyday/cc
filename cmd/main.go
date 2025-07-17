package main

import (
	"github.com/lonelyday/cc/internal/api"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/rates", api.Rates)
	err := router.Run(":9090")
	if err != nil {
		log.Fatal(err)
	}

}
