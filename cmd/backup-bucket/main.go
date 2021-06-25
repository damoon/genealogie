package main

import (
	"fmt"
	"log"
	"os"

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

	os, err := setupObjectStore("localhost:9000", "minio", "minio123", false, "us-west-1", "test")
	if err != nil {
		return err
	}

	ctx := c.Context

	os.Client.CreateBucketWithContext(ctx, &s3.CreateBucketInput{
		Bucket: aws.String("test"),
	})

	buckets, err := os.Client.ListBucketsWithContext(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return err
	}

	for _, bucket := range buckets.Buckets {
		fmt.Printf("bucket: %s\n", *bucket.Name)

		_, err := os.Client.PutBucketVersioningWithContext(ctx, &s3.PutBucketVersioningInput{
			Bucket: bucket.Name,
			VersioningConfiguration: &s3.VersioningConfiguration{
				//				MFADelete: aws.String("Disabled"),
				//Status: aws.String("Enabled"),
				Status: aws.String(s3.BucketVersioningStatusEnabled),
				//				MFADelete: aws.String(s3.MFADeleteStatusDisabled),
			},
		})
		if err != nil {
			return err
		}

		objects, err := os.Client.ListObjectsV2WithContext(ctx, &s3.ListObjectsV2Input{
			Bucket: bucket.Name,
		})
		if err != nil {
			return err
		}

		for _, obj := range objects.Contents {
			fmt.Printf(" %s\n", *obj.Key)
			fmt.Printf(" %s\n", *obj.ETag)
		}

		versions, err := os.Client.ListObjectVersionsWithContext(ctx, &s3.ListObjectVersionsInput{
			Bucket: bucket.Name,
		})
		if err != nil {
			return err
		}

		for _, version := range versions.Versions {
			fmt.Printf(" %s\n", *version.Key)
			fmt.Printf(" %s\n", *version.ETag)
		}
	}

	//_, err = os.Client.GetObjectWithContext(ctx, &s3.GetObjectInput{
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

type ObjectStore struct {
	Client   *s3.S3
	Uploader *s3manager.Uploader
	Bucket   string
}

func setupObjectStore(
	endpoint, accessKey, secretKey string,
	useSSL bool,
	region, bucket string,
) (*ObjectStore, error) {
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

	s3Client := s3.New(sess)

	return &ObjectStore{
		Client:   s3Client,
		Uploader: s3manager.NewUploader(sess),
		Bucket:   bucket,
	}, nil
}
