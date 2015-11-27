package legolegends

import (
	"fmt"
)

type Summoner struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	ProfileIconId int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}

type MasteryBook struct {
	Pages      []MasteryPage `json:"pages"`
	SummonerId int64         `json:"summonerId"`
}

type MasteryPage struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Current   bool      `json:"current"`
	Masteries []Mastery `json:"masteries"`
}

type Mastery struct {
	Id   int `json:"id"`
	Rank int `json:"rank"`
}

type RuneBook struct {
	Pages      []RunePage `json:"pages"`
	SummonerId int        `json:"summonerId"`
}

type RunePage struct {
	Current bool       `json:"current"`
	Id      int64      `json:"id"`
	Name    string     `json:"name"`
	Slots   []RuneSlot `json:"slots"`
}

type RuneSlot struct {
	RuneId     int `json:"runeId"`
	RuneSlotId int `json:"runeSlotId"`
}

func getSummonerByNames(names ...string) map[string]Summoner {
	url := fmt.Sprintf("https://%v.%v/api/lol/%v/v1.4/summoner/by-name/%v?%v", Region, BASE, Region, nameString, API_STR)

	summoners := make(map[string]Summoner)
	reqBody := SendRequest(url)
	summoners = DecodeRequest(reqBody, &summoners)

	return summoners
}

func getSummonerByIds() {

}

func getSummonerMasteries() {

}

func getSummonerRunes() {

}

func getSummonerNames() {

}
