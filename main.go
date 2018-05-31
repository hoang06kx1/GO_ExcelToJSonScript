package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/tealeg/xlsx"
)

type station struct {
	lineName   string
	position   string
	ward       string
	district   string
	province   string
	latitude   string
	longtitude string
	powerLevel string
	zone       string
	team       string
	height     float64
	columnType string
	box        string
	note       string
}

func main() {
	excelFileName := "tram.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		panic(err)
	}

	var stations []station
	// create station list from first sheet
	for i, row := range xlFile.Sheets[0].Rows[1:] {
		// iterate through all struct fields and get corresponding data from excel
		if len(row.Cells[0].String()) == 0 {
			continue
		}
		station := station{}
		station.lineName = row.Cells[0].String()
		station.position = row.Cells[1].String()
		station.ward = row.Cells[2].String()
		station.district = row.Cells[3].String()
		station.province = row.Cells[4].String()
		station.latitude = row.Cells[5].String()
		station.longtitude = row.Cells[6].String()
		station.powerLevel = row.Cells[7].String()
		station.zone = row.Cells[8].String()
		station.team = row.Cells[9].String()
		station.height, err = row.Cells[10].Float()
		if err != nil {
			fmt.Printf("Error parsing height of station number %d: %s \n", i, row.Cells[10].String())
		}
		station.columnType = row.Cells[11].String()
		station.box = row.Cells[12].String()
		station.note = row.Cells[13].String()
		stations = append(stations, station)
	}

	jsonStations, err := json.Marshal(stations)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("stations.json", jsonStations, 0644)
	if err != nil {
		panic(err)
	}
	println(len(stations))
}
