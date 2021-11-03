package middleware

import (
	"log"
	"time"

	"github.com/hzy/web/framework"
)

func Cost() framework.ControllerHandler {
	return func(c *framework.Context) error {
		start := time.Now()

		c.Next()

		end := time.Now()
		cost := end.Sub(start)
		log.Printf("api uri:%v, cost:%v", c.Request().RequestURI, cost.Seconds())
		return nil
	}
}
