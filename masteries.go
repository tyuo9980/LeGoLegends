package legolegends

import (
    "fmt"
    "strings"
)

type MasteryPages struct {
    Pages      []MasteryPage `json:"pages"`
    SummonerId int64         `json:"summonerId"`
}

type MasteryPage struct {
    Id        int64         `json:"id"`
    Name      string        `json:"name"`
    Current   bool          `json:"current"`
    Masteries []MasterySlot `json:"masteries"`
}

type MasterySlot struct {
    Id   int `json:"id"`
    Rank int `json:"rank"`
}

func GetSummonerMasteries(region string, summonerId string) (MasteryPages, error) {
    url := createApiUrl(MASTERIES, region, summonerId)

    var masteries MasteryPages
    err := decodeRequest(region, url, &masteries)

    return masteries, err
}
