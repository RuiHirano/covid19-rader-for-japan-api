package types

type Sex uint64

const (
	SexMALE   Sex = 0
	SexFEMALE Sex = 1
)

type Prefecture struct {
	ID        string `json:"id"`
	NameJa    string `json:"name_ja"`
	NameEn    string `json:"name_en"`
	Regions   string `json:"regions"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

func NewPrefecture() Prefecture {
	p := Prefecture{
		ID:        "",
		NameJa:    "",
		NameEn:    "",
		Regions:   "",
		Longitude: "",
		Latitude:  "",
	}
	return p
}

type Patient struct {
	ID                  string `json:"id"`
	Date                string `json:"date"`
	Prefecture          string `json:"prefecture"`
	Residence           string `json:"residence"`
	Age                 string `json:"age"`
	Sex                 string `json:"sex"`
	Attribute           string `json:"attribute"`
	PrefectureNumber    string `json:"prefecture_number"`
	TravelOrContact     string `json:"travel_or_contact"`
	Detail              string `json:"detail"`
	Src                 string `json:"src"`
	Onset               string `json:"onset"`
	Symptom             string `json:"symptom"`
	DeathOrDischageDate string `json:"death_or_discharge_date"`
	Comment1            string `json:"comment1"`
	Comment2            string `json:"comment2"`
	Outcome             string `json:"outcome"`
	OutcomeSrc          string `json:"outcome_src"`
}

func NewPatient() Patient {
	p := Patient{
		ID:                  "",
		Date:                "",
		Prefecture:          "",
		Residence:           "",
		Age:                 "",
		Sex:                 "",
		Attribute:           "",
		PrefectureNumber:    "",
		TravelOrContact:     "",
		Detail:              "",
		Src:                 "",
		Onset:               "",
		Symptom:             "",
		DeathOrDischageDate: "",
		Comment1:            "",
		Comment2:            "",
		Outcome:             "",
		OutcomeSrc:          "",
	}
	return p
}

type Stat struct {
	Date            string  `json:"date"`
	Prefecture      string  `json:"prefecture"`
	Cases           int     `json:"cases"`
	TotalCases      int     `json:"total_cases"` // 累計
	Hospitals       int     `json:"hospital"`
	TotalHospitals  int     `json:"total_hospitals"` // 累計
	Discharges      int     `json:"discharge"`
	TotalDischarges int     `json:"total_discharges"` // 累計
	Deaths          int     `json:"deaths"`
	TotalDeaths     int     `json:"total_deaths"` // 累計
	SexData         SexData `json:"sex_data"`
	TotalSexData    SexData `json:"total_sex_data"` // 累計
	AgeData         AgeData `json:"age_data"`
	TotalAgeData    AgeData `json:"total_age_data"` // 累計
}

func NewStat() Stat {
	s := Stat{
		Date:            "",
		Prefecture:      "",
		Cases:           0,
		TotalCases:      0,
		Hospitals:       0,
		TotalHospitals:  0,
		Discharges:      0,
		TotalDischarges: 0,
		Deaths:          0,
		TotalDeaths:     0,
		SexData:         NewSexData(),
		TotalSexData:    NewSexData(),
		AgeData:         NewAgeData(),
		TotalAgeData:    NewAgeData(),
	}
	return s
}

type SexData struct {
	Female  int `json:"female"`
	Male    int `json:"male"`
	Unknown int `json:"unknown"`
}

func NewSexData() SexData {
	s := SexData{
		Female:  0,
		Male:    0,
		Unknown: 0,
	}
	return s
}

type AgeData struct {
	Age10   int `json:"age_10"`
	Age20   int `json:"age_20"`
	Age30   int `json:"age_30"`
	Age40   int `json:"age_40"`
	Age50   int `json:"age_50"`
	Age60   int `json:"age_60"`
	Age70   int `json:"age_70"`
	Age80   int `json:"age_80"`
	Age90   int `json:"age_90"`
	Unknown int `json:"age_unknown"`
}

func NewAgeData() AgeData {
	a := AgeData{
		Age10:   0,
		Age20:   0,
		Age30:   0,
		Age40:   0,
		Age50:   0,
		Age60:   0,
		Age70:   0,
		Age80:   0,
		Age90:   0,
		Unknown: 0,
	}
	return a
}

type DateSummary struct {
	Date       string `csv:"日付"`
	Prefecture string `csv:"都道府県名"`
	Positives  string `csv:"陽性者"`
	Pcrs       string `csv:"検査人数"`
	Hospitals  string `csv:"入院中"`
	Severes    string `csv:"重症者"`
	Discharges string `csv:"退院者"`
	Deaths     string `csv:"死亡者"`
	Checking   string `csv:"確認中"`
	Class      string `csv:"分類"`
}

type DateWorldSummary struct {
	Date            string `csv:"date"`
	Update          string `csv:"update"`
	Cases           string `csv:"cases"`
	NewCases        string `csv:"new_cases"`
	Deaths          string `csv:"deaths"`
	Country         string `csv:"country"`
	LastUpdate      string `csv:"last_update"`
	Population      string `csv:"population"`
	CasesPer100k    string `csv:"casesPer100k"`
	NewTests        string `csv:"new_tests"`
	TestsCumulative string `csv:"tests_cumulative"`
	TestsPer100k    string `csv:"testsPer100k"`
	PositiveRate    string `csv:"positiveRate"`
}

type DateReport struct {
	Date                 string `json: "date"`
	PcrD                 string `json: "pcr_d"`
	PositiveD            string `json: "positive_d"`
	SymptomD             string `json: "symptom_d"`
	SymptomlessD         string `json: "symptomless_d"`
	SymptomConfirmingD   string `json: "symptom_confirming_d"`
	HospitalizeD         string `json: "hospitalize_d"`
	MildD                string `json: "mild_d"`
	SevereD              string `json: "severe_d"`
	ConfirmingD          string `json: "confirming_d"`
	WaitingD             string `json: "waiting_d"`
	DischargeD           string `json: "discharge_d"`
	DeathD               string `json: "death_d"`
	PcrF                 string `json: "pcr_f"`
	PositiveF            string `json: "positive_f"`
	SymptomF             string `json: "symptom_f"`
	SymptomlessF         string `json: "symptomless_f"`
	SymptomConfirmingF   string `json: "symptom_confirming_f"`
	HospitalizeF         string `json: "hospitalize_f"`
	MildF                string `json: "mild_f"`
	SevereF              string `json: "severe_f"`
	ConfirmingF          string `json: "confirming_f"`
	WaitingF             string `json: "waiting_f"`
	DischargeF           string `json: "discharge_f"`
	DeathF               string `json: "death_f"`
	PcrX                 string `json: "pcr_x"`
	PositiveX            string `json: "positive_x"`
	Symptom              string `json: "symptom"`
	Symptomless          string `json: "symptomless"`
	SymptomConfirming    string `json: "symptom_confirming"`
	Hospitalized         string `json: "hospitalized"`
	Mild                 string `json: "mild"`
	SevereX              string `json: "severe_x"`
	Confirming           string `json: "confirming"`
	Waiting              string `json: "waiting"`
	DischargeX           string `json: "discharge_x"`
	DeathX               string `json: "death_x"`
	PcrY                 string `json: "pcr_y"`
	PositiveY            string `json: "positive_y"`
	DischargeY           string `json: "discharge_y"`
	SymptomlessDischarge string `json: "symptomless_discharge"`
	SymptomDischarge     string `json: "symptomDischarge"`
	SevereY              string `json: "severe_y"`
	DeathY               string `json: "death_y"`
	Pcr                  string `json: "pcr"`
	Discharge            string `json: "discharge"`
	PcrDiff              string `json: "pcr_diff"`
	DischargeDiff        string `json: "discharge_diff"`
}

func NewDateReport() DateReport {
	s := DateReport{
		Date:                 "",
		PcrD:                 "0",
		PositiveD:            "0",
		SymptomD:             "0",
		SymptomlessD:         "0",
		SymptomConfirmingD:   "0",
		HospitalizeD:         "0",
		MildD:                "0",
		SevereD:              "0",
		ConfirmingD:          "0",
		WaitingD:             "0",
		DischargeD:           "0",
		DeathD:               "0",
		PcrF:                 "0",
		PositiveF:            "0",
		SymptomF:             "0",
		SymptomlessF:         "0",
		SymptomConfirmingF:   "0",
		HospitalizeF:         "0",
		MildF:                "0",
		SevereF:              "0",
		ConfirmingF:          "0",
		WaitingF:             "0",
		DischargeF:           "0",
		DeathF:               "0",
		PcrX:                 "0",
		PositiveX:            "0",
		Symptom:              "0",
		Symptomless:          "0",
		SymptomConfirming:    "0",
		Hospitalized:         "0",
		Mild:                 "0",
		SevereX:              "0",
		Confirming:           "0",
		Waiting:              "0",
		DischargeX:           "0",
		DeathX:               "0",
		PcrY:                 "0",
		PositiveY:            "0",
		DischargeY:           "0",
		SymptomlessDischarge: "0",
		SymptomDischarge:     "0",
		SevereY:              "0",
		DeathY:               "0",
		Pcr:                  "0",
		Discharge:            "0",
		PcrDiff:              "0",
		DischargeDiff:        "0",
	}
	return s
}

type News struct {
	ID         string `json:"id" csv:"id"`
	Date       string `json:"date" csv:"date"`
	Title      string `json:"title" csv:"title"`
	Link       string `json:"link" csv:"link"`
	Prefecture string `json:"pre" csv:"pre"`
	Resident   string `json:"resident" csv:"resident"`
}

func NewNews() News {
	s := News{
		ID:    "",
		Date:  "",
		Title: "",
		Link:  "",
	}
	return s
}

type DateShip struct {
	Date                 string `json:"date"`
	Pcr                  string `json:"pcr"`
	Positive             string `json:"positive"`
	Discharge            string `json:"discharge"` // 累計
	SymotomlessDischarge string `json:"symotomless_discharge"`
	SymotomDischarge     string `json:"symotom_discharge"` // 累計
	Severe               string `json:"severe"`
	Death                string `json:"death"` // 累計
}

func NewDateShip() DateShip {
	s := DateShip{
		Date:                 "",
		Pcr:                  "",
		Positive:             "",
		Discharge:            "",
		SymotomlessDischarge: "",
		SymotomDischarge:     "",
		Severe:               "",
		Death:                "",
	}
	return s
}

type DateSexByPref struct {
	Date    string `json:"date"`
	Female  string `json:"Female"`
	Male    string `json:"Male"`
	Unknown string `json:"unknown"` // 累計
}

func NewDateSexByPref() DateSexByPref {
	s := DateSexByPref{
		Date:    "",
		Female:  "",
		Male:    "",
		Unknown: "",
	}
	return s
}

type RegionDetail struct {
	Date       string `json:"date"`
	Prefecture string `json:"prefecture"`
	Cases      int    `json:"cases"`
	Hospitals  int    `json:"hospital"`
	Discharges int    `json:"discharge"`
	Deaths     int    `json:"deaths"`
}

func NewRegionDetail() RegionDetail {
	s := RegionDetail{
		Date:       "",
		Prefecture: "",
		Cases:      0,
		Hospitals:  0,
		Discharges: 0,
		Deaths:     0,
	}
	return s
}

type DateAgeByPref struct {
	Date    string `json:"date"`
	Age10   string `json:"age_10"`
	Age20   string `json:"age_20"`
	Age30   string `json:"age_30"`
	Age40   string `json:"age_40"`
	Age50   string `json:"age_50"`
	Age60   string `json:"age_60"`
	Age70   string `json:"age_70"`
	Age80   string `json:"age_80"`
	Age90   string `json:"age_90"`
	Unknown string `json:"age_unknown"`
}

func NewDateAgeByPref() DateAgeByPref {
	s := DateAgeByPref{
		Date:    "",
		Age10:   "",
		Age20:   "",
		Age30:   "",
		Age40:   "",
		Age50:   "",
		Age60:   "",
		Age70:   "",
		Age80:   "",
		Age90:   "",
		Unknown: "",
	}
	return s
}

type DateCallcenter struct {
	Date string `csv:"date"`
	Call string `csv:"call"`
	Fax  string `csv:"fax"`
	Mail string `csv:"mail"`
	Line string `csv:"line"`
}

// Finish
type DateDeathByPref struct {
	Date              string `csv:"date"`
	Hokkaido          string `csv:"北海道"`
	Aomori            string `csv:"青森"`
	Iwate             string `csv:"岩手"`
	Miyagi            string `csv:"宮城"`
	Akita             string `csv:"秋田"`
	Yamagata          string `csv:"山形"`
	Fukushima         string `csv:"福島"`
	Ibaraki           string `csv:"茨城"`
	Tochigi           string `csv:"栃木"`
	Gunma             string `csv:"群馬"`
	Saitama           string `csv:"埼玉"`
	Chiba             string `csv:"千葉"`
	Tokyo             string `csv:"東京"`
	Kanagawa          string `csv:"神奈川"`
	Niigata           string `csv:"新潟"`
	Toyama            string `csv:"富山"`
	Ishikawa          string `csv:"石川"`
	Fukui             string `csv:"福井"`
	Yamanashi         string `csv:"山梨"`
	Nagano            string `csv:"長野"`
	Gifu              string `csv:"岐阜"`
	Shizuoka          string `csv:"静岡"`
	Aichi             string `csv:"愛知"`
	Mie               string `csv:"三重"`
	Shiga             string `csv:"滋賀"`
	Kyoto             string `csv:"京都"`
	Osaka             string `csv:"大阪"`
	Hyogo             string `csv:"兵庫"`
	Nara              string `csv:"奈良"`
	Wakayama          string `csv:"和歌山"`
	Tottori           string `csv:"鳥取"`
	Shimane           string `csv:"島根"`
	Okayama           string `csv:"岡山"`
	Hiroshima         string `csv:"広島"`
	Yamaguchi         string `csv:"山口"`
	Tokushima         string `csv:"徳島"`
	Kagawa            string `csv:"香川"`
	Ehime             string `csv:"愛媛"`
	Kochi             string `csv:"高知"`
	Fukuoka           string `csv:"福岡"`
	Saga              string `csv:"佐賀"`
	Nagasaki          string `csv:"長崎"`
	Kumamoto          string `csv:"熊本"`
	Oita              string `csv:"大分"`
	Miyazaki          string `csv:"宮崎"`
	Kagoshima         string `csv:"鹿児島"`
	Okinawa           string `csv:"沖縄"`
	Charter           string `csv:"チャーター便"`
	QuarantineOfficer string `csv:"検疫職員"`
	Cruise            string `csv:"クルーズ船"`
}

type DatePositiveByPref struct {
	Date              string `csv:"date"`
	Hokkaido          string `csv:"北海道"`
	Aomori            string `csv:"青森"`
	Iwate             string `csv:"岩手"`
	Miyagi            string `csv:"宮城"`
	Akita             string `csv:"秋田"`
	Yamagata          string `csv:"山形"`
	Fukushima         string `csv:"福島"`
	Ibaraki           string `csv:"茨城"`
	Tochigi           string `csv:"栃木"`
	Gunma             string `csv:"群馬"`
	Saitama           string `csv:"埼玉"`
	Chiba             string `csv:"千葉"`
	Tokyo             string `csv:"東京"`
	Kanagawa          string `csv:"神奈川"`
	Niigata           string `csv:"新潟"`
	Toyama            string `csv:"富山"`
	Ishikawa          string `csv:"石川"`
	Fukui             string `csv:"福井"`
	Yamanashi         string `csv:"山梨"`
	Nagano            string `csv:"長野"`
	Gifu              string `csv:"岐阜"`
	Shizuoka          string `csv:"静岡"`
	Aichi             string `csv:"愛知"`
	Mie               string `csv:"三重"`
	Shiga             string `csv:"滋賀"`
	Kyoto             string `csv:"京都"`
	Osaka             string `csv:"大阪"`
	Hyogo             string `csv:"兵庫"`
	Nara              string `csv:"奈良"`
	Wakayama          string `csv:"和歌山"`
	Tottori           string `csv:"鳥取"`
	Shimane           string `csv:"島根"`
	Okayama           string `csv:"岡山"`
	Hiroshima         string `csv:"広島"`
	Yamaguchi         string `csv:"山口"`
	Tokushima         string `csv:"徳島"`
	Kagawa            string `csv:"香川"`
	Ehime             string `csv:"愛媛"`
	Kochi             string `csv:"高知"`
	Fukuoka           string `csv:"福岡"`
	Saga              string `csv:"佐賀"`
	Nagasaki          string `csv:"長崎"`
	Kumamoto          string `csv:"熊本"`
	Oita              string `csv:"大分"`
	Miyazaki          string `csv:"宮崎"`
	Kagoshima         string `csv:"鹿児島"`
	Okinawa           string `csv:"沖縄"`
	Charter           string `csv:"チャーター便"`
	QuarantineOfficer string `csv:"検疫職員"`
	Cruise            string `csv:"クルーズ船"`
}

type DateDetailByPref struct {
	Date                     string `csv:"date"`
	Tests                    string `csv:"tests"`
	Confirmed                string `csv:"confirmed"`
	Deaths                   string `csv:"deaths"`
	Recovered                string `csv:"recovered"`
	Hosp                     string `csv:"hosp"`
	Vent                     string `csv:"vent"`
	Icu                      string `csv:"icu"`
	Severe                   string `csv:"severe"`
	Population               string `csv:"population"`
	AdministrativeAreaLevel  string `csv:"administrative_area_level"`
	AdministrativeAreaLevel1 string `csv:"administrative_area_level_1"`
	AdministrativeAreaLevel2 string `csv:"administrative_area_level_2"`
	JisCode                  string `csv:"jis_code"`
}
