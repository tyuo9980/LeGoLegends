package legolegends

import (
    "fmt"
)

type Match struct {
    SeasonId                int                       `json:"seasonId"`
    QueueId                 int                       `json:"queueId"`
    GameId                  int64                     `json:"gameId"`
    ParticipantIdentities   []ParticipantIdentity  `json:"participantIdentities"`
    GameVersion             string                    `json:"gameVersion"`
    PlatformId              string                    `json:"platformId"`
    GameMode                string                    `json:"gameMode"`
    MapId                   int                       `json:"mapId"`
    GameType                string                    `json:"gameType"`
    Teams                   []TeamStats            `json:"teams"`
    Participants            []Participant          `json:"participants"`
    GameCreation            int64                     `json:"gameCreation"`
    GameDuration            int64                     `json:"gameDuration"`
}

type ParticipantIdentity struct {
    ParticipantId int        `json:"participantId"`
    Player        Player  `json:"player"`
}

type Player struct {
    CurrentPlatformId string  `json:"currentPlatformId"`
    SummonerName      string  `json:"summonerName"`
    MatchHistoryURI   string  `json:"matchHistoryUri"`
    PlatformId        string  `json:"platformId"`
    CurrentAccountId  int64   `json:"currentAccountId"`
    ProfileIcon       int     `json:"profileIcon"`
    SummonerId        int64   `json:"summonerId"`
    AccountId         int64   `json:"accountId"`
}

type TeamStats struct {
    Bans                 []TeamBans    `json:"bans"`
    BaronKills           int              `json:"baronKills"`
    DominionVictoryScore int              `json:"dominionVictoryScore"`
    DragonKills          int              `json:"dragonKills"`
    FirstBaron           bool             `json:"firstBaron"`
    FirstBlood           bool             `json:"firstBlood"`
    FirstDragon          bool             `json:"firstDragon"`
    FirstInhibitor       bool             `json:"firstInhibitor"`
    FirstRiftHerald      bool             `json:"firstRiftHerald"`
    FirstTower           bool             `json:"firstTower"`
    InhibitorKills       int              `json:"inhibitorKills"`
    RiftHeraldKills      int              `json:"riftHeraldKills"`
    TeamId               int              `json:"teamId"`
    TowerKills           int              `json:"towerKills"`
    VilemawKills         int              `json:"vilemawKills"`
    Win                  string           `json:"win"`
}

type TeamBans struct {
    ChampionId int `json:"championId"`
    PickTurn   int `json:"pickTurn"`
}

type Participant struct {
    ChampionId                int                    `json:"championId"`
    HighestAchievedSeasonTier string                 `json:"highestAchievedSeasonTier"`
    Masteries                 []Mastery           `json:"masteries"`
    ParticipantId             int                    `json:"participantId"`
    Runes                     []Rune              `json:"runes"`
    Spell1Id                  int                    `json:"spell1Id"`
    Spell2Id                  int                    `json:"spell2Id"`
    Stats                     ParticipantStats    `json:"stats"`
    TeamId                    int                    `json:"teamId"`
    Timeline                  ParticipantTimeline `json:"timeline"`
}

