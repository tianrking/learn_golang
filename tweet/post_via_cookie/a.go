package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("socks5://127.0.0.1:12999")
	}

	httpTransport := &http.Transport{
		Proxy: proxy,
	}

	client := &http.Client{
		Transport: httpTransport,
	}

	var data = strings.NewReader(`{"url":"https://api.twitter.com/2/tweets","method":"POST","body":"{\"text\":\"1\"}"}`)
	req, err := http.NewRequest("POST", "https://oauth-playground.glitch.me/request", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "oauth-playground.glitch.me")
	// req.Header.Set("sec-ch-ua", "" Not A;Brand";v="99", "Chromium";v="99", "Microsoft Edge";v="99"")
	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36 Edg/99.0.1150.39")
	req.Header.Set("sec-ch-ua-platform", "Windows")
	req.Header.Set("origin", "https://oauth-playground.glitch.me")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("referer", "https://oauth-playground.glitch.me/?id=createTweet&params=%28%29_&body=%27%28*text%5C%21*1*%29%27*%5C%27%01*_")
	req.Header.Set("accept-language", "zh-TW,zh-HK;q=0.9,zh;q=0.8,en;q=0.7,zh-CN;q=0.6,en-GB;q=0.5,en-US;q=0.4")
	req.Header.Set("cookie", "token=j%3A%7B%22token_type%22%3A%22bearer%22%2C%22expires_in%22%3A7200%2C%22access_token%22%3A%22U3NhX3MyX19HTUdSOFFqTXhsLUtTbS1CZFRTMjlsMnMtSUtVck1YSHY4WS1jOjE2NDc1Njc0MDE3OTQ6MToxOmF0OjE%22%2C%22scope%22%3A%22tweet.write%20users.read%20tweet.read%20offline.access%22%2C%22refresh_token%22%3A%22cWJjZnVHcGpPY24xOTk0cEhoVmphTHVLSVZTMURYY1RtRDBmaEhnVW1ybXpmOjE2NDc1Njc0MDE3OTQ6MTowOnJ0OjE%22%2C%22expires_at%22%3A1647574601812%7D")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
