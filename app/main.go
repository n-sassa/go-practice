package main

import (
	"app/app/config"
	"app/app/controller"
	"app/app/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"hello": "world!",
		})
	})
	// Connect To Database
	config.DatabaseInit()
	gorm := config.DB()
	gorm.AutoMigrate(&models.User{})

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()

	bookRoute := e.Group("/user")
	bookRoute.POST("/", controller.CreateUser)
	bookRoute.GET("/:id", controller.GetUser)
	bookRoute.PUT("/:id", controller.UpdateUser)
	bookRoute.DELETE("/:id", controller.DeleteUser)

	e.Logger.Fatal(e.Start(":3000"))
}
