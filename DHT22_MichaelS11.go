package main

import (
	"fmt"
	"log"
	//"time"
	"github.com/MichaelS11/go-dht"
)

func main() {
	err := dht.HostInit()
	if err != nil {
		fmt.Println("HostInit error:", err)
		return
	}

	dht, err := dht.NewDHT("GPIO4", dht.Celsius, "")
	if err != nil {
		fmt.Println("NewDHT error:", err)
		return
	}
	
	for {
		humidity, temperature, err := dht.ReadRetry(11)
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}
		log.Printf("humidity: %v\n", humidity)
		log.Printf("temperature: %v\n", temperature)
		
		//time.Sleep (1 * time.Second)
	}
}
