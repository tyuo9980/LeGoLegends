package legolegends

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "strings"
    "time"
)

var (
    PlatformId string
    ApiKey     string
    Debug      bool = false

    pidMap = map[string]string{
        "NA":   "NA1",
        "EUW":  "EUW1",
        "EUNE": "EUN1",
        "KR":   "KR",
        "BR":   "BR1",
        "LAN":  "LA1",
        "LAS":  "LA2",
        "OCE":  "OC1",
        "RU":   "RU",
        "TR":   "TR1",
    }

    requestChannelMap = make(map[string]chan bool)
)

const (
    CDN_ROOT string = "https://ddragon.leagueoflegends.com/cdn"

    MATCH_LIST string = "v2.2/matchlist/by-summoner/"

    MATCH string = "v2.2/match/"

    GAME string = "v1.3/game/by-summoner/"

    CURRENT_GAME string = "observer-mode/rest/consumer/getSpectatorGameInfo/"

    LEAGUE_BY_SUMMONER string = "v2.5/league/by-summoner/"
    LEAGUE_BY_TEAM     string = "v2.5/league/by-team/"
    LEAGUE_CHALLENGER  string = "v2.5/league/challenger/"
    LEAGUE_MASTER      string = "v2.5/league/master"

    SUMMONER_BY_NAME   string = "v1.4/summoner/by-name/"
    SUMMONER_BY_ID     string = "v1.4/summoner/"
    SUMMONER_NAME      string = "v1.4/summoner/"
    SUMMONER_RUNES     string = "v1.4/summoner/"
    SUMMONER_MASTERIES string = "v1.4/summoner/"

    STATS_RANKED  string = "v1.3/stats/by-summoner/"
    STATS_SUMMARY string = "v1.3/stats/by-summoner/"

    STATIC_VERSION  string = "v1.2/versions/"
    STATIC_CHAMPION string = "v1.2/champion/"
    STATIC_ITEM     string = "v1.2/item/"
)

func NormalizeName(names ...string) []string {
    var nameList []string

    for i, name := range names {
        name = strings.Replace(name, " ", "", -1)
        name = strings.ToLower(name)
        nameList[i] = name
    }

    return nameList
}

func createApiUrl(endpoint string, region string) string {
    return fmt.Sprintf("https://%v.api.pvp.net/api/lol/%v/%v", region, region, endpoint)
}

func createStaticUrl(endpoint string, region string) string {
    return fmt.Sprintf("https://%v.api.pvp.net/api/lol/static-data/%v/%v", region, region, endpoint)
}

func createObserverUrl(endpoint string, region string) string {
    return fmt.Sprintf("https://%v.api.pvp.net/observer-mode/rest/%v", region, endpoint)
}

func createShardUrl(endpoint string, region string) string {
    return fmt.Sprintf("https://%v.api.pvp.net/shards/%v", region, endpoint)
}

func createArgs(keys string, argVals ...interface{}) string {
    keyArray := strings.Split(keys, ",")
    args := "api_key=" + ApiKey

    for i, arg := range argVals {
        var val string = ""
        switch arg.(type) {
        case bool:
            val = strconv.FormatBool(arg.(bool))
        case int:
            if arg.(int) != -1 {
                val = strconv.Itoa(arg.(int))
            }
        case int64:
            if arg.(int64) != -1 {
                val = strconv.FormatInt(arg.(int64), 10)
            }
        default:
            val = arg.(string)
        }

        args += fmt.Sprintf("&%v=%v", keyArray[i], val)
    }

    return args
}

func decodeRequest(region string, url string, v interface{}) error {
    pid := pidMap[region]
    requestChannelMap[pid] <- true

    if Debug {
        log.Println("decodeRequest: " + url)
    }

    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    err = json.NewDecoder(resp.Body).Decode(&v)
    if err != nil {
        return err
    }

    return err
}

func EncodeStruct(v interface{}) (string, error) {
    jsonString, err := json.Marshal(v)

    return string(jsonString), err
}

func SetApiKey(key string) {
    ApiKey = key
}

func SetDebug(debug bool) {
    Debug = debug
}

func SetRateLimit(rps int) {
    for _, pid := range pidMap {
        log.Println(pid)
        requestChannel := make(chan bool, rps)
        requestChannelMap[pid] = requestChannel
        go rateLimitHandler(requestChannel)
    }
}

func rateLimitHandler(requestChannel chan bool) {
    for {
        time.Sleep(time.Second)
        for i := 0; i < len(requestChannel); i++ {
            <-requestChannel
        }
    }
}
