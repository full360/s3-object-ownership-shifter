package main //

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func requestHandler(S3Event events.S3Event) error {

	return nil
}

func main() {
	fmt.Println("Starting Function")
	lambda.Start(requestHandler)
}
