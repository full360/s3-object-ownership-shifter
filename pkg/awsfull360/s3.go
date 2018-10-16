package awsfull360

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3 struct {
	client *s3.S3
}

// Create New S3 Client
func NewS3Client() (*S3, error) {
	sess := session.Must(session.NewSession())

	return &S3{
		client: s3.New(sess),
	}, nil
}

func (svc *S3) PutObjectACL(bucket string, key string, fullControlUser string) (*s3.PutObjectAclOutput, error) {

	input := s3.PutObjectAclInput{
		Bucket:           &bucket,
		Key:              &key,
		GrantFullControl: &fullControlUser}

	fmt.Println("the input")
	fmt.Println(input)
	return svc.client.PutObjectAcl()
}

func (svc *S3) CopyObjectToBucket(targetBucket string, srcBucket string, key string) (*s3.CopyObjectOutput, error) {
	source := srcBucket + "/" + key
	input := s3.CopyObjectInput{
		Bucket:     &targetBucket,
		CopySource: &source,
		Key:        &key}

	return svc.client.CopyObject(&input)
}
