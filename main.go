package main

import (
	//"context"

	//"bytes"

	"encoding/json"
	"fmt"
	"strconv"

	"io/ioutil"

	"handler"
	"log"
	"os"

	"types"

	"github.com/PuerkitoBio/goquery"

	"net/http"

	"github.com/carlescere/scheduler"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	//fetchData()
}

func fetchDataSceduler() {
	scheduler.Every(2).Hours().Run(fetchData)
}

func fetchData() {
	// patientsData
	/*patients := fetchPatients()
	handler.APIData.PatientsData = patients

	// totalStatsData
	totalData := fetchTotalData()
	totalData.SexData = calcSexData(patients)
	totalData.AgeData = calcAgeData(patients)
	handler.APIData.TotalData = totalData

	// prefsStatsData
	// prefsPatientsData
	prefsData := fetchPrefsData()
	prefsPatients := []*types.PrefPatientsData{}
	for i, prefData := range prefsData {
		prefPatients := []*types.Patient{}
		for _, patient := range patients {
			if patient.Prefecture == prefData.NameJa {
				prefPatients = append(prefPatients, patient)
			}
		}
		prefsData[i].SexData = calcSexData(prefPatients)
		prefsData[i].AgeData = calcAgeData(prefPatients)
		prefsPatients = append(prefsPatients, &types.PrefPatientsData{
			Id:       prefData.Id,
			NameJa:   prefData.NameJa,
			NameEn:   prefData.NameEn,
			Lat:      prefData.Lat,
			Lng:      prefData.Lng,
			Patients: prefPatients,
		})
	}
	handler.APIData.PrefsPatientsData = prefsPatients
	handler.APIData.PrefsData = prefsData

	// prefsStatsData
	// prefsPatientsData
	dailyData := []*types.DateData{}
	dailyPatients := []*types.DatePatientsData{}
	dailyPatientsMap := map[string]*types.DatePatientsData{}
	for _, patient := range patients {
		date := patient.Onset
		if dailyPatientsMap[date] == nil { // 初期値がない場合代入
			dailyPatientsMap[date] = &types.DatePatientsData{
				Date:     date,
				Patients: []*types.Patient{patient},
			}
		} else {

		}

	}
	handler.APIData.DailyPatientsData = dailyPatients
	handler.APIData.DailyData = dailyData

	//fetchDailyPrefsData(dailyPatients)
	//fetchPrefsDailyData(prefsPatients)*/

	patients := fetchPatients()
	prefectures := fetchPrefectures()
	handler.APIData.Patients = patients
	handler.APIData.Prefectures = prefectures

	stats := calcStats(patients)
	handler.APIData.Stats = stats
}

