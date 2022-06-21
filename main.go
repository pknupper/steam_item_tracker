package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Item struct {
	Name     string
	HashName string
}

type Response struct {
	Success     bool   `json:"success"`
	LowestPrice string `json:"lowest_price"`
	Volume      string `json:"volume"`
	MedianPrice string `json:"median_price"`
}

func main() {
	const baseUrl = "https://steamcommunity.com/market/priceoverview/?appid=730&currency=3&market_hash_name="

	c := http.Client{Timeout: time.Duration(1) * time.Second}

	var items = []Item{
		{
			Name:     "Prisma 2 Case",
			HashName: "Prisma%202%20Case",
		},
		{
			Name:     "Danger Zone Case",
			HashName: "Danger%20Zone%20Case",
		},
		{
			Name:     "Fracture Case",
			HashName: "Fracture%20Case",
		},
		{
			Name:     "Prisma Case",
			HashName: "Prisma%20Case",
		},
		{
			Name:     "Snakebite Case",
			HashName: "Snakebite%20Case",
		},
		{
			Name:     "Horizon Case",
			HashName: "Horizon%20Case",
		},
		{
			Name:     "CS20 Case",
			HashName: "CS20%20Case",
		},
		{
			Name:     "Revolver Case",
			HashName: "Revolver%20Case",
		},
		{
			Name:     "Shadow Case",
			HashName: "Shadow%20Case",
		},
		{
			Name:     "Falchion Case",
			HashName: "Falchion%20Case",
		},
		{
			Name:     "Clutch Case",
			HashName: "Clutch%20Case",
		},
		{
			Name:     "Chroma 3 Case",
			HashName: "Chroma%203%20Case",
		},
		{
			Name:     "Spectrum 2 Case",
			HashName: "Spectrum%202%20Case",
		},
		{
			Name:     "Gamma Case",
			HashName: "Gamma%20Case",
		},
		{
			Name:     "Chroma 2 Case",
			HashName: "Chroma%202%20Case",
		},
		{
			Name:     "Dreams And Nightmares Case",
			HashName: "Dreams%20%26%20Nightmares%20Case",
		},
		{
			Name:     "Gamma 2 Case",
			HashName: "Gamma%202%20Case",
		},
		{
			Name:     "Operation Vanguard Case",
			HashName: "Operation%20Vanguard%20Weapon%20Case",
		},
		{
			Name:     "Spectrum Case",
			HashName: "Spectrum%20Case",
		},
		{
			Name:     "Chroma Case",
			HashName: "Chroma%20Case",
		},
		{
			Name:     "Shattered Web Case",
			HashName: "Shattered%20Web%20Case",
		},
		{
			Name:     "Operation Phoenix Case",
			HashName: "Operation%20Phoenix%20Weapon%20Case",
		},
		{
			Name:     "Operation Broken Fang Case",
			HashName: "Operation%20Broken%20Fang%20Casee",
		},
		{
			Name:     "Glove Case",
			HashName: "Glove%20Case",
		},
		{
			Name:     "Operation Breakout Case",
			HashName: "Operation%20Breakout%20Weapon%20Case",
		},
		{
			Name:     "Huntsman Case",
			HashName: "Huntsman%20Weapon%20Case",
		},
		{
			Name:     "eSports Summer Case 2014",
			HashName: "eSports%202014%20Summer%20Case",
		},
		{
			Name:     "Winter Offensive Case",
			HashName: "Winter%20Offensive%20Weapon%20Case",
		},
		{
			Name:     "eSports Winter Case 2013",
			HashName: "eSports%202013%20Winter%20Case",
		},
		{
			Name:     "Weapon Case 3",
			HashName: "CS%3AGO%20Weapon%20Case%203",
		},
		{
			Name:     "Weapon Case 2",
			HashName: "CS%3AGO%20Weapon%20Case%202",
		},
		{
			Name:     "Operation Hydra Case",
			HashName: "Operation%20Hydra%20Case",
		},
		{
			Name:     "eSports 2013 Case",
			HashName: "eSports%202013%20Case",
		},
		{
			Name:     "Operation Bravo Case",
			HashName: "Operation%20Bravo%20Case",
		},
		{
			Name:     "Weapon Case",
			HashName: "CS%3AGO%20Weapon%20Case",
		},
	}

	for _, item := range items {
		resp, err := c.Get(baseUrl + item.HashName)

		if err != nil {
			fmt.Printf("Error %s", err)
			return
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		var response Response
		json.Unmarshal([]byte(body), &response)

		fmt.Println(item.Name + ": Lowest Price: " + response.LowestPrice + ", Volume: " + response.Volume + ", Median Price: " + response.MedianPrice)

		time.Sleep(5 * time.Second)
	}

}
