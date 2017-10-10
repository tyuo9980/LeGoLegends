package legolegends

import (
    "fmt"
    "strings"
)

type LeagueList struct {
    Entries       []LeaguePosition `json:"entries"`
    Name          string        `json:"name"`
    Queue         string        `json:"queue"`
    Tier          string        `json:"tier"`
}

type LeaguePosition struct {
    FreshBlood     bool   `json:"freshBlood"`
    HotStreak      bool   `json:"hotStreak"`
    Inactive       bool   `json:"inactive"`
    Veteran        bool   `json:"veteran"`
    LeaguePoints     int    `json:"leaguePoints"`
    Losses           int    `json:"losses"`
    MiniSeries       MiniSeries `json:"miniSeries"`
    PlayerOrTeamId   string `json:"playerOrTeamId"`
    PlayerOrTeamName string `json:"playerOrTeamName"`
    Rank             string `json:"rank"`
    Wins             int    `json:"wins"`
}

type MiniSeries struct {
    Losses   int    `json:"losses"`
    Progress string `json:"progress"`
    Target   int    `json:"target"`
    Wins     int    `json:"wins"`
}

func GetLeagues(region string, summonerId string) ([]LeagueList, error) {
    url := createApiUrl(LEAGUE_BY_SUMMONER, region, summonerId)

    var leagues []LeagueList
    err := decodeRequest(region, url, &leagues)

    return leagues, err
}

func GetLeaguePositions(region string, summonerId string) ([]LeaguePosition, error) {
    url := createApiUrl(POSITION_BY_SUMMONER, region, summonerId)

    var positions []LeaguePosition
    err := decodeRequest(region, url, &positions)

    return positions, err
}

func GetChallengerLeagues(region string, queue string) (LeagueList, error) {
    args := createArgs("type", queue)
    url := createApiUrl(CHALLENGER_LEAGUE, region) + fmt.Sprintf("?%v", args)

    var leagues LeagueList
    err := decodeRequest(region, url, &leagues)

    return leagues, err
}

func GetMasterLeagues(region string, queue string) (LeagueList, error) {
    args := createArgs("type", queue)
    url := createApiUrl(MASTER_LEAGUE, region) + fmt.Sprintf("?%v", args)

    var leagues LeagueList
    err := decodeRequest(region, url, &leagues)

    return leagues, err
}
