package backup

import (
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func init() {
	obj := &GetObjectOutput{}
	version := &ObjectVersion{}
	head := &HeadObjectOutput{}
	tagging := &s3.GetBucketTaggingOutput{}

	_, _, _, _ = obj, version, head, tagging

	upload := &s3manager.UploadInput{
		ACL:                       nil,
		Body:                      obj.Body,
		Bucket:                    nil, // static
		BucketKeyEnabled:          obj.BucketKeyEnabled,
		CacheControl:              head.CacheControl,
		ContentDisposition:        head.ContentDisposition,
		ContentEncoding:           head.ContentEncoding,
		ContentLanguage:           head.ContentLanguage,
		ContentMD5:                nil, // calculate
		ContentType:               head.ContentType,
		ExpectedBucketOwner:       version.Owner.ID,
		Expires:                   nil, // head.Expires
		GrantFullControl:          nil,
		GrantRead:                 nil,
		GrantReadACP:              nil,
		GrantWriteACP:             nil,
		Key:                       version.Key,
		Metadata:                  map[string]*string{"a": aws.String("b")},
		ObjectLockLegalHoldStatus: obj.ObjectLockLegalHoldStatus,
		ObjectLockMode:            obj.ObjectLockMode,
		ObjectLockRetainUntilDate: obj.ObjectLockRetainUntilDate,
		RequestPayer:              nil,
		SSECustomerAlgorithm:      obj.SSECustomerAlgorithm,
		SSECustomerKey:            nil,
		SSECustomerKeyMD5:         obj.SSECustomerKeyMD5,
		SSEKMSEncryptionContext:   nil,
		SSEKMSKeyId:               obj.SSEKMSKeyId,
		ServerSideEncryption:      obj.ServerSideEncryption,
		StorageClass:              obj.StorageClass,
		Tagging:                   nil, // url encode tagging,
		WebsiteRedirectLocation:   obj.WebsiteRedirectLocation,
	}

	_ = upload
}

type GetObjectOutput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// Indicates that a range of bytes was specified.
	AcceptRanges *string `location:"header" locationName:"accept-ranges" type:"string"`

	// Object data.
	Body io.ReadCloser `type:"blob"`

	// Indicates whether the object uses an S3 Bucket Key for server-side encryption
	// with AWS KMS (SSE-KMS).
	BucketKeyEnabled *bool `location:"header" locationName:"x-amz-server-side-encryption-bucket-key-enabled" type:"boolean"`

	// Specifies caching behavior along the request/reply chain.
	CacheControl *string `location:"header" locationName:"Cache-Control" type:"string"`

	// Specifies presentational information for the object.
	ContentDisposition *string `location:"header" locationName:"Content-Disposition" type:"string"`

	// Specifies what content encodings have been applied to the object and thus
	// what decoding mechanisms must be applied to obtain the media-type referenced
	// by the Content-Type header field.
	ContentEncoding *string `location:"header" locationName:"Content-Encoding" type:"string"`

	// The language the content is in.
	ContentLanguage *string `location:"header" locationName:"Content-Language" type:"string"`

	// Size of the body in bytes.
	ContentLength *int64 `location:"header" locationName:"Content-Length" type:"long"`

	// The portion of the object returned in the response.
	ContentRange *string `location:"header" locationName:"Content-Range" type:"string"`

	// A standard MIME type describing the format of the object data.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	// Specifies whether the object retrieved was (true) or was not (false) a Delete
	// Marker. If false, this response header does not appear in the response.
	DeleteMarker *bool `location:"header" locationName:"x-amz-delete-marker" type:"boolean"`

	// An ETag is an opaque identifier assigned by a web server to a specific version
	// of a resource found at a URL.
	ETag *string `location:"header" locationName:"ETag" type:"string"`

	// If the object expiration is configured (see PUT Bucket lifecycle), the response
	// includes this header. It includes the expiry-date and rule-id key-value pairs
	// providing object expiration information. The value of the rule-id is URL
	// encoded.
	Expiration *string `location:"header" locationName:"x-amz-expiration" type:"string"`

	// The date and time at which the object is no longer cacheable.
	Expires *string `location:"header" locationName:"Expires" type:"string"`

	// Creation date of the object.
	LastModified *time.Time `location:"header" locationName:"Last-Modified" type:"timestamp"`

	// A map of metadata to store with the object in S3.
	//
	// By default unmarshaled keys are written as a map keys in following canonicalized format:
	// the first letter and any letter following a hyphen will be capitalized, and the rest as lowercase.
	// Set `aws.Config.LowerCaseHeaderMaps` to `true` to write unmarshaled keys to the map as lowercase.
	Metadata map[string]*string `location:"headers" locationName:"x-amz-meta-" type:"map"`

	// This is set to the number of metadata entries not returned in x-amz-meta
	// headers. This can happen if you create metadata using an API like SOAP that
	// supports more flexible metadata than the REST API. For example, using SOAP,
	// you can create metadata whose values are not legal HTTP headers.
	MissingMeta *int64 `location:"header" locationName:"x-amz-missing-meta" type:"integer"`

	// Indicates whether this object has an active legal hold. This field is only
	// returned if you have permission to view an object's legal hold status.
	ObjectLockLegalHoldStatus *string `location:"header" locationName:"x-amz-object-lock-legal-hold" type:"string" enum:"ObjectLockLegalHoldStatus"`

	// The Object Lock mode currently in place for this object.
	ObjectLockMode *string `location:"header" locationName:"x-amz-object-lock-mode" type:"string" enum:"ObjectLockMode"`

	// The date and time when this object's Object Lock will expire.
	ObjectLockRetainUntilDate *time.Time `location:"header" locationName:"x-amz-object-lock-retain-until-date" type:"timestamp" timestampFormat:"iso8601"`

	// The count of parts this object has.
	PartsCount *int64 `location:"header" locationName:"x-amz-mp-parts-count" type:"integer"`

	// Amazon S3 can return this if your request involves a bucket that is either
	// a source or destination in a replication rule.
	ReplicationStatus *string `location:"header" locationName:"x-amz-replication-status" type:"string" enum:"ReplicationStatus"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged *string `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"RequestCharged"`

	// Provides information about object restoration action and expiration time
	// of the restored object copy.
	Restore *string `location:"header" locationName:"x-amz-restore" type:"string"`

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm
	// used.
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round-trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`

	// If present, specifies the ID of the AWS Key Management Service (AWS KMS)
	// symmetric customer managed customer master key (CMK) that was used for the
	// object.
	SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string" sensitive:"true"`

	// The server-side encryption algorithm used when storing this object in Amazon
	// S3 (for example, AES256, aws:kms).
	ServerSideEncryption *string `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"ServerSideEncryption"`

	// Provides storage class information of the object. Amazon S3 returns this
	// header for all objects except for S3 Standard storage class objects.
	StorageClass *string `location:"header" locationName:"x-amz-storage-class" type:"string" enum:"StorageClass"`

	// The number of tags, if any, on the object.
	TagCount *int64 `location:"header" locationName:"x-amz-tagging-count" type:"integer"`

	// Version of the object.
	VersionId *string `location:"header" locationName:"x-amz-version-id" type:"string"`

	// If the bucket is configured as a website, redirects requests for this object
	// to another object in the same bucket or to an external URL. Amazon S3 stores
	// the value of this header in the object metadata.
	WebsiteRedirectLocation *string `location:"header" locationName:"x-amz-website-redirect-location" type:"string"`
}

