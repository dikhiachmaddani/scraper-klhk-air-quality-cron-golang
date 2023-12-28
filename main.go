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

	"github.com/robfig/cron/v3"
)

func main() {
	dbConfig := config.DatabaseConfig{
		User:     "root",
		Password: "root",
		Host:     "localhost",
		Port:     "3306",
		DBName:   "cron-ispu",
	}
	db, err := config.SetupDB(dbConfig)
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&models.QualityAir{})

	r := routes.SetupRoutes(db)
	go func() {
		err := r.Run(":2000")
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
