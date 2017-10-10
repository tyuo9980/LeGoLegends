package legolegends

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

    return summoner, err
}

func GetSummonerByAccountId(region string, accountId int64) (Summoner, error) {
    url := createApiUrl(SUMMONER_BY_ACCOUNT, region, accountId)

    var summoner Summoner
    err := decodeRequest(region, url, &summoner)

    return summoner, err
}

func GetSummonerBySummonerId(region string, summonerId int64) (Summoner, error) {
    url := createApiUrl(SUMMONER_BY_ID, region, summonerId)

    var summoner Summoner
    err := decodeRequest(region, url, &summoner)

    return summoner, err
}
