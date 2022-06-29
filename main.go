package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/vcokltfre/volcan/src/config"
	"github.com/vcokltfre/volcan/src/core"
	"github.com/vcokltfre/volcan/src/database"
)

func init() {
	godotenv.Load()

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {
	err := database.SetupDB()
	if err != nil {
		logrus.Fatal(err)
	}

	err = config.LoadConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Println(config.Config)

	err = core.SetupSession()
	if err != nil {
		logrus.Fatal(err)
	}

	defer logrus.Info("Shutdown complete.")
	defer core.Session.Close()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	<-exit
}
