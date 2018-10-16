package main //gitlab.full360.com/experiments/s3-copier

import (
	"fmt"
	"log"
	"os"

	awsF360 "gitlab.full360.com/experiments/s3-copier/pkg/awsfull360"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func requestHandler(S3Event events.S3Event) error {
	targetBucket := os.Getenv("TARGET_S3_BUCKET")
	grantFullControl := os.Getenv("OWNERSHIP_FULL_CONTROL")

	s3, err := awsF360.NewS3Client()
	if err != nil {
		log.Fatal("Error Getting S3 Client: ", err)
	}

	for _, rec := range S3Event.Records {
		_, err := s3.PutObjectACL(rec.S3.Bucket.Name, rec.S3.Object.Key, grantFullControl)
		if err != nil {
			log.Fatal("Error Add ACL: ", err)
		} else {
			_, errCP := s3.CopyObjectToBucket(targetBucket, rec.S3.Bucket.Name, rec.S3.Object.Key)
			if errCP != nil {
				log.Fatal("Error Copying file ", errCP)
			}
		}
	}

	fmt.Println("Done")
	return nil
}

func main() {
	fmt.Println("Starting Function")
	lambda.Start(requestHandler)
}
