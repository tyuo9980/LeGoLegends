package legolegends

import (
    "fmt"
    "log"
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

func GetSummonerByNames(names ...string) (map[string]Summoner, error) {
    nameString := strings.Join([]string(names), ",")
    log.Println(nameString)

    args := createArgs("")
    url := createApiUrl(SUMMONER_BY_NAME) + fmt.Sprintf("%d?%v", nameString, args)

    log.Println(url)
    summoners := make(map[string]Summoner)
    err := decodeRequest(url, &summoners)

    return summoners, err
}

func GetSummonerByIds(ids ...string) map[string]Summoner {
    idString := strings.Join([]string(ids), ",")
    args := createArgs("")
    url := createApiUrl(SUMMONER_BY_ID) + fmt.Sprintf("%d?%v", idString, args)

    summoners := make(map[string]Summoner)
    decodeRequest(url, &summoners)

    return summoners
}

func GetSummonerMasteries(ids ...string) map[string]MasteryPages {
    idString := strings.Join([]string(ids), ",")
    args := createArgs("")
    url := createApiUrl(SUMMONER_MASTERIES) + fmt.Sprintf("%d?%v", idString, args)

    masteries := make(map[string]MasteryPages)
    decodeRequest(url, &masteries)

    return masteries
}

func GetSummonerRunes(ids ...string) map[string]RunePages {
    idString := strings.Join([]string(ids), ",")
    args := createArgs("")
    url := createApiUrl(SUMMONER_RUNES) + fmt.Sprintf("%d?%v", idString, args)

    runes := make(map[string]RunePages)
    decodeRequest(url, &runes)

    return runes
}

func GetSummonerName(ids ...string) map[string]string {
    idString := strings.Join([]string(ids), ",")
    args := createArgs("")
    url := createApiUrl(SUMMONER_NAME) + fmt.Sprintf("%d?%v", idString, args)

    names := make(map[string]string)
    decodeRequest(url, &names)

    return names
}
