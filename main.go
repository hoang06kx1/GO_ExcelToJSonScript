package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/tealeg/xlsx"
)

type station struct {
	LineName   string
	Position   string
	Ward       string
	District   string
	Province   string
	Latitude   string
	Longtitude string
	PowerLevel string
	Zone       string
	Team       string
	Height     float64
	ColumnType string
	Box        string
	Note       string
}

type stationsType []station

func main() {
	excelFileName := "tram.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		panic(err)
	}

	var stations stationsType
	// create station list from first sheet
	for i, row := range xlFile.Sheets[0].Rows[1:] {
		// iterate through all struct fields and get corresponding data from excel
		if len(row.Cells[0].String()) == 0 {
			continue
		}
		station := station{}
		station.LineName = row.Cells[0].String()
		station.Position = row.Cells[1].String()
		station.Ward = row.Cells[2].String()
		station.District = row.Cells[3].String()
		station.Province = row.Cells[4].String()
		station.Latitude = row.Cells[5].String()
		station.Longtitude = row.Cells[6].String()
		station.PowerLevel = row.Cells[7].String()
		station.Zone = row.Cells[8].String()
		station.Team = row.Cells[9].String()
		station.Height, err = row.Cells[10].Float()
		if err != nil {
			station.Height = 0
			fmt.Printf("Error parsing height of station number %d: %s \n", i, row.Cells[10].String())
		}
		station.ColumnType = row.Cells[11].String()
		station.Box = row.Cells[12].String()
		station.Note = row.Cells[13].String()
		stations = append(stations, station)
	}

	stationsJSON, err := json.Marshal(stations)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", stationsJSON)
	err = ioutil.WriteFile("stations.json", stationsJSON, 0644)
	if err != nil {
		panic(err)
	}
	println(len(stations))
}
