# COVID-19 Rader for Japan Web API

<img src="https://user-images.githubusercontent.com/43264434/78895987-b781fd80-7aaa-11ea-874b-9c49d801e693.png" width=800>

🦠 Web API to get COVID-19(coronavirus) information of each prefecture in Japan

# Features

* 🔁 Update data every 2 hours
* 🚀 Provide REST API
* 🇯🇵 Get info of each prefecture in Japan

# Project using this API

if your project is not listed here,let us know!

## Covid19-rader-for-japan
[https://covid19-rader-for-japan.com](https://covid19-rader-for-japan.com)


# Usage

1. [Patients Data](#anchor1)
2. [Prefecture Data](#anchor2)
3. [DetailByRegion Data](#anchor3)
4. [DailyReport Data](#anchor4)
5. [DailyPositiveByPref Data](#anchor5)
6. [DailyDeathByPref Data](#anchor6)
7. [DailyCallcenter Data](#anchor7)
8. [DailyShip Data](#anchor8)
9. [News Data](#anchor9)
10. [Statistics Data(Duplicated)](#anchor10)

<a id="anchor1"></a>
## Patients Data

**Dataset**

[https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/positiveDetail.csv](https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/positiveDetail.csv)

**Endpont**: [https://covid19-rader-for-japan.appspot.com/api/v1/patients](https://covid19-rader-for-japan.appspot.com/api/v1/patients)
```bash
$ curl https://covid19-rader-for-japan.appspot.com/api/v1/patients
```

**Response:**
```json
[
    {
        "id":"ID01001",
        "date":"2020/01/28",
        "prefecture":"北海道",
        "residence":"国外（武漢市）",
        "age":"40代",
        "sex":"女性",
        "attribute":"来日観光客",
        "prefecture_number":"北海道1",
        "travel_or_contact":"渡航歴",
        "detail":"中国（武漢）",
        "src":"https://www.mhlw.go.jp/stf/newpage_09158.html",
        "onset":"2020/01/26",
        "symptom":"1",
        "death_or_discharge_date":"",
        "comment1":"",
        "comment2":"",
        "outcome":"",
        "outcome_src":""},{"id":"ID01002",
    },
    {
        "date":"2020/02/14",
        "prefecture":"北海道",
        "residence":"札幌市",
        "age":"50代",
        "sex":"男性",
        "attribute":"来日観光客",
        "prefecture_number":"北海道2",
        "travel_or_contact":"",
        "detail":"",
        "src":"http://www.pref.hokkaido.lg.jp/hf/kth/kak/hasseijoukyou.htm",
        "onset":"2020/01/31",
        "symptom":"1",
        "death_or_discharge_date":"",
        "comment1":"",
        "comment2":"",
        "outcome":"",
        "outcome_src":""
    },
...
```

<a id="anchor2"></a>
## Prefecture Data

**Dataset**

[https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/prefectures.csv](https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/prefectures.csv)

**Endpont**: [https://covid19-rader-for-japan.appspot.com/api/v1/prefectures](https://covid19-rader-for-japan.appspot.com/api/v1/prefectures)
```bash
$ curl https://covid19-rader-for-japan.appspot.com/api/v1/prefectures
```

**Response:**
```json
[
  {
    "id": 1,
    "name_ja": "北海道",
    "name_en": "Hokkaido",
    "regions": "0", 
    "lat": "43.46722222",
    "lng": "142.8277778",
  },
  {
    "id": 2,
    "name_ja": "青森",
    "name_en": "Aomori",
    "regions": "1", 
    "lat": "40.78027778",
    "lng": "140.83194440000003",
  },
  ...
]
```

<a id="anchor3"></a>
## DetailByRegion Data

**Dataset**

[https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/prefectures.csv](https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/prefectures.csv)

**Endpont**: [https://covid19-rader-for-japan.appspot.com/api/v1/prefectures](https://covid19-rader-for-japan.appspot.com/api/v1/prefectures)
```bash
$ curl https://covid19-rader-for-japan.appspot.com/api/v1/prefectures
```

**Response:**
```json
[
  {
    "id": 1,
    "name_ja": "北海道",
    "name_en": "Hokkaido",
    "regions": "0", 
    "lat": "43.46722222",
    "lng": "142.8277778",
  },
  {
    "id": 2,
    "name_ja": "青森",
    "name_en": "Aomori",
    "regions": "1", 
    "lat": "40.78027778",
    "lng": "140.83194440000003",
  },
  ...
]
```

<a id="anchor4"></a>
## DailyReport Data

**Dataset**

[https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/resultDailyReport.csv](https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/resultDailyReport.csv)

**Endpont**: [https://covid19-rader-for-japan.appspot.com/api/v1/daily/report](https://covid19-rader-for-japan.appspot.com/api/v1/daily/report)
```bash
$ curl https://covid19-rader-for-japan.appspot.com/api/v1/daily/report
```

**Response:**
```json
[
    {
        "Date":"2020-02-06",
        "PcrD":"132",
        "PositiveD":"16",
        "SymptomD":"16",
        "SymptomlessD":"0",
        "SymptomConfirmingD":"0",
        "HospitalizeD":"9",
        "MildD":"0",
        "SevereD":"0",
        "ConfirmingD":"3",
        "WaitingD":"0",
        "DischargeD":"4",
        "DeathD":"0",
        "PcrF":"566",
        "PositiveF":"9",
        "SymptomF":"5",
        "SymptomlessF":"4",
        "SymptomConfirmingF":"0",
        "HospitalizeF":"9",
        "MildF":"7",
        "SevereF":"0",
        "ConfirmingF":"2",
        "WaitingF":"0",
        "DischargeF":"0",
        "DeathF":"0",
        "PcrX":"",
        "PositiveX":"",
        "Symptom":"",
        "Symptomless":"",
        "SymptomConfirming":"",
        "Hospitalized":"",
        "Mild":"",
        "SevereX":"",
        "Confirming":"",
        "Waiting":"",
        "DischargeX":"",
        "DeathX":"",
        "PcrY":"31",
        "PositiveY":"10",
        "DischargeY":"",
        "SymptomlessDischarge":"",
        "SymptomDischarge":"",
        "SevereY":"",
        "DeathY":"0",
        "Pcr":"729",
        "Discharge":"4",
        "PcrDiff":"",
        "DischargeDiff":""
    }
  ...
]
```

<a id="anchor5"></a>
## DailyPositiveByPref Data

**Dataset**

[https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/byDate.csv](https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/byDate.csv)

**Endpont**: [https://covid19-rader-for-japan.appspot.com/api/v1/daily/positive-by-pref](https://covid19-rader-for-japan.appspot.com/api/v1/daily/positive-by-pref)
```bash
$ curl https://covid19-rader-for-japan.appspot.com/api/v1/daily/positive-by-pref
```

**Response:**
```json
[
    {
        "date":"20200330",
        "hokkaido":"",
        "aomori":"",
        "iwate":"",
        ...
        "tokyo":"1",
        "kanagawa":"1",
        "niigata":"",
        "oita":"",
        "charter":"",
        "quarantine_officer":"",
        "cruise":"1"
    },
  ...
]
```

<a id="anchor6"></a>

## DailyDeathByPref Data

**Dataset**

[https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/death.csv](https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/death.csv)

**Endpont**: [https://covid19-rader-for-japan.appspot.com/api/v1/daily/death-by-pref](https://covid19-rader-for-japan.appspot.com/api/v1/daily/death-by-pref)
```bash
$ curl https://covid19-rader-for-japan.appspot.com/api/v1/daily/death-by-pref
```

**Response:**
```json
[
    {
        "date":"20200330",
        "hokkaido":"",
        "aomori":"",
        "iwate":"",
        ...
        "tokyo":"1",
        "kanagawa":"1",
        "niigata":"",
        "oita":"",
        "charter":"",
        "quarantine_officer":"",
        "cruise":"1"
    },
  ...
]
```

<a id="anchor7"></a>
## DailyCallCenter Data

**Dataset**

[https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/callCenter.csv](https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/callCenter.csv)

**Endpont**: [https://covid19-rader-for-japan.appspot.com/api/v1/daily/callcenter](https://covid19-rader-for-japan.appspot.com/api/v1/daily/callcenter)
```bash
$ curl https://covid19-rader-for-japan.appspot.com/api/v1/daily/callcenter
```

**Response:**
```json
[
    {
        "date":"20200128",
        "call":"99",
        "fax":"",
        "mail":"",
        "line":"3"
    },
    {
        "date":"20200129",
        "call":"250",
        "fax":"",
        "mail":"",
        "line":"3"
    },
  ...
]
```

<a id="anchor8"></a>
## DailyShip Data

**Dataset**

[https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/shipDailyReport.csv](https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/shipDailyReport.csv)

**Endpont**: [https://covid19-rader-for-japan.appspot.com/api/v1/daily/ship](https://covid19-rader-for-japan.appspot.com/api/v1/daily/ship)
```bash
$ curl https://covid19-rader-for-japan.appspot.com/api/v1/daily/ship
```

**Response:**
```json
[
    {
        "date":"20200205",
        "pcr":"31",
        "positive":"10",
        "discharge":"",
        "symotomless_discharge":"",
        "symotom_discharge":"",
        "severe":"",
        "death":"0"
    },
    {
        "date":"20200206",
        "pcr":"",
        "positive":"",
        "discharge":"",
        "symotomless_discharge":"",
        "symotom_discharge":"",
        "severe":"",
        "death":"0"
    },
    ...
]
```

<a id="anchor9"></a>
## DailyNews Data

**Dataset**

[https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/mhlw_houdou.csv](https://github.com/swsoyee/2019-ncov-japan/blob/master/Data/mhlw_houdou.csv)

**Endpont**: [https://covid19-rader-for-japan.appspot.com/api/v1/news](https://covid19-rader-for-japan.appspot.com/api/v1/news)
```bash
$ curl https://covid19-rader-for-japan.appspot.com/api/v1/news
```

**Response:**
```json
[
    {
        "id":"1",
        "date":"20200116",
        "title":"新型コロナウイルスに関連した肺炎の患者の発生について（1例目）",
        "Link":"https://www.mhlw.go.jp/stf/newpage_08906.html"
    },
    {
        "id":"2",
        "date":"20200124",
        "title":"新型コロナウイルスに関連した肺炎の患者の発生について（2例目）",
        "Link":"https://www.mhlw.go.jp/stf/newpage_09079.html"
    },
...
]
```

<a id="anchor10"></a>
## Statistics Data

### Caution!
```
This is duplicated because calculation of total data may be wrong.
Please use detailByRegion data.
```

**Endpont**: [https://covid19-rader-for-japan.appspot.com/api/v1/stats](https://covid19-rader-for-japan.appspot.com/api/v1/stats)
```bash
$ curl https://covid19-rader-for-japan.appspot.com/api/v1/stats
```

**Response:**
```json
[
    {
        "date":"20200319",
        "prefecture":"北海道",
        "cases":0,
        "total_cases":153,
        "hospital":0,
        "total_hospitals":77,
        "discharge":0,
        "total_discharges":69,
        "deaths":0,
        "total_deaths":7,
        "sex_data":{
            "female":0,
            "male":0,
            "unknown":0
        },
        "total_sex_data":{
            "female":0,
            "male":0,
            "unknown":0
        },
        "age_data":{
            "age_10":0,
            "age_20":0,
            "age_30":0,
            "age_40":0,
            "age_50":0,
            "age_60":0,
            "age_70":0,
            "age_80":0,
            "age_90":0,
            "age_unknown":0
        },
        "total_age_data":{
            "age_10":0,
            "age_20":0,
            "age_30":0,
            "age_40":0,
            "age_50":0,
            "age_60":0,
            "age_70":0,
            "age_80":0,
            "age_90":0,
            "age_unknown":0
        }
    },
    {
        "date":"20200319",
        "prefecture":"愛知県",
        "cases":0,
        "total_cases":123,
        "hospital":0,
        "total_hospitals":106,
        "discharge":0,
        "total_discharges":3,
        "deaths":0,
        "total_deaths":14,
        "sex_data":{
            "female":0,
            "male":0,
            "unknown":0
        },
        "total_sex_data":{
            "female":0,
            "male":0,
            "unknown":0
        },
        "age_data":{
            "age_10":0,
            "age_20":0,
            "age_30":0,
            "age_40":0,
            "age_50":0,
            "age_60":0,
            "age_70":0,
            "age_80":0,
            "age_90":0,
            "age_unknown":0
        },
        "total_age_data":{
            "age_10":0,
            "age_20":0,
            "age_30":0,
            "age_40":0,
            "age_50":0,
            "age_60":0,
            "age_70":0,
            "age_80":0,
            "age_90":0,
            "age_unknown":0
        }
    }

...
```

# Data Sources

* [swsoyee/2019-ncov-japan](https://github.com/swsoyee/2019-ncov-japan)


# How to contribute
## Rule of branch
Please pull-request to development branch only. 
Don't pull-request to master and staging branch.

If you send pull-request, please follow the roles below.
1. function addition: feature/#{ISSUE_ID}-#{branch_title_name}
2. hotfix: hotfix/#{ISSUE_ID}-#{branch_title_name}

## Basic branch
| purpose | branch | remarks |
| ---- | -------- | ---- |
| Development | development | base branch. Basically send a Pull Request here |
| Staging | staging | For final confirmation before production. Non-admin pull requests are forbidden |
| Production | master | Non-admin pull requests are forbidden |


# Contributers
| [inductor](https://github.com/inductor) | [mattn](https://github.com/mattn) | [Yoshiteru Nagata](https://github.com/nagata-yoshiteru) | [otokunaga](https://github.com/otokunaga) |
|:---:|:---:|:---:|:---:|
<img src="https://avatars3.githubusercontent.com/u/20236173?s=400&u=d8dda91e4bc2bdc7736f607b36fa53c9e82e08db&v=4" width=100> |<img src="https://avatars3.githubusercontent.com/u/10111?s=400&u=52c03ac58f0027d43f6708fcbc3c2913f195439c&v=4" width=100> |<img src="https://avatars0.githubusercontent.com/u/38305549?s=400&v=4" width=100> |<img src="https://avatars3.githubusercontent.com/u/36445214?s=400&v=4" width=100> |
