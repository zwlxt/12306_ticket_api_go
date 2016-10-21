package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"./query_parser"
	"./stationname_parser"
	"github.com/tealeg/xlsx"
)

const (
	configFile       = "./config.json"
	stationNamesFile = "./stationname_parser/station_name.js"
	queryUrlPattern  = "https://kyfw.12306.cn/otn/lcxxcx/query?purpose_codes=ADULT&queryDate=%s&from_station=%s&to_station=%s"
	excelFileName    = "record.xlsx"
	refreshInterval  = 2 * time.Minute
)

type Config struct {
	StationTrainCodes []string
	FromStationName,
	ToStationName,
	Date,
	Type string
}

func readConfigFile() []*Config {
	f, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal("readConfigFile", err)
		return nil
	}
	var configs []*Config

	dec := json.NewDecoder(bytes.NewReader(f))
	for {
		var c Config
		if err := dec.Decode(&c); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("readConfigFile2", err)
		}
		configs = append(configs, &c)
	}
	return configs
}

func httpGet(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("httpGet", err)
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("httpGet2", err)
		return "", err
	}

	return string(body), nil
}

func buildQueryUrl(from, to, date string) string {
	queryUrl := fmt.Sprintf(queryUrlPattern,
		date,
		from,
		to)
	log.Println("query url", queryUrl)
	return queryUrl
}

func getAvailableSeats(typeName string, query *query_parser.QueryResponse) string {
	switch typeName {
	case "商务座":
		return query.SwzNum
	case "特等座":
		return query.TzNum
	case "一等座":
		return query.ZyNum
	case "二等座":
		return query.ZeNum
	case "软卧":
		return query.RwNum
	case "硬卧":
		return query.YwNum
	case "无座":
		return query.WzNum
	default:
		return ""
	}
}

func main() {
	// Read config
	configs := readConfigFile()
	if configs == nil {
		log.Fatal("main no config")
	}
	// Read station names&codes
	file, err := os.Open(stationNamesFile)
	if err != nil {
		log.Fatal("open station names file", err)
	}
	// Open record file
	excelFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		excelFile = xlsx.NewFile()
	}
	stationNames := stationname_parser.ParseStations(file)
	c := time.Tick(refreshInterval)
	for {
		for _, config := range configs {
			fromStationCode := stationNames[config.FromStationName]
			toStationCode := stationNames[config.ToStationName]
			queryBytes, err := httpGet(buildQueryUrl(
				fromStationCode,
				toStationCode,
				config.Date))
			//queryBytes, err := ioutil.ReadFile("./query.json")
			queries := query_parser.ParseQuery(string(queryBytes))
			if err != nil {
				log.Fatal("main", err)
			}
			for _, query := range queries {
				for _, v := range config.StationTrainCodes {
					// variety of trains
					if v == query.StationTrainCode {
						sheetName := v + "|" + config.FromStationName +
							"-" + config.ToStationName
						sheet, err := excelFile.AddSheet(sheetName)
						if err != nil {
							sheet = excelFile.Sheet[sheetName]
						}
						row := sheet.AddRow()
						cell_formattedTime := row.AddCell()
						cell_unixTime := row.AddCell()
						cell_val := row.AddCell()
						cell_formattedTime.Value = time.Now().Format(
							"2006-01-02 15:04:05")
						cell_unixTime.SetDateTime(time.Now())
						val, err := strconv.Atoi(
							getAvailableSeats(config.Type, query))
						if err != nil {
							log.Fatal("main atoi", err)
						}
						cell_val.SetInt(val)
						excelFile.Save(excelFileName)
						log.Println(query.StationTrainCode,
							query.FromStationName,
							query.ToStationName,
							getAvailableSeats(config.Type, query))
					}
				}
			}
			excelFile.Save(excelFileName)
		}
		<-c // delay
	}
}
