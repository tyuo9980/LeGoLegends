package legolegends

import (
	"fmt"
)

type ChampionMastery struct {
	ChampionId                   int64 `json:"championId"`
	ChampionLevel                int   `json:"championLevel`
	ChampionPoints               int   `json:"championPoints`
	ChampionPointsSinceLastLevel int64 `json:"championPointsSinceLastLevel`
	ChampionPointsUntilNextLevel int64 `json:"championPointsUntilNextLevel`
	ChestGranted                 bool  `json:"chestGranted`
	LastPlayTime                 int64 `json:"lastPlayTime`
	PlayerId                     int64 `json:"playerId`
}

func getChampionMastery(region string, summonerId int64, championId int64) (ChampionMastery, error) {
	url := createApiUrl(CHAMPION_MASTERY, region, fmt.Sprintf("%d/by-champion/%d", summonerId, championId))

	var championMastery ChampionMastery
	err := decodeRequest(region, url, &championMastery)

	return championMastery, err
}

func getChampionMasteries(region string, summonerId int64) ([]ChampionMastery, error) {
	url := createApiUrl(CHAMPION_MASTERY, region, summonerId)

	var championMastery []ChampionMastery
	err := decodeRequest(region, url, &championMastery)

	return championMastery, err
}