func calcStats(patients []types.Patient) []types.Stat {
	result := []types.Stat{}
	rawStats := fetchDetailByRegion()
	//dfStats := dataframe.LoadStructs(rawStats)
	dfPatients := dataframe.LoadStructs(patients)

	//fmt.Printf("dfStats: ", dfStats)
	fmt.Printf("dfPatients: ", dfPatients)
	firstDate := rawStats[0].Date
	for i, stat := range rawStats {
		if stat.Date == firstDate {
			// 初日はスキップ
			continue
		}
		date := fmt.Sprintf("%s/%s/%s", string([]rune(stat.Date)[:4]), string([]rune(stat.Date)[4:6]), string([]rune(stat.Date)[6:8]))
		pref := stat.Prefecture
		// その日のpatientsデータ
		dfTgt := dfPatients.Filter(
			dataframe.F{"Prefecture", series.Eq, pref},
		).Filter(
			dataframe.F{"Date", series.Eq, date},
		)
		// その日以前全てのpatientsデータ
		dfLtTgt := dfPatients.Filter(
			dataframe.F{"Prefecture", series.Eq, pref},
		).Filter(
			dataframe.F{"Date", "<=", date},
		)

		// その日以前全てのstatsデータ
		/*dfLtStats := dfStats.Filter(
			dataframe.F{"Prefecture", series.Eq, pref},
		).Filter(
			dataframe.F{"Date", "<=", stat.Date},
		)*/

		// その日の発生件数と累計発生件数
		cases := dfTgt.Nrow()
		totalCases := dfLtTgt.Nrow()

		// 患者数等
		deaths := 0
		discharges := 0
		hospitals := 0
		prevStat := types.NewStat()
		for _, stat_ := range rawStats {
			date, _ := strconv.Atoi(stat.Date)
			if stat_.Prefecture == pref && stat_.Date == strconv.Itoa(date-1) {
				prevStat = stat_
			}
		}
		deaths = stat.TotalDeaths - prevStat.TotalDeaths
		discharges = stat.TotalDischarges - prevStat.TotalDischarges
		hospitals = stat.TotalHospitals - prevStat.TotalHospitals

		// 性別
		maleNum := dfTgt.Filter(dataframe.F{"Sex", series.Eq, "男性"}).Nrow()
		femaleNum := dfTgt.Filter(dataframe.F{"Sex", series.Eq, "女性"}).Nrow()
		unknownNum := cases - femaleNum - maleNum
		totalMaleNum := dfLtTgt.Filter(dataframe.F{"Sex", series.Eq, "男性"}).Nrow()
		totalFemaleNum := dfLtTgt.Filter(dataframe.F{"Sex", series.Eq, "女性"}).Nrow()
		totalUnknownNum := cases - totalFemaleNum - totalMaleNum

		// 年齢
		age10 := dfTgt.Filter(dataframe.F{"Age", series.Eq, "10代"}).Nrow()
		age20 := dfTgt.Filter(dataframe.F{"Age", series.Eq, "20代"}).Nrow()
		age30 := dfTgt.Filter(dataframe.F{"Age", series.Eq, "30代"}).Nrow()
		age40 := dfTgt.Filter(dataframe.F{"Age", series.Eq, "40代"}).Nrow()
		age50 := dfTgt.Filter(dataframe.F{"Age", series.Eq, "50代"}).Nrow()
		age60 := dfTgt.Filter(dataframe.F{"Age", series.Eq, "60代"}).Nrow()
		age70 := dfTgt.Filter(dataframe.F{"Age", series.Eq, "70代"}).Nrow()
		age80 := dfTgt.Filter(dataframe.F{"Age", series.Eq, "80代"}).Nrow()
		age90 := dfTgt.Filter(dataframe.F{"Age", series.Eq, "90代"}).Nrow()
		ageUnknown := cases - age10 - age20 - age30 - age40 - age50 - age60 - age70 - age80 - age90

		totalAge10 := dfLtTgt.Filter(dataframe.F{"Age", series.Eq, "10代"}).Nrow()
		totalAge20 := dfLtTgt.Filter(dataframe.F{"Age", series.Eq, "20代"}).Nrow()
		totalAge30 := dfLtTgt.Filter(dataframe.F{"Age", series.Eq, "30代"}).Nrow()
		totalAge40 := dfLtTgt.Filter(dataframe.F{"Age", series.Eq, "40代"}).Nrow()
		totalAge50 := dfLtTgt.Filter(dataframe.F{"Age", series.Eq, "50代"}).Nrow()
		totalAge60 := dfLtTgt.Filter(dataframe.F{"Age", series.Eq, "60代"}).Nrow()
		totalAge70 := dfLtTgt.Filter(dataframe.F{"Age", series.Eq, "70代"}).Nrow()
		totalAge80 := dfLtTgt.Filter(dataframe.F{"Age", series.Eq, "80代"}).Nrow()
		totalAge90 := dfLtTgt.Filter(dataframe.F{"Age", series.Eq, "90代"}).Nrow()
		totalAgeUnknown := totalCases - totalAge10 - totalAge20 - totalAge30 - totalAge40 - totalAge50 - totalAge60 - totalAge70 - totalAge80 - totalAge90

		// 代入
		rawStats[i].Cases = cases
		rawStats[i].Deaths = deaths
		rawStats[i].Hospitals = hospitals
		rawStats[i].Discharges = discharges
		rawStats[i].SexData = types.SexData{
			Female:  femaleNum,
			Male:    maleNum,
			Unknown: unknownNum,
		}
		rawStats[i].TotalSexData = types.SexData{
			Female:  totalFemaleNum,
			Male:    totalMaleNum,
			Unknown: totalUnknownNum,
		}
		rawStats[i].AgeData = types.AgeData{
			Age10:   age10,
			Age20:   age20,
			Age30:   age30,
			Age40:   age40,
			Age50:   age50,
			Age60:   age60,
			Age70:   age70,
			Age80:   age80,
			Age90:   age90,
			Unknown: ageUnknown,
		}
		rawStats[i].TotalAgeData = types.AgeData{
			Age10:   totalAge10,
			Age20:   totalAge20,
			Age30:   totalAge30,
			Age40:   totalAge40,
			Age50:   totalAge50,
			Age60:   totalAge60,
			Age70:   totalAge70,
			Age80:   totalAge80,
			Age90:   totalAge90,
			Unknown: totalAgeUnknown,
		}

		//fmt.Printf("date %v pref %v cases: %v deaths %v discharges %v hospitals %v male %v femail %v unknown %v\n", date, pref, cases, deaths, discharges, hospitals, maleNum, femaleNum, unknownNum)
		//fmt.Printf("age10 %v age20 %v age30 %v age40 %v age50 %v age60 %v age70 %v age80 %v age90 %v unknown %v\n", age10, age20, age30, age40, age50, age60, age70, age80, age90, ageUnknown)

	}

	result = rawStats
	return result
}

