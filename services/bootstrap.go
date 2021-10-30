package services

import (
	"BackLight/database"
	"github.com/joho/godotenv"
	"log"
)

func Bootstrap()  {
	initDotEnv()
	database.Init()
	database.InitMigrations()
}

func initDotEnv()  {
	err := godotenv.Load()
	if err != nil{
		log.Fatal(".env Error: ", err)
	}
}