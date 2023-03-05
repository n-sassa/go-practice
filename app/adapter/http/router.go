package http

import (
	"app/app/domain/service"
	"app/app/infrastructure/postgresql/repository"
	"app/app/usecase"
	"database/sql"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter(conn *sql.DB) *echo.Echo {
	e := echo.New()
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORS(),
	)

	healthCheckGroup := e.Group("")
	{
		relativePath := ""
		healthCheckGroup.GET(relativePath, healthCheck)
	}

	userRepository := repository.NewUserRepository(conn)
	userService := service.NewUserService(userRepository)
	userUsecase := usecase.NewUserUsecase(userService)

	// ログイン
	loginGroup := e.Group("/login")
	{
		handler := NewUserHandler(userUsecase)
		loginGroup.POST("", handler.Login())
	}

	// ユーザ
	userGroup := e.Group("/user")
	{
		handler := NewUserHandler(userUsecase)
		// 新規登録
		userGroup.POST("", handler.Create())
		// ユーザ一覧 ※clientでは使わない
		relativePath := fmt.Sprintf("/:%s", "id")
		userGroup.GET(relativePath, handler.FindByID())
	}

	todoRepository := repository.NewTodoRepository(conn)
	todoService := service.NewTodoService(todoRepository)
	todoUsecase := usecase.NewTodoUsecase(todoService)

	// Todoリスト
	todoGroup := e.Group("/todos")
	{
		handler := NewTodoHandler(todoUsecase)
		// 認証周りは実装しないため全ユーザのTodoを取得
		todoGroup.GET("", handler.FindAll())
		todoGroup.POST("", handler.Create())
		todoGroup.PUT("", handler.Update())
		relativePath := fmt.Sprintf("/:%s", "id")
		todoGroup.DELETE(relativePath, handler.Delete())
	}

	return e
}
