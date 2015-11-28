package legolegends

type MatchList struct {
    TotalGames int              `json:"totalGames"`
    StartIndex int              `json:"startIndex"`
    EndIndex   int              `json:"endIndex"`
    Matches    []MatchReference `json:"matches"`
}

type MatchReference struct {
    Timestamp  int64  `json:"timestamp"`
    Champion   string `json:"champion"`
    Region     string `json:"region"`
    Queue      string `json:"queue"`
    Season     string `json:"season"`
    MatchId    int64  `json:"matchId"`
    Role       string `json:"role"`
    PlatformId int64  `json:"platformId"`
    Lane       string `json:"lane"`
}

func GetMatchlist(
    summonerId string,
    championIds string,
    rankedQueues string,
    seasons string,
    beginTime int64,
    endTime int64,
    beginIndex int,
    endIndex int) MatchList {

    idString := strings.Join([]string(championIds), ",")
    argString := "championIds,rankedQueues,seasons,beginTime,endTime,beginIndex,endIndex"
    args := createArgs(argString, championIds, rankedQueues, seasons, beginTime, endTime, beginIndex, endIndex)
    url := createApiUrl(SUMMONER_BY_NAME) + fmt.Sprintf("%d?%v", summonerId, args)

    var matchList MatchList
    reqBody := sendRequest(url)
    decodeRequest(reqBody, &matchList)

    return matchList
}
