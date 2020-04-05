package types

type Sex uint64

const (
	Sex_MALE   Sex = 0
	Sex_FEMALE Sex = 1
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

type PrefData struct {
	Id      int      `json:"id"`
	NameJa  string   `json:"name_ja"`
	NameEn  string   `json:"name_en"`
	Lat     float64  `json:"lat"`
	Lng     float64  `json:"lng"`
	Cases   int      `json:"cases"`
	Deaths  int      `json:"deaths"`
	SexData *SexData `json:"sex_data"`
	AgeData *AgeData `json:"age_data"`
}

type DateData struct {
	Id      int      `json:"id"`
	Date    int      `json:"date"`
	Cases   int      `json:"cases"`
	Deaths  int      `json:"deaths"`
	SexData *SexData `json:"sex_data"`
	AgeData *AgeData `json:"age_data"`
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

type DatePatientsData struct {
	Id       int        `json:"id"`
	Date     string     `json:"date"`
	Patients []*Patient `json:"patients"`
}

type PrefPatientsData struct {
	Id       int        `json:"id"`
	NameJa   string     `json:"name_ja"`
	NameEn   string     `json:"name_en"`
	Lat      float64    `json:"lat"`
	Lng      float64    `json:"lng"`
	Patients []*Patient `json:"patients"`
}

type DatePrefPatientsData struct {
	Id       int        `json:"id"`
	Date     string     `json:"date"`
	Patients []*Patient `json:"patients"`
}

type TotalData struct {
	Date             int      `json:"date"`
	Pcr              int      `json:"pcr"`
	Positive         int      `json:"positive"`
	Symptom          int      `json:"symptom"`
	Symptomless      int      `json:"symptomless"`
	SymtomConfirming int      `json:"symtomConfirming"`
	Hospitalize      int      `json:"hospitalize"`
	Mild             int      `json:"mild"`
	Severe           int      `json:"severe"`
	Confirming       int      `json:"confirming"`
	Waiting          int      `json:"waiting"`
	Discharge        int      `json:"discharge"`
	Death            int      `json:"death"`
	SexData          *SexData `json:"sex_data"`
	AgeData          *AgeData `json:"age_data"`
}
