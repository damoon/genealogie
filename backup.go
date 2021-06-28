package backup

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type BackupLocation struct {
	Storage
	EncrytionKey string
}

type Storage struct {
	Sess       *session.Session
	BucketName string
}

func (s BackupLocation) Snapshot(ctx context.Context, source Storage) error {
	client := s3.New(source.Sess)

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

	return nil
}

func (s BackupLocation) ListSnapshots(bucket string) error {
	return nil
}

func (s BackupLocation) ListFiles(bucket, snapshot string, timestamp *time.Time) error {
	return nil
}

func (s BackupLocation) ListFileSnapshots(bucket, path string) error {
	return nil
}

func (s BackupLocation) GarbageCollection() error {
	return nil
}

func (s BackupLocation) Scrub() error {
	return nil
}

func (s BackupLocation) Restore(bucket, prefix string, timestamp *time.Time, target Storage) error {
	return nil
}
