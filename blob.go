package backup

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"gocloud.dev/blob/s3blob"
)

func gocloud(ctx context.Context) error {
	// Establish an AWS session.
	// See https://docs.aws.amazon.com/sdk-for-go/api/aws/session/ for more info.
	// The region must match the region for "my-bucket".
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-1"),
	})
	if err != nil {
		return err
	}

	// Create a *blob.Bucket.
	bucket, err := s3blob.OpenBucket(ctx, sess, "my-bucket", nil)
	if err != nil {
		return err
	}
	defer bucket.Close()

	att, _ := bucket.Attributes(ctx, "key")
	_ = att.CacheControl

	return nil
}
