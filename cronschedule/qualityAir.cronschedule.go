package cronschedule

import (
	"fmt"
	"quality-air-golang/controllers"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func QualityCronSchedule(scheduler *cron.Cron, db *gorm.DB) error {
	// Schedule cron job to get quality air data every minute
	scheduler.AddFunc("2 * * * *", func() {
		err := controllers.GetQualityAir(db)
		if err != nil {
			fmt.Println("Error getting quality air:", err)
		} else {
			fmt.Println("Data berhasil ditambahkan")
		}
	})

	return nil
}
