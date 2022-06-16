package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get("https://steamcommunity.com/market/priceoverview/?appid=730&currency=3&market_hash_name=Operation%20Breakout%20Weapon%20Case")
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body : %s", body)
}
