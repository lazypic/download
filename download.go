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
	projectPtr := flag.String("project", "", "project name")
	episodePtr := flag.Int("ep", 0, "episode number")
	scenePtr := flag.Int("s", 0, "scene number")
	cutPtr := flag.Int("c", 0, "cut number")
	flag.Parse()
	if *projectPtr == "" {
		flag.PrintDefaults()
		return
	}
	key := *projectPtr
	if *episodePtr != 0 && *scenePtr == 0 && *cutPtr == 0 {
		key = fmt.Sprintf("%s/%d", key, *episodePtr)
	}
	if *episodePtr != 0 && *scenePtr != 0 && *cutPtr == 0 {
		key = fmt.Sprintf("%s/%d/%d", key, *episodePtr, *scenePtr)
	}
	if *episodePtr != 0 && *scenePtr != 0 && *cutPtr != 0 {
		key = fmt.Sprintf("%s/%d/%d/%d", key, *episodePtr, *scenePtr, *cutPtr)
	}

	svc := s3.New(session.New(), &aws.Config{Region: aws.String(*regionPtr)})
	params := &s3.ListObjectsInput{
		Bucket: aws.String(*bucketPtr),
		Prefix: aws.String(key),
	}
	resp, _ := svc.ListObjects(params)
	for _, key := range resp.Contents {
		fmt.Println(*key.Key)
	}
}