type ParticipantStats struct {
    Assists                         int   `json:"assists"`
    AltarsCaptured                  int   `json:"altarsCaptured"`
    AltarsNeutralized               int   `json:"altarsNeutralized"`
    ChampLevel                      int   `json:"champLevel"`
    CombatPlayerScore               int   `json:"combatPlayerScore"`
    DamageDealtToObjectives         int64 `json:"damageDealtToObjectives"`
    DamageDealtToTurrets            int64 `json:"damageDealtToTurrets"`
    DamageSelfMitigated             int64 `json:"damageSelfMitigated"`
    Deaths                          int   `json:"deaths"`
    DoubleKills                     int   `json:"doubleKills"`
    FirstBloodAssist                bool  `json:"firstBloodAssist"`
    FirstBloodKill                  bool  `json:"firstBloodKill"`
    FirstInhibitorAssist            bool  `json:"firstInhibitorAssist"`
    FirstInhibitorKill              bool  `json:"firstInhibitorKill"`
    FirstTowerAssist                bool  `json:"firstTowerAssist"`
    FirstTowerKill                  bool  `json:"firstTowerKill"`
    GoldEarned                      int   `json:"goldEarned"`
    GoldSpent                       int   `json:"goldSpent"`
    InhibitorKills                  int   `json:"inhibitorKills"`
    Item0                           int   `json:"item0"`
    Item1                           int   `json:"item1"`
    Item2                           int   `json:"item2"`
    Item3                           int   `json:"item3"`
    Item4                           int   `json:"item4"`
    Item5                           int   `json:"item5"`
    Item6                           int   `json:"item6"`
    KillingSprees                   int   `json:"killingSprees"`
    Kills                           int   `json:"kills"`
    LargestCriticalStrike           int   `json:"largestCriticalStrike"`
    LargestKillingSpree             int   `json:"largestKillingSpree"`
    LargestMultiKill                int   `json:"largestMultiKill"`
    LongestTimeSpentLiving          int   `json:"longestTimeSpentLiving"`
    MagicDamageDealt                int64 `json:"magicDamageDealt"`
    MagicDamageDealtToChampions     int64 `json:"magicDamageDealtToChampions"`
    MagicalDamageTaken              int64 `json:"magicalDamageTaken"`
    NeutralMinionsKilled            int   `json:"neutralMinionsKilled"`
    NeutralMinionsKilledEnemyJungle int   `json:"neutralMinionsKilledEnemyJungle"`
    NeutralMinionsKilledTeamJungle  int   `json:"neutralMinionsKilledTeamJungle"`
    NodeCapture                     int   `json:"nodeCapture"`
    NodeCaptureAssist               int   `json:"nodeCaptureAssist"`
    NodeNeutralize                  int   `json:"nodeNeutralize"`
    NodeNeutralizeAssist            int   `json:"nodeNeutralizeAssist"`
    ObjectivePlayerScore            int   `json:"objectivePlayerScore"`
    ParticipantId                   int   `json:"participantId"`
    PentaKills                      int   `json:"pentaKills"`
    PhysicalDamageDealt             int64 `json:"physicalDamageDealt"`
    PhysicalDamageDealtToChampions  int64 `json:"physicalDamageDealtToChampions"`
    PhysicalDamageTaken             int64 `json:"physicalDamageTaken"`
    QuadraKills                     int   `json:"quadraKills"`
    SightWardsBoughtInGame          int   `json:"sightWardsBoughtInGame"`
    TeamObjective                   int   `json:"teamObjective"`
    TimeCCingOthers                 int64 `json:"timeCCingOthers"`
    TotalDamageDealt                int64 `json:"totalDamageDealt"`
    TotalDamageDealtToChampions     int64 `json:"totalDamageDealtToChampions"`
    TotalDamageTaken                int64 `json:"totalDamageTaken"`
    TotalHeal                       int64 `json:"totalHeal"`
    TotalMinionsKilled              int   `json:"totalMinionsKilled"`
    TotalPlayerScore                int   `json:"totalPlayerScore"`
    TotalScoreRank                  int   `json:"totalScoreRank"`
    TotalTimeCrowdControlDealt      int   `json:"totalTimeCrowdControlDealt"`
    TotalUnitsHealed                int   `json:"totalUnitsHealed"`
    TurretKills                     int   `json:"turretKills"`
    TripleKills                     int   `json:"tripleKills"`
    TrueDamageDealt                 int64 `json:"trueDamageDealt"`
    TrueDamageDealtToChampions      int64 `json:"trueDamageDealtToChampions"`
    TrueDamageTaken                 int64 `json:"trueDamageTaken"`
    UnrealKills                     int   `json:"unrealKills"`
    VisionScore                     int64 `json:"visionScore"`
    VisionWardsBoughtInGame         int   `json:"visionWardsBoughtInGame"`
    WardsKilled                     int   `json:"wardsKilled"`
    WardsPlaced                     int64 `json:"wardsPlaced"`
    Win                             bool  `json:"win"`
}

type Rune struct {
    RuneId int `json:"runeId"`
    Rank   int `json:"rank"`
}

type ParticipantTimeline struct {
    CreepsPerMinDeltas            map[string]float64 `json:"creepsPerMinDeltas"`
    CsDiffPerMinDeltas            map[string]float64 `json:"csDiffPerMinDeltas"`
    DamageTakenDiffPerMinDeltas   map[string]float64 `json:"damageTakenDiffPerMinDeltas"`
    DamageTakenPerMinDeltas       map[string]float64 `json:"damageTakenPerMinDeltas"`
    GoldPerMinDeltas              map[string]float64 `json:"goldPerMinDeltas"`
    ParticipantId                 int                     `json:"participantId"`
    Lane                          string                  `json:"lane"`
    Role                          string                  `json:"role"`
    XpDiffPerMinDeltas            map[string]float64 `json:"xpDiffPerMinDeltas"`
    XpPerMinDeltas                map[string]float64 `json:"xpPerMinDeltas"`
}

