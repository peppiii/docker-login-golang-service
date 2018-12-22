package main

import (
	"fmt"
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
//	"log"
	"crypto/md5"
)


func Database() *gorm.DB{

	//open a db connection
	db_user   := os.Getenv("DB_USER")
	db_pass   := os.Getenv("DB_PASS")
	db_host   := os.Getenv("DB_HOST")
	db_port   := os.Getenv("DB_PORT")
	db, err := gorm.Open("mysql", db_user+":"+db_pass+"@tcp("+db_host+":"+db_port+")/kubernetes")
	if err != nil{
		fmt.Println(err.Error())
	}
	return db
}

// Binding from JSON
type logins struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func GetLogin(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"status": "Selamat Datang Harap Login"})
}

func digestString(s string) string {
    return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func PostLogin(c *gin.Context) {
	var json logins
		if err := c.ShouldBindJSON(&json); err != nil {
                        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                        return
                }
		db := Database()
		if err := db.Where("username = ? and password = ?", json.Username, digestString(json.Password)).Find(&json).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),"status" : "404"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"messages": "You are logged in", "status": "200"})
                }


func main() {

	//router := gin.Default()
        router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.POST("/login", PostLogin)
	router.GET("/login", GetLogin)
	router.Run(":8082")
}
