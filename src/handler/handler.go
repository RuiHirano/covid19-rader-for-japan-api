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
	DailyPositiveByPref []types.DatePositiveByPref
	DailyDeathByPref    []types.DateDeathByPref
	DailyDetailByPref   []types.DateDetailByPref
	DailyCallcenter     []types.DateCallcenter
	News                []types.News
	DailySummary        []types.DateSummary
	DailyWorldSummary   []types.DateWorldSummary
}

func NewData() *Data {
	d := &Data{
		DailyPositiveByPref: []types.DatePositiveByPref{},
		DailyDeathByPref:    []types.DateDeathByPref{},
		DailyDetailByPref:   []types.DateDetailByPref{},
		DailySummary:        []types.DateSummary{},
		DailyWorldSummary:   []types.DateWorldSummary{},
		DailyCallcenter:     []types.DateCallcenter{},
		News:                []types.News{},
	}
	return d
}

var (
	APIData *Data
)

func init() {
	APIData = NewData()
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

func SendDailyDetailByPref() echo.HandlerFunc {
	return func(c echo.Context) error {
		dailyDetailByPref := APIData.DailyDetailByPref
		dailyDetailByPrefjson, _ := json.Marshal(dailyDetailByPref)
		return c.String(http.StatusOK, string(dailyDetailByPrefjson))
	}
}

func SendDailySummary() echo.HandlerFunc {
	return func(c echo.Context) error {
		dailySummary := APIData.DailySummary
		dailySummaryjson, _ := json.Marshal(dailySummary)
		return c.String(http.StatusOK, string(dailySummaryjson))
	}
}

func SendDailyWorldSummary() echo.HandlerFunc {
	return func(c echo.Context) error {
		dailyWorldSummary := APIData.DailyWorldSummary
		dailyWorldSummaryjson, _ := json.Marshal(dailyWorldSummary)
		return c.String(http.StatusOK, string(dailyWorldSummaryjson))
	}
}

func SendDailyCallcenter() echo.HandlerFunc {
	return func(c echo.Context) error {
		dailyCallcenter := APIData.DailyCallcenter
		dailyCallcenterjson, _ := json.Marshal(dailyCallcenter)
		return c.String(http.StatusOK, string(dailyCallcenterjson))
	}
}

func SendNews() echo.HandlerFunc {
	return func(c echo.Context) error {
		news := APIData.News
		newsjson, _ := json.Marshal(news)
		return c.String(http.StatusOK, string(newsjson))
	}
}

func SendSummary() echo.HandlerFunc {
	return func(c echo.Context) error {
		summary := APIData.DailySummary
		summaryjson, _ := json.Marshal(summary)
		return c.String(http.StatusOK, string(summaryjson))
	}
}

func SendWorldSummary() echo.HandlerFunc {
	return func(c echo.Context) error {
		worldSummary := APIData.DailyWorldSummary
		worldSummaryjson, _ := json.Marshal(worldSummary)
		return c.String(http.StatusOK, string(worldSummaryjson))
	}
}
