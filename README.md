# UPDATE

Starting from [version 1.26](https://github.com/serverless/serverless/releases/tag/v1.26.0) Serverless Framework includes two Golang templates:

* `aws-go` - basic template with two functions
* `aws-go-dep` - **recommended** template using [`dep`](https://github.com/golang/dep) package manager

You can use them with `create` command:

```
serverless create -t aws-go-dep
```

Original README below.

---

# Serverless Template for Golang

This repository contains template for creating serverless services written in Golang.

## Quick Start

1. Create a new service based on this template

```
serverless create -u https://github.com/serverless/serverless-golang/ -p myservice
```

2. Compile function

```
cd myservice
GOOS=linux go build -o bin/main
```

3. Deploy!

```
serverless deploy
```

## Example

**You need csv file on root directory if you run it.**

<details><summary>run & result</summary>

```
$ serverless invoke -f parse
{
    "version": "1st&substream",
    "title": "GRADIUSIC CYBER",
    "genre": "DIGI-ROCK",
    "artist": "TAKA",
    "playcount": 1,
    "normal": {
        "level": 5,
        "exscore": 0,
        "pgreat": 0,
        "great": 0,
        "miss": 0,
        "cleartype": "NO PLAY",
        "djlevel": "---"
    },
    "hyper": {
        "level": 6,
        "exscore": 0,
        "pgreat": 0,
        "great": 0,
        "miss": 0,
        "cleartype": "NO PLAY",
        "djlevel": "---"
    },
    "another": {
        "level": 7,
        "exscore": 697,
        "pgreat": 260,
        "great": 177,
        "miss": 27,
        "cleartype": "CLEAR",
        "djlevel": "B"
    },
    "lastplayeddate": "2017-10-12T13:24:00Z"
}
```

</details>
