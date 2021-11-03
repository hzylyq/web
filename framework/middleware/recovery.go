package middleware

import (
	"net/http"

	"github.com/hzy/web/framework"
)

func Recovery() framework.ControllerHandler {
	return func(c *framework.Context) error {
		defer func() {
			if err := recover(); err != nil {
				c.Json(http.StatusInternalServerError, err)
			}
		}()

		c.Next()

		return nil
	}
}
