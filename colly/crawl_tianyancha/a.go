package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

var count int = 0

type _Result struct {
	Name                 string
	Legal_representative string
	Leader               string
	Capital              string
	Date_Establishment   string
	State                string
	Phone                string
	Email                string
	Address              string
	Text                 string
	Href                 string
	Contact              string
}

// 公司名称
// 法定代表人
// 注册资本
// 成立日期
// 经营状况
// 联系电话
// 邮箱
// 地址

func main() {
	GG := _Result{}
	c := colly.NewCollector()              // 在colly中使用 Collector 這類物件 來做事情
	c.OnResponse(func(r *colly.Response) { // 當Visit訪問網頁後，網頁響應(Response)時候執行的事情
		// fmt.Println(string(r.Body)) // 返回的Response物件r.Body 是[]Byte格式，要再轉成字串
		// fmt.Println(reflect.TypeOf(r))
	})
	c.OnRequest(func(r *colly.Request) { // 需要寫這一段 User-Agent 降低被ban 可能性
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})
	// 當Visit訪問網頁後，在網頁響應(Response)之後、發現這是HTML格式 執行的事情
	// F12 OnHTML 支持 class
	c.OnHTML("div[class='search-item sv-search-company  ']", func(e *colly.HTMLElement) { // 每找到一個符合 goquerySelector字樣的結果，便會進這個OnHTML一次

		// GG.Href = e.Attr("href")
		// ch := e.DOM.Children()
		//.Eq(2).Children().Eq(0).Children().Eq(0).Attr("href")
		// GG.Href, _ = e.DOM.Find("a[class='name select-none ").Attr("href")
		GG.Name = e.ChildText("a[class='name select-none ']")
		GG.Href = e.ChildAttr("a[class='name select-none ']", "href")
		GG.Leader = e.ChildText("div[class='title -wider text-ellipsis']")
		// GG.Phone = e.ChildText("span[tyc-event-ch='CompanySearch.MoreTel']")
		// GG.Email = e.ChildText("div[class='contact row ']")
		GG.Contact = e.ChildText("div[class='contact row ']")
		GG.Capital = e.ChildText("div[class='title -narrow text-ellipsis']")
		GG.Date_Establishment = e.ChildText("div[class='title  text-ellipsis']")

		fmt.Println(GG.Name)
		fmt.Println(GG.Leader)
		fmt.Println(GG.Capital)
		// fmt.Println(GG.Phone)
		// fmt.Println(GG.Email)
		fmt.Println(GG.Contact)
		fmt.Println(GG.Date_Establishment)
		fmt.Println(GG.Href)
		count++
	})

	c.Visit("https://www.tianyancha.com/search?key=%E4%BA%AC%E4%B8%9C") // Visit 要放最後
}
