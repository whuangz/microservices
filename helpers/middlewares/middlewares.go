package middlewares

import "github.com/labstack/echo/v4"

type CoreMiddlewares struct{}

func InitMiddleware() *CoreMiddlewares {
	return &CoreMiddlewares{}
}

func (c *CoreMiddlewares) CORS(next echo.HandleFunc) echo.HandleFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}
