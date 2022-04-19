package main

import (
	"APIEmployee/employee"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database Connection Failed")
	}

	db.AutoMigrate(&employee.Employee{})

	repositoryEmployee := employee.NewRepositoy(db)
	serviceEmployee := employee.NewService(repositoryEmployee)
	handlerEmployee := employee.NewHandler(serviceEmployee)

	route := gin.Default()

	route.GET("/employee", handlerEmployee.Get)
	route.GET("/employee/:id", handlerEmployee.Find)
	route.POST("/employee", handlerEmployee.New)
	route.PUT("/employee/:id", handlerEmployee.Put)
	route.DELETE("/employee/:id", handlerEmployee.Destroy)

	route.Run()
}
