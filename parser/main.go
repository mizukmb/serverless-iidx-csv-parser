package main

import (
	"encoding/csv"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mizukmb/serverless-iidx-csv-parser/iidx"
	"os"
)

func handler() (events.APIGatewayProxyResponse, error) {
	file, _ := os.Open("./iidx24_sinobuz_score.csv")
	defer file.Close()

	r := csv.NewReader(file)

	// Read header line (not use)
	r.Read()

	record, _ := r.Read()
	j, _ := json.Marshal(iidx.NewIidx(record))

	return events.APIGatewayProxyResponse{Body: string(j), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}
