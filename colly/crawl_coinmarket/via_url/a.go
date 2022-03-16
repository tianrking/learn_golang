package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
)

func main() {
	// fName := "cryptocoinmarketcap.csv"

	// file, err := os.Create(fName)
	// if err != nil {
	// 	log.Fatalf("Cannot create file %q: %s\n", fName, err)
	// 	return
	// }
	// defer file.Close()
	// writer := csv.NewWriter(file)
	// defer writer.Flush()

	// Write CSV header
	// writer.Write([]string{"Name", "Symbol", "Price (USD)", "Volume (USD)", "Market capacity (USD)", "Change (1h)", "Change (24h)", "Change (7d)"})

	// Instantiate default collector
	c := colly.NewCollector()
	// c := colly.NewCollector(colly.AllowURLRevisit())

	// Rotate two socks5 proxies
	rp, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:12999", "socks5://127.0.0.1:12999")
	if err != nil {
		log.Fatal(err)
	}
	// 【设置代理IP】 ，这里使用的是轮询ip方式
	c.SetProxyFunc(rp)

	c.OnRequest(func(r *colly.Request) { // 需要寫這一段 User-Agent 降低被ban 可能性
		//r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")

		r.Headers.Set("Authority", "coinmarketcap.com")
		r.Headers.Set("Cache-Control", "max-age=0")
		r.Headers.Set("Sec-Ch-Ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"99\", \"Microsoft Edge\";v=\"99\"")
		r.Headers.Set("Sec-Ch-Ua-Mobile", "?0")
		r.Headers.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36 Edg/99.0.1150.39")
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		r.Headers.Set("Sec-Fetch-Site", "same-origin")
		r.Headers.Set("Sec-Fetch-Mode", "navigate")
		r.Headers.Set("Sec-Fetch-User", "?2")
		r.Headers.Set("Sec-Fetch-Dest", "document")
		r.Headers.Set("Referer", "https://coinmarketcap.com/all/views/all/")
		r.Headers.Set("Accept-Language", "zh-TW,zh-HK;q=0.9,zh;q=0.8,en;q=0.7,zh-CN;q=0.6,en-GB;q=0.5,en-US;q=0.4")

	})

	c.OnResponse(func(r *colly.Response) { // 當Visit訪問網頁後，網頁響應(Response)時候執行的事情
		// fmt.Println(string(r.Body)) // 返回的Response物件r.Body 是[]Byte格式，要再轉成字串
		// fmt.Println(reflect.TypeOf(r))
	})

	c.OnHTML("tr[class='cmc-table-row']", func(e *colly.HTMLElement) {
		// writer.Write([]string{
		// 	e.ChildText(".currency-name-container"),
		// 	e.ChildText(".col-symbol"),
		// 	e.ChildAttr("a.price", "data-usd"),
		// 	e.ChildAttr("a.volume", "data-usd"),
		// 	e.ChildAttr(".market-cap", "data-usd"),
		// 	e.ChildText(".percent-1h"),
		// 	e.ChildText(".percent-24h"),
		// 	e.ChildText(".percent-7d"),
		// })
		fmt.Println(e.ChildText("a[class='cmc-table__column-name--name cmc-link']"))
	})

	c.Visit("https://coinmarketcap.com/all/views/all/")
	// c.Visit("https://coinmarketcap.com/currencies/bitcoin/")

	// log.Printf("Scraping finished, check file %q for results\n", fName)
}

// curl 'https://api.coinmarketcap.com/data-api/v3/sector/w/lite-list' \
// -H 'authority: api.coinmarketcap.com' \
// -H 'sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="99", "Microsoft Edge";v="99"' \
// -H 'accept: application/json, text/plain, */*' \
// -H 'platform: web' \
// -H 'sec-ch-ua-mobile: ?0' \
// -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36 Edg/99.0.1150.39' \
// -H 'x-request-id: c4b05ffe-228d-4573-a207-b9e50184620f' \
// -H 'sec-ch-ua-platform: "Windows"' \
// -H 'origin: https://coinmarketcap.com' \
// -H 'sec-fetch-site: same-site' \
// -H 'sec-fetch-mode: cors' \
// -H 'sec-fetch-dest: empty' \
// -H 'referer: https://coinmarketcap.com/' \
// -H 'accept-language: zh-TW,zh-HK;q=0.9,zh;q=0.8,en;q=0.7,zh-CN;q=0.6,en-GB;q=0.5,en-US;q=0.4' \
// -H 'if-modified-since: Wed, 16 Mar 2022 06:52:32 GMT' \
// --compressed

