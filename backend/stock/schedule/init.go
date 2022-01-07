package schedule

import (
	"stock/services"

	"github.com/robfig/cron/v3"
)

var scheduler *cron.Cron

// Init Init
func Init() (err error) {
	scheduler = cron.New()
	defer scheduler.Start()

	stockService := services.CreateStockService()
	scheduler.AddFunc("0 16 * * 1-5", func() {
		stockService.DailyParser()
	})
	// stockService.DailyCalc()
	// stockService.DailyParser()

	return
}

// GetScheduler GetScheduler
func GetScheduler() *cron.Cron {
	if scheduler == nil {
		Init()
	}

	return scheduler
}
