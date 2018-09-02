package main

import (
		"log"
		"github.com/morus12/dht22"
)

func main(){
	sensor := dht22.New("GPIO_4")
	temperature, err := sensor.Temperature()
	if err != nil {
		log.Fatal(err)
	}else{
		log.Printf("temperature:%v\n",temperature)
	}
		
	humidity, err := sensor.Humidity()
	if err != nil {
		log.Fatal(err)
	}else{
		log.Printf("humidity:%v\n",humidity)
	}


	
}
