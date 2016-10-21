package stationname_parser

import (
	"bufio"
	"io"
	"os"
)

func ParseStations(file *os.File) map[string]string {
	var count, temp int = 0, 10
	reader := bufio.NewReader(file)
	stations := make(map[string]string)
	var k, v string
	for {
		segment, err := reader.ReadString('|')
		if err != nil || err == io.EOF {
			break
		}
		count++
		if count == 2 || count == temp+5 {
			k = segment[:len(segment)-1]
			temp = count
		} else if count == temp+1 {
			v = segment[:len(segment)-1]
			stations[k] = v
		} else {
			continue
		}
	}
	return stations
}

//func main() {
//	f, err := os.Open("./station_name.js")
//	if err != nil {
//		panic(err)
//	}
//	defer f.Close()
//	stations := parseStations(f)
//	for k, v := range stations {
//		fmt.Println(k, "\t", v)
//	}
//}
