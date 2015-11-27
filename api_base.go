package legolegends

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var (
	Region  string = "na"
	Api_Key string = "6bf45188-e99f-46c1-9695-1a035454bc4c"
)

const (
	API_PARAM string = "api_key=" + API_KEY

	BASE          string = "api.pvp.net"
	CDN_ROOT_PATH string = "https://ddragon.leagueoflegends.com/cdn"
	API_STR       string = "/api/lol"
	API_STR_DATA  string = API_STR + "/static-data"

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

func SendRequest(url string) io.Reader {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	printError(err)

	return resp.Body
}

func DecodeRequest(r io.Reader, v interface{}) interface{} {
	err := json.NewDecoder(r).Decode(v)
	printError(err)

	return v
}

func printError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func setRegion(region string) {
	Region = region
}

func setAPIKey(key string) {
	Api_Key = key
}
