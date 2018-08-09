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
	prefixPtr := flag.String("prefix", "", "prefix name")
	flag.Parse()
	if *prefixPtr == "" {
		flag.PrintDefaults()
		return
	}
	svc := s3.New(session.New(), &aws.Config{Region: aws.String(*regionPtr)})
	params := &s3.ListObjectsInput{
		Bucket: aws.String(*bucketPtr),
		Prefix: aws.String(*prefixPtr),
	}
	resp, _ := svc.ListObjects(params)
	if len(resp.Contents) == 0 {
		fmt.Println("no data to download.")
		return
	}
	for _, key := range resp.Contents {
		fmt.Println(*key.Key)
	}
}
