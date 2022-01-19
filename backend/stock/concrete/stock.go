package concrete

import (
	"fmt"
	"stock/config"
	"stock/models"
	"stock/services"
	"stock/utils"
	"strconv"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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
	end := time.Now()
	start := end.AddDate(0, -13, 0)
	type StockInfo struct {
		HasStock bool
		Cost     float64
		BuyDate  time.Time
	}
	stockInfo := &StockInfo{
		HasStock: false,
		Cost:     0,
	}
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		stockTechVal, err := c.stockServ.Calc("0050", &d)
		if err != nil {
			break
		}
		// fmt.Println(fmt.Sprintf("Date %s, Action %s", d.String(), stockTechVal.Action))
		if stockTechVal.Action == "buy" {
			stockInfo.HasStock = true
			stockInfo.Cost = stockTechVal.PriceOnClose
			stockInfo.BuyDate = d
		}
		if stockTechVal.Action == "sell" && stockInfo.HasStock {
			stockInfo.HasStock = false
			earn := stockTechVal.PriceOnClose - stockInfo.Cost
			printString := fmt.Sprintf("Buy on %s, Sell on %s, Earn %f", stockInfo.BuyDate.String(), d.String(), earn)
			fmt.Println(printString)
		}
	}
	return nil, nil
}

func (c *StockConcrete) DailyCalc() error {
	now := time.Now()
	html := fmt.Sprintf(`
	<meta charset="utf-8">
	<style type="text/css">
	.tg  {border-collapse:collapse;border-spacing:0;}
	.tg td{border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;
	  overflow:hidden;padding:10px 5px;word-break:normal;}
	.tg th{border-color:black;border-style:solid;border-width:1px;font-family:Arial, sans-serif;font-size:14px;
	  font-weight:normal;overflow:hidden;padding:10px 5px;word-break:normal;}
	.tg .tg-0pky{border-color:inherit;text-align:left;vertical-align:top}
	</style>
	<table class="tg">
	<thead>
		<tr>
			<th class="tg-0lax" colspan="7">Date %s</th>
		</tr>
	</thead>
	<tbody>
	  <tr>
		<td class="tg-0pky">股票編號</td>
		<td class="tg-0pky">收盤價</td>
		<td class="tg-0pky">K</td>
		<td class="tg-0pky">D</td>
		<td class="tg-0pky">上漲交叉</td>
		<td class="tg-0pky">下跌交叉</td>
		<td class="tg-0pky">動作</td>
	  </tr>
	</tbody>
	<tbody>
	`, now.Format("2006/01/02"))

	// calc stock
	stockList := []string{
		"0050", "2330", "0056", "2412",
	}
	m := sync.Map{}
	wg := sync.WaitGroup{}
	fmt.Println("Start calc")
	for _, stockNumber := range stockList {
		wg.Add(1)
		go func(stockNumber string) {
			stockTechVal, err := c.stockServ.Calc(stockNumber, &now)
			if err != nil {
				fmt.Println("error:", err)
				return
			}
			m.Store(stockNumber, stockTechVal)
			defer wg.Done()
		}(stockNumber)
	}
	wg.Wait()
	for _, stockNumber := range stockList {
		stockTechVal, ok := m.Load(stockNumber)
		if !ok {
			continue
		}
		// Mapping html
		stockHtml := fmt.Sprintf(
			`
					<tr>
						<td class="tg-0pky">%s</td>
						<td class="tg-0pky">%f</td>
						<td class="tg-0pky">%f</td>
						<td class="tg-0pky">%f</td>
						<td class="tg-0pky">%t</td>
						<td class="tg-0pky">%t</td>
						<td class="tg-0pky">%s</td>
					</tr>
					`,
			stockNumber,
			stockTechVal.(models.StockTechVal).PriceOnClose,
			stockTechVal.(models.StockTechVal).KDVal.KVal,
			stockTechVal.(models.StockTechVal).KDVal.DVal,
			stockTechVal.(models.StockTechVal).KDVal.Uppercross,
			stockTechVal.(models.StockTechVal).KDVal.Undercross,
			stockTechVal.(models.StockTechVal).Action,
		)
		html = fmt.Sprintf("%s%s", html, stockHtml)
	}

	// Last html append
	html = fmt.Sprintf("%s</tbody></table>", html)
	// Generate photo
	imgPath, err := utils.HtmlToImage(html)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(imgPath)
	// TG bot
	bot, err := tgbotapi.NewBotAPI(config.Telegram.Token)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}
	chatId, err := strconv.Atoi(config.Telegram.Chats[0])
	if err != nil {
		fmt.Println("error:", err)
		return err
	}
	tgPhoto := tgbotapi.NewPhoto(int64(chatId), tgbotapi.FilePath(imgPath))
	_, err = bot.Send(tgPhoto)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}
	// Delete file
	utils.HtmlToImageDelete(imgPath)
	return nil
}
