package main

import (
	adaptorHTTP "app/app/adapter/http"
	"app/app/infrastructure/postgresql"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Hello, World!"})
}

func main() {
	boil.DebugMode = true

	postgresConn := postgresql.NewPostgreSQLConnector()
	defer postgresConn.Conn.Close()

	e := adaptorHTTP.InitRouter(postgresConn.Conn)
	e.Logger.Fatal(e.Start(":3000"))
}
