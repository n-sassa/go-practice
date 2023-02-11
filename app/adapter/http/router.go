package http

import (
	"app/app/domain/service"
	"app/app/infrastructure/postgresql"
	"app/app/infrastructure/postgresql/repository"
	"app/app/usecase"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter() *echo.Echo {
	e := echo.New()
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	healthCheckGroup := e.Group("/health_check")
	{
		relativePath := ""
		healthCheckGroup.GET(relativePath, healthCheck)
	}

	postgresConn := postgresql.NewPostgreSQLConnector()
	UserRepository := repository.NewUserRepository(postgresConn.Conn)
	userService := service.NewUserService(UserRepository)
	userUsecase := usecase.NewUserUsecase(userService)

	userGroup := e.Group("/user")
	{
		handler := NewUserHandler(userUsecase)
		relativePath := fmt.Sprintf("/:%s", "id")
		userGroup.GET(relativePath, handler.FindByID())
	}
	return e
}
