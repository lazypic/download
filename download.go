package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	regionPtr := flag.String("region", "ap-northeast-2", "AWS region name.")
	bucketPtr := flag.String("bucket", "lazypic", "S3 bucket name.")
	keyPtr := flag.String("key", "", "key name")
	subdirPtr := flag.Bool("subdir", false, "sub directory fuction.")
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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Download?(y/n): ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		return
	}
	text = strings.Replace(text, "\n", "", -1)
	if strings.ToLower(text) != "y" {
		fmt.Println("Download Cancel")
		return
	}
	for _, key := range resp.Contents {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error()+"\n")
			continue
		}
		keyStr := *key.Key
		if !*subdirPtr {
			keyStr = strings.Replace(keyStr, "/", "_", -1)
		}
		path := dir + "/" + keyStr
		err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error()+"\n")
			continue
		}
		file, err := os.Create(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error()+"\n")
			continue
		}
		defer file.Close()
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String(*regionPtr)},
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error()+"\n")
			continue
		}
		downloader := s3manager.NewDownloader(sess)
		numBytes, err := downloader.Download(file,
			&s3.GetObjectInput{
				Bucket: aws.String(*bucketPtr),
				Key:    aws.String(*key.Key),
			})
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error()+"\n")
			continue
		}
		fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
	}
}
