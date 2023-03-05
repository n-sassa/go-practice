package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type healthCheckResponse struct {
	Message string `json:"message"`
}

// for health check
//
//	{
//	    "message": "Hello, C Team. you've requested: /health_check"
//	}
//
// will return
func healthCheck(c echo.Context) error {
	now := time.Now()
	message := fmt.Sprintf("Hello! you've requested: %s, time: %s", c.Path(), now)
	return c.JSON(
		http.StatusOK,
		healthCheckResponse{
			Message: message,
		},
	)
}
