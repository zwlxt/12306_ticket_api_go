package query_parser

import (
	"encoding/json"
	"io"
	"log"
	"strings"
)

const (
	STATION_TRAIN_CODE = "station_train_code"
	START_STATION_NAME = "start_station_name"
	FROM_STATION_NAME  = "from_station_name"
	TO_STATION_NAME    = "to_station_name"
	END_STATION_NAME   = "end_station_name"
	START_TIME         = "start_time"
	LISHI              = "lishi"
	DAY_DIFFERENCE     = "day_difference"
	ARRIVE_TIME        = "arrive_time"
	SWZ_NUM            = "swz_num"
	TZ_NUM             = "tz_num"
	ZY_NUM             = "zy_num"
	ZE_NUM             = "ze_num"
	RW_NUM             = "rw_num"
	YW_NUM             = "yw_num"
	WZ_NUM             = "wz_num"
)

type QueryResponse struct {
	StationTrainCode,
	StartStationName,
	FromStationName,
	ToStationName,
	EndStationName,
	StartTime,
	Lishi,
	DayDifferece,
	ArriveTime,
	SwzNum,
	TzNum,
	ZyNum,
	ZeNum,
	RwNum,
	YwNum,
	WzNum string
}

func ParseQuery(jsonStream string) []*QueryResponse {
	var qrs []*QueryResponse
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	isValue := false
	isAfterValue := false
	var outputPrefix string
	var qr *QueryResponse
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("parse query", err)
		}
		if tk, ok := t.(json.Delim); ok {
			if tk.String() == "}" {
				if isAfterValue {
					qrs = append(qrs, qr)
					isAfterValue = false
				}
				continue
			} else if tk.String() == "{" {
				qr = new(QueryResponse)
			}
		}
		if ts, ok := t.(string); ok {
			if isValue == true {
				//log.Println(outputPrefix, "\t", ts)
				switch outputPrefix {
				case STATION_TRAIN_CODE:
					qr.StationTrainCode = ts
					isAfterValue = true
				case START_TIME:
					qr.StartTime = ts
					isAfterValue = true
				case ARRIVE_TIME:
					qr.ArriveTime = ts
					isAfterValue = true
				case START_STATION_NAME:
					qr.StartStationName = ts
					isAfterValue = true
				case FROM_STATION_NAME:
					qr.FromStationName = ts
					isAfterValue = true
				case TO_STATION_NAME:
					qr.ToStationName = ts
					isAfterValue = true
				case END_STATION_NAME:
					qr.EndStationName = ts
					isAfterValue = true
				case LISHI:
					qr.Lishi = ts
					isAfterValue = true
				case DAY_DIFFERENCE:
					qr.DayDifferece = ts
					isAfterValue = true
				case SWZ_NUM:
					qr.SwzNum = ts
					isAfterValue = true
				case TZ_NUM:
					qr.TzNum = ts
					isAfterValue = true
				case ZY_NUM:
					qr.ZyNum = ts
					isAfterValue = true
				case ZE_NUM:
					qr.ZeNum = ts
					isAfterValue = true
				case RW_NUM:
					qr.RwNum = ts
					isAfterValue = true
				case YW_NUM:
					qr.YwNum = ts
					isAfterValue = true
				case WZ_NUM:
					qr.WzNum = ts
					isAfterValue = true
				default:

				}
				isValue = false
			}
			switch ts {
			case STATION_TRAIN_CODE:
				fallthrough
			case START_TIME:
				fallthrough
			case ARRIVE_TIME:
				fallthrough
			case START_STATION_NAME:
				fallthrough
			case FROM_STATION_NAME:
				fallthrough
			case TO_STATION_NAME:
				fallthrough
			case END_STATION_NAME:
				fallthrough
			case LISHI:
				fallthrough
			case DAY_DIFFERENCE:
				fallthrough
			case SWZ_NUM:
				fallthrough
			case TZ_NUM:
				fallthrough
			case ZY_NUM:
				fallthrough
			case ZE_NUM:
				fallthrough
			case RW_NUM:
				fallthrough
			case YW_NUM:
				fallthrough
			case WZ_NUM:
				outputPrefix = ts
				isValue = true

			default:
			}
		}
	}
	return qrs
}

//func main() {
//	str, err := ioutil.ReadFile("../query.json")
//	if err != nil {
//		log.Fatal("Read error")
//	}
//	qrs := ParseQuery(string(str))
//	for _, qr := range qrs {
//		log.Println(qr.StationTrainCode, qr.StartStationName)
//	}
//}
