package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"go-aws-eb/model"

	"github.com/joho/godotenv"
)

func serviceEnv() (properties model.Properties) {
	isLocalDev := flag.Bool("local", false, "=(true/false)")
	flag.Parse()
	if *isLocalDev {
		return loadEnvFile()
	} else {
		return loadEnv()
	}
}

func loadEnvFile() model.Properties {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(err)
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
	properties := loadEnv()
	return properties
}

func loadEnv() model.Properties {
	timestart := time.Now()
	fmt.Printf("Starting Load Config %s\n", timestart.Format("2006-01-02 15:04:05"))
	properties := model.Properties{
		ServicePort: os.Getenv("SERVICE_PORT"),
		LogPath:     os.Getenv("LOG_PATH"),
	}
	timefinish := time.Now()
	fmt.Printf("Finish Load Config %s\n", timefinish.Format("2006-01-02 15:04:05"))
	return properties
}
