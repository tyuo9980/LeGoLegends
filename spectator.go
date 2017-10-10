package legolegends

import (
    "fmt"
)

type CurrentGameInfo struct {
    BannedChampions   []BannedChampion `json:"bannedChampion"`
    GameId            int64                       `json:"gameId"`
    GameLength        int64                       `json:"gameLength"`
    GameMode          string                      `json:"gameMode"`
    GameQueueConfigId int64                       `json:"gameQueueConfigId"`
    GameStartTime     int64                       `json:"gameStartTime"`
    GameType          string                      `json:"gameType"`
    MapId             int64                       `json:"mapId"`
    Observers         Observer                    `json:"observers"`
    Participants      []CurrentGameParticipant    `json:"participants"`
    PlatformId        string                      `json:"platformId"`
}

type BannedChampion struct {
    ChampionId int64 `json:"championId"`
    PickTurn   int   `json:"pickTurn"`
    TeamId     int64 `json:"teamId"`
}

type CurrentGameParticipant struct {
    Bot           bool                 `json:"bot"`
    ChampionId    int64                `json:"championId"`
    Masteries     []CurrentGameMastery `json:"masteries"`
    ProfileIconId int64                `json:"profileIconId"`
    Runes         []CurrentGameRune    `json:"runes"`
    Spell1Id      int64                `json:"spell1Id"`
    Spell2Id      int64                `json:"spell2Id"`
    SummonerId    int64                `json:"summonerId"`
    SummonerName  string               `json:"summonerName"`
    TeamId        int64                `json:"teamId"`
}

type Observer struct {
    EncryptionKey string `json:"encryptionKey"`
}

type CurrentGameMastery struct {
    MasteryId int64 `json:"masteryId"`
    Rank      int   `json:"rank"`
}

type CurrentGameRune struct {
    Count  int   `json:"count"`
    RuneId int64 `json:"runeId"`
}

func GetActiveGame(region string, summonerId int64) (CurrentGameInfo, error) {
    url := createApiUrl(SPECTATOR_ACTIVE, region, summonerId)

    var currentGameInfo CurrentGameInfo
    err := decodeRequest(region, url, &currentGameInfo)

    return currentGameInfo, err
}
