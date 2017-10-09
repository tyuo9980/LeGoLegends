package legolegends

import (
	"encoding/json"
	"errors"
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
		"na":   "NA1",
		"euw":  "EUW1",
		"eune": "EUN1",
		"kr":   "KR",
		"br":   "BR1",
		"lan":  "LA1",
		"las":  "LA2",
		"oce":  "OC1",
		"ru":   "RU",
		"tr":   "TR1",
	}

	requestChannelMap = make(map[string]chan bool)
)

const (
	CDN_ROOT string = "https://ddragon.leagueoflegends.com/cdn"
	API_ROOT string = ".api.riotgames.com/lol"

	STATUS string = "/status/v3/shard-data"

	SPECTATOR_ACTIVE string = "/spectator/v3/active-games/by-summoner/"

	SUMMONER_BY_NAME    string = "/summoner/v3/summoners/by-name/"
	SUMMONER_BY_ACCOUNT string = "/summoner/v3/summoners/by-account/"
	SUMMONER_BY_ID      string = "/summoner/v3/summoners/"

	RUNES     string = "/platform/v3/runes/by-summoner/"
	MASTERIES string = "/platform/v3/masteries/by-summoner/"

	MATCH      string = "/match/v3/matches/"
	MATCHLISTS string = "/match/v3/matchlists/by-account/"
	TIMELINES  string = "/match/v3/timelines/by-match/"

	CHAMPION_MASTERY string = "champion-mastery/v3/champion-masteries/by-summoner/"

	CHALLENGER_LEAGUE_BY_QUEUE string = "/league/v3/challengerleagues/by-queue/"
	MASTER_LEAGUE_BY_QUEUE     string = "/league/v3/masterleagues/by-queue/"
	LEAGUE_BY_SUMMONER         string = "/league/v3/leagues/by-summoner/"
	POSITION_BY_SUMMONER       string = "/league/v3/positions/by-summoner/"

	STATIC_CHAMPIONS        string = "/static-data/v3/champions/"
	STATIC_ITEMS            string = "/static-data/v3/items/"
	STATIC_LANGUAGE_STRINGS string = "/static-data/v3/language-strings"
	STATIC_LANGUAGES        string = "/static-data/v3/languages"
	STATIC_MAPS             string = "/static-data/v3/maps"
	STATIC_MASTERIES        string = "/static-data/v3/masteries/"
	STATIC_ICONS            string = "/static-data/v3/profile-icons"
	STATIC_RUNES            string = "/static-data/v3/runes/"
	STATIC_SUMMONERS        string = "/static-data/v3/summoner-spells"
	STATIC_VERSIONS         string = "/static-data/v3/versions"
)

func NormalizeNames(names ...string) []string {
	nameList := []string(names)

	for i, name := range nameList {
		name = strings.Replace(name, " ", "", -1)
		name = strings.ToLower(name)
		nameList[i] = name
	}

	return nameList
}

func createApiUrl(endpoint string, region string, arg string) string {
	return fmt.Sprintf("https://%v.api.riotgames.com/lol/%v/%v?api_key=%v", region, endpoint, arg, ApiKey)
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

	if resp.StatusCode == http.StatusBadRequest {
		log.Println("400 Bad request")
		return errors.New("400")
	} else if resp.StatusCode == http.StatusUnauthorized {
		log.Println("401 Unauthorized - API key may not be set")
		return errors.New("401")
	} else if resp.StatusCode == http.StatusNotFound {
		log.Println("404 No summoner data found for any specified inputs")
		return errors.New("404")
	} else if resp.StatusCode == 429 {
		log.Println("429 Rate limit exceeded")
		return errors.New("429")
	} else if resp.StatusCode == http.StatusInternalServerError {
		log.Println("500 Rito pls")
		return errors.New("500")
	} else if resp.StatusCode == http.StatusServiceUnavailable {
		log.Println("503 Rito servers unavailable")
		return errors.New("503")
	}

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

func DecodeString(jsonString string, v interface{}) error {
	err := json.Unmarshal([]byte(jsonString), v)

	return err
}

func SetApiKey(key string) {
	ApiKey = key
}

func SetDebug(debug bool) {
	Debug = debug
}

func SetRateLimit(rate int, delay int) {
	for _, pid := range pidMap {
		requestChannel := make(chan bool, rate)
		requestChannelMap[pid] = requestChannel
		go rateLimitHandler(requestChannel, delay)
	}
}

func rateLimitHandler(requestChannel chan bool, delay int) {
	for {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		for i := 0; i < len(requestChannel); i++ {
			<-requestChannel
		}
	}
}
