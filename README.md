# LeGoLegends
Go library for Riot's League of Legends API

Contains support for:
Endpoint  | Version
--------- | -------
Match     | v2.2
MatchList | v2.2
Summoner  | v1.4

| Tables        | Are           | Cool  |
| ------------- |:-------------:| -----:|
| col 3 is      | right-aligned | $1600 |
| col 2 is      | centered      |   $12 |
| zebra stripes | are neat      |    $1 |

Import using:
```
go get github.com/tyuo9980/legolegends
```

# Getting started:
Before using, make sure an api key and region is set.
```
legolegends.SetApiKey(key)
legolegends.SetRegion(region)
```

Debug mode is available by setting Debug to true
```
legolegends.SetDebug(bool)
```
