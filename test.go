package legolegends

import (
    "fmt"
)

var (
    region string = "NA1"
    accountId int64 = 33206532
    summonerId int64 = 20297715
    championId int = 1
    queue string = "RANKED_SOLO_5x5"
    name string = "xxCode"
)

func Tests() {
    //champ mastery
    res, err := GetChampionMastery(region, summonerId, 1)
    res, err := GetChampionMasteries(region, summonerId)

    //league
    res, err := GetLeagues(region, summonerId)
    res, err := GetLeaguePositions(region, summonerId)
    res, err := GetChallengerLeagues(region, queue)
    res, err := GetMasterLeagues(region, queue)

    //masteries
    res, err := GetSummonerMasteries(region, summonerId)

    //match
    res, err := GetMatchlist(region, accountId, nil, nil, nil, nil, nil, 1, 2)
    res, err := GetMatch(region, matchId)
    res, err := GetMatchTimeline(region, matchId)

    //runes
    res, err := GetSummonerRunes(region, summonerId)

    //spectator
    res, err := GetActiveGame(region, summonerId)

    //status
    res, err := GetShardStatus(region)

    //summoner
    res, err := GetSummonerByName(region, name)
    res, err := GetSummonerByAccountId(region, accountId)
    res, err := GetSummonerBySummonerId(region, summonerId)
}
