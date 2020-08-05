package main

import "database/sql"

var (
	StationMasterDictID   = make(map[int]Station)
	StationMasterDictName = make(map[string]Station)
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
