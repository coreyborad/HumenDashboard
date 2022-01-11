package concrete

import "stock/services"

// StockConcrete StockConcrete
type StockConcrete struct {
	stockServ *services.StockService
}

// NewStockConcrete New StockConcrete
func NewStockConcrete(
	stockServ *services.StockService,
) *StockConcrete {
	return &StockConcrete{
		stockServ: stockServ,
	}
}

func (c *StockConcrete) CheckHistory() (interface{}, error) {

	return nil, nil
}
