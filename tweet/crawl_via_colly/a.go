package main

import (
	"fmt"
	"log"

	colly "github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
)

func main() {

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

		r.Headers.Set("authority", "twitter.com")
		// r.Headers.Set("sec-ch-ua", "Not A;Brand";v="99", "Chromium";v="99", "Microsoft Edge";v="99"")
		r.Headers.Set("x-twitter-client-language", "zh-tw")
		r.Headers.Set("x-csrf-token", "e8d03e24e35e5193e5da2262c8aa1237")
		r.Headers.Set("sec-ch-ua-mobile", "?0")
		r.Headers.Set("authorization", "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA")
		r.Headers.Set("content-type", "application/json")
		r.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36 Edg/99.0.1150.39")
		r.Headers.Set("x-guest-token", "1504647851560734720")
		r.Headers.Set("x-twitter-active-user", "yes")
		r.Headers.Set("sec-ch-ua-platform", "Windows")
		r.Headers.Set("accept", "*/*")
		r.Headers.Set("sec-fetch-site", "same-origin")
		r.Headers.Set("sec-fetch-mode", "cors")
		r.Headers.Set("sec-fetch-dest", "empty")
		r.Headers.Set("referer", "https://twitter.com/elonmusk")
		r.Headers.Set("accept-language", "zh-TW,zh;q=0.9")
		r.Headers.Set("cookie", "guest_id_marketing=v1%3A164757097289679256; guest_id_ads=v1%3A164757097289679256; personalization_id='v1_OFbNgcu51v01tVdezN5P1w=='; guest_id=v1%3A164757097289679256; ct0=e8d03e24e35e5193e5da2262c8aa1237; gt=1504647851560734720; _ga=GA1.2.1601084745.1647570912; _gid=GA1.2.765857587.1647570912")
	})

	c.OnResponse(func(r *colly.Response) { // 當Visit訪問網頁後，網頁響應(Response)時候執行的事情
		fmt.Println(string(r.Body)) // 返回的Response物件r.Body 是[]Byte格式，要再轉成字串
		// fmt.Println(reflect.TypeOf(r))
	})

	c.OnHTML("div[class='css-901oao r-1nao33i r-37j5jr r-a023e6 r-16dba41 r-rjixqe r-bcqeeo r-bnwqim r-qvutc0']", func(e *colly.HTMLElement) {
		fmt.Println(e.ChildText("span[class='css-901oao css-16my406 r-poiln3 r-bcqeeo r-qvutc0']"))
	})

	c.Visit("https://twitter.com/i/api/graphql/CDDPst9A-AHg6Q0k9-wo7w/UserTweets?variables=%7B%22userId%22%3A%2244196397%22%2C%22count%22%3A40%2C%22includePromotedContent%22%3Atrue%2C%22withQuickPromoteEligibilityTweetFields%22%3Atrue%2C%22withSuperFollowsUserFields%22%3Atrue%2C%22withDownvotePerspective%22%3Afalse%2C%22withReactionsMetadata%22%3Afalse%2C%22withReactionsPerspective%22%3Afalse%2C%22withSuperFollowsTweetFields%22%3Atrue%2C%22withVoice%22%3Atrue%2C%22withV2Timeline%22%3Atrue%2C%22__fs_dont_mention_me_view_api_enabled%22%3Afalse%2C%22__fs_interactive_text_enabled%22%3Atrue%2C%22__fs_responsive_web_uc_gql_enabled%22%3Afalse%7D")

	// log.Printf("Scraping finished, check file %q for results\n", fName)
}

