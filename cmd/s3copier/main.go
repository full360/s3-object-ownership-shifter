package main //github.com/full360/s3-object-ownership-shifter

import (
	"fmt"
	"log"
	"os"
	"strings"

	awsF360 "github.com/full360/s3-object-ownership-shifter/pkg/awsfull360"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func filterFile(fileFilter string, fileName string) bool {
	if len(fileFilter) == 0 {
		return true
	}

	return strings.Contains(fileName, fileFilter)
}

func requestHandler(S3Event events.S3Event) error {
	targetBucket := os.Getenv("TARGET_S3_BUCKET")
	grantFullControl := os.Getenv("OWNERSHIP_FULL_CONTROL")
	fileFilter := os.Getenv("FILE_FILTER")

	s3, err := awsF360.NewS3Client()
	if err != nil {
		log.Fatal("Error Getting S3 Client: ", err)
	}

	for _, rec := range S3Event.Records {
		fmt.Println("Processing file ", rec.S3.Object.Key)
		if filterFile(fileFilter, rec.S3.Object.Key) {
			_, err := s3.CopyObjectToBucket(targetBucket, rec.S3.Bucket.Name, rec.S3.Object.Key, grantFullControl)
			if err != nil {
				log.Fatal("Error Copying file ", err)
			}
		} else {
			log.Println("S3 Key does not part of the filter: ", rec.S3.Object.Key)
		}
	}

	fmt.Println("Done")
	return nil
}

func main() {
	lambda.Start(requestHandler)
}
