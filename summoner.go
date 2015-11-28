package legolegends

import (
    "fmt"
    "strings"
)

type Summoner struct {
    Id            int64  `json:"id"`
    Name          string `json:"name"`
    ProfileIconId int    `json:"profileIconId"`
    RevisionDate  int64  `json:"revisionDate"`
    SummonerLevel int    `json:"summonerLevel"`
}

type MasteryBook struct {
    Pages      []MasteryPage `json:"pages"`
    SummonerId int64         `json:"summonerId"`
}

type MasteryPage struct {
    Id        int64     `json:"id"`
    Name      string    `json:"name"`
    Current   bool      `json:"current"`
    Masteries []Mastery `json:"masteries"`
}

type Mastery struct {
    Id   int `json:"id"`
    Rank int `json:"rank"`
}

type RuneBook struct {
    Pages      []RunePage `json:"pages"`
    SummonerId int        `json:"summonerId"`
}

type RunePage struct {
    Current bool       `json:"current"`
    Id      int64      `json:"id"`
    Name    string     `json:"name"`
    Slots   []RuneSlot `json:"slots"`
}

type RuneSlot struct {
    RuneId     int `json:"runeId"`
    RuneSlotId int `json:"runeSlotId"`
}

func GetSummonerByNames(names ...string) map[string]Summoner {
    nameString := strings.Join([]string(names), ",")
    args := createArgs("")
    url := createApiUrl(SUMMONER_BY_NAME) + fmt.Sprintf("%d?%v", nameString, args)

    summoners := make(map[string]Summoner)
    reqBody := sendRequest(url)
    decodeRequest(reqBody, &summoners)

    return summoners
}

func getSummonerByIds() {

}

func getSummonerMasteries() {

}

func getSummonerRunes() {

}

func getSummonerNames() {

}
