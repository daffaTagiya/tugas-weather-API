package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

const apiKey = "b17e6581d71642dbba8ddeeffbd24769" // Ganti dengan kunci API cuaca Anda

func main() {
	city := "Jakarta" // Ganti dengan kota yang ingin Anda periksa cuacanya

	// URL API cuaca (contoh: OpenWeatherMap)
	apiURL := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	// Membuat instance dari resty.Client
	client := resty.New()

	// Melakukan permintaan HTTP GET ke API cuaca
	resp, err := client.R().Get(apiURL)
	if err != nil {
		log.Fatal(err)
	}

	// Memeriksa apakah permintaan berhasil (status kode 200)
	if resp.StatusCode() == 200 {
		// Menguraikan data JSON respons
		var weatherData map[string]interface{}
		err := json.Unmarshal(resp.Body(), &weatherData)
		if err != nil {
			log.Fatal(err)
		}

		// Menampilkan data cuaca
		fmt.Println("Weather in", city)
		fmt.Println("Description:", weatherData["weather"].([]interface{})[0].(map[string]interface{})["description"])
		fmt.Println("Temperature:", weatherData["main"].(map[string]interface{})["temp"])
		fmt.Println("Humidity:", weatherData["main"].(map[string]interface{})["humidity"])
	} else {
		log.Printf("Error: %v", resp.Status())
	}
}