///  这个 就是 数据源
// curl 'https://api.coinmarketcap.com/data-api/v3/topsearch/rank' \
//   -H 'authority: api.coinmarketcap.com' \
//   -H 'sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="99", "Microsoft Edge";v="99"' \
//   -H 'accept: application/json, text/plain, */*' \
//   -H 'platform: web' \
//   -H 'sec-ch-ua-mobile: ?0' \
//   -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36 Edg/99.0.1150.39' \
//   -H 'x-request-id: 94ae8373-239b-4296-a353-15536d721619' \
//   -H 'sec-ch-ua-platform: "Windows"' \
//   -H 'origin: https://coinmarketcap.com' \
//   -H 'sec-fetch-site: same-site' \
//   -H 'sec-fetch-mode: cors' \
//   -H 'sec-fetch-dest: empty' \
//   -H 'referer: https://coinmarketcap.com/' \
//   -H 'accept-language: zh-TW,zh-HK;q=0.9,zh;q=0.8,en;q=0.7,zh-CN;q=0.6,en-GB;q=0.5,en-US;q=0.4' \
//   -H 'if-modified-since: Wed, 16 Mar 2022 06:50:24 GMT' \
//   --compressed
///  数据源

// curl 'https://api.coinmarketcap.com/data-api/v3/topsearch/rank' \
//   -H 'authority: api.coinmarketcap.com' \
//   -H 'sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="99", "Google Chrome";v="99"' \
//   -H 'accept: application/json, text/plain, */*' \
//   -H 'platform: web' \
//   -H 'sec-ch-ua-mobile: ?0' \
//   -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36' \
//   -H 'x-request-id: 62f24a92-35f9-4cda-be48-3fe381aa1b27' \
//   -H 'sec-ch-ua-platform: "Windows"' \
//   -H 'origin: https://coinmarketcap.com' \
//   -H 'sec-fetch-site: same-site' \
//   -H 'sec-fetch-mode: cors' \
//   -H 'sec-fetch-dest: empty' \
//   -H 'referer: https://coinmarketcap.com/' \
//   -H 'accept-language: zh-TW,zh-HK;q=0.9,zh;q=0.8,en;q=0.7,zh-CN;q=0.6,en-GB;q=0.5,en-US;q=0.4' \
//   -H 'if-modified-since: Wed, 16 Mar 2022 15:34:16 GMT' \
//   --compressed

// curl 'https://api.coinmarketcap.com/data-api/v3/map/all?listing_status=active,untracked&exchangeAux=is_active,status&cryptoAux=is_active,status&start=10001&limit=10000' \
//   -H 'authority: api.coinmarketcap.com' \
//   -H 'sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="99", "Microsoft Edge";v="99"' \
//   -H 'accept: application/json, text/plain, */*' \
//   -H 'platform: web' \
//   -H 'sec-ch-ua-mobile: ?0' \
//   -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36 Edg/99.0.1150.39' \
//   -H 'x-request-id: 6813b015-430c-432d-876f-207bd8a555ec' \
//   -H 'sec-ch-ua-platform: "Windows"' \
//   -H 'origin: https://coinmarketcap.com' \
//   -H 'sec-fetch-site: same-site' \
//   -H 'sec-fetch-mode: cors' \
//   -H 'sec-fetch-dest: empty' \
//   -H 'referer: https://coinmarketcap.com/' \
//   -H 'accept-language: zh-TW,zh-HK;q=0.9,zh;q=0.8,en;q=0.7,zh-CN;q=0.6,en-GB;q=0.5,en-US;q=0.4' \
//   -H 'if-modified-since: Wed, 16 Mar 2022 06:48:08 GMT' \
//   --compressed
