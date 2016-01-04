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

type MasteryPages struct {
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

type RunePages struct {
    Pages      []RunePage `json:"pages"`
    SummonerId int64      `json:"summonerId"`
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

func GetSummonerByNames(region string, names ...string) (map[string]Summoner, error) {
    nameString := strings.Join([]string(names), ",")

    args := createArgs("")
    url := createApiUrl(SUMMONER_BY_NAME, region) + fmt.Sprintf("%v?%v", nameString, args)

    summoners := make(map[string]Summoner)
    err := decodeRequest(url, &summoners)

    return summoners, err
}

func GetSummonerByIds(region string, ids ...string) (map[string]Summoner, error) {
    idString := strings.Join([]string(ids), ",")

    args := createArgs("")
    url := createApiUrl(SUMMONER_BY_ID, region) + fmt.Sprintf("%v?%v", idString, args)

    summoners := make(map[string]Summoner)
    err := decodeRequest(url, &summoners)

    return summoners, err
}

func GetSummonerMasteries(region string, ids ...string) (map[string]MasteryPages, error) {
    idString := strings.Join([]string(ids), ",")

    args := createArgs("")
    url := createApiUrl(SUMMONER_MASTERIES, region) + fmt.Sprintf("%v/masteries?%v", idString, args)

    masteries := make(map[string]MasteryPages)
    err := decodeRequest(url, &masteries)

    return masteries, err
}

func GetSummonerRunes(region string, ids ...string) (map[string]RunePages, error) {
    idString := strings.Join([]string(ids), ",")

    args := createArgs("")
    url := createApiUrl(SUMMONER_RUNES, region) + fmt.Sprintf("%v/runes?%v", idString, args)

    runes := make(map[string]RunePages)
    err := decodeRequest(url, &runes)

    return runes, err
}

func GetSummonerName(region string, ids ...string) (map[string]string, error) {
    idString := strings.Join([]string(ids), ",")

    args := createArgs("")
    url := createApiUrl(SUMMONER_NAME, region) + fmt.Sprintf("%v?%v", idString, args)

    names := make(map[string]string)
    err := decodeRequest(url, &names)

    return names, err
}
