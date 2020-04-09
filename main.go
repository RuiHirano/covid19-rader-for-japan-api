package main

import (
	//"context"

	//"bytes"

	//"encoding/json"
	"fmt"
	"strconv"

	//"io/ioutil"

	"handler"
	"log"
	"os"

	"types"

	"github.com/PuerkitoBio/goquery"

	//"net/http"

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

	// patients data
	patients := fetchPatients()
	handler.APIData.Patients = patients

	// prefectures data
	prefectures := fetchPrefectures()
	handler.APIData.Prefectures = prefectures

	// dailyReport data
	dailyReport := fetchDailyReport()
	handler.APIData.DailyReport = dailyReport

	// stats data
	stats := calcStats(patients)
	handler.APIData.Stats = stats
}

func makeF(colname string, comparator series.Comparator, comparando interface{}) dataframe.F {
	return dataframe.F{
		Colname:    colname,
		Comparator: comparator,
		Comparando: comparando,
	}
}

func calcStats(patients []types.Patient) []types.Stat {
	result := []types.Stat{}
	rawStats := fetchDetailByRegion()
	//dfStats := dataframe.LoadStructs(rawStats)
	dfPatients := dataframe.LoadStructs(patients)

	//fmt.Printf("dfStats: ", dfStats)
	//fmt.Printf("dfPatients: ", dfPatients)
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
			makeF("Prefecture", series.Eq, pref),
		).Filter(
			makeF("Date", series.Eq, date),
		)
		// その日以前全てのpatientsデータ
		dfLtTgt := dfPatients.Filter(
			makeF("Prefecture", series.Eq, pref),
		).Filter(
			makeF("Date", "<=", date),
		)

		// その日以前全てのstatsデータ
		/*dfLtStats := dfStats.Filter(
			makeF("Prefecture", series.Eq, pref),
		).Filter(
			makeF("Date", "<=", stat.Date),
		)*/

		// その日の発生件数と累計発生件数
		cases := dfTgt.Nrow()
		totalCases := dfLtTgt.Nrow()

		// 患者数等
		deaths := 0
		discharges := 0
		hospitals := 0
		prevStat := types.NewStat()
		for _, st := range rawStats {
			date, _ := strconv.Atoi(stat.Date)
			if st.Prefecture == pref && st.Date == strconv.Itoa(date-1) {
				prevStat = st
			}
		}
		deaths = stat.TotalDeaths - prevStat.TotalDeaths
		discharges = stat.TotalDischarges - prevStat.TotalDischarges
		hospitals = stat.TotalHospitals - prevStat.TotalHospitals

		// 性別
		maleNum := dfTgt.Filter(makeF("Sex", series.Eq, "男性")).Nrow()
		femaleNum := dfTgt.Filter(makeF("Sex", series.Eq, "女性")).Nrow()
		unknownNum := cases - femaleNum - maleNum
		totalMaleNum := dfLtTgt.Filter(makeF("Sex", series.Eq, "男性")).Nrow()
		totalFemaleNum := dfLtTgt.Filter(makeF("Sex", series.Eq, "女性")).Nrow()
		totalUnknownNum := cases - totalFemaleNum - totalMaleNum

		// 年齢
		age10 := dfTgt.Filter(makeF("Age", series.Eq, "10代")).Nrow()
		age20 := dfTgt.Filter(makeF("Age", series.Eq, "20代")).Nrow()
		age30 := dfTgt.Filter(makeF("Age", series.Eq, "30代")).Nrow()
		age40 := dfTgt.Filter(makeF("Age", series.Eq, "40代")).Nrow()
		age50 := dfTgt.Filter(makeF("Age", series.Eq, "50代")).Nrow()
		age60 := dfTgt.Filter(makeF("Age", series.Eq, "60代")).Nrow()
		age70 := dfTgt.Filter(makeF("Age", series.Eq, "70代")).Nrow()
		age80 := dfTgt.Filter(makeF("Age", series.Eq, "80代")).Nrow()
		age90 := dfTgt.Filter(makeF("Age", series.Eq, "90代")).Nrow()
		ageUnknown := cases - age10 - age20 - age30 - age40 - age50 - age60 - age70 - age80 - age90

		totalAge10 := dfLtTgt.Filter(makeF("Age", series.Eq, "10代")).Nrow()
		totalAge20 := dfLtTgt.Filter(makeF("Age", series.Eq, "20代")).Nrow()
		totalAge30 := dfLtTgt.Filter(makeF("Age", series.Eq, "30代")).Nrow()
		totalAge40 := dfLtTgt.Filter(makeF("Age", series.Eq, "40代")).Nrow()
		totalAge50 := dfLtTgt.Filter(makeF("Age", series.Eq, "50代")).Nrow()
		totalAge60 := dfLtTgt.Filter(makeF("Age", series.Eq, "60代")).Nrow()
		totalAge70 := dfLtTgt.Filter(makeF("Age", series.Eq, "70代")).Nrow()
		totalAge80 := dfLtTgt.Filter(makeF("Age", series.Eq, "80代")).Nrow()
		totalAge90 := dfLtTgt.Filter(makeF("Age", series.Eq, "90代")).Nrow()
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
		//fmt.Printf("gender %v\n", patient.Sex, patient, patient.Age)
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
	//fmt.Println(" doc", doc)
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
	return prefectures
}

