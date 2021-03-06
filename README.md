# LeGoLegends
Go library for Riot's League of Legends API

Contains support for:

| Endpoint        | Version |
| --------------- | ------- |
| Match           | v2.2    |
| MatchList       | v2.2    |
| Summoner        | v1.4    |
| Game            | v1.3    |
| Stats           | v1.3    |
| CurrentGame     | v1.0    |
| League          | v2.5    |
| ChampionMastery |         |

Install using:
```
go get github.com/tyuo9980/legolegends
```

# Getting started
Before using, make sure an api key and rate limit is set.
```
legolegends.SetApiKey(key)
legolegends.SetRateLimit(rate, delay)
```

Debug mode can be used to track call execution and errors.
```
legolegends.SetDebug(flag)
```
