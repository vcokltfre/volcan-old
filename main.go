package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/vcokltfre/volcan/src/database"
)

func init() {
	godotenv.Load()
}

func main() {
	err := database.SetupDB()
	if err != nil {
		logrus.Fatal(err)
	}
}
