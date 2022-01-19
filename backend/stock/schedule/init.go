package schedule

import (
	"stock/concrete"
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
	stockConcrete := concrete.CreateStockConcrete()
	scheduler.AddFunc("15 16 * * 1-5", func() {
		stockConcrete.DailyCalc()
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