type ObjectVersion struct {
	_ struct{} `type:"structure"`

	// The entity tag is an MD5 hash of that version of the object.
	ETag *string `type:"string"`

	// Specifies whether the object is (true) or is not (false) the latest version
	// of an object.
	IsLatest *bool `type:"boolean"`

	// The object key.
	Key *string `min:"1" type:"string"`

	// Date and time the object was last modified.
	LastModified *time.Time `type:"timestamp"`

	// Specifies the owner of the object.
	Owner *Owner `type:"structure"`

	// Size in bytes of the object.
	Size *int64 `type:"integer"`

	// The class of storage used to store the object.
	StorageClass *string `type:"string" enum:"ObjectVersionStorageClass"`

	// Version ID of an object.
	VersionId *string `type:"string"`
}

type Owner struct {
	_ struct{} `type:"structure"`

	// Container for the display name of the owner.
	DisplayName *string `type:"string"`

	// Container for the ID of the owner.
	ID *string `type:"string"`
}

type HeadObjectOutput struct {
	_ struct{} `type:"structure"`

	// Indicates that a range of bytes was specified.
	AcceptRanges *string `location:"header" locationName:"accept-ranges" type:"string"`

	// The archive state of the head object.
	ArchiveStatus *string `location:"header" locationName:"x-amz-archive-status" type:"string" enum:"ArchiveStatus"`

	// Indicates whether the object uses an S3 Bucket Key for server-side encryption
	// with AWS KMS (SSE-KMS).
	BucketKeyEnabled *bool `location:"header" locationName:"x-amz-server-side-encryption-bucket-key-enabled" type:"boolean"`

	// Specifies caching behavior along the request/reply chain.
	CacheControl *string `location:"header" locationName:"Cache-Control" type:"string"`

	// Specifies presentational information for the object.
	ContentDisposition *string `location:"header" locationName:"Content-Disposition" type:"string"`

	// Specifies what content encodings have been applied to the object and thus
	// what decoding mechanisms must be applied to obtain the media-type referenced
	// by the Content-Type header field.
	ContentEncoding *string `location:"header" locationName:"Content-Encoding" type:"string"`

	// The language the content is in.
	ContentLanguage *string `location:"header" locationName:"Content-Language" type:"string"`

	// Size of the body in bytes.
	ContentLength *int64 `location:"header" locationName:"Content-Length" type:"long"`

	// A standard MIME type describing the format of the object data.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	// Specifies whether the object retrieved was (true) or was not (false) a Delete
	// Marker. If false, this response header does not appear in the response.
	DeleteMarker *bool `location:"header" locationName:"x-amz-delete-marker" type:"boolean"`

	// An ETag is an opaque identifier assigned by a web server to a specific version
	// of a resource found at a URL.
	ETag *string `location:"header" locationName:"ETag" type:"string"`

	// If the object expiration is configured (see PUT Bucket lifecycle), the response
	// includes this header. It includes the expiry-date and rule-id key-value pairs
	// providing object expiration information. The value of the rule-id is URL
	// encoded.
	Expiration *string `location:"header" locationName:"x-amz-expiration" type:"string"`

	// The date and time at which the object is no longer cacheable.
	Expires *string `location:"header" locationName:"Expires" type:"string"`

	// Creation date of the object.
	LastModified *time.Time `location:"header" locationName:"Last-Modified" type:"timestamp"`

	// A map of metadata to store with the object in S3.
	//
	// By default unmarshaled keys are written as a map keys in following canonicalized format:
	// the first letter and any letter following a hyphen will be capitalized, and the rest as lowercase.
	// Set `aws.Config.LowerCaseHeaderMaps` to `true` to write unmarshaled keys to the map as lowercase.
	Metadata map[string]*string `location:"headers" locationName:"x-amz-meta-" type:"map"`

	// This is set to the number of metadata entries not returned in x-amz-meta
	// headers. This can happen if you create metadata using an API like SOAP that
	// supports more flexible metadata than the REST API. For example, using SOAP,
	// you can create metadata whose values are not legal HTTP headers.
	MissingMeta *int64 `location:"header" locationName:"x-amz-missing-meta" type:"integer"`

	// Specifies whether a legal hold is in effect for this object. This header
	// is only returned if the requester has the s3:GetObjectLegalHold permission.
	// This header is not returned if the specified version of this object has never
	// had a legal hold applied. For more information about S3 Object Lock, see
	// Object Lock (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html).
	ObjectLockLegalHoldStatus *string `location:"header" locationName:"x-amz-object-lock-legal-hold" type:"string" enum:"ObjectLockLegalHoldStatus"`

	// The Object Lock mode, if any, that's in effect for this object. This header
	// is only returned if the requester has the s3:GetObjectRetention permission.
	// For more information about S3 Object Lock, see Object Lock (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html).
	ObjectLockMode *string `location:"header" locationName:"x-amz-object-lock-mode" type:"string" enum:"ObjectLockMode"`

	// The date and time when the Object Lock retention period expires. This header
	// is only returned if the requester has the s3:GetObjectRetention permission.
	ObjectLockRetainUntilDate *time.Time `location:"header" locationName:"x-amz-object-lock-retain-until-date" type:"timestamp" timestampFormat:"iso8601"`

	// The count of parts this object has.
	PartsCount *int64 `location:"header" locationName:"x-amz-mp-parts-count" type:"integer"`

	// Amazon S3 can return this header if your request involves a bucket that is
	// either a source or a destination in a replication rule.
	//
	// In replication, you have a source bucket on which you configure replication
	// and destination bucket or buckets where Amazon S3 stores object replicas.
	// When you request an object (GetObject) or object metadata (HeadObject) from
	// these buckets, Amazon S3 will return the x-amz-replication-status header
	// in the response as follows:
	//
	//    * If requesting an object from the source bucket — Amazon S3 will return
	//    the x-amz-replication-status header if the object in your request is eligible
	//    for replication. For example, suppose that in your replication configuration,
	//    you specify object prefix TaxDocs requesting Amazon S3 to replicate objects
	//    with key prefix TaxDocs. Any objects you upload with this key name prefix,
	//    for example TaxDocs/document1.pdf, are eligible for replication. For any
	//    object request with this key name prefix, Amazon S3 will return the x-amz-replication-status
	//    header with value PENDING, COMPLETED or FAILED indicating object replication
	//    status.
	//
	//    * If requesting an object from a destination bucket — Amazon S3 will
	//    return the x-amz-replication-status header with value REPLICA if the object
	//    in your request is a replica that Amazon S3 created and there is no replica
	//    modification replication in progress.
	//
	//    * When replicating objects to multiple destination buckets the x-amz-replication-status
	//    header acts differently. The header of the source object will only return
	//    a value of COMPLETED when replication is successful to all destinations.
	//    The header will remain at value PENDING until replication has completed
	//    for all destinations. If one or more destinations fails replication the
	//    header will return FAILED.
	//
	// For more information, see Replication (https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html).
	ReplicationStatus *string `location:"header" locationName:"x-amz-replication-status" type:"string" enum:"ReplicationStatus"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged *string `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"RequestCharged"`

	// If the object is an archived object (an object whose storage class is GLACIER),
	// the response includes this header if either the archive restoration is in
	// progress (see RestoreObject (https://docs.aws.amazon.com/AmazonS3/latest/API/API_RestoreObject.html)
	// or an archive copy is already restored.
	//
	// If an archive copy is already restored, the header value indicates when Amazon
	// S3 is scheduled to delete the object copy. For example:
	//
	// x-amz-restore: ongoing-request="false", expiry-date="Fri, 21 Dec 2012 00:00:00
	// GMT"
	//
	// If the object restoration is in progress, the header returns the value ongoing-request="true".
	//
	// For more information about archiving objects, see Transitioning Objects:
	// General Considerations (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html#lifecycle-transition-general-considerations).
	Restore *string `location:"header" locationName:"x-amz-restore" type:"string"`

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm
	// used.
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round-trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`

	// If present, specifies the ID of the AWS Key Management Service (AWS KMS)
	// symmetric customer managed customer master key (CMK) that was used for the
	// object.
	SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string" sensitive:"true"`

	// If the object is stored using server-side encryption either with an AWS KMS
	// customer master key (CMK) or an Amazon S3-managed encryption key, the response
	// includes this header with the value of the server-side encryption algorithm
	// used when storing this object in Amazon S3 (for example, AES256, aws:kms).
	ServerSideEncryption *string `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"ServerSideEncryption"`

	// Provides storage class information of the object. Amazon S3 returns this
	// header for all objects except for S3 Standard storage class objects.
	//
	// For more information, see Storage Classes (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html).
	StorageClass *string `location:"header" locationName:"x-amz-storage-class" type:"string" enum:"StorageClass"`

	// Version of the object.
	VersionId *string `location:"header" locationName:"x-amz-version-id" type:"string"`

	// If the bucket is configured as a website, redirects requests for this object
	// to another object in the same bucket or to an external URL. Amazon S3 stores
	// the value of this header in the object metadata.
	WebsiteRedirectLocation *string `location:"header" locationName:"x-amz-website-redirect-location" type:"string"`
}
