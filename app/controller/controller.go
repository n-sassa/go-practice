package controller

import (
	"app/app/config"
	"app/app/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	b := new(models.User)
	db := config.DB()

	// Binding data
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	user := &models.User{
		Name:     b.Name,
		Password: b.Password,
	}

	if err := db.Create(&user).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": b,
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	b := new(models.User)
	db := config.DB()

	// Binding data
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existing_user := new(models.User)

	if err := db.First(&existing_user, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusNotFound, data)
	}

	existing_user.Name = b.Name
	existing_user.Password = b.Password
	if err := db.Save(&existing_user).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": existing_user,
	}

	return c.JSON(http.StatusOK, response)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	var users []*models.User

	if res := db.Find(&users, id); res.Error != nil {
		data := map[string]interface{}{
			"message": res.Error.Error(),
		}

		return c.JSON(http.StatusOK, data)
	}

	response := map[string]interface{}{
		"data": users[0],
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	user := new(models.User)

	err := db.Delete(&user, id).Error
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "a user has been deleted",
	}
	return c.JSON(http.StatusOK, response)
}
