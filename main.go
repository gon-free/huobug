package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	huoUrl := "https://api.huobi.pro/market/history/kline?period=60min&size=200&symbol=btcusdt"
	resp, err := http.Get(huoUrl)
	if err != nil {
		fmt.Printf("get data from huo fail.err:%v", err)
		return
	}
	defer resp.Body.Close()

	rspBody, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("get data:%s", string(rspBody))
}
