package main

import (
	"fmt"
	"log"
	"strconv"

	colly "github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
	excelize "github.com/xuri/excelize/v2"
)

// var _url string = ""

type GG struct {
	_url       string
	_name      string
	_now_price string
}

// DOGE := GG{}

func main() {

	num := 0
	f := excelize.NewFile()
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

	// c.OnHTML("tr[class='s395gx-1 eChPfw cmc-table-row']", func(e *colly.HTMLElement) {
	// 	// writer.Write([]string{\
	// 	// fmt.Print(string(e.Text))

	// 	fmt.Print(e.ChildText("td[class='name-cell']"), "       ")
	// 	fmt.Println(e.ChildAttr("a[class='cmc-link']", "href"))
	// })

	c.OnHTML(".cmc-table__table-wrapper-outer tbody tr", func(e *colly.HTMLElement) {
		// writer.Write([]string{\
		// fmt.Print(string(e.Text))
		DOGE := GG{}
		if len(e.ChildText("td[class='name-cell']")) == 0 {
			DOGE._url = "https://coinmarketcap.com" + e.ChildAttr("a[class='cmc-link']", "href")
			DOGE._name = e.ChildText("a[class='cmc-table__column-name--name cmc-link']")
			// fmt.Print(e.ChildText("a[class='cmc-table__column-name--name cmc-link']"), "       ")
		} else {
			DOGE._url = "https://coinmarketcap.com" + e.ChildAttr("a[class='cmc-link']", "href")
			DOGE._name = e.ChildText("td[class='name-cell']")
			// fmt.Print(e.ChildText("td[class='name-cell']"), "       ")
		}

		// fmt.Print(DOGE._name, " ")
		// fmt.Print(DOGE._url, " ")

		c_coin := colly.NewCollector()
		c_coin.SetProxyFunc(rp)
		c_coin.OnRequest(func(r *colly.Request) {
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
		c_coin.OnResponse(func(r *colly.Response) {
		})
		c_coin.OnHTML("div[class='priceValue ']", func(r *colly.HTMLElement) {
			DOGE._now_price = string(r.Text)
		})
		c_coin.Visit(DOGE._url)

		fmt.Print(DOGE._name, "     ")
		fmt.Print(DOGE._now_price, "     ")
		fmt.Println(DOGE._url)

		// Create a new sheet.
		// index := f.NewSheet("Sheet2")
		// Set value of a cell.
		// f.SetCellValue("Sheet2", "A2", "Hello world.")

		var A_num string = "A" + strconv.Itoa(num)
		B_price := "B" + strconv.Itoa(num)
		C_url := "C" + strconv.Itoa(num)

		// fmt.Println(A_num)
		// fmt.Println(B_price)

		// fmt.Println(reflect.TypeOf(A_num))
		// fmt.Println(reflect.TypeOf(B_price))

		f.SetCellValue("Sheet1", A_num, DOGE._name)
		f.SetCellValue("Sheet1", B_price, DOGE._now_price)
		f.SetCellValue("Sheet1", C_url, DOGE._url)
		// f.SetCellValue("Sheet1", "A", 300)
		// f.SetCellValue("Sheet1", "B5", 200)
		// Set active sheet of the workbook.
		// f.SetActiveSheet(index)
		// Save spreadsheet by the given path.
		if err := f.SaveAs("Coin_price.xlsx"); err != nil {
			fmt.Println(err)
		}
		num += 1

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
