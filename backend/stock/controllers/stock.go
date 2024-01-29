package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"stock/config"
	"stock/database"
	"stock/models"
	"stock/services"
	"stock/websocket"
	"strings"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StockController struct {
	serv *services.StockService
}

// NewStockController NewStockController
func NewStockController(service *services.StockService) *StockController {
	return &StockController{
		serv: service,
	}
}

// GetInfo GetInfo
func (c *StockController) StockWs(ctx *gin.Context) {
	if user, ok := ctx.MustGet("user").(*models.User); ok {
		client, err := websocket.NewClient(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 傳送歡迎訊息
		msg := &models.WsToClientData{}
		msg.Data = fmt.Sprintf("Welcome, %s", user.Name)
		message, _ := json.Marshal(&msg)
		client.Send(message)

		// 設定接收Channel
		go c.serv.ReadWs(client.Context(), client)
	}
}

func (c *StockController) TgMessage(ctx *gin.Context) {
	defer ctx.Request.Body.Close()
	bot, err := tgbotapi.NewBotAPI(config.Telegram.Token)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	bytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var update tgbotapi.Update
	err = json.Unmarshal(bytes, &update)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	if update.Message != nil {
		if update.Message.From != nil {
			if !update.Message.From.IsBot {
				if update.Message.Chat != nil {
					// 開始解析指令
					originString := strings.TrimSpace(update.Message.Text)
					_, i := utf8.DecodeRuneInString(originString)
					fullCommand := originString[i:]

					replyMsg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
					replyMsg.ReplyToMessageID = update.Message.MessageID

					// 用空格來切
					// [0]新增/刪除/清單
					// [1]看板名稱
					// [2]關鍵字
					commands := strings.Split(fullCommand, " ")
					if len(commands) >= 1 {
						db := database.GetDB()
						action := commands[0]
						if action == "清單" {
							tgPttAlerts := []*models.TgPttAlert{}
							whereCond := map[string]interface{}{
								"chart_id": fmt.Sprintf("%d", update.Message.Chat.ID),
							}
							err := db.Where(whereCond).Find(&tgPttAlerts).Error
							replyMsg.Text = ""
							if err != nil {
								replyMsg.Text = `訂閱清單獲取失敗`
							} else {
								for _, tgPttAlert := range tgPttAlerts {
									replyMsg.Text += fmt.Sprintf("看板:%s,關鍵字:%s\n", tgPttAlert.Kanban, tgPttAlert.KeyWord)
								}
								fmt.Println(replyMsg.Text, "LIST")
								replyMsg.ParseMode = tgbotapi.ModeMarkdownV2
							}
							_, err = bot.Send(replyMsg)
							if err != nil {
								fmt.Println(err.Error())
								ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
								return
							}

							ctx.JSON(http.StatusOK, nil)
							return
						}
						if len(commands) >= 3 {
							kanban := commands[1]
							keyword := commands[2]
							if action == "新增" {
								tgPttAlert := &models.TgPttAlert{
									ChartID: fmt.Sprintf("%d", update.Message.Chat.ID),
									Kanban:  kanban,
									KeyWord: keyword,
								}
								err := db.Create(&tgPttAlert).Error
								replyMsg.Text = ""
								if err != nil {
									replyMsg.Text = `新增失敗`
								} else {
									replyMsg.Text = `新增成功`
								}
								_, err = bot.Send(replyMsg)
								if err != nil {
									fmt.Println(err.Error())
									ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
									return
								}

								ctx.JSON(http.StatusOK, nil)
								return
							}
							if action == "刪除" {
								tgPttAlert := &models.TgPttAlert{}
								whereCond := map[string]interface{}{
									"chart_id": fmt.Sprintf("%d", update.Message.Chat.ID),
									"kanban":   kanban,
									"key_word": keyword,
								}
								err := db.Where(whereCond).First(&tgPttAlert).Error
								replyMsg.Text = ""
								if err != nil {
									replyMsg.Text = `刪除失敗,找不到相關設定`
								} else {
									err = db.Delete(&tgPttAlert).Error
									if err != nil {
										replyMsg.Text = `刪除失敗,內部問題`
									}
								}
								replyMsg.Text = `刪除成功`
								_, err = bot.Send(replyMsg)
								if err != nil {
									fmt.Println(err.Error())
									ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
									return
								}

								ctx.JSON(http.StatusOK, nil)
								return
							}
						}
					}
					replyMsg.Text = "新增指令\n`/新增 看板名稱 關鍵字`\n刪除指令\n`/刪除 看板名稱 關鍵字`\n清單指令\n`/清單`"
					replyMsg.ParseMode = tgbotapi.ModeMarkdownV2
					_, err := bot.Send(replyMsg)
					if err != nil {
						fmt.Println(err.Error())
						ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
						return
					}
				}
			}
		}
	}
	fmt.Println(string(bytes))
	ctx.JSON(http.StatusOK, nil)
}
