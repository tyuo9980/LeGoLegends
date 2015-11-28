package legolegends

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "strconv"
    "strings"
)

var (
    Region string
    ApiKey string
)

const (
    CDN_ROOT string = "https://ddragon.leagueoflegends.com/cdn"

    MATCH_LIST       string = "/v2.2/matchlist/by-summoner"
    MATCH            string = "/v2.2/match"
    SUMMONER_BY_NAME string = "/v1.4/summoner/by-name"

    STATIC_VERSION  string = "/v1.2/versions"
    STATIC_CHAMPION string = "/v1.2/champion"
    STATIC_ITEM     string = "/v1.2/item"
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
    args := "?api_key=" + Api_Key

    for i, arg := range argVals {
        var val string
        switch arg.(type) {
        case bool:
            val = strconv.FormatBool(arg.(bool))
        case int:
            val = strconv.Itoa(arg.(int))
        case int64:
            val = strconv.FormatInt(arg.(int64), 10)
        default:
            val = arg.(string)
        }

        args += fmt.Sprintf("?%v=%v", keyArray[i], val)
    }

    return args
}

func sendRequest(url string) io.Reader {
    resp, err := http.Get(url)
    defer resp.Body.Close()
    printError(err)

    return resp.Body
}

func decodeRequest(r io.Reader, v interface{}) {
    err := json.NewDecoder(r).Decode(&v)
    printError(err)
}

func printError(err error) {
    if err != nil {
        log.Println(err)
    }
}

func SetRegion(region string) {
    Region = region
}

func SetApiKey(key string) {
    ApiKey = key
}
