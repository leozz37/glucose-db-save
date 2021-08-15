package db

import (
	"context"
	"os"
	"strconv"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func Save(glucoseLevel, glucoseGap string) {
	token := os.Getenv("INFLUXDB_TOKEN")
	bucket := os.Getenv("INFLUXDB_BUCKET")
	org := os.Getenv("INFLUXDB_ORG")

	client := influxdb2.NewClient(os.Getenv("INFLUXDB_URL"), token)
	writeAPI := client.WriteAPIBlocking(org, bucket)

	glucose, _ := strconv.ParseFloat(glucoseLevel, 64)
	gap, _ := strconv.ParseFloat(glucoseGap, 64)

	p := influxdb2.NewPointWithMeasurement("measure").
		AddTag("unit", "glucose").
		AddField("value", glucose).
		AddField("gap", gap).
		SetTime(time.Now())
	writeAPI.WritePoint(context.Background(), p)
	client.Close()
}
