package http

import (
	"app/app/infrastructure/postgresql/models"
	"app/app/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	usecase usecase.IUserUsecase
}

func NewUserHandler(usecase usecase.IUserUsecase) *userHandler {
	return &userHandler{
		usecase: usecase,
	}
}

func (h *userHandler) FindByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		user, err := h.usecase.FindByID(c.Request().Context(), id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, user)
	}
}

func (h *userHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		p := new(models.User)
		if err := c.Bind(p); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		err := h.usecase.Create(c.Request().Context(), p)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, "created")
	}
}

func (h *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		p := new(models.User)
		if err := c.Bind(p); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		isExists, err := h.usecase.Login(c.Request().Context(), p.Name, p.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		if isExists {
			return c.JSON(http.StatusOK, true)
		}
		return c.JSON(http.StatusBadRequest, "ユーザが存在しないかパスワードが誤っています。")
	}
}
