package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}

func PostHostConfiguration(hostNames []byte) error {

	var hostInstances HostInstances
	err := json.Unmarshal(hostNames, &hostInstances)
	if err != nil {
		fmt.Println("ERROR WITH JSON")
		return errors.New("bad json")
	}

	LoadEnv()
	session := ConnectAws()
	svc := s3.New(session)

	bucket := GetEnvWithKey("BUCKET_NAME")
	key := GetEnvWithKey("KEY")

	var timeout time.Duration
	flag.DurationVar(&timeout, "d", 0, "Upload timeout.")
	flag.Parse()

	ctx := context.Background()
	var cancelFn func()
	if timeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, timeout)
	}
	if cancelFn != nil {
		defer cancelFn()
	}

	b, err := json.Marshal(hostInstances)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}

	_, err = svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(b),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			fmt.Fprintf(os.Stderr, "upload canceled due to timeout, %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
		}
		return err
	}

	fmt.Printf("successfully uploaded file to %s/%s\n", bucket, key)
	return nil
}

func GetHostConfiguration() []byte {
	LoadEnv()
	session := ConnectAws()
	svc := s3.New(session)

	bucket := GetEnvWithKey("BUCKET_NAME")
	key := GetEnvWithKey("KEY")

	rawObject, _ := svc.GetObject(
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})

	buf := new(bytes.Buffer)
	buf.ReadFrom(rawObject.Body)
	return buf.Bytes()
}

var AccessKeyID string
var SecretAccessKey string
var MyRegion string

func ConnectAws() *session.Session {
	AccessKeyID = GetEnvWithKey("AWS_ACCESS_KEY_ID")
	SecretAccessKey = GetEnvWithKey("AWS_SECRET_ACCESS_KEY")
	MyRegion = GetEnvWithKey("AWS_REGION")
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(MyRegion),
			Credentials: credentials.NewStaticCredentials(
				AccessKeyID,
				SecretAccessKey,
				"",
			),
		})
	if err != nil {
		panic(err)
	}
	return sess
}
