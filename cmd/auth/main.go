package main

import (
	"github.com/engageapp/pkg/controllers"
)

func main() {

	var base controllers.Base

	base.Init()
	base.RunAuth()
}
