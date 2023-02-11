package http

import (
	"app/app/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type post struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

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
		p := new(post)
		if err := c.Bind(p); err != nil {
			log.Printf("err %v", err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, "created")
	}
}
