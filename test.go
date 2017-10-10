package legolegends

import (
    "testing"
    l "github.com/tyuo9980/LeGoLegends"
)

var (
    region string = "NA1"
    accountId int64 = 33206532
    summonerId int64 = 20297715
    championId int = 1
    matchId int64 = 2616900038
    queue string = "RANKED_SOLO_5x5"
    name string = "xxCode"
)

func TestAll(t *testing.T) {
    //champ mastery
    res, err := l.GetChampionMastery(region, summonerId, 1)
    res, err = l.GetChampionMasteries(region, summonerId)

    //league
    res, err = l.GetLeagues(region, summonerId)
    res, err = l.GetLeaguePositions(region, summonerId)
    res, err = l.GetChallengerLeagues(region, queue)
    res, err = l.GetMasterLeagues(region, queue)

    //masteries
    res, err = l.GetSummonerMasteries(region, summonerId)

    //match
    res, err = l.GetMatchlist(region, accountId, nil, nil, nil, nil, nil, 1, 2)
    res, err = l.GetMatch(region, matchId)

    res, err = l.GetMatchTimeline(region, matchId)
    //runes
    res, err = l.GetSummonerRunes(region, summonerId)

    //spectator
    res, err = l.GetActiveGame(region, summonerId)

    //status
    res, err = l.GetShardStatus(region)

    //summoner
    res, err = l.GetSummonerByName(region, name)
    res, err = l.GetSummonerByAccountId(region, accountId)
    res, err = l.GetSummonerBySummonerId(region, summonerId)
}
