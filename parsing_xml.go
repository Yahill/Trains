package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Trains struct {
	XMLName xml.Name `xml:"TrainLegs"`
	Train   []Train  `xml:"TrainLeg"`
}

type Train struct {
	TrainId             int     `xml:"TrainId,attr"`
	DepartureStationId  int     `xml:"DepartureStationId,attr"`
	ArrivalStationId    int     `xml:"ArrivalStationId,attr"`
	Price               float32 `xml:"Price,attr"`
	ArrivalTimeString   string  `xml:"ArrivalTimeString,attr"`
	DepartureTimeString string  `xml:"DepartureTimeString,attr"`
}

func readXML() (trains Trains) {

	//open file
	xmlFile, err := os.Open("data.xml")
	//handle error if it happens
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Your xml file - successfully opened.")
	//closing xml
	defer xmlFile.Close()
	//read xmlFile as a byte array
	byteValue, _ := ioutil.ReadAll(xmlFile)

	//initializing of the Trains array
	//unmarshalling byteArray
	err = xml.Unmarshal(byteValue, &trains)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	return
}
