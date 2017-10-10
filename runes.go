package legolegends

type RunePages struct {
    Pages      []RunePage `json:"pages"`
    SummonerId int64      `json:"summonerId"`
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

func GetSummonerRunes(region string, summonerId int64) (RunePages, error) {
    url := createApiUrl(RUNES, region, summonerId)

    var runes RunePages
    err := decodeRequest(region, url, &runes)

    return runes, err
}
