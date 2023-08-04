package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/patrickmn/go-cache"
)

var localCache *cache.Cache

func GetPttContent(url string) *http.Response {
	ageCookie := &http.Cookie{
		Name:   "over18",
		Value:  "1",
		MaxAge: 300,
		Domain: "www.ptt.cc",
		Path:   "/",
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Got error while creating cookie jar %s", err.Error())
	}
	client := &http.Client{
		Jar: jar,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
		// handle error
	}
	req.AddCookie(ageCookie)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return resp
}

func GetPttImages(url string) []string {
	// 取得圖片
	respContent := GetPttContent(url)
	defer respContent.Body.Close()
	targetContent, err := goquery.NewDocumentFromReader(respContent.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	hrefList := []string{}
	targetContent.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			if strings.Contains(href, "https://i.imgur.com") {
				hrefList = append(hrefList, href)
			}
		}
	})
	return hrefList
}

type PttImage struct {
	ImgUrl  []string
	PageUrl string
	Nrec    uint64
}

func StartPtt() {
	resp := GetPttContent("https://www.ptt.cc/bbs/Beauty/index.html")
	defer resp.Body.Close()
	var doc *goquery.Document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := []*PttImage{}
	for {
		for _, node := range doc.Find(".r-ent").Nodes {
			if len(result) >= 3 {
				break
			}
			s := goquery.NewDocumentFromNode(node)
			titleNode := s.Find(".title")
			title := strings.TrimSpace(titleNode.Text())
			if !strings.Contains(title, "[正妹]") {
				continue
			}
			nrecNode := s.Find(".nrec span")
			nrecString := nrecNode.Text()
			if nrecString == "" {
				continue
			}
			thisGirl := &PttImage{}
			if nrecString == "爆" {
				pageHref, exists := titleNode.Find("a").Attr("href")
				if exists {
					thisGirl.PageUrl = pageHref
					thisGirl.ImgUrl = GetPttImages(fmt.Sprintf("https://www.ptt.cc%s", pageHref))
					thisGirl.Nrec = 999
					result = append(result, thisGirl)
					continue
				}
			}
			nrec, err := strconv.ParseUint(nrecString, 10, 64)
			if err != nil {
				continue
			}
			if nrec > 50 {
				pageHref, exists := titleNode.Find("a").Attr("href")
				if exists {
					thisGirl.PageUrl = fmt.Sprintf("https://www.ptt.cc%s", pageHref)
					thisGirl.ImgUrl = GetPttImages(thisGirl.PageUrl)
					thisGirl.Nrec = nrec
					result = append(result, thisGirl)
					continue
				}
			}
		}
		if len(result) >= 3 {
			break
		} else {
			found := false
			for _, n := range doc.Find(`a`).Nodes {
				ns := goquery.NewDocumentFromNode(n)
				pageLink, exists := ns.Attr("href")
				if exists && ns.Text() == "‹ 上頁" {
					resp := GetPttContent(fmt.Sprintf("https://www.ptt.cc%s", pageLink))
					defer resp.Body.Close()
					doc, err = goquery.NewDocumentFromReader(resp.Body)
					if err != nil {
						fmt.Println(err)
						return
					}
					found = true
					break
				}
			}

			if !found {
				break
			}
		}
	}

	localCache.Set("img", result, time.Minute*10)
}

type CarPage struct {
	ImgUrl  string
	PageUrl string
	Title   string
}

func StartCar() {
	doc, err := goquery.NewDocument("https://gallery.u-car.com.tw/galleries")
	if err != nil {
		fmt.Println(err)
		return
	}
	result := []*CarPage{}

	doc.Find(".cell_album_item").Each(func(i int, s *goquery.Selection) {
		carPage := &CarPage{}
		pageNode := s.Find("a")
		pageHref, exists := pageNode.Attr("href")
		if !exists {
			return
		}
		carPage.PageUrl = fmt.Sprintf("https://gallery.u-car.com.tw%s", pageHref)
		carPage.Title = s.Find(".title").Text()
		imgHref, exists := s.Find(".thumb img").Attr("src")
		if !exists {
			return
		}
		carPage.ImgUrl = fmt.Sprintf("https:%s", imgHref)
		result = append(result, carPage)
	})
	localCache.Set("cars", result, time.Minute*60)
}

func main() {
	localCache = cache.New(5*time.Minute, 10*time.Minute)
	router := gin.Default()

	bot, err := linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_ACCESSTOKEN"),
	)
	if err != nil {
		panic(err)
	}

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})
	router.POST("/api", func(c *gin.Context) {
		events, err := bot.ParseRequest(c.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid signature"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			}
			return
		}

		for _, event := range events {
			switch event.Type {
			case linebot.EventTypeMessage:
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if message.Text == "5566" {
						hrefList := []*PttImage{}
						hrefListInterface, found := localCache.Get("img")
						if found {
							hrefList = hrefListInterface.([]*PttImage)
						} else {
							StartPtt()
						}

						hrefListInterface, found = localCache.Get("img")
						if found {
							hrefList = hrefListInterface.([]*PttImage)
						}
						s1 := rand.NewSource(time.Now().UnixNano())
						r1 := rand.New(s1)
						index := r1.Intn(len(hrefList))
						if len(hrefList) > 0 {
							imgIndex := r1.Intn(len(hrefList[index].ImgUrl))
							bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(hrefList[index].ImgUrl[imgIndex], hrefList[index].ImgUrl[imgIndex]), linebot.NewTextMessage(hrefList[index].PageUrl)).Do()
						}
					}
					if message.Text == "7788" {
						hrefList := []*CarPage{}
						hrefListInterface, found := localCache.Get("cars")
						if found {
							hrefList = hrefListInterface.([]*CarPage)
						} else {
							StartCar()
						}

						hrefListInterface, found = localCache.Get("cars")
						if found {
							hrefList = hrefListInterface.([]*CarPage)
						}
						s1 := rand.NewSource(time.Now().UnixNano())
						r1 := rand.New(s1)
						index := r1.Intn(len(hrefList))
						if len(hrefList) > 0 {
							bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(hrefList[index].ImgUrl, hrefList[index].ImgUrl), linebot.NewTextMessage(hrefList[index].Title+" "+hrefList[index].PageUrl)).Do()
						}
					}
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	router.Run(":80")
}