// curl 'https://twitter.com/i/api/graphql/CDDPst9A-AHg6Q0k9-wo7w/UserTweets?variables=%7B%22userId%22%3A%2244196397%22%2C%22count%22%3A40%2C%22includePromotedContent%22%3Atrue%2C%22withQuickPromoteEligibilityTweetFields%22%3Atrue%2C%22withSuperFollowsUserFields%22%3Atrue%2C%22withDownvotePerspective%22%3Afalse%2C%22withReactionsMetadata%22%3Afalse%2C%22withReactionsPerspective%22%3Afalse%2C%22withSuperFollowsTweetFields%22%3Atrue%2C%22withVoice%22%3Atrue%2C%22withV2Timeline%22%3Atrue%2C%22__fs_dont_mention_me_view_api_enabled%22%3Afalse%2C%22__fs_interactive_text_enabled%22%3Atrue%2C%22__fs_responsive_web_uc_gql_enabled%22%3Afalse%7D' \
//   -H 'authority: twitter.com' \
//   -H 'sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="99", "Microsoft Edge";v="99"' \
//   -H 'x-twitter-client-language: zh-tw' \
//   -H 'x-csrf-token: 00403b4a88494594e1e2ae52e2e2fa286a85976a1b44392ba1149d3e64c0b001038ead76d036377cafad77dc4f51e81e06cc7daed2bd0dd1b86ac35f589b084798de3b9023596cc116b95e818a0c21b3' \
//   -H 'sec-ch-ua-mobile: ?0' \
//   -H 'authorization: Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA' \
//   -H 'content-type: application/json' \
//   -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36 Edg/99.0.1150.39' \
//   -H 'x-twitter-auth-type: OAuth2Session' \
//   -H 'x-twitter-active-user: yes' \
//   -H 'sec-ch-ua-platform: "Windows"' \
//   -H 'accept: */*' \
//   -H 'sec-fetch-site: same-origin' \
//   -H 'sec-fetch-mode: cors' \
//   -H 'sec-fetch-dest: empty' \
//   -H 'referer: https://twitter.com/elonmusk' \
//   -H 'accept-language: zh-TW,zh-HK;q=0.9,zh;q=0.8,en;q=0.7,zh-CN;q=0.6,en-GB;q=0.5,en-US;q=0.4' \
//   -H 'cookie: guest_id=v1%3A163055344729664497; dnt=1; ads_prefs="HBESAAA="; kdt=l6Uox8y9VKFcT8uF2lMti0j35OVX8J3OjaXFgW5y; auth_token=c4da4b011cce167c087ff0cff564357bbd204c38; twid=u%3D2405890146; ct0=00403b4a88494594e1e2ae52e2e2fa286a85976a1b44392ba1149d3e64c0b001038ead76d036377cafad77dc4f51e81e06cc7daed2bd0dd1b86ac35f589b084798de3b9023596cc116b95e818a0c21b3; eu_cn=1; d_prefs=MToxLGNvbnNlbnRfdmVyc2lvbjoyLHRleHRfdmVyc2lvbjoxMDAw; guest_id_ads=v1%3A163055344729664497; guest_id_marketing=v1%3A163055344729664497; personalization_id="v1_azM6Ij0GTuTZDtSAyzJzNw=="; des_opt_in=Y; _ga_BYKEBDM7DS=GS1.1.1646634012.2.0.1646634012.0; _gid=GA1.2.388915402.1647185658; lang=zh-tw; at_check=true; external_referer=padhuUp37zhE6DuBcHcZNU%2BMzEzAMVh3szyKSjaVzV8lWiTXhlGhdA%3D%3D|0|8e8t2xd8A2w%3D; mbox=PC#308a0da8f7d64f6a8f98005f43c387d0.34_0#1710812838|session#d59b1d703c2d45959c1ac63cf7809dfd#1647569898; _ga=GA1.2.1836147314.1644653558; _ga_34PHSZMC42=GS1.1.1647567366.4.1.1647570038.0' \
//   --compressed

// curl 'https://twitter.com/i/api/graphql/CDDPst9A-AHg6Q0k9-wo7w/UserTweets?variables=%7B%22userId%22%3A%2244196397%22%2C%22count%22%3A40%2C%22includePromotedContent%22%3Atrue%2C%22withQuickPromoteEligibilityTweetFields%22%3Atrue%2C%22withSuperFollowsUserFields%22%3Atrue%2C%22withDownvotePerspective%22%3Afalse%2C%22withReactionsMetadata%22%3Afalse%2C%22withReactionsPerspective%22%3Afalse%2C%22withSuperFollowsTweetFields%22%3Atrue%2C%22withVoice%22%3Atrue%2C%22withV2Timeline%22%3Atrue%2C%22__fs_dont_mention_me_view_api_enabled%22%3Afalse%2C%22__fs_interactive_text_enabled%22%3Atrue%2C%22__fs_responsive_web_uc_gql_enabled%22%3Afalse%7D' \
//   -H 'authority: twitter.com' \
//   -H 'sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="99", "Microsoft Edge";v="99"' \
//   -H 'x-twitter-client-language: zh-tw' \
//   -H 'x-csrf-token: e8d03e24e35e5193e5da2262c8aa1237' \
//   -H 'sec-ch-ua-mobile: ?0' \
//   -H 'authorization: Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA' \
//   -H 'content-type: application/json' \
//   -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36 Edg/99.0.1150.39' \
//   -H 'x-guest-token: 1504647851560734720' \
//   -H 'x-twitter-active-user: yes' \
//   -H 'sec-ch-ua-platform: "Windows"' \
//   -H 'accept: */*' \
//   -H 'sec-fetch-site: same-origin' \
//   -H 'sec-fetch-mode: cors' \
//   -H 'sec-fetch-dest: empty' \
//   -H 'referer: https://twitter.com/elonmusk' \
//   -H 'accept-language: zh-TW,zh;q=0.9' \
//   -H 'cookie: guest_id_marketing=v1%3A164757097289679256; guest_id_ads=v1%3A164757097289679256; personalization_id="v1_OFbNgcu51v01tVdezN5P1w=="; guest_id=v1%3A164757097289679256; ct0=e8d03e24e35e5193e5da2262c8aa1237; gt=1504647851560734720; _ga=GA1.2.1601084745.1647570912; _gid=GA1.2.765857587.1647570912' \
//   --compressed
