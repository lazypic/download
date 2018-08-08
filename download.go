package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	regionPtr := flag.String("region", "ap-northeast-2", "AWS region name.")
	bucketPtr := flag.String("bucket", "lazypic", "S3 bucket name.")
	keyPtr := flag.String("key", "", "key name")
	flag.Parse()
	if *keyPtr == "" {
		flag.PrintDefaults()
		return
	}
	svc := s3.New(session.New(), &aws.Config{Region: aws.String(*regionPtr)})
	params := &s3.ListObjectsInput{
		Bucket: aws.String(*bucketPtr),
		Prefix: aws.String(*keyPtr),
	}
	resp, _ := svc.ListObjects(params)
	for _, key := range resp.Contents {
		fmt.Println(*key.Key)
	}
}
