package iidx

import (
	"strconv"
	"strings"
	"time"
)

type Iidx struct {
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

func NewIidx(record []string) Iidx {
	iidx := Iidx{}

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

func (p Iidx) ScrapboxTitle(difficulty string) string {
	return p.Title + " (" + difficultyShort(difficulty) + ")"
}

func (p Iidx) ScrapboxArticle(difficulty string) []string {
	pD := p.difficulty(difficulty)
	var ret []string

	ret = append(ret, p.Title+" ("+difficultyShort(difficulty)+")")
	ret = append(ret, "")
	ret = append(ret, "バージョン: ["+p.Version+"]")
	ret = append(ret, "ジャンル: ["+p.Genre+"]")
	ret = append(ret, "アーティスト: ["+p.Artist+"]")
	ret = append(ret, "プレイ回数: "+strconv.Itoa(p.PlayCount))
	ret = append(ret, "")

	ret = append(ret, " [LEVEL: "+strconv.Itoa(pD.Level)+"]")
	ret = append(ret, " EXSCORE: "+strconv.Itoa(pD.ExScore))
	ret = append(ret, " PGREAT: "+strconv.Itoa(pD.PGreat))
	ret = append(ret, " GREAT: "+strconv.Itoa(pD.Great))
	ret = append(ret, " MISS: "+strconv.Itoa(pD.Miss))
	if pD.ClearType == "NO PLAY" {
		ret = append(ret, " CLEARTYPE: "+pD.ClearType)
	} else {
		ret = append(ret, " [CLEARTYPE: "+pD.ClearType+"]")
	}
	if pD.DjLevel == "---" {
		ret = append(ret, " DJLEVEL: "+pD.DjLevel)
	} else {
		ret = append(ret, " [DJLEVEL: "+pD.DjLevel+"]")
	}

	return ret
}

func (p *Iidx) difficulty(d string) difficulty {
	var ret difficulty
	switch strings.ToLower(d) {
	case "normal":
		ret = p.Normal
	case "hyper":
		ret = p.Hyper
	case "another":
		ret = p.Another
	}

	return ret
}

func difficultyShort(d string) string {
	var ret string

	switch strings.ToLower(d) {
	case "normal":
		ret = "N"
	case "hyper":
		ret = "H"
	case "another":
		ret = "A"
	}

	return ret
}
