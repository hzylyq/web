package main

import "github.com/hzy/web/framework"

func registerRouter(core *framework.Core) {
	core.Set("foo", FooControlHandler)
}
