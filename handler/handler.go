package handler

import (

	//"fmt"
	"encoding/json"
	"net/http"

	//"os"
	//"strconv"
	//"strings"

	//"encoding/csv"
	"types"

	"github.com/labstack/echo"
)

type Data struct {
	Prefectures []types.Prefecture
	Patients    []types.Patient
	Stats       []types.Stat
	DailyReport []types.DateReport
}

func NewData() *Data {
	d := &Data{
		Prefectures: []types.Prefecture{},
		Patients:    []types.Patient{},
		Stats:       []types.Stat{},
		DailyReport: []types.DateReport{},
	}
	return d
}

var (
	APIData *Data
)

func init() {
	APIData = NewData()
}

func SendPatients() echo.HandlerFunc {
	return func(c echo.Context) error {
		patients := APIData.Patients
		patientsjson, _ := json.Marshal(patients)
		return c.String(http.StatusOK, string(patientsjson))
	}
}

func SendPrefectures() echo.HandlerFunc {
	return func(c echo.Context) error {
		prefectures := APIData.Prefectures
		prefecturesjson, _ := json.Marshal(prefectures)
		return c.String(http.StatusOK, string(prefecturesjson))
	}
}

func SendStats() echo.HandlerFunc {
	return func(c echo.Context) error {
		stats := APIData.Stats
		statsjson, _ := json.Marshal(stats)
		return c.String(http.StatusOK, string(statsjson))
	}
}

func SendDailyReport() echo.HandlerFunc {
	return func(c echo.Context) error {
		dailyReport := APIData.DailyReport
		dailyReportjson, _ := json.Marshal(dailyReport)
		return c.String(http.StatusOK, string(dailyReportjson))
	}
}
