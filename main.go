package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Item struct {
	Name     string
	HashName string
	Stock    int
}

type Response struct {
	Success     bool   `json:"success"`
	LowestPrice string `json:"lowest_price"`
	Volume      string `json:"volume"`
	MedianPrice string `json:"median_price"`
}

type EmbedField struct {
	Name  string
	Value string
}

type Message struct {
	Title  string
	Fields []EmbedField
}

var (
	Channel  = flag.String("channel", "", "Channel ID")
	BotToken = flag.String("token", "", "Bot token")
)

func init() { flag.Parse() }

func main() {

	fields := getSteamItems()

	message := buildMessage(fields)

	session, err := discordgo.New("Bot " + *BotToken)

	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	sendDiscordMessage(session, message)

}

func getSteamItems() []EmbedField {
	const baseUrl = "https://steamcommunity.com/market/priceoverview/?appid=730&currency=3&market_hash_name="

	c := http.Client{Timeout: time.Duration(1) * time.Second}

	var items = []Item{
		{
			Name:     "Recoil Case",
			HashName: "Recoil%20Case",
			Stock:    0,
		},
		{
			Name:     "Prisma 2 Case",
			HashName: "Prisma%202%20Case",
			Stock:    0,
		},
		{
			Name:     "Danger Zone Case",
			HashName: "Danger%20Zone%20Case",
			Stock:    0,
		},
		{
			Name:     "Fracture Case",
			HashName: "Fracture%20Case",
			Stock:    0,
		},
		{
			Name:     "Prisma Case",
			HashName: "Prisma%20Case",
			Stock:    0,
		},
		{
			Name:     "Snakebite Case",
			HashName: "Snakebite%20Case",
			Stock:    0,
		},
		{
			Name:     "Horizon Case",
			HashName: "Horizon%20Case",
			Stock:    0,
		},
		{
			Name:     "CS20 Case",
			HashName: "CS20%20Case",
			Stock:    0,
		},
		{
			Name:     "Revolver Case",
			HashName: "Revolver%20Case",
			Stock:    25,
		},
		{
			Name:     "Shadow Case",
			HashName: "Shadow%20Case",
			Stock:    0,
		},
		{
			Name:     "Falchion Case",
			HashName: "Falchion%20Case",
			Stock:    0,
		},
		{
			Name:     "Clutch Case",
			HashName: "Clutch%20Case",
			Stock:    10,
		},
		{
			Name:     "Chroma 3 Case",
			HashName: "Chroma%203%20Case",
			Stock:    0,
		},
		{
			Name:     "Spectrum 2 Case",
			HashName: "Spectrum%202%20Case",
			Stock:    0,
		},
		{
			Name:     "Gamma Case",
			HashName: "Gamma%20Case",
			Stock:    0,
		},
		{
			Name:     "Chroma 2 Case",
			HashName: "Chroma%202%20Case",
			Stock:    0,
		},
		{
			Name:     "Dreams And Nightmares Case",
			HashName: "Dreams%20%26%20Nightmares%20Case",
			Stock:    0,
		},
		{
			Name:     "Gamma 2 Case",
			HashName: "Gamma%202%20Case",
			Stock:    0,
		},
		{
			Name:     "Operation Vanguard Case",
			HashName: "Operation%20Vanguard%20Weapon%20Case",
			Stock:    0,
		},
		{
			Name:     "Spectrum Case",
			HashName: "Spectrum%20Case",
			Stock:    0,
		},
		{
			Name:     "Chroma Case",
			HashName: "Chroma%20Case",
			Stock:    0,
		},
		{
			Name:     "Shattered Web Case",
			HashName: "Shattered%20Web%20Case",
			Stock:    0,
		},
		{
			Name:     "Operation Phoenix Case",
			HashName: "Operation%20Phoenix%20Weapon%20Case",
			Stock:    0,
		},
		{
			Name:     "Operation Broken Fang Case",
			HashName: "Operation%20Broken%20Fang%20Casee",
			Stock:    0,
		},
		{
			Name:     "Glove Case",
			HashName: "Glove%20Case",
			Stock:    0,
		},
		{
			Name:     "Operation Breakout Case",
			HashName: "Operation%20Breakout%20Weapon%20Case",
			Stock:    0,
		},
		{
			Name:     "Huntsman Case",
			HashName: "Huntsman%20Weapon%20Case",
			Stock:    0,
		},
		{
			Name:     "eSports Summer Case 2014",
			HashName: "eSports%202014%20Summer%20Case",
			Stock:    0,
		},
		{
			Name:     "Winter Offensive Case",
			HashName: "Winter%20Offensive%20Weapon%20Case",
			Stock:    0,
		},
		{
			Name:     "eSports Winter Case 2013",
			HashName: "eSports%202013%20Winter%20Case",
			Stock:    0,
		},
		{
			Name:     "Weapon Case 3",
			HashName: "CS%3AGO%20Weapon%20Case%203",
			Stock:    0,
		},
		{
			Name:     "Weapon Case 2",
			HashName: "CS%3AGO%20Weapon%20Case%202",
			Stock:    0,
		},
		{
			Name:     "Operation Hydra Case",
			HashName: "Operation%20Hydra%20Case",
			Stock:    0,
		},
		{
			Name:     "eSports 2013 Case",
			HashName: "eSports%202013%20Case",
			Stock:    0,
		},
		{
			Name:     "Operation Bravo Case",
			HashName: "Operation%20Bravo%20Case",
			Stock:    0,
		},
		{
			Name:     "Weapon Case",
			HashName: "CS%3AGO%20Weapon%20Case",
			Stock:    0,
		},
	}

	var fields []EmbedField

	for _, item := range items {

		resp, err := c.Get(baseUrl + item.HashName)

		if err != nil {
			fmt.Printf("Error %s", err)
			return fields
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		var response Response
		json.Unmarshal([]byte(body), &response)

		var newField EmbedField
		
		newField.Name = item.Name
		
		lowestPriceFloat, err := strconv.ParseFloat(normalizeGermanFloatString(strings.TrimSuffix(response.LowestPrice, "€")), 32)
		
		log.Printf("Price for %s is %f", newField.Name, lowestPriceFloat)
		
		itemValue := lowestPriceFloat * float32(item.Stock)
		
		newField.Value = fmt.Sprintf("%f", itemValue)
		
		log.Printf("Field: %s, Value: %s", newField.Name, newField.Value)
		fields = append(fields, newField)

		time.Sleep(5 * time.Second)
	}	
	return fields
}

func buildMessage(fields []EmbedField) Message {
	var message Message
	message.Fields = fields

	totalValue := 0

	for _, field := range fields {
		value, err := strconv.Atoi(field.Value)
		if err != nil {
			fmt.Printf("Error %s", err)
		}
		
		totalValue = totalValue + value
		
		log.Printf("Current value: %d", totalValue)
	}

	message.Title = fmt.Sprintf("Your inventory has a value of %d€", totalValue)

	return message
}

func sendDiscordMessage(session *discordgo.Session, message Message) {

	_, err := session.ChannelMessageSendEmbed(*Channel, &discordgo.MessageEmbed{
		Title: message.Title,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "test1",
				Value: "test123",
			},
			{
				Name:  "test2",
				Value: "test1234",
			},
		},
	})
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}

func normalizeGermanFloatString(old string) string {
    s := strings.Replace(old, ",", ".", -1)
    return strings.Replace(s, ".", "", 1)
}
