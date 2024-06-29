package main

import (
	"github.com/engageapp/pkg/controllers"
	"github.com/engageapp/pkg/utils"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		utils.Log("ERR", "logs", "Error loading logs because of %v", err)
	}

	var base controllers.Base

	base.Init()
	base.RunComment()

}
