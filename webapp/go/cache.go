package main

import (
	"database/sql"
	"fmt"
)

var (
	StationMasterDictID   = make(map[int]Station)
	StationMasterDictName = make(map[string]Station)

	FareMasterDict = make(map[string][]Fare)
)

func initStationMasterDict() {
	var stations []Station
	err := dbx.Select(&stations, "SELECT * FROM station_master")
	if err != nil {
		panic(err)
	}

	for _, v := range stations {
		StationMasterDictID[v.ID] = v
		StationMasterDictName[v.Name] = v
	}

	return
}

func FetchStationMasterByID(id int) (Station, error) {
	v, ok := StationMasterDictID[id]
	if ok {
		return v, nil
	}
	return Station{}, sql.ErrNoRows
}

func FetchStationMasterByName(name string) (Station, error) {
	v, ok := StationMasterDictName[name]
	if ok {
		return v, nil
	}
	return Station{}, sql.ErrNoRows
}

func initFareMaster() {
	fareList := []Fare{}
	query := "SELECT * FROM fare_master ORDER BY start_date"
	err := dbx.Select(&fareList, query)
	if err != nil {
		panic(err)
	}

	for _, v := range fareList {
		key := fmt.Sprintf("%s-%s", v.TrainClass, v.SeatClass)
		FareMasterDict[key] = append(FareMasterDict[key], v)
	}
}

func FetchFare(trainClass, seatClass string) ([]Fare, error) {
	key := fmt.Sprintf("%s-%s", trainClass, seatClass)
	v, ok := FareMasterDict[key]
	if ok {
		return v, nil
	}
	return nil, sql.ErrNoRows
}
