package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

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

type distanceType struct {
	LineName            string
	Position            string
	Distance            int
	Incremental         int
	ReversedIncremental int
}

type stationsType []station
type distancesType []distanceType

func main() {
	excelFileName := "tram.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		panic(err)
	}

	var stations stationsType
	// 1. Create station list from first sheet
	for i, row := range xlFile.Sheets[0].Rows[1:] {
		// iterate through all struct fields and get corresponding data from excel
		if len(strings.TrimSpace(row.Cells[0].String())) == 0 {
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
	err = ioutil.WriteFile("stations.json", stationsJSON, 0644)
	if err != nil {
		panic(err)
	}

	// 2. Create distance from second sheet
	var distances distancesType
	var distance distanceType
	for _, row := range xlFile.Sheets[1].Rows[1:] {
		// iterate through all struct fields and get corresponding data from excel

		if len(strings.TrimSpace(row.Cells[0].String())) == 0 && len(strings.TrimSpace(row.Cells[2].String())) == 0 {
			continue
		}
		if len(strings.TrimSpace(row.Cells[0].String())) > 0 { // name row
			distance = distanceType{}
			distance.LineName = row.Cells[0].String()
			distance.Position = row.Cells[1].String()
			distances = append(distances, distance)
		} else {
			distances[len(distances)-1].Distance, err = row.Cells[2].Int()
			if err != nil {
				panic(err)
			}
		}
	}

	distancesJSON, err := json.Marshal(distances)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", distancesJSON)
	err = ioutil.WriteFile("distances.json", distancesJSON, 0644)
	if err != nil {
		panic(err)
	}
}
