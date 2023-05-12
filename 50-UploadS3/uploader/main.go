package main

import (
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/flpgst/golang-studies/50-UploadS3/configs"
)

var (
	s3Client *s3.S3
	s3Bucket string
)

func init() {

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(configs.AWSRegion),
			Credentials: credentials.NewStaticCredentials(
				configs.AWSPublicKey,
				configs.AWSPrivateKey,
				"",
			),
		},
	)
	if err != nil {
		panic(err)
	}
	s3Client = s3.New(sess)
	s3Bucket = "filipe-goexpert-bucket"
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	for {
		files, err := dir.Readdir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %v\n", err)
			continue
		}
		uploadFile(files[0].Name())
	}

}

func uploadFile(filename string) {
	completeFileName := "./tmp/" + filename
	fmt.Printf("Pploading file: %s to bocket: %s\n", filename, s3Bucket)
	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", filename)
		return
	}
	defer f.Close()
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		fmt.Printf("Error uploading file: %s\n", filename)
		return
	}
	fmt.Printf("Uploaded %s\n", filename)
}
