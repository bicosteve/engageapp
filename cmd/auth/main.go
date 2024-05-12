package main

import (
	"os"

	"github.com/engageapp/pkg/controllers"
	"github.com/engageapp/pkg/utils"
	"github.com/joho/godotenv"
)

var base controllers.Base

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		utils.Log("ERROR", "env", "error loading env file because of %v", err)
		os.Exit(1)
	}

	base.RunAuth()

}

func main() {
	base.Init()
}
