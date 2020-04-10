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
	Prefectures         []types.Prefecture
	Patients            []types.Patient
	Stats               []types.Stat
	DailyReport         []types.DateReport
	DailyPositiveByPref []types.DatePositiveByPref
	DailyDeathByPref    []types.DateDeathByPref
	DailySexByPref      []types.DateSexByPref
	DailyAgeByPref      []types.DateAgeByPref
	DailyCallcenter     []types.DateCallcenter
	DailyShip           []types.DateShip
	News                []types.News
	DetailByRegion      []types.RegionDetail
}

func NewData() *Data {
	d := &Data{
		Prefectures:         []types.Prefecture{},
		Patients:            []types.Patient{},
		Stats:               []types.Stat{},
		DailyReport:         []types.DateReport{},
		DailyPositiveByPref: []types.DatePositiveByPref{},
		DailyDeathByPref:    []types.DateDeathByPref{},
		DailySexByPref:      []types.DateSexByPref{},
		DailyAgeByPref:      []types.DateAgeByPref{},
		DailyCallcenter:     []types.DateCallcenter{},
		DailyShip:           []types.DateShip{},
		News:                []types.News{},
		DetailByRegion:      []types.RegionDetail{},
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

func SendDailyPositiveByPref() echo.HandlerFunc {
	return func(c echo.Context) error {
		dailyPositiveByPref := APIData.DailyPositiveByPref
		dailyPositiveByPrefjson, _ := json.Marshal(dailyPositiveByPref)
		return c.String(http.StatusOK, string(dailyPositiveByPrefjson))
	}
}

func SendDailyDeathByPref() echo.HandlerFunc {
	return func(c echo.Context) error {
		dailyDeathByPref := APIData.DailyDeathByPref
		dailyDeathByPrefjson, _ := json.Marshal(dailyDeathByPref)
		return c.String(http.StatusOK, string(dailyDeathByPrefjson))
	}
}

func SendDailySexByPref() echo.HandlerFunc {
	return func(c echo.Context) error {
		dailySexByPref := APIData.DailySexByPref
		dailySexByPrefjson, _ := json.Marshal(dailySexByPref)
		return c.String(http.StatusOK, string(dailySexByPrefjson))
	}
}

func SendDailyAgeByPref() echo.HandlerFunc {
	return func(c echo.Context) error {
		dailyAgeByPref := APIData.DailyAgeByPref
		dailyAgeByPrefjson, _ := json.Marshal(dailyAgeByPref)
		return c.String(http.StatusOK, string(dailyAgeByPrefjson))
	}
}

func SendDailyCallcenter() echo.HandlerFunc {
	return func(c echo.Context) error {
		dailyCallcenter := APIData.DailyCallcenter
		dailyCallcenterjson, _ := json.Marshal(dailyCallcenter)
		return c.String(http.StatusOK, string(dailyCallcenterjson))
	}
}

func SendDailyShip() echo.HandlerFunc {
	return func(c echo.Context) error {
		dailyShip := APIData.DailyShip
		dailyShipjson, _ := json.Marshal(dailyShip)
		return c.String(http.StatusOK, string(dailyShipjson))
	}
}

func SendNews() echo.HandlerFunc {
	return func(c echo.Context) error {
		news := APIData.News
		newsjson, _ := json.Marshal(news)
		return c.String(http.StatusOK, string(newsjson))
	}
}

func SendDetailByRegion() echo.HandlerFunc {
	return func(c echo.Context) error {
		detailByRegion := APIData.DetailByRegion
		detailByRegionjson, _ := json.Marshal(detailByRegion)
		return c.String(http.StatusOK, string(detailByRegionjson))
	}
}
