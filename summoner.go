package legolegends

import (
    "fmt"
    "strings"
)

type Summoner struct {
    AccountId     int64  `json:"accountId"`
    Id            int64  `json:"id"`
    Name          string `json:"name"`
    ProfileIconId int    `json:"profileIconId"`
    RevisionDate  int64  `json:"revisionDate"`
    SummonerLevel int64  `json:"summonerLevel"`
}

func GetSummonerByName(region string, name string) (Summoner, error) {
    url := createApiUrl(SUMMONER_BY_NAME, region, name)

    var summoner Summoner
    err := decodeRequest(region, url, &summoner)

    return summoners, err
}

func GetSummonerByAccountId(region string, accountId string) (Summoner, error) {
    url := createApiUrl(SUMMONER_BY_ACCOUNT, region, accountId)

    var summoners Summoner
    err := decodeRequest(region, url, &summoner)

    return summoners, err
}

func GetSummonerBySummonerId(region string, summonerId string) (Summoner, error) {
    url := createApiUrl(SUMMONER_BY_ID, region, summonerId)

    var summoners Summoner
    err := decodeRequest(region, url, &summoner)

    return summoners, err
}
