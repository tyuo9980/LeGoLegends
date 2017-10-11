
import (
    "testing"
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
    res, err := GetChampionMastery(region, summonerId, 1)
    if err != nil {
        t.Errorf("Expected no error, got %s", err)
    }
    /*
    GetChampionMasteries(region, summonerId)

    //league
    GetLeagues(region, summonerId)
    GetLeaguePositions(region, summonerId)
    GetChallengerLeagues(region, queue)
    GetMasterLeagues(region, queue)

    //masteries
    GetSummonerMasteries(region, summonerId)

    //match
    //res7, err := GetMatchlist(region, accountId, nil, nil, nil, nil, nil, 1, 2)
    GetMatch(region, matchId)
    GetMatchTimeline(region, matchId)

    //runes
    GetSummonerRunes(region, summonerId)

    //spectator
    GetActiveGame(region, summonerId)

    //status
    GetShardStatus(region)

    //summoner
    GetSummonerByName(region, name)
    GetSummonerByAccountId(region, accountId)
    GetSummonerBySummonerId(region, summonerId)
    */
}
