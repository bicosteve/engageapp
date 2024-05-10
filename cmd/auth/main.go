package main

import (
	"os"

	"github.com/engageapp/pkg/controllers"
	"github.com/engageapp/pkg/utils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		utils.Log("ERROR", "env", "error loading env file because of %v", err)
		os.Exit(1)
	}

	var base controllers.Base

	base.Init()
	base.RunAuth()
}
