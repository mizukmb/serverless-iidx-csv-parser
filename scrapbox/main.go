package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mizukmb/serverless-iidx-csv-parser/iidx"
	"github.com/mizukmb/serverless-iidx-csv-parser/scrapbox/scrapbox"
	"io"
	"os"
	"strconv"
)

func articleText(iidx iidx.Iidx) []string {
	return []string{
		iidx.Title,
		"",
		"[" + iidx.Version + "]",
		"[" + iidx.Genre + "]",
		"[" + iidx.Artist + "]",
		"プレイ回数: " + strconv.Itoa(iidx.PlayCount),
		"",
		"[** NORMAL]",
		" [LEVEL: " + strconv.Itoa(iidx.Normal.Level) + "]",
		" EXSCORE: " + strconv.Itoa(iidx.Normal.ExScore),
		" PGREAT: " + strconv.Itoa(iidx.Normal.PGreat),
		" GREAT: " + strconv.Itoa(iidx.Normal.Great),
		" MISS: " + strconv.Itoa(iidx.Normal.Miss),
		" CLEARTYPE: " + iidx.Normal.ClearType,
		" DJLEVEL: " + iidx.Normal.DjLevel,
		"",
		"[** HYPER]",
		" [LEVEL: " + strconv.Itoa(iidx.Hyper.Level) + "]",
		" EXSCORE: " + strconv.Itoa(iidx.Hyper.ExScore),
		" PGREAT: " + strconv.Itoa(iidx.Hyper.PGreat),
		" GREAT: " + strconv.Itoa(iidx.Hyper.Great),
		" MISS: " + strconv.Itoa(iidx.Hyper.Miss),
		" CLEARTYPE: " + iidx.Hyper.ClearType,
		" DJLEVEL: " + iidx.Hyper.DjLevel,
		"",
		"[** ANOTHER]",
		" [LEVEL: " + strconv.Itoa(iidx.Another.Level) + "]",
		" EXSCORE: " + strconv.Itoa(iidx.Another.ExScore),
		" PGREAT: " + strconv.Itoa(iidx.Another.PGreat),
		" GREAT: " + strconv.Itoa(iidx.Another.Great),
		" MISS: " + strconv.Itoa(iidx.Another.Miss),
		" CLEARTYPE: " + iidx.Another.ClearType,
		" DJLEVEL: " + iidx.Another.DjLevel,
	}
}

func printScrapbox(j []byte) {
	fmt.Println(string(j))
}

func handler() (events.APIGatewayProxyResponse, error) {
	file, _ := os.Open("./iidx24_sinobuz_score.csv")
	defer file.Close()

	r := csv.NewReader(file)
	r.LazyQuotes = true

	// Read header line (not use)
	r.Read()

	var articles []scrapbox.Article

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			// 読み込みエラー発生
			fmt.Println("Read error: ", err)
			break
		}

		iidx := iidx.NewIidx(record)

		title := iidx.Title
		article := scrapbox.NewArticle(title, articleText(iidx))

		articles = append(articles, article)
	}

	scrapbox := scrapbox.NewScrapbox(articles)
	j, _ := json.Marshal(scrapbox)
	// for debug
	// printScrapbox(j)

	return events.APIGatewayProxyResponse{Body: string(j), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}
