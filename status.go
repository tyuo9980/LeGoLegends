package legolegends

import (
    "fmt"
)

type ShardStatus struct {
    SummonerId int64      `json:"summonerId"`

    Name        string      `json:"name"`
    RegionTag   string      `json:"region_tag"`
    Hostname    string      `json:"hostname"`
    Services    []Service   `json:"services"`
    Slug        string      `json:"slug"`
    Locales     []string    `json:"locales"`
}

type Service struct {
    Status      string      `json:"status"`
    Incidents   []Incident  `json:"incidents"`
    Name        string      `json:"name"`
    Slug        string      `json:"slug"`
}

type Incident struct {
    Active      bool        `json:"active"`
    CreatedAt   string      `json:"created_at"`
    Id          int64       `json:"id"`
    Updates     []Message   `json:"updates"`
}

type Message struct {
    Severity        string          `json:"severity"`
    Author          string          `json:"author"`
    CreatedAt       string          `json:"created_at"`
    Translations    []Translation   `json:"translations"`
    UpdatedAt       string          `json:"updated_at"`
    Content         string          `json:"content"`
    Id              string          `json:"id"`
}

type Translation struct {
    Locale          string      `json:"locale"`
    Content         string  `json:"content"`
    UpdatedAt       string  `json:"updated_at"`
}

func GetSummonerRunes(region string) (ShardStatus, error) {
    url := createApiUrl(STATUS, region)

    var status ShardStatus
    err := decodeRequest(region, url, &status)

    return runes, err
}
