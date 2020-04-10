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

	// dailyPositiveByPref data
	dailyPositiveByPref := fetchDailyPositiveByPref()
	handler.APIData.DailyPositiveByPref = dailyPositiveByPref

	// dailyDeathByPref data
	dailyDeathByPref := fetchDailyDeathByPref()
	handler.APIData.DailyDeathByPref = dailyDeathByPref

	// dailyCallcenter data
	dailyCallcenter := fetchDailyCallcenter()
	handler.APIData.DailyCallcenter = dailyCallcenter

	// dailyShip data
	dailyShip := fetchDailyShip()
	handler.APIData.DailyShip = dailyShip

	// news data
	news := fetchNews()
	handler.APIData.News = news

	// detailByRegion data
	detailByRegion := fetchDetailByRegion()
	handler.APIData.DetailByRegion = detailByRegion

	// stats data
	stats := calcStats(patients)
	handler.APIData.Stats = stats

	// dailySexByPref data
	//dailySexByPref := calcDailySexByPref(patients)
	//handler.APIData.DailySexByPref = dailySexByPref

	// dailyAgeByPref data
	//dailyAgeByPref := calcDailyAgeByPref(patients)
	//handler.APIData.DailyAgeByPref = dailyAgeByPref
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
	doc, err := goquery.NewDocument("https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/detailByRegion.csv")
	if err != nil {
		panic(err)
	}
	rawStats := []types.Stat{}
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
		rawStats = append(rawStats, stat)
	})
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

/*func calcDailySexByPref(patients []types.Patient) []types.Stat {
	result := []types.DateSexByPref{}
	//dfStats := dataframe.LoadStructs(rawStats)
	dfPatients := dataframe.LoadStructs(patients)

	for i, stat := range rawStats {
		dateSexByPref := types.DateSexByPref{}
		date := fmt.Sprintf("%s/%s/%s", string([]rune(stat.Date)[:4]), string([]rune(stat.Date)[4:6]), string([]rune(stat.Date)[6:8]))
		pref := stat.Prefecture
		// その日のpatientsデータ
		dfTgt := dfPatients.Filter(
			makeF("Prefecture", series.Eq, pref),
		).Filter(
			makeF("Date", series.Eq, date),
		)

		// その日の発生件数
		cases := dfTgt.Nrow()

		// 性別
		maleNum := dfTgt.Filter(makeF("Sex", series.Eq, "男性")).Nrow()
		femaleNum := dfTgt.Filter(makeF("Sex", series.Eq, "女性")).Nrow()
		unknownNum := cases - femaleNum - maleNum

	}

	result = rawStats
	return result
}*/

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

