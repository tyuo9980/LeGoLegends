# LeGoLegends
Go library for Riot's League of Legends API

Contains support for:

| Endpoint  | Version |
| --------- | ------- |
| Match     | v2.2    |
| MatchList | v2.2    |
| Summoner  | v1.4    |
| Game      | v1.3    |
| Stats     | v1.3    |

Install using:
```
go get github.com/tyuo9980/legolegends
```

# Getting started
Before using, make sure an api key and region is set.
```
legolegends.SetApiKey(key)
legolegends.SetRegion(region)
```

Debug mode is available by setting Debug to true
```
legolegends.SetDebug(flag)
```
