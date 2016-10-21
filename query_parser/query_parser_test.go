package query_parser

import (
	"io/ioutil"
	"testing"
)

func TestParseQuery(t *testing.T) {
	str, _ := ioutil.ReadFile("../query.json")
	qrs := ParseQuery(string(str))
	for _, qr := range qrs {
		if qr.StationTrainCode == "" {
			t.Errorf("StationTrainCode NO OUTPUT")
		} else {
			t.Log(qr.StationTrainCode)
		}
		if qr.StartStationName == "" {
			t.Errorf("StartStationName NO OUTPUT")
		} else {
			t.Log(qr.StartStationName)
		}
		if qr.ToStationName == "" {
			t.Errorf("ToStationName NO OUTPUT")
		} else {
			t.Log(qr.ToStationName)
		}
		if qr.StartTime == "" {
			t.Errorf("StartTime NO OUTPUT")
		} else {
			t.Log(qr.StartTime)
		}
		if qr.Lishi == "" {
			t.Errorf("Lishi NO OUTPUT")
		} else {
			t.Log(qr.Lishi)
		}
		if qr.DayDifferece == "" {
			t.Errorf("DayDifferece NO OUTPUT")
		} else {
			t.Log(qr.DayDifferece)
		}
		if qr.ArriveTime == "" {
			t.Errorf("ArriveTime NO OUTPUT")
		} else {
			t.Log(qr.ArriveTime)
		}
		if qr.SwzNum == "" {
			t.Errorf("SwzNum NO OUTPUT")
		} else {
			t.Log(qr.SwzNum)
		}
		if qr.TzNum == "" {
			t.Errorf("TzNum NO OUTPUT")
		} else {
			t.Log(qr.TzNum)
		}
		if qr.ZyNum == "" {
			t.Errorf("ZyNum NO OUTPUT")
		} else {
			t.Log(qr.ZyNum)
		}
		if qr.ZeNum == "" {
			t.Errorf("ZeNum NO OUTPUT")
		} else {
			t.Log(qr.ZeNum)
		}
		if qr.RwNum == "" {
			t.Errorf("RwNum NO OUTPUT")
		} else {
			t.Log(qr.RwNum)
		}
		if qr.YwNum == "" {
			t.Errorf("YwNum NO OUTPUT")
		} else {
			t.Log(qr.YwNum)
		}
		if qr.WzNum == "" {
			t.Errorf("WzNum NO OUTPUT")
		} else {
			t.Log(qr.WzNum)
		}
	}
}
