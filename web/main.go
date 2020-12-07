package main

import (
	"github.com/ytz4178/BCALib-go/router"
)

func main() {
	rt := router.RouterInit()

	rt.Run(":9999")
}
