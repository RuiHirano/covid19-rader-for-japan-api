package main

import (
	"fmt"
	"handler"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"types"

	"github.com/carlescere/scheduler"
	"github.com/jszwec/csvutil"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	SRC_DIR string
)

func init() {
	// csv読み込み
	p, _ := os.Getwd()
	fmt.Println(p)
	SRC_DIR = p

	// for test
	/*dpp := fetchDailyPositiveByPref()
	for _, v := range dpp {
		fmt.Printf("debug %+v\n", v)
	}
	houdou := fetchNews()
	for _, v := range houdou {
		fmt.Printf("debug %+v\n", v)
	}
	deaths := fetchDailyDeathByPref()
	for _, v := range deaths {
		fmt.Printf("debug %+v\n", v)
	}
	details := fetchDailyDetailByPref()
	for _, v := range details {
		fmt.Printf("debug %+v\n", v)
	}
	callCenters := fetchDailyCallcenter()
	for _, v := range callCenters {
		fmt.Printf("debug %+v\n", v)
	}
	summary := fetchDailySummary()
	for _, v := range summary {
		fmt.Printf("debug %+v\n", v)
	}
	worldsummary := fetchDailyWorldSummary()
	for _, v := range worldsummary {
		fmt.Printf("debug %+v\n", v)
	}*/
}

func fetchDataSceduler() {
	scheduler.Every(2).Hours().Run(fetchData)
}

func fetchData() {

	// dailyPositiveByPref data
	dailyPositiveByPref := fetchDailyPositiveByPref()
	handler.APIData.DailyPositiveByPref = dailyPositiveByPref

	// dailyDeathByPref data
	dailyDeathByPref := fetchDailyDeathByPref()
	handler.APIData.DailyDeathByPref = dailyDeathByPref

	// dailyDetailByPref data
	dailyDetailByPref := fetchDailyDetailByPref()
	handler.APIData.DailyDetailByPref = dailyDetailByPref

	// dailyCallcenter data
	dailyCallcenter := fetchDailyCallcenter()
	handler.APIData.DailyCallcenter = dailyCallcenter

	// news data
	news := fetchNews()
	handler.APIData.News = news

	dailySummary := fetchDailySummary()
	handler.APIData.DailySummary = dailySummary

	dailyWorldSummary := fetchDailyWorldSummary()
	handler.APIData.DailyWorldSummary = dailyWorldSummary
}

func insertCommaBeforeDate(b []byte, date_name string, md_pattern string) []byte {
	byteStr := string(b)
	rep := regexp.MustCompile("2020" + md_pattern + ",")
	byteStr = rep.ReplaceAllString(byteStr, ",2020$1$2,")
	rep = regexp.MustCompile(date_name + `,`)
	byteStr = rep.ReplaceAllString(byteStr, ","+date_name+",")
	return []byte(byteStr)
}

func fetchDailyDetailByPref() []types.DateDetailByPref {
	var dailyDetailByPref []types.DateDetailByPref
	// バイト列を読み込む
	b, _ := ioutil.ReadFile(SRC_DIR + "/2019-ncov-japan/50_Data/covid19_jp.csv")
	b = insertCommaBeforeDate(b, "date", "-(..)-(..)")
	// ユーザー定義型スライスにマッピング
	if err := csvutil.Unmarshal(b, &dailyDetailByPref); err != nil {
		fmt.Println("error:", err)
	}
	return dailyDetailByPref
}

func fetchDailyPositiveByPref() []types.DatePositiveByPref {
	var dailyPositiveByPref []types.DatePositiveByPref
	// バイト列を読み込む
	b, _ := ioutil.ReadFile(SRC_DIR + "/2019-ncov-japan/50_Data/byDate.csv")
	b = insertCommaBeforeDate(b, "date", "(..)(..)")
	// ユーザー定義型スライスにマッピング
	if err := csvutil.Unmarshal(b, &dailyPositiveByPref); err != nil {
		fmt.Println("error:", err)
	}
	return dailyPositiveByPref
}

