package main

import (
	"encoding/csv"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mizukmb/serverless-iidx-csv-parser/iidx"
	"os"
)

func handler() (iidx.Iidx, error) {
	file, _ := os.Open("./iidx24_sinobuz_score.csv")
	defer file.Close()

	r := csv.NewReader(file)

	// Read header line (not use)
	r.Read()

	record, _ := r.Read()

	return iidx.NewIidx(record), nil
}

func main() {
	lambda.Start(handler)
}