func calcSexData(patients []*types.Patient) *types.SexData {
	sexData := &types.SexData{Female: 0, Male: 0, Unknown: 0}
	for _, patient := range patients {
		fmt.Printf("gender %v\n", patient.Sex, patient, patient.Age)
		if patient.Sex == "女性" {
			sexData.Female++
		} else if patient.Sex == "男性" {
			sexData.Male++
		} else {
			sexData.Unknown++
		}
	}
	return sexData
}

func calcAgeData(patients []*types.Patient) *types.AgeData {
	ageData := &types.AgeData{Age10: 0, Age20: 0, Age30: 0, Age40: 0, Age50: 0, Age60: 0, Age70: 0, Age80: 0, Age90: 0, Unknown: 0}
	for _, patient := range patients {
		switch patient.Age {
		case "10代":
			ageData.Age10++
		case "20代":
			ageData.Age20++
		case "30代":
			ageData.Age30++
		case "40代":
			ageData.Age40++
		case "50代":
			ageData.Age50++
		case "60代":
			ageData.Age60++
		case "70代":
			ageData.Age70++
		case "80代":
			ageData.Age80++
		case "90代":
			ageData.Age90++
		default:
			ageData.Unknown++
		}
	}
	return ageData
}

func fetchPatients() []types.Patient {
	doc, err := goquery.NewDocument("https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/positiveDetail.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println(" doc", doc)
	patients := []types.Patient{}
	selection := doc.Find("tbody")
	innerSelection := selection.Find("tr")
	innerSelection.Each(func(i int, s *goquery.Selection) {
		patient := types.NewPatient()
		s.Find("td").Each(func(k int, s2 *goquery.Selection) {
			switch k {
			case 1:
				patient.ID = s2.Text()
			case 2:
				patient.Date = s2.Text()
			case 4:
				patient.Prefecture = s2.Text()
			case 5:
				patient.Residence = s2.Text()
			case 6:
				patient.Age = s2.Text()
			case 7:
				patient.Sex = s2.Text()
			case 8:
				patient.Attribute = s2.Text()
			case 9:
				patient.PrefectureNumber = s2.Text()
			case 10:
				patient.TravelOrContact = s2.Text()
			case 11:
				patient.Detail = s2.Text()
			case 13:
				patient.Src = s2.Text()
			case 14:
				patient.Onset = s2.Text()
			case 15:
				patient.Symptom = s2.Text()
			case 16:
				patient.DeathOrDischageDate = s2.Text()
			case 17:
				patient.Comment1 = s2.Text()
			case 20:
				patient.Comment2 = s2.Text()
			case 18:
				patient.Outcome = s2.Text()
			case 19:
				patient.OutcomeSrc = s2.Text()

			}
		})
		patients = append(patients, patient)
	})

	for _, patient := range patients {
		fmt.Println("patient %v\n", patient)
	}
	return patients
}

func fetchDetailByRegion() []types.Stat {
	doc, err := goquery.NewDocument("https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/detailByRegion.csv")
	if err != nil {
		panic(err)
	}
	stats := []types.Stat{}
	selection := doc.Find("tbody")
	innerSelection := selection.Find("tr")
	innerSelection.Each(func(i int, s *goquery.Selection) {
		stat := types.NewStat()
		s.Find("td").Each(func(k int, s2 *goquery.Selection) {
			switch k {
			case 1:
				stat.Date = s2.Text()
			case 2:
				stat.Prefecture = s2.Text()
			case 3:
				totalCases, _ := strconv.Atoi(s2.Text())
				stat.TotalCases = totalCases
			case 4:
				totalHospitals, _ := strconv.Atoi(s2.Text())
				stat.TotalHospitals = totalHospitals
			case 5:
				totalDischarges, _ := strconv.Atoi(s2.Text())
				stat.TotalDischarges = totalDischarges
			case 6:
				totalDeaths, _ := strconv.Atoi(s2.Text())
				stat.TotalDeaths = totalDeaths
			}
		})
		stats = append(stats, stat)
	})

	for _, stat := range stats {
		fmt.Println("stat %v\n", stat)
	}
	return stats
}

func fetchPrefectures() []types.Prefecture {
	doc, err := goquery.NewDocument("https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/prefectures.csv")
	if err != nil {
		panic(err)
	}
	prefectures := []types.Prefecture{}
	selection := doc.Find("tbody")
	innerSelection := selection.Find("tr")
	innerSelection.Each(func(i int, s *goquery.Selection) {
		prefecture := types.NewPrefecture()
		s.Find("td").Each(func(k int, s2 *goquery.Selection) {
			switch k {
			case 1:
				prefecture.ID = s2.Text()
			case 2:
				prefecture.NameJa = s2.Text()
			case 3:
				prefecture.NameEn = s2.Text()
			case 4:
				prefecture.Regions = s2.Text()
			case 5:
				prefecture.Latitude = s2.Text()
			case 6:
				prefecture.Longitude = s2.Text()
			}
		})
		prefectures = append(prefectures, prefecture)
	})

	for _, prefecture := range prefectures {
		fmt.Println("prefecture %v\n", prefecture)
	}
	return prefectures
}

func fetchTotalData() *types.TotalData {
	// file取得
	fmt.Printf("get total data\n")

	// 感染者一覧
	url := "https://covid19-japan-web-api.now.sh/api/v1/total"

	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Error2 occur...")
		return nil
		//log.Printf(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error3 occur...")
		return nil
		//log.Printf(err)
	}

	//fmt.Printf("data is %v\n", body)
	var totalData *types.TotalData
	err = json.Unmarshal(body, &totalData)
	if err != nil {
		log.Printf("Error4 occur...")
		return nil
		//log.Printf(err)
	}

	fmt.Printf("totalData is %v\n", totalData)
	handler.APIData.TotalData = totalData
	return totalData

}

