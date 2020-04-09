package handler

import (

	//"fmt"
	"encoding/json"
	"net/http"

	//"os"
	"strconv"
	"strings"

	//"encoding/csv"
	"types"

	"github.com/labstack/echo"
)

type Data struct {
	Prefectures           []types.Prefecture
	Patients              []types.Patient
	Stats                 []types.Stat
	PrefsData             []*types.PrefData
	TotalData             *types.TotalData
	DailyData             []*types.DateData
	DailyPrefData         []*types.DateData
	PatientsData          []*types.Patient
	DailyPatientsData     []*types.DatePatientsData
	PrefsPatientsData     []*types.PrefPatientsData
	DailyPrefPatientsData []*types.DatePrefPatientsData
}

func NewData() *Data {
	d := &Data{
		Prefectures:           []types.Prefecture{},
		Patients:              []types.Patient{},
		Stats:                 []types.Stat{},
		PrefsData:             []*types.PrefData{},
		TotalData:             &types.TotalData{},
		DailyData:             []*types.DateData{},
		DailyPrefData:         []*types.DateData{},
		PatientsData:          []*types.Patient{},
		DailyPatientsData:     []*types.DatePatientsData{},
		PrefsPatientsData:     []*types.PrefPatientsData{},
		DailyPrefPatientsData: []*types.DatePrefPatientsData{},
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

// Finish
// SendPrefsData: 都道府県ごとのデータ
func SendPrefsData() echo.HandlerFunc {
	return func(c echo.Context) error {
		prefsData := APIData.PrefsData
		prefsjson, _ := json.Marshal(prefsData)
		return c.String(http.StatusOK, string(prefsjson))
	}
}

// SendDailyData: 都道府県ごとのデータ
func SendDailyData() echo.HandlerFunc {
	return func(c echo.Context) error {
		dailyData := APIData.DailyData
		dailyjson, _ := json.Marshal(dailyData)
		return c.String(http.StatusOK, string(dailyjson))
	}
}

// SendDailyPrefData: 都道府県ごとのデータ
func SendDailyPrefData() echo.HandlerFunc {
	return func(c echo.Context) error {
		dailyPrefData := APIData.DailyPrefData
		dailyPrefjson, _ := json.Marshal(dailyPrefData)
		return c.String(http.StatusOK, string(dailyPrefjson))
	}
}

// Finish
// SendTotalData: 都道府県ごとのデータ
func SendTotalData() echo.HandlerFunc {
	return func(c echo.Context) error {
		totalData := APIData.TotalData
		patientsjson, _ := json.Marshal(totalData)
		return c.String(http.StatusOK, string(patientsjson))
	}
}

// Finish
// SendPatientsData: 都道府県ごとのデータ
func SendPatientsData() echo.HandlerFunc {
	return func(c echo.Context) error {
		patientsData := APIData.PatientsData
		patientsjson, _ := json.Marshal(patientsData)
		return c.String(http.StatusOK, string(patientsjson))
	}
}

// SendDailyPatientsData: 都道府県ごとのデータ
func SendDailyPatientsData() echo.HandlerFunc {
	return func(c echo.Context) error {
		dailyPatientsData := APIData.DailyPatientsData
		dailyPatientsjson, _ := json.Marshal(dailyPatientsData)
		return c.String(http.StatusOK, string(dailyPatientsjson))
	}
}

// SendPrefPatientsData: 都道府県ごとのデータ
func SendPrefPatientsData() echo.HandlerFunc {
	return func(c echo.Context) error {
		prefPatientsData := APIData.PrefsPatientsData
		prefPatientsjson, _ := json.Marshal(prefPatientsData)
		return c.String(http.StatusOK, string(prefPatientsjson))
	}
}

// SendDailyPrefPatientsData: 都道府県ごとのデータ
func SendDailyPrefPatientsData() echo.HandlerFunc {
	return func(c echo.Context) error {
		dailyPrefPatientsData := APIData.DailyPrefPatientsData
		dailyPrefPatientsjson, _ := json.Marshal(dailyPrefPatientsData)
		return c.String(http.StatusOK, string(dailyPrefPatientsjson))
	}
}

func convertAge(age string) uint64 {
	ageNum := uint64(999) // Unknown
	if strings.Contains(age, "代") {
		ageNum = numCheck(age)
		//ageNum = 20
	}
	return ageNum
}

func strToFloat64(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		f = 0
	}
	return f
}

func convertSex(sexStr string) types.Sex {
	switch sexStr {
	case "男性":
		return types.SexMALE
	case "女性":
		return types.SexFEMALE
	default:
		return types.SexMALE
	}
}

func convertFlag(flag string) bool {
	if flag == "1" {
		return true
	}
	return false
}

func numCheck(s string) uint64 {
	n := 0
	for _, r := range s {
		if '0' <= r && r <= '9' {
			n = n*10 + int(r-'0')
		} else {
			break
		}
	}
	return uint64(n)
}
