package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hzy/web/framework"
)

func FooControlHandler(c *framework.Context) error {
	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	durationCtx, cancel := context.WithTimeout(c.BaseContext(), 1*time.Second)
	defer cancel()

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()

		time.Sleep(10 * time.Second)
		c.Json(200, "ok")

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
