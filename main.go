package main

import (
	"encoding/csv"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
	"strconv"
	"time"
)

type iidx struct {
	Version        string     `json:"version"`
	Title          string     `json:"title"`
	Genre          string     `json:"genre"`
	Artist         string     `json:"artist"`
	PlayCount      int        `json:"playcount"`
	Normal         difficulty `json:"normal"`
	Hyper          difficulty `json:"hyper"`
	Another        difficulty `json:"another"`
	LastPlayedDate time.Time  `json:"lastplayeddate"`
}

type difficulty struct {
	Level     int    `json:"level"`
	ExScore   int    `json:"exscore"`
	PGreat    int    `json:"pgreat"`
	Great     int    `json:"great"`
	Miss      int    `json:"miss"`
	ClearType string `json:"cleartype"`
	DjLevel   string `json:"djlevel"`
}

func newIidx(record []string) iidx {
	iidx := iidx{}

	iidx.Version = record[0]
	iidx.Title = record[1]
	iidx.Genre = record[2]
	iidx.Artist = record[3]

	var err error

	iidx.PlayCount, err = strconv.Atoi(record[4])
	if err != nil {
		panic("Parse ERROR")
	}

	// normal
	nLevel, _ := strconv.Atoi(record[5])
	nExScore, _ := strconv.Atoi(record[6])
	nPGreat, _ := strconv.Atoi(record[7])
	nGreat, _ := strconv.Atoi(record[8])
	nMiss, _ := strconv.Atoi(record[9])
	nClearType := record[10]
	nDjLevel := record[11]

	iidx.Normal = difficulty{
		Level:     nLevel,
		ExScore:   nExScore,
		PGreat:    nPGreat,
		Great:     nGreat,
		Miss:      nMiss,
		ClearType: nClearType,
		DjLevel:   nDjLevel,
	}

	// hyper
	hLevel, _ := strconv.Atoi(record[12])
	hExScore, _ := strconv.Atoi(record[13])
	hPGreat, _ := strconv.Atoi(record[14])
	hGreat, _ := strconv.Atoi(record[15])
	hMiss, _ := strconv.Atoi(record[16])
	hClearType := record[17]
	hDjLevel := record[18]

	iidx.Hyper = difficulty{
		Level:     hLevel,
		ExScore:   hExScore,
		PGreat:    hPGreat,
		Great:     hGreat,
		Miss:      hMiss,
		ClearType: hClearType,
		DjLevel:   hDjLevel,
	}

	// another
	aLevel, _ := strconv.Atoi(record[19])
	aExScore, _ := strconv.Atoi(record[20])
	aPGreat, _ := strconv.Atoi(record[21])
	aGreat, _ := strconv.Atoi(record[22])
	aMiss, _ := strconv.Atoi(record[23])
	aClearType := record[24]
	aDjLevel := record[25]

	iidx.Another = difficulty{
		Level:     aLevel,
		ExScore:   aExScore,
		PGreat:    aPGreat,
		Great:     aGreat,
		Miss:      aMiss,
		ClearType: aClearType,
		DjLevel:   aDjLevel,
	}

	err = nil
	layout := "2006-01-02 15:04"
	iidx.LastPlayedDate, err = time.Parse(layout, record[26])
	if err != nil {
		panic("Parse ERROR")
	}

	return iidx
}

func handler() (iidx, error) {
	file, _ := os.Open("./iidx24_sinobuz_score.csv")
	defer file.Close()

	r := csv.NewReader(file)

	// Read header line (not use)
	r.Read()

	record, _ := r.Read()

	return newIidx(record), nil
}

func main() {
	lambda.Start(handler)
}