type Mastery struct {
    MasteryId int `json:"masteryId"`
    Rank      int `json:"rank"`
}

type MatchList struct {
    Matches     []MatchReference `json:"matches"`
    TotalGames  int                 `json:"totalGames"`
    StartIndex  int                 `json:"startIndex"`
    EndIndex    int                 `json:"endIndex"`
}

type MatchReference struct {
    Lane        string  `json:"lane"`
    GameId      int64   `json:"gameId"`
    Champion    int     `json:"champion"`
    PlatformId  string  `json:"platformId"`
    Season      int     `json:"season"`
    Queue       int     `json:"queue"`
    Role        string  `json:"role"`
    Timestamp   int64   `json:"timestamp"`
}

type MatchTimeline struct {
    FrameInterval int64           `json:"frameInterval"`
    Frames        []MatchFrame `json:"frames"`
}

type MatchFrame struct {
    Events            []MatchEvent                     `json:"events"`
    ParticipantFrames map[string]MatchParticipantFrame `json:"participantFrames"`
    Timestamp         int64                               `json:"timestamp"`
}

type MatchParticipantFrame struct {
    CurrentGold         int              `json:"currentGold"`
    DominionScore       int              `json:"dominionScore"`
    JungleMinionsKilled int              `json:"jungleMinionsKilled"`
    Level               int              `json:"level"`
    MinionsKilled       int              `json:"minionsKilled"`
    ParticipantId       int              `json:"participantId"`
    Position            MatchPosition `json:"position"`
    TeamScore           int              `json:"teamScore"`
    TotalGold           int              `json:"totalGold"`
    Xp                  int              `json:"xp"`
}

type MatchPosition struct {
    X   int `json:"x"`
    Y   int `json:"y"`
}

type MatchEvent struct {
    AscendedType            string           `json:"ascendedType"`
    AssistingParticipantIds []int            `json:"assistingParticipantIds"`
    BuildingType            string           `json:"buildingType"`
    CreatorId               int              `json:"creatorId"`
    EventType               string           `json:"eventType"`
    AfterId                 int              `json:"afterId"`
    BeforeId                int              `json:"beforeId"`
    ItemId                  int              `json:"itemId"`
    KillerId                int              `json:"killerId"`
    LaneType                string           `json:"laneType"`
    LevelUpType             string           `json:"levelUpType"`
    MonsterType             string           `json:"monsterType"`
    MonsterSubType          string           `json:"monsterSubType"`
    ParticipantId           int              `json:"participantId"`
    PointCaptured           string           `json:"pointCaptured"`
    Position                MatchPosition `json:"position"`
    SkillSlot               int              `json:"skillSlot"`
    TeamId                  int              `json:"teamId"`
    Timestamp               int64            `json:"timestamp"`
    TowerType               string           `json:"towerType"`
    Type                    string           `json:"type"`
    VictimId                int              `json:"victimId"`
    WardType                string           `json:"wardType"`
}

type MatchList struct {
    TotalGames int              `json:"totalGames"`
    StartIndex int              `json:"startIndex"`
    EndIndex   int              `json:"endIndex"`
    Matches    []MatchReference `json:"matches"`
}

type MatchReference struct {
    Timestamp  int64  `json:"timestamp"`
    Champion   int    `json:"champion"`
    Queue      string `json:"queue"`
    Season     string `json:"season"`
    GameId     int64  `json:"gameId"`
    Role       string `json:"role"`
    PlatformId string `json:"platformId"`
    Lane       string `json:"lane"`
}

func GetMatchlist(
    region string,
    accountId int64,
    championIds []int,
    rankedQueues []int,
    seasons []int,
    beginTime int64,
    endTime int64,
    beginIndex int,
    endIndex int) (MatchList, error) {

    argString := "championIds,rankedQueues,seasons,beginTime,endTime,beginIndex,endIndex"
    args := createArgs(argString, championIds, rankedQueues, seasons, beginTime, endTime, beginIndex, endIndex)
    url := createApiUrl(MATCHLISTS, region, accountId) + fmt.Sprintf("?%v", args)

    var matchList MatchList
    err := decodeRequest(region, url, &matchList)

    return matchList, err
}

func GetMatch(region string, matchId int64) (Match, error) {
    url := createApiUrl(MATCH, region, matchId)

    var match Match
    err := decodeRequest(region, url, &match)

    return match, err
}

func GetMatchTimeline(region string, matchId int64) (MatchTimeline, error) {
    url := createApiUrl(TIMELINES, region, matchId)

    var matchTimeline MatchTimeline
    err := decodeRequest(region, url, &matchTimeline)

    return matchTimeline, err
}
