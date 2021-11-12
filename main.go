package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


var products string= "https://reserve-prime.apple.com/CN/zh_CN/reserve/A/availability?iUP=N"

func main() {

	var availability_url = "https://reserve-prime.apple.com/CN/zh_CN/reserve/A/availability.json"
	var stores map[string]string
	stores = make(map[string]string)
	stores["R643"] = "虹悦城"
	stores["R493"] = "南京艾尚天地"
	stores["R703"] = "玄武湖"

	products := map[string]string{"MLT93CH/A":"iPhone 13 Pro 256GB 石墨色","MLTE3CH/A":"iPhone 13 Pro 256GB 远峰蓝色"}

	availability,_ := http.Get(availability_url)
	for store := range stores{
		for product := range products{
			result_body, _ := ioutil.ReadAll(availability.Body)
			var jsonBody interface{}

			json.Unmarshaler(result_body,&jsonBody)
			product_availability := jsonBody["stores"][stores[store]][products[product]]
			fmt.Println()
		}
	}

	fmt.Println(string(result_body))
}