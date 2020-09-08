package main

import (
	"database/sql"
	"fmt"
)

var (
	StationMasterDictID   = make(map[int]Station)
	StationMasterDictName = make(map[string]Station)

	DistanceFareMasterArray = []DistanceFare{}

	FareMasterDict = make(map[string][]Fare)

	SeatMasterDict          = make(map[string][]Seat)
	SeatMasterDictBySmoking = make(map[string][]Seat)
)

func initStationMasterDict() error {
	var stations []Station
	err := dbx.Select(&stations, "SELECT * FROM station_master")
	if err != nil {
		return err
	}

	for _, v := range stations {
		StationMasterDictID[v.ID] = v
		StationMasterDictName[v.Name] = v
	}
	return nil
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

func initDistanceFareMaster() error {
	query := "SELECT * FROM distance_fare_master ORDER BY distance"
	return dbx.Select(&DistanceFareMasterArray, query)
}

func initFareMasterDict() error {
	FareMasterDict = make(map[string][]Fare)
	query := "SELECT * FROM fare_master ORDER BY start_date"
	var fares []Fare
	err := dbx.Select(&fares, query)
	if err != nil {
		return err
	}

	for _, v := range fares {
		key := fmt.Sprintf("%s-%s", v.TrainClass, v.SeatClass)
		FareMasterDict[key] = append(FareMasterDict[key], v)
	}
	return nil
}

func initSeatMasterDict() error {
	SeatMasterDict = make(map[string][]Seat)
	SeatMasterDictBySmoking = make(map[string][]Seat)
	query := "SELECT * FROM seat_master ORDER BY seat_row, seat_column;"
	var seats []Seat
	err := dbx.Select(&seats, query)
	if err != nil {
		return err
	}

	for _, v := range seats {
		key := fmt.Sprintf("%s-%d", v.TrainClass, v.CarNumber)
		SeatMasterDict[key] = append(SeatMasterDict[key], v)
		key = fmt.Sprintf("%s-%s-%t", v.TrainClass, v.SeatClass, v.IsSmokingSeat)
		SeatMasterDictBySmoking[key] = append(SeatMasterDictBySmoking[key], v)
	}
	return nil
}
