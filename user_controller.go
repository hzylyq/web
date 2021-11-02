package main

import (
	"github.com/hzy/web/framework"
	"net/http"
)

func UserLoginController(c *framework.Context) error {
	c.Json(http.StatusOK, "ok, UserLoginController")
	return nil
}
