package legolegends

import (
	"fmt"
)

type ChampionMasteryDto struct {
	ChampionId                   int64  `json:"championId"`
	ChampionLevel                int    `json:"championLevel`
	ChampionPoints               int    `json:"championPoints`
	ChampionPointsSinceLastLevel int64  `json:"championPointsSinceLastLevel`
	ChampionPointsUntilNextLevel int64  `json:"championPointsUntilNextLevel`
	ChestGranted                 bool   `json:"chestGranted`
	HighestGrade                 string `json:"highestGrade`
	LastPlayTime                 int64  `json:"lastPlayTime`
	PlayerId                     int64  `json:"playerId`
}

func getChampionMastery(region string, summonerId int64, championId int64) (ChampionMastery, error) {
	args := createArgs("")
	url := createChampMasteryUrl(CHAMPION_MASTERY, region) + fmt.Sprintf("%d/champion/%d?%v", summonerId, championId, args)

	var championMastery ChampionMastery
	err := decodeRequest(region, url, &championMastery)

	return championMastery, err
}

func getChampionMasteries(region string, summonerId int64) ([]ChampionMastery, error) {
	args := createArgs("")
	url := createChampMasteryUrl(CHAMPION_MASTERY, region) + fmt.Sprintf("%d/champions?%v", summonerId, args)

	var championMastery []ChampionMastery
	err := decodeRequest(region, url, &championMastery)

	return championMastery, err
}

func getSpecificChampionMastery(region string, summonerId int64, count int) ([]ChampionMastery, error) {
	args := createArgs("count", count)
	url := createChampMasteryUrl(CHAMPION_MASTERY, region) + fmt.Sprintf("%d/topchampions?%v", summonerId, args)

	var championMastery []ChampionMastery
	err := decodeRequest(region, url, &championMastery)

	return championMastery, err
}