func fetchDetailByRegion() []types.RegionDetail {
	doc, err := goquery.NewDocument("https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/detailByRegion.csv")
	if err != nil {
		panic(err)
	}
	regionDetails := []types.RegionDetail{}
	selection := doc.Find("tbody")
	innerSelection := selection.Find("tr")
	innerSelection.Each(func(i int, s *goquery.Selection) {
		regionDetail := types.NewRegionDetail()
		s.Find("td").Each(func(k int, s2 *goquery.Selection) {
			switch k {
			case 1:
				regionDetail.Date = s2.Text()
			case 2:
				regionDetail.Prefecture = s2.Text()
			case 3:
				cases, _ := strconv.Atoi(s2.Text())
				regionDetail.Cases = cases
			case 4:
				hospitals, _ := strconv.Atoi(s2.Text())
				regionDetail.Hospitals = hospitals
			case 5:
				discharges, _ := strconv.Atoi(s2.Text())
				regionDetail.Discharges = discharges
			case 6:
				deaths, _ := strconv.Atoi(s2.Text())
				regionDetail.Deaths = deaths
			}
		})
		regionDetails = append(regionDetails, regionDetail)
	})

	return regionDetails
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

func fetchDailyPositiveByPref() []types.DatePositiveByPref {
	doc, err := goquery.NewDocument("https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/byDate.csv")
	if err != nil {
		panic(err)
	}
	dailyPositiveByPref := []types.DatePositiveByPref{}
	selection := doc.Find("tbody")
	innerSelection := selection.Find("tr")
	innerSelection.Each(func(i int, s *goquery.Selection) {
		datePositiveByPref := types.NewDatePositiveByPref()
		s.Find("td").Each(func(k int, s2 *goquery.Selection) {
			switch k {
			case 1:
				datePositiveByPref.Date = s2.Text()
			case 2:
				datePositiveByPref.Hokkaido = s2.Text()
			case 3:
				datePositiveByPref.Aomori = s2.Text()
			case 4:
				datePositiveByPref.Iwate = s2.Text()
			case 5:
				datePositiveByPref.Miyagi = s2.Text()
			case 6:
				datePositiveByPref.Akita = s2.Text()
			case 7:
				datePositiveByPref.Yamagata = s2.Text()
			case 8:
				datePositiveByPref.Fukushima = s2.Text()
			case 9:
				datePositiveByPref.Ibaraki = s2.Text()
			case 10:
				datePositiveByPref.Tochigi = s2.Text()
			case 11:
				datePositiveByPref.Gunma = s2.Text()
			case 12:
				datePositiveByPref.Saitama = s2.Text()
			case 13:
				datePositiveByPref.Chiba = s2.Text()
			case 14:
				datePositiveByPref.Tokyo = s2.Text()
			case 15:
				datePositiveByPref.Kanagawa = s2.Text()
			case 16:
				datePositiveByPref.Niigata = s2.Text()
			case 17:
				datePositiveByPref.Toyama = s2.Text()
			case 18:
				datePositiveByPref.Ishikawa = s2.Text()
			case 19:
				datePositiveByPref.Fukui = s2.Text()
			case 20:
				datePositiveByPref.Yamanashi = s2.Text()
			case 21:
				datePositiveByPref.Nagano = s2.Text()
			case 22:
				datePositiveByPref.Gifu = s2.Text()
			case 23:
				datePositiveByPref.Shizuoka = s2.Text()
			case 24:
				datePositiveByPref.Aichi = s2.Text()
			case 25:
				datePositiveByPref.Mie = s2.Text()
			case 26:
				datePositiveByPref.Shiga = s2.Text()
			case 27:
				datePositiveByPref.Kyoto = s2.Text()
			case 28:
				datePositiveByPref.Osaka = s2.Text()
			case 29:
				datePositiveByPref.Hyogo = s2.Text()
			case 30:
				datePositiveByPref.Nara = s2.Text()
			case 31:
				datePositiveByPref.Wakayama = s2.Text()
			case 32:
				datePositiveByPref.Tottori = s2.Text()
			case 33:
				datePositiveByPref.Shimane = s2.Text()
			case 34:
				datePositiveByPref.Okayama = s2.Text()
			case 35:
				datePositiveByPref.Hiroshima = s2.Text()
			case 36:
				datePositiveByPref.Yamaguchi = s2.Text()
			case 37:
				datePositiveByPref.Tokushima = s2.Text()
			case 38:
				datePositiveByPref.Kagawa = s2.Text()
			case 39:
				datePositiveByPref.Ehime = s2.Text()
			case 40:
				datePositiveByPref.Kochi = s2.Text()
			case 41:
				datePositiveByPref.Fukuoka = s2.Text()
			case 42:
				datePositiveByPref.Saga = s2.Text()
			case 43:
				datePositiveByPref.Nagasaki = s2.Text()
			case 44:
				datePositiveByPref.Kumamoto = s2.Text()
			case 45:
				datePositiveByPref.Oita = s2.Text()
			case 46:
				datePositiveByPref.Miyazaki = s2.Text()
			case 47:
				datePositiveByPref.Kagoshima = s2.Text()
			case 48:
				datePositiveByPref.Okinawa = s2.Text()
			case 49:
				datePositiveByPref.Charter = s2.Text()
			case 50:
				datePositiveByPref.QuarantineOfficer = s2.Text()
			case 51:
				datePositiveByPref.Cruise = s2.Text()
			}
		})
		dailyPositiveByPref = append(dailyPositiveByPref, datePositiveByPref)
	})
	return dailyPositiveByPref
}

func fetchDailyDeathByPref() []types.DateDeathByPref {
	doc, err := goquery.NewDocument("https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/death.csv")
	if err != nil {
		panic(err)
	}
	dailyDeathByPref := []types.DateDeathByPref{}
	selection := doc.Find("tbody")
	innerSelection := selection.Find("tr")
	innerSelection.Each(func(i int, s *goquery.Selection) {
		dateDeathByPref := types.NewDateDeathByPref()
		s.Find("td").Each(func(k int, s2 *goquery.Selection) {
			switch k {
			case 1:
				dateDeathByPref.Date = s2.Text()
			case 2:
				dateDeathByPref.Hokkaido = s2.Text()
			case 3:
				dateDeathByPref.Aomori = s2.Text()
			case 4:
				dateDeathByPref.Iwate = s2.Text()
			case 5:
				dateDeathByPref.Miyagi = s2.Text()
			case 6:
				dateDeathByPref.Akita = s2.Text()
			case 7:
				dateDeathByPref.Yamagata = s2.Text()
			case 8:
				dateDeathByPref.Fukushima = s2.Text()
			case 9:
				dateDeathByPref.Ibaraki = s2.Text()
			case 10:
				dateDeathByPref.Tochigi = s2.Text()
			case 11:
				dateDeathByPref.Gunma = s2.Text()
			case 12:
				dateDeathByPref.Saitama = s2.Text()
			case 13:
				dateDeathByPref.Chiba = s2.Text()
			case 14:
				dateDeathByPref.Tokyo = s2.Text()
			case 15:
				dateDeathByPref.Kanagawa = s2.Text()
			case 16:
				dateDeathByPref.Niigata = s2.Text()
			case 17:
				dateDeathByPref.Toyama = s2.Text()
			case 18:
				dateDeathByPref.Ishikawa = s2.Text()
			case 19:
				dateDeathByPref.Fukui = s2.Text()
			case 20:
				dateDeathByPref.Yamanashi = s2.Text()
			case 21:
				dateDeathByPref.Nagano = s2.Text()
			case 22:
				dateDeathByPref.Gifu = s2.Text()
			case 23:
				dateDeathByPref.Shizuoka = s2.Text()
			case 24:
				dateDeathByPref.Aichi = s2.Text()
			case 25:
				dateDeathByPref.Mie = s2.Text()
			case 26:
				dateDeathByPref.Shiga = s2.Text()
			case 27:
				dateDeathByPref.Kyoto = s2.Text()
			case 28:
				dateDeathByPref.Osaka = s2.Text()
			case 29:
				dateDeathByPref.Hyogo = s2.Text()
			case 30:
				dateDeathByPref.Nara = s2.Text()
			case 31:
				dateDeathByPref.Wakayama = s2.Text()
			case 32:
				dateDeathByPref.Tottori = s2.Text()
			case 33:
				dateDeathByPref.Shimane = s2.Text()
			case 34:
				dateDeathByPref.Okayama = s2.Text()
			case 35:
				dateDeathByPref.Hiroshima = s2.Text()
			case 36:
				dateDeathByPref.Yamaguchi = s2.Text()
			case 37:
				dateDeathByPref.Tokushima = s2.Text()
			case 38:
				dateDeathByPref.Kagawa = s2.Text()
			case 39:
				dateDeathByPref.Ehime = s2.Text()
			case 40:
				dateDeathByPref.Kochi = s2.Text()
			case 41:
				dateDeathByPref.Fukuoka = s2.Text()
			case 42:
				dateDeathByPref.Saga = s2.Text()
			case 43:
				dateDeathByPref.Nagasaki = s2.Text()
			case 44:
				dateDeathByPref.Kumamoto = s2.Text()
			case 45:
				dateDeathByPref.Oita = s2.Text()
			case 46:
				dateDeathByPref.Miyazaki = s2.Text()
			case 47:
				dateDeathByPref.Kagoshima = s2.Text()
			case 48:
				dateDeathByPref.Okinawa = s2.Text()
			case 49:
				dateDeathByPref.Charter = s2.Text()
			case 50:
				dateDeathByPref.QuarantineOfficer = s2.Text()
			case 51:
				dateDeathByPref.Cruise = s2.Text()
			}
		})
		dailyDeathByPref = append(dailyDeathByPref, dateDeathByPref)
	})
	return dailyDeathByPref
}

func fetchDailyCallcenter() []types.DateCallcenter {
	doc, err := goquery.NewDocument("https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/callCenter.csv")
	if err != nil {
		panic(err)
	}
	dailyCallcenter := []types.DateCallcenter{}
	selection := doc.Find("tbody")
	innerSelection := selection.Find("tr")
	innerSelection.Each(func(i int, s *goquery.Selection) {
		dateCallcenter := types.NewDateCallcenter()
		s.Find("td").Each(func(k int, s2 *goquery.Selection) {
			switch k {
			case 1:
				dateCallcenter.Date = s2.Text()
			case 2:
				dateCallcenter.Call = s2.Text()
			case 3:
				dateCallcenter.Fax = s2.Text()
			case 4:
				dateCallcenter.Mail = s2.Text()
			case 5:
				dateCallcenter.Line = s2.Text()
			}
		})
		dailyCallcenter = append(dailyCallcenter, dateCallcenter)
	})
	return dailyCallcenter
}

func fetchDailyShip() []types.DateShip {
	doc, err := goquery.NewDocument("https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/shipDailyReport.csv")
	if err != nil {
		panic(err)
	}
	dailyShip := []types.DateShip{}
	selection := doc.Find("tbody")
	innerSelection := selection.Find("tr")
	innerSelection.Each(func(i int, s *goquery.Selection) {
		dateShip := types.NewDateShip()
		s.Find("td").Each(func(k int, s2 *goquery.Selection) {
			switch k {
			case 1:
				dateShip.Date = s2.Text()
			case 2:
				dateShip.Pcr = s2.Text()
			case 3:
				dateShip.Positive = s2.Text()
			case 4:
				dateShip.Discharge = s2.Text()
			case 5:
				dateShip.SymotomlessDischarge = s2.Text()
			case 6:
				dateShip.SymotomDischarge = s2.Text()
			case 7:
				dateShip.Severe = s2.Text()
			case 8:
				dateShip.Death = s2.Text()
			}
		})
		dailyShip = append(dailyShip, dateShip)
	})
	return dailyShip
}

func fetchNews() []types.News {
	doc, err := goquery.NewDocument("https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/mhlw_houdou.csv")
	if err != nil {
		panic(err)
	}
	news := []types.News{}
	selection := doc.Find("tbody")
	innerSelection := selection.Find("tr")
	innerSelection.Each(func(i int, s *goquery.Selection) {
		ns := types.NewNews()
		s.Find("td").Each(func(k int, s2 *goquery.Selection) {
			switch k {
			case 1:
				ns.ID = s2.Text()
			case 2:
				ns.Date = s2.Text()
			case 3:
				ns.Title = s2.Text()
			case 4:
				ns.Link = s2.Text()
			}
		})
		news = append(news, ns)
	})
	return news
}

func main() {
	fmt.Printf("Starting...")

	go fetchDataSceduler()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/api/v1/stats", handler.SendStats())                                // 統計データ
	e.GET("/api/v1/patients", handler.SendPatients())                          // 罹患者データ
	e.GET("/api/v1/prefectures", handler.SendPrefectures())                    // 都道府県データ
	e.GET("/api/v1/detail-by-region", handler.SendDetailByRegion())            // 地域ごとの詳細データ
	e.GET("/api/v1/daily/report", handler.SendDailyReport())                   // 日ごとの統計データ
	e.GET("/api/v1/daily/positive-by-pref", handler.SendDailyPositiveByPref()) // 日ごと、都道府県ごとの感染者データ
	e.GET("/api/v1/daily/death-by-pref", handler.SendDailyDeathByPref())       // 日ごと、都道府県ごとの死亡者データ
	e.GET("/api/v1/daily/callcenter", handler.SendDailyCallcenter())           // 日ごと、コールセンターの感染者データ
	e.GET("/api/v1/daily/ship", handler.SendDailyShip())                       // 日ごと、クルーズ船の感染者データ
	//e.GET("/api/v1/daily/sex-by-pref", handler.SendDailySexByPref())           // 日ごと、クルーズ船の感染者データ
	//e.GET("/api/v1/daily/age-by-pref", handler.SendDailyAgeByPref())           // 日ごと、クルーズ船の感染者データ
	e.GET("/api/v1/news", handler.SendNews()) // 日ごと、都道府県ごとの感染者データ

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
		log.Printf("Defaulting to port %s", port)
	}
	e.Start(":" + port)

}
