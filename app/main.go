package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	// "database/sql"
	"fmt"
	"net/http"
)

var db *gorm.DB
var err error

func main() {

	db, err = gorm.Open("mysql", "docker:docker@tcp(db:3306)/test_db?parseTime=true")
	if err != nil {
		http.HandleFunc("/", handlerErr)
		panic("failed to connect database")
	}
	createTable()
	r := gin.Default()
	users := r.Group("/users")
	{
		users.GET("/:id", getUser)
		users.POST("/edit", func(c *gin.Context) {
			c.JSON(200, gin.H{"name": "user.Name", "age": "user.Age"})
		})
	}
	r.Run()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Kết nối CDSL thành công ")
}

func handlerErr(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Lỗi kết nối CSDL ")
}

type User struct {
	gorm.Model
	Name string
	Age  int64
}

func createTable() {
	db.Exec("TRUNCATE TABLE users;")
	db.AutoMigrate(&User{})
	db.Create(&User{Name: "phule", Age: 25})
	db.Create(&User{Name: "dungdt", Age: 28})
	db.Create(&User{Name: "tuantt", Age: 25})
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	db.First(&user, id)
	log.Println(user)
	log.Println(user.Name)
	log.Println(user.Age)
	c.JSON(200, gin.H{"name": user.Name, "age": user.Age})
}