func fetchDailyReport() []types.DateReport {
	doc, err := goquery.NewDocument("https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/resultDailyReport.csv")
	if err != nil {
		panic(err)
	}
	dailyReport := []types.DateReport{}
	selection := doc.Find("tbody")
	innerSelection := selection.Find("tr")
	innerSelection.Each(func(i int, s *goquery.Selection) {
		dateReport := types.NewDateReport()
		s.Find("td").Each(func(k int, s2 *goquery.Selection) {
			switch k {
			case 1:
				dateReport.Date = s2.Text()
			case 2:
				dateReport.PcrD = s2.Text()
			case 3:
				dateReport.PositiveD = s2.Text()
			case 4:
				dateReport.SymptomD = s2.Text()
			case 5:
				dateReport.SymptomlessD = s2.Text()
			case 6:
				dateReport.SymptomConfirmingD = s2.Text()
			case 7:
				dateReport.HospitalizeD = s2.Text()
			case 8:
				dateReport.MildD = s2.Text()
			case 9:
				dateReport.SevereD = s2.Text()
			case 10:
				dateReport.ConfirmingD = s2.Text()
			case 11:
				dateReport.WaitingD = s2.Text()
			case 12:
				dateReport.DischargeD = s2.Text()
			case 13:
				dateReport.DeathD = s2.Text()
			case 14:
				dateReport.PcrF = s2.Text()
			case 15:
				dateReport.PositiveF = s2.Text()
			case 16:
				dateReport.SymptomF = s2.Text()
			case 17:
				dateReport.SymptomlessF = s2.Text()
			case 18:
				dateReport.SymptomConfirmingF = s2.Text()
			case 19:
				dateReport.HospitalizeF = s2.Text()
			case 20:
				dateReport.MildF = s2.Text()
			case 21:
				dateReport.SevereF = s2.Text()
			case 22:
				dateReport.ConfirmingF = s2.Text()
			case 23:
				dateReport.WaitingF = s2.Text()
			case 24:
				dateReport.DischargeF = s2.Text()
			case 25:
				dateReport.DeathF = s2.Text()
			case 26:
				dateReport.PcrX = s2.Text()
			case 27:
				dateReport.PositiveX = s2.Text()
			case 28:
				dateReport.Symptom = s2.Text()
			case 29:
				dateReport.Symptomless = s2.Text()
			case 30:
				dateReport.SymptomConfirming = s2.Text()
			case 31:
				dateReport.Hospitalized = s2.Text()
			case 32:
				dateReport.Mild = s2.Text()
			case 33:
				dateReport.SevereX = s2.Text()
			case 34:
				dateReport.Confirming = s2.Text()
			case 35:
				dateReport.Waiting = s2.Text()
			case 36:
				dateReport.DischargeX = s2.Text()
			case 37:
				dateReport.DeathX = s2.Text()
			case 38:
				dateReport.PcrY = s2.Text()
			case 39:
				dateReport.PositiveY = s2.Text()
			case 40:
				dateReport.DischargeY = s2.Text()
			case 41:
				dateReport.SymptomlessDischarge = s2.Text()
			case 42:
				dateReport.SymptomDischarge = s2.Text()
			case 43:
				dateReport.SevereY = s2.Text()
			case 44:
				dateReport.DeathY = s2.Text()
			case 45:
				dateReport.Pcr = s2.Text()
			case 46:
				dateReport.Discharge = s2.Text()
			case 47:
				dateReport.PcrDiff = s2.Text()
			case 48:
				dateReport.DischargeDiff = s2.Text()
			}
		})
		dailyReport = append(dailyReport, dateReport)
	})
	return dailyReport
}

func main() {
	fmt.Printf("Starting...")

	go fetchDataSceduler()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/api/v1/stats", handler.SendStats())              // 統計データ
	e.GET("/api/v1/patients", handler.SendPatients())        // 統計データ
	e.GET("/api/v1/prefectures", handler.SendPrefectures())  // 統計データ
	e.GET("/api/v1/daily-report", handler.SendDailyReport()) // 日ごとの統計データ

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
		log.Printf("Defaulting to port %s", port)
	}
	e.Start(":" + port)

}
