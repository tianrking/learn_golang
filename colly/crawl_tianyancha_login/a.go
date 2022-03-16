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
		r.Headers.Set("Cookies", "TYCID=31ce63707e4711ecba526ddcaca6ce5d; ssuid=2877677358; _ga=GA1.2.1065970361.1643160481; jsid=SEO-GOOGLE-ALL-SY-000001; bad_id658cce70-d9dc-11e9-96c6-833900356dc6=4553d751-96f1-11ec-b406-811e0c9b6895; creditGuide=1; aliyungf_tc=d388b8c2e3fccd26957870591f2d8fb7d0b3c79c8739d97d07711eb245bb0f64; csrfToken=seNjJViJs6I4qs9ugWvFpyMw; bannerFlag=true; Hm_lvt_e92c8d65d92d534b0fc290df538b4758=1645872350,1647396463; _gid=GA1.2.887464488.1647396463; _bl_uid=wvlnt0C6sz9xCmcwO45XqhIbO8n7; RTYCID=ef076b9b337148bf97494c2aef10c85d; CT_TYCID=ba295de73e56433c861712bbc3438014; searchSessionId=1647396884.87386679; relatedHumanSearchGraphId=19191697; relatedHumanSearchGraphId.sig=hql_k7yuqYFYUPOdL00DMQ9MtLWCjvJ7iAqLJr9-nQw; cloud_token=a025968730c74fb4ae1b8e3fd41c1ec4; acw_tc=781bad4a16474057099172095e3c440b44eda090d3ed9140a8c735df001935; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%22206974836%22%2C%22first_id%22%3A%2217e93ff234dd4a-0755e316cf2088-f791539-1327104-17e93ff234e2f2%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E8%87%AA%E7%84%B6%E6%90%9C%E7%B4%A2%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC%22%2C%22%24latest_referrer%22%3A%22https%3A%2F%2Fwww.google.com%2F%22%7D%2C%22%24device_id%22%3A%2217e93ff234dd4a-0755e316cf2088-f791539-1327104-17e93ff234e2f2%22%7D; tyc-user-info={%22state%22:%220%22%2C%22vipManager%22:%220%22%2C%22mobile%22:%2215773211225%22}; tyc-user-info-save-time=1647405693934; auth_token=eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIxNTc3MzIxMTIyNSIsImlhdCI6MTY0NzQwNTc1NywiZXhwIjoxNjQ5OTk3NzU3fQ.5PjC44WhxpveAN9XwlidAkew2kUBDQdwUqEPMNGEsR9C4R9l2nlQE9Vbaf9N23IoW_bpJe12_RM0piGKbSNVTA; tyc-user-phone=%255B%252215773211225%2522%255D; Hm_lpvt_e92c8d65d92d534b0fc290df538b4758=1647405733")

		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Cache-Control", "max-age=0")
		// r.Headers.Set("sec-ch-ua", "" Not A;Brand";v="99", "Chromium";v="99", "Microsoft Edge";v="99"")
		r.Headers.Set("sec-ch-ua-mobile", "?0")
		r.Headers.Set("sec-ch-ua-platform", "Windows")
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36 Edg/99.0.1150.39")
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		r.Headers.Set("Sec-Fetch-Site", "same-origin")
		r.Headers.Set("Sec-Fetch-Mode", "navigate")
		r.Headers.Set("Sec-Fetch-User", "?1")
		r.Headers.Set("Sec-Fetch-Dest", "document")
		r.Headers.Set("Referer", "https://www.tianyancha.com/search?key=%E4%BA%AC%E4%B8%9C")
		r.Headers.Set("Accept-Language", "zh-TW,zh-HK;q=0.9,zh;q=0.8,en;q=0.7,zh-CN;q=0.6,en-GB;q=0.5,en-US;q=0.4")
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
