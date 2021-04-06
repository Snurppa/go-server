package main

import "encoding/json"

type SensorData struct {
	From     string
	Location string
	Message  string
	Value    int
}

func ParseSensorData(bytes []byte) (SensorData, error) {
	var data SensorData
	if err := json.Unmarshal(bytes, &data); err != nil {
		return data, err
	}
	return data, nil
}
