package framework

import (
	"context"
	"fmt"
	"log"
	"time"
)

func TimeoutHandler(fun ControllerHandler, d time.Duration) ControllerHandler {
	return func(c *Context) error {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		durationCtx, cancel := context.WithTimeout(c.BaseContext(), d)
		defer cancel()

		c.req.WithContext(durationCtx)

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()

			fun(c)

			finish <- struct{}{}
		}()

		select {
		case p := <-panicChan:
			c.WriteMux().Lock()
			defer c.WriteMux().Unlock()

			log.Println(p)
			c.Json(500, "panic")
		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			c.WriteMux().Lock()
			defer c.WriteMux().Unlock()

			c.Json(500, "time out")
			c.SetHasTimeout()
		}

		return nil
	}
}
