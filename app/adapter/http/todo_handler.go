package http

import (
	"app/app/infrastructure/postgresql/models"
	"app/app/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type todoHandler struct {
	usecase usecase.ITodoUsecase
}

func NewTodoHandler(usecase usecase.ITodoUsecase) *todoHandler {
	return &todoHandler{
		usecase: usecase,
	}
}

func (h *todoHandler) FindAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		todo, err := h.usecase.FindAll(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, todo)
	}
}

func (h *todoHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		p := new(models.Todo)
		if err := c.Bind(p); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		todo, err := h.usecase.Create(c.Request().Context(), p)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, todo)
	}
}

func (h *todoHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		p := new(models.Todo)
		if err := c.Bind(p); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		todo, err := h.usecase.Update(c.Request().Context(), p)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, todo)
	}
}

func (h *todoHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		err = h.usecase.Delete(c.Request().Context(), id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, "deleted")
	}
}
