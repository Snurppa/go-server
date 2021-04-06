package db

import (
	"fmt"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func get_client(token string) influxdb2.Client {
	return influxdb2.NewClient("https://europe-west1-1.gcp.cloud2.influxdata.com", token)
}

func WriteMoisture(org string, bucket string, sensor string, value int) {
	// get non-blocking write client
	token := os.Getenv("INFLUX_TOKEN")

	var client = get_client(token)
	// You can generate a Token from the "Tokens Tab" in the UI
	writeAPI := client.WriteAPI(org, bucket)

	// write line protocol
	writeAPI.WriteRecord(fmt.Sprintf("moisture,sensor=%s value=%d %d", sensor, value, time.Now().UnixNano()))
	// Flush writes
	writeAPI.Flush()
	// always close client at the end

	defer client.Close()
}
