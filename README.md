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

# Getting Started
```
$ git clone https://github.com/RuiHirano/covid19-rader-for-japan-api.git
$ cd covid19-rader
$ git submodule init
$ git submodule update
$ cd src
$ go run main.go
```
You can get information at http://localhost:5000/api/v2/....

# Usage
1. [DailyPositiveByPref Data](#anchor1)
2. [Daitailf Data](#anchor2)
3. [DailyDetailByPref Data](#anchor3)
4. [DailyCallcenter Data](#anchor4)
5. [News Data](#anchor5)
6. [Summary Data](#anchor6)
7. [WorldSummary Data](#anchor7)

<a id="anchor1"></a>
## 1. DailyPositiveByPref Data

**Dataset**

[https://github.com/swsoyee/2019-ncov-japan/blob/master/50_Data/byDate.csv](https://github.com/swsoyee/2019-ncov-japan/blob/master/50_Data/byDate.csv)

**Endpont**: [https://covid19-rader-for-japan.appspot.com/api/v2/daily/positive_by_pref](https://covid19-rader-for-japan.appspot.com/api/v2/daily/positive_by_pref)
```bash
// $ curl https://covid19-rader-for-japan.appspot.com/api/v2/daily/positive_by_pref
// or at local
$ curl https://localhost:5000/api/v2/daily/positive_by_pref
```

**Response:**
```json
[
    {
        "date":"20200330",          // 日付
        "hokkaido":"",              // 北海道における死亡者数
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

<a id="anchor2"></a>

## 2. DailyDeathByPref Data

**Dataset**

[https://github.com/swsoyee/2019-ncov-japan/blob/master/50_Data/death.csv](https://github.com/swsoyee/2019-ncov-japan/blob/master/50_Data/death.csv)

**Endpont**: [https://covid19-rader-for-japan.appspot.com/api/v2/daily/death_by_pref](https://covid19-rader-for-japan.appspot.com/api/v2/daily/death_by_pref)
```bash
// $ curl https://covid19-rader-for-japan.appspot.com/api/v2/daily/death_by_pref
// or at local
$ curl https://localhost:5000/api/v2/daily/death_by_pref
```

**Response:**
```json
[
    {
        "date":"20200330",          // 日付
        "hokkaido":"",              // 北海道における死亡者数
        "aomori":"",
        "iwate":"",
        ...
        "tokyo":"1",
        "kanagawa":"1",
        "niigata":"",
        "oita":"",
        "charter":"",
        "quarantine_officer":"",    // 検疫官の死亡者数
        "cruise":"1"                // クルーズ船での死亡者数
    },
  ...
]
```

<a id="anchor3"></a>

## 3. DailyDetailByPref Data

**Dataset**

[https://github.com/swsoyee/2019-ncov-japan/blob/master/50_Data/covid19_jp.csv](https://github.com/swsoyee/2019-ncov-japan/blob/master/50_Data/covid19_jp.csv)

**Endpont**: [https://covid19-rader-for-japan.appspot.com/api/v2/daily/detail_by_pref](https://covid19-rader-for-japan.appspot.com/api/v2/daily/detail_by_pref)
```bash
// $ curl https://covid19-rader-for-japan.appspot.com/api/v2/daily/detail_by_pref
// or at local
$ curl https://localhost:5000/api/v2/daily/detail_by_pref
```

**Response:**
```json
[
    {
        "date":"20200206",                      // 日付
        "tests":"698",                          // 検査数
        "confirmed":"25",                       // 確認済み
        "deaths":"0",                           // 死亡者数
        "recovered":"4",                        // 回復者数
        "hosp":"18",                            // 入院者数
        "vent":"",                              // 
        "icu":"",                               // 
        "severe":"0",                           // 重症者数
        "population":"126216142",               // 人口
        "administrative_area_level":"1",        // 管理エリアレベル
        "administrative_area_level_1":"Japan",  // 
        "administrative_area_level_2":"",       // 
        "jis_code":""                           // JISコード
    },

  ...
]
```

<a id="anchor4"></a>
## 4. DailyCallCenter Data

**Dataset**

[https://github.com/swsoyee/2019-ncov-japan/blob/master/50_Data/MHLW/callCenter.csv](https://github.com/swsoyee/2019-ncov-japan/blob/master/50_Data/MHLW/callCenter.csv)

**Endpont**: [https://covid19-rader-for-japan.appspot.com/api/v2/daily/callcenter](https://covid19-rader-for-japan.appspot.com/api/v2/daily/callcenter)
```bash
// $ curl https://covid19-rader-for-japan.appspot.com/api/v2/daily/callcenter
// or at local
$ curl https://localhost:5000/api/v2/daily/callcenter
```

**Response:**
```json
[
    {
        "date":"20200128",     // 日付
        "call":"99",           // 電話問合せ件数
        "fax":"",              // fax件数
        "mail":"",             // mail件数
        "line":"3"             // 
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

<a id="anchor5"></a>
## 5. DailyNews Data

**Dataset**

[https://github.com/swsoyee/2019-ncov-japan/blob/master/50_Data/mhlw_houdou.csv](https://github.com/swsoyee/2019-ncov-japan/blob/master/50_Data/mhlw_houdou.csv)

**Endpont**: [https://covid19-rader-for-japan.appspot.com/api/v2/news](https://covid19-rader-for-japan.appspot.com/api/v2/news)
```bash
// $ curl https://covid19-rader-for-japan.appspot.com/api/v2/news
// or at local
$ curl https://localhost:5000/api/v2/news
```

**Response:**
```json
[
    {                                                      
        "date":"20200116",                                              // 日付
        "title":"新型コロナウイルスに関連した肺炎の患者の発生について（1例目）",  // タイトル
        "link":"https://www.mhlw.go.jp/stf/newpage_08906.html",         // URL
        "pre":"神奈川",                                                  // 都道府県名
        "resident":"日本"                                               // 国名
    },
    {
        "date":"20200124",
        "title":"新型コロナウイルスに関連した肺炎の患者の発生について（2例目）",
        "link":"https://www.mhlw.go.jp/stf/newpage_09079.html",
        "pre":"神奈川",
        "resident":"日本"
    },
...
]
```

<a id="anchor6"></a>
## 6. Summary Data

**Dataset**

[https://github.com/swsoyee/2019-ncov-japan/blob/master/50_Data/MHLW/summary.csv](https://github.com/swsoyee/2019-ncov-japan/blob/master/50_Data/MHLW/summary.csv)

**Endpont**: [https://covid19-rader-for-japan.appspot.com/api/v2/summary](https://covid19-rader-for-japan.appspot.com/api/v2/summary)
```bash
// $ curl https://covid19-rader-for-japan.appspot.com/api/v2/summary
// or at local
$ curl https://localhost:5000/api/v2/summary
```

**Response:**
```json
[
    {
        "date":"20200205",          // 日付
        "prefecture":"クルーズ船",    // 都道府県
        "positives":"10",           // 陽性者数
        "pcrs":"31",                // 検査人数
        "hospitals":"",             // 入院者数
        "severes":"",               // 重症者数
        "discharges":"",            // 退院者数
        "deaths":"0",               // 死亡者数
        "checking":"",              // 確認中
        "class":"3"                 // 分類
    },
...
]
```

<a id="anchor7"></a>
## 7. World Summary Data

**Dataset**

[https://github.com/swsoyee/2019-ncov-japan/blob/master/FIND/worldSummary.csv](https://github.com/swsoyee/2019-ncov-japan/blob/master/FIND/worldSummary.csv)

**Endpont**: [https://covid19-rader-for-japan.appspot.com/api/v2/world_summary](https://covid19-rader-for-japan.appspot.com/api/v2/world_summary)
```bash
// $ curl https://covid19-rader-for-japan.appspot.com/api/v2/world_summary
// or at local
$ curl https://localhost:5000/api/v2/world_summary
```

**Response:**
```json
[
    {
        "date":"20200224",          // 日付
        "update":"34",              // 
        "cases":"1",                // 件数
        "new_cases":"1",            // 新規件数
        "deaths":"0",               // 死亡者数
        "country":"Afghanistan",    // 国名
        "last_update":"",           // 
        "population":"38928000",    // 人口
        "casesPer100k":"0",         // 10万人あたりの件数割合
        "new_tests":"",             // 新規検査人数
        "tests_cumulative":"",      // 累計検査人数
        "testsPer100k":"",          // 10万人あたりの検査人数割合
        "positiveRate":""           // 陽性者率
    },

...
]
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
