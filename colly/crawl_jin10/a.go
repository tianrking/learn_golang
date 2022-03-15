package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

var count int = 0

type GG struct {
	text string
	herf string
}

func main() {

	c := colly.NewCollector() // 在colly中使用 Collector 這類物件 來做事情

	c.OnResponse(func(r *colly.Response) { // 當Visit訪問網頁後，網頁響應(Response)時候執行的事情
		// fmt.Println(string(r.Body)) // 返回的Response物件r.Body 是[]Byte格式，要再轉成字串
	})

	c.OnRequest(func(r *colly.Request) { // 需要寫這一段 User-Agent 降低被ban 可能性
		// r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	// 當Visit訪問網頁後，在網頁響應(Response)之後、發現這是HTML格式 執行的事情
	// F12 OnHTML 支持 class

	c.OnHTML(".qa-list__title-link", func(e *colly.HTMLElement) { // 每找到一個符合 goquerySelector字樣的結果，便會進這個OnHTML一次
		_GG := GG{}
		_GG.text = strings.TrimSpace(e.Text)
		_GG.herf = strings.TrimSpace(e.Attr("href"))
		fmt.Println(_GG.text, _GG.herf)
		// count++
	})

	// c.OnHTML("meta[name]", func(e *colly.HTMLElement) {
	// 	fmt.Println(e)
	// })
	// c.OnHTML(".qa-condition__count", func(e *colly.HTMLElement) { // 每找到一個符合 goquerySelector字樣的結果，便會進這個OnHTML一次
	// 	BB := strings.TrimSpace(e.Text)
	// 	fmt.Println(BB)
	// 	// count++
	// })

	// c.Visit("https://ithelp.ithome.com.tw/users/20125192/ironman/3155") // Visit 要放最後
	// c.Visit("https://ithelp.ithome.com.tw/users/20125192/ironman/3155?page=1")
	// c.Visit("https://ithelp.ithome.com.tw/users/20125192/ironman/3155?page=2")
	// c.Visit("https://ithelp.ithome.com.tw/users/20125192/ironman/3155?page=3")
	// c.Visit("https://ithelp.ithome.com.tw/users/20125192/ironman/3155?page=4")
	for i := 0; i < 5; i++ {
		url := "https://ithelp.ithome.com.tw/?page=" + string(strconv.Itoa(i))
		// strings.Join(url, string(i))
		// fmt.Println(url)
		c.Visit(url)
	}
}
