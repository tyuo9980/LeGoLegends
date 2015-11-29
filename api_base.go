package legolegends

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"
    "strings"
)

var (
    Region string
    ApiKey string
    Debug  bool = false
)

const (
    CDN_ROOT string = "https://ddragon.leagueoflegends.com/cdn"

    MATCH_LIST string = "v2.2/matchlist/by-summoner/"

    MATCH string = "v2.2/match/"

    SUMMONER_BY_NAME   string = "v1.4/summoner/by-name/"
    SUMMONER_BY_ID     string = "v1.4/summoner/"
    SUMMONER_NAME      string = "v1.4/summoner/"
    SUMMONER_RUNES     string = "v1.4/summoner/"
    SUMMONER_MASTERIES string = "v1.4/summoner/"

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

func createApiUrl(endpoint string) string {
    return fmt.Sprintf("https://%v.api.pvp.net/api/lol/%v/%v", Region, Region, endpoint)
}

func createStaticUrl(endpoint string) string {
    return fmt.Sprintf("https://%v.api.pvp.net/api/lol/static-data/%v/%v", Region, Region, endpoint)
}

func createObserverUrl(endpoint string) string {
    return fmt.Sprintf("https://%v.api.pvp.net/observer-mode/rest/%v", Region, endpoint)
}

func createShardUrl(endpoint string) string {
    return fmt.Sprintf("https://%v.api.pvp.net/shards/%v", Region, endpoint)
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

func decodeRequest(url string, v interface{}) error {
    if Debug {
        log.Println("decodeRequest: " + url)
    }

    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    //err = json.NewDecoder(resp.Body).Decode(&v)
    err = json.Unmarshal(body, v)
    if err != nil {
        return err
    }

    return err
}

func SetRegion(region string) {
    Region = region
}

func SetApiKey(key string) {
    ApiKey = key
}

func SetDebug(debug bool) {
    Debug = debug
}