func fetchDailyDeathByPref() []types.DateDeathByPref {
	var dailyDeathByPref []types.DateDeathByPref
	// バイト列を読み込む
	b, _ := ioutil.ReadFile(SRC_DIR + "/2019-ncov-japan/50_Data/death.csv")
	b = insertCommaBeforeDate(b, "date", "(..)(..)")
	// ユーザー定義型スライスにマッピング
	if err := csvutil.Unmarshal(b, &dailyDeathByPref); err != nil {
		fmt.Println("error:", err)
	}
	return dailyDeathByPref
}

func fetchDailyCallcenter() []types.DateCallcenter {
	var dailyCallcenter []types.DateCallcenter
	// バイト列を読み込む
	b, _ := ioutil.ReadFile(SRC_DIR + "/2019-ncov-japan/50_Data/MHLW/callCenter.csv")
	b = insertCommaBeforeDate(b, "date", "(..)(..)")
	// ユーザー定義型スライスにマッピング
	if err := csvutil.Unmarshal(b, &dailyCallcenter); err != nil {
		fmt.Println("error:", err)
	}
	return dailyCallcenter
}

func fetchDailySummary() []types.DateSummary {
	var dailySummary []types.DateSummary
	// バイト列を読み込む
	b, _ := ioutil.ReadFile(SRC_DIR + "/2019-ncov-japan/50_Data/MHLW/summary.csv")
	//fmt.Println("debug:", string(b))
	b = insertCommaBeforeDate(b, "日付", "(..)(..)")
	// ユーザー定義型スライスにマッピング
	if err := csvutil.Unmarshal(b, &dailySummary); err != nil {
		fmt.Println("error:", err)
	}
	return dailySummary
}

func fetchDailyWorldSummary() []types.DateWorldSummary {
	var dailyWorldSummary []types.DateWorldSummary
	// バイト列を読み込む
	b, _ := ioutil.ReadFile(SRC_DIR + "/2019-ncov-japan/50_Data/FIND/worldSummary.csv")
	//fmt.Println("debug:", string(b))
	b = insertCommaBeforeDate(b, ",date", "-(..)-(..)")
	// ユーザー定義型スライスにマッピング
	if err := csvutil.Unmarshal(b, &dailyWorldSummary); err != nil {
		fmt.Println("error:", err)
	}
	return dailyWorldSummary
}

func fetchNews() []types.News {
	news := []types.News{}
	// バイト列を読み込む
	b, _ := ioutil.ReadFile(SRC_DIR + "/2019-ncov-japan/50_Data/mhlw_houdou.csv")
	//b = insertCommaBeforeDate(b)
	// ユーザー定義型スライスにマッピング
	if err := csvutil.Unmarshal(b, &news); err != nil {
		fmt.Println("error:", err)
	}

	return news
}

func main() {
	fmt.Printf("Starting...")

	go fetchDataSceduler()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/api/v2/daily/positive_by_pref", handler.SendDailyPositiveByPref()) // 日ごと、都道府県ごとの感染者データ
	e.GET("/api/v2/daily/death_by_pref", handler.SendDailyDeathByPref())       // 日ごと、都道府県ごとの死亡者データ
	e.GET("/api/v2/daily/detail_by_pref", handler.SendDailyDetailByPref())     // 日ごと、都道府県ごとの死亡者データ
	e.GET("/api/v2/daily/callcenter", handler.SendDailyCallcenter())           // 日ごと、コールセンターの感染者データ
	e.GET("/api/v2/news", handler.SendNews())                                  // 日ごと、都道府県ごとの感染者データ
	e.GET("/api/v2/summary", handler.SendSummary())                            //
	e.GET("/api/v2/world_summary", handler.SendWorldSummary())                 //

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
		log.Printf("Defaulting to port %s", port)
	}
	e.Start(":" + port)

}
