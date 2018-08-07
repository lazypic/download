package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	svc := s3.New(session.New(), &aws.Config{Region: aws.String("ap-northeast-2")})
	params := &s3.ListObjectsInput{
		Bucket: aws.String("lazypic"),
		Prefix: aws.String("mamma/"),
	}
	resp, _ := svc.ListObjects(params)
	for _, key := range resp.Contents {
		fmt.Println(*key.Key)
	}
}
