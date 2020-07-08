package main

import (
//	"fmt"
//	"os"
	"net/http"
//	"log"
        "github.com/gin-gonic/gin"
)


func GetLogin(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"status": "Selamat Datang Harap Login"})
}


func main() {

	//router := gin.Default()
        router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/login", GetLogin)
	router.Run(":8082")
}
