package main

import (
	"fmt"
	"os"
	"os/signal"
	"quality-air-golang/config"
	"quality-air-golang/cronschedule"
	"quality-air-golang/models"
	"quality-air-golang/routes"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	godotenv.Load(".env")
	dbConfig := config.DatabaseConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
	}
	db, err := config.SetupDB(dbConfig)
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&models.QualityAir{})

	r := routes.SetupRoutes(db)
	go func() {
		err := r.Run(":8080")
		if err != nil {
			fmt.Println("Error running routes:", err)
		}
	}()

	// set scheduler berdasarkan zona waktu sesuai kebutuhan
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron.New(cron.WithLocation(jakartaTime))

	// stop scheduler tepat sebelum fungsi berakhir
	defer scheduler.Stop()
	cronschedule.QualityCronSchedule(scheduler, db)

	// start scheduler
	go func() {
		scheduler.Start()
	}()

	// trap SIGINT untuk trigger shutdown.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

}
