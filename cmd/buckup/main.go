package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/urfave/cli/v2"
)

func main() {
	err := run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	app := &cli.App{
		Action: doctor,
	}

	return app.Run(os.Args)
}

func doctor(c *cli.Context) error {
	fmt.Println("Hello, world.")

	sess, err := setupObjectStore("localhost:9000", "minio", "minio123", false, "us-west-1")
	if err != nil {
		return err
	}

	ctx := c.Context

	client := s3.New(sess)

	uploader := s3manager.NewUploader(sess)

	_ = uploader

	_, err = client.CreateBucketWithContext(ctx, &s3.CreateBucketInput{
		Bucket: aws.String("test"),
	})
	if err != nil && !strings.HasPrefix(err.Error(), s3.ErrCodeBucketAlreadyOwnedByYou) {
		return err
	}

	obj := &s3manager.UploadInput{
		ACL:                       nil,
		Body:                      strings.NewReader("abc"),
		Bucket:                    aws.String("test"),
		BucketKeyEnabled:          nil,
		CacheControl:              nil,
		ContentDisposition:        nil,
		ContentEncoding:           nil,
		ContentLanguage:           nil,
		ContentMD5:                nil,
		ContentType:               nil,
		ExpectedBucketOwner:       nil,
		Expires:                   aws.Time(time.Now().Add(time.Hour)),
		GrantFullControl:          nil,
		GrantRead:                 nil,
		GrantReadACP:              nil,
		GrantWriteACP:             nil,
		Key:                       aws.String("pa/th"),
		Metadata:                  map[string]*string{"a": aws.String("b")},
		ObjectLockLegalHoldStatus: nil,
		ObjectLockMode:            nil,
		ObjectLockRetainUntilDate: nil,
		RequestPayer:              nil,
		SSECustomerAlgorithm:      nil,
		SSECustomerKey:            nil,
		SSECustomerKeyMD5:         nil,
		SSEKMSEncryptionContext:   nil,
		SSEKMSKeyId:               nil,
		ServerSideEncryption:      nil,
		StorageClass:              nil,
		Tagging:                   nil,
		WebsiteRedirectLocation:   nil,
	}
	_ = obj
	uploader.UploadWithContext(ctx, obj)

	//uploader.UploadWithContext(ctx, &s3manager.UploadInput{
	//	Bucket: aws.String("test"),
	//	Key:    aws.String("pa/th"),
	//	Body:   strings.NewReader("abc"),
	//})

	//_, err = client.PutBucketVersioningWithContext(ctx, &s3.PutBucketVersioningInput{
	//	Bucket: aws.String("test"),
	//	VersioningConfiguration: &s3.VersioningConfiguration{
	//		//				MFADelete: aws.String("Disabled"),
	//		//Status: aws.String("Enabled"),
	//		Status: aws.String(s3.BucketVersioningStatusEnabled),
	//		//				MFADelete: aws.String(s3.MFADeleteStatusDisabled),
	//	},
	//})
	//if err != nil {
	//	return err
	//}

	buckets, err := client.ListBucketsWithContext(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return err
	}

	for _, bucket := range buckets.Buckets {
		fmt.Printf("bucket: %s\n", *bucket.Name)

		err = client.ListObjectVersionsPagesWithContext(ctx, &s3.ListObjectVersionsInput{
			Bucket: bucket.Name,
		}, func(page *s3.ListObjectVersionsOutput, lastPage bool) bool {
			for _, obj := range page.Versions {
				fmt.Printf("%v\n", *obj)

				head, err := client.HeadObjectWithContext(ctx, &s3.HeadObjectInput{
					Bucket: bucket.Name,
					Key:    obj.Key,
				})
				if err != nil {
					log.Println(err)
					return false
				}

				fmt.Printf("%v\n", *head)
			}

			return lastPage
		})
		if err != nil {
			return err
		}
	}

	//_, err = client.GetObjectWithContext(ctx, &s3.GetObjectInput{
	//	Bucket: aws.String(os.Bucket),
	//	Key:    aws.String("path"),
	//})
	//if err != nil {
	//	return err
	//}

	return nil
}

/*
func backup(c *cli.Context) error {

}

func verify(c *cli.Context) error {

}

func restore(c *cli.Context) error {

}
*/

func setupObjectStore(endpoint, accessKey, secretKey string, useSSL bool, region string) (*session.Session, error) {
	endpointProtocol := "http"
	if useSSL {
		endpointProtocol = "https"
	}

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:         aws.String(fmt.Sprintf("%s://%s", endpointProtocol, endpoint)),
		Region:           aws.String(region),
		DisableSSL:       aws.Bool(!useSSL),
		S3ForcePathStyle: aws.Bool(true),
	}

	sess, err := session.NewSession(s3Config)
	if err != nil {
		return nil, fmt.Errorf("set up aws session: %v", err)
	}

	return sess, nil
}
