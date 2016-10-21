package stationname_parser

import (
	"log"
	"os"
	"testing"
)

func TestParseStations(t *testing.T) {
	file, err := os.Open("./station_name.js")
	if err != nil {
		panic(err)
	}
	stations := ParseStations(file)
	i := 0
	for k, v := range stations {
		if k == "" || v == "" {
			t.Errorf(k, v)
		} else {
			log.Println(k, v)
		}
		i++
	}
	if i != 2562 {
		t.Errorf("Count error")
	}
}
