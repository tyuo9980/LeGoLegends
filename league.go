package legolegends

import (
    "fmt"
    "strings"
)

type League struct {
    Entries       []LeagueEntry `json:"entries"`
    Name          string        `json:"name"`
    ParticipantId string        `json:"participantId"`
    Queue         string        `json:"queue"`
    Tier          string        `json:"tier"`
}

type LeagueEntry struct {
    Division         string `json:"division"`
    IsFreshBlood     bool   `json:"isFreshBlood"`
    IsHotStreak      bool   `json:"isHotStreak"`
    IsInactive       bool   `json:"isInactive"`
    IsVeteran        bool   `json:"isVeteran"`
    LeaguePoints     int    `json:"leaguePoints"`
    Losses           int    `json:"losses"`
    MiniSeries       Series `json:"miniSeries"`
    PlayerOrTeamId   string `json:"playerOrTeamId"`
    PlayerOrTeamName string `json:"playerOrTeamName"`
    Wins             int    `json:"wins"`
}

type Series struct {
    Losses   int    `json:"losses"`
    Progress string `json:"progress"`
    Target   int    `json:"target"`
    Wins     int    `json:"wins"`
}

func GetLeagues(region string, ids ...string) (map[string][]League, error) {
    nameString := strings.Join([]string(ids), ",")

    args := createArgs("")
    url := createApiUrl(LEAGUE_BY_SUMMONER, region) + fmt.Sprintf("%v?%v", nameString, args)

    leagues := make(map[string][]League)
    err := decodeRequest(region, url, &leagues)

    return leagues, err
}

func GetLeagueEntries(region string, ids ...string) (map[string][]League, error) {
    nameString := strings.Join([]string(ids), ",")

    args := createArgs("")
    url := createApiUrl(LEAGUE_BY_SUMMONER, region) + fmt.Sprintf("%v/entry?%v", nameString, args)

    leagues := make(map[string][]League)
    err := decodeRequest(region, url, &leagues)

    return leagues, err
}

func GetTeamLeagues(region string, teamids ...string) (map[string][]League, error) {
    nameString := strings.Join([]string(teamids), ",")

    args := createArgs("")
    url := createApiUrl(LEAGUE_BY_TEAM, region) + fmt.Sprintf("%v?%v", nameString, args)

    leagues := make(map[string][]League)
    err := decodeRequest(region, url, &leagues)

    return leagues, err
}

func GetTeamLeagueEntries(region string, teamids ...string) (map[string][]League, error) {
    nameString := strings.Join([]string(teamids), ",")

    args := createArgs("")
    url := createApiUrl(LEAGUE_BY_TEAM, region) + fmt.Sprintf("%v/entry?%v", nameString, args)

    leagues := make(map[string][]League)
    err := decodeRequest(region, url, &leagues)

    return leagues, err
}

func GetChallengerLeagues(region string, rankType string) (League, error) {
    args := createArgs("type", rankType)
    url := createApiUrl(LEAGUE_CHALLENGER, region) + fmt.Sprintf("?%v", args)

    var leagues League
    err := decodeRequest(region, url, &leagues)

    return leagues, err
}

func GetMasterLeagues(region string, rankType string) (League, error) {
    args := createArgs("type", rankType)
    url := createApiUrl(LEAGUE_MASTER, region) + fmt.Sprintf("?%v", args)

    var leagues League
    err := decodeRequest(region, url, &leagues)

    return leagues, err
}