// 県ごとの感染者数、死者数
func fetchPrefsData() []*types.PrefData {
	// file取得
	fmt.Printf("get prefs data\n")

	// 感染者一覧
	url := "https://covid19-japan-web-api.now.sh/api/v1/prefectures"

	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Error2 occur...")
		return nil
		//log.Printf(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error3 occur...")
		return nil
		//log.Printf(err)
	}

	fmt.Printf("data is %v\n", body)
	var prefs []*types.PrefData
	err = json.Unmarshal(body, &prefs)
	if err != nil {
		log.Printf("Error4 occur...")
		return nil
	}
	for _, pref := range prefs {
		fmt.Printf("pref is %v\n", pref)
	}
	fmt.Printf("prefNum is %v\n", len(prefs))

	handler.APIData.PrefsData = prefs
	return prefs
}

func main() {
	fmt.Printf("Starting...")

	go fetchDataSceduler()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/api/v1/stats", handler.SendStats())             // 統計データ
	e.GET("/api/v1/patients", handler.SendPatients())       // 統計データ
	e.GET("/api/v1/prefectures", handler.SendPrefectures()) // 統計データ
	//e.GET("/api/v1/stats/daily", handler.SendDailyData())                            // 日ごとの統計データ
	e.GET("/api/v1/stats/prefectures", handler.SendPrefsData()) // 都道府県ごとの統計データ
	//e.GET("/api/v1/stats/daily-prefectures", handler.SendDailyPrefData())                // 日時、都道府県ごとの統計データ
	//e.GET("/api/v1/stats/prefectures-daily", handler.SendPrefsDailyData())                // 日時、都道府県ごとの統計データ
	e.GET("/api/v1/stats/total", handler.SendTotalData())       // 日本全体としての統計データ
	e.GET("/api/v1/patients/total", handler.SendPatientsData()) // 患者一人一人のデータ
	//e.GET("/api/v1/patients/daily", handler.SendDailyPatientsData())                 // 日付ごとの患者データ
	//e.GET("/api/v1/patients/prefecture", handler.SendPrefPatientsData())             // 都道府県ごとの患者データ
	//e.GET("/api/v1/patients/daily-prefectures", handler.SendDailyPrefPatientsData()) // 日付,都道府県ごとの患者データ
	//e.GET("/api/v1/patients/prefectures-daily", handler.SendPrefsDailyPatientsData()) // 日付,都道府県ごとの患者データ

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
		log.Printf("Defaulting to port %s", port)
	}
	e.Start(":" + port)

}
