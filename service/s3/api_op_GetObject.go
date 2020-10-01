// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
	"time"
)

// Retrieves objects from Amazon S3. To use GET, you must have READ access to the
// object. If you grant READ access to the anonymous user, you can return the
// object without using an authorization header.  <p>An Amazon S3 bucket has no
// directory hierarchy such as you would find in a typical computer file system.
// You can, however, create a logical hierarchy by using object key names that
// imply a folder structure. For example, instead of naming an object
// <code>sample.jpg</code>, you can name it
// <code>photos/2006/February/sample.jpg</code>.</p> <p>To get an object from such
// a logical hierarchy, specify the full key name for the object in the
// <code>GET</code> operation. For a virtual hosted-style request example, if you
// have the object <code>photos/2006/February/sample.jpg</code>, specify the
// resource as <code>/photos/2006/February/sample.jpg</code>. For a path-style
// request example, if you have the object
// <code>photos/2006/February/sample.jpg</code> in the bucket named
// <code>examplebucket</code>, specify the resource as
// <code>/examplebucket/photos/2006/February/sample.jpg</code>. For more
// information about request types, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html#VirtualHostingSpecifyBucket">HTTP
// Host Header Bucket Specification</a>.</p> <p>To distribute large files to many
// people, you can save bandwidth costs by using BitTorrent. For more information,
// see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/S3Torrent.html">Amazon S3
// Torrent</a>. For more information about returning the ACL of an object, see
// <a>GetObjectAcl</a>.</p> <p>If the object you are retrieving is stored in the
// GLACIER or DEEP_ARCHIVE storage classes, before you can retrieve the object you
// must first restore a copy using . Otherwise, this operation returns an
// <code>InvalidObjectStateError</code> error. For information about restoring
// archived objects, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/restoring-objects.html">Restoring
// Archived Objects</a>.</p> <p>Encryption request headers, like
// <code>x-amz-server-side-encryption</code>, should not be sent for GET requests
// if your object uses server-side encryption with CMKs stored in AWS KMS (SSE-KMS)
// or server-side encryption with Amazon S3–managed encryption keys (SSE-S3). If
// your object does use these types of keys, you’ll get an HTTP 400 BadRequest
// error.</p> <p>If you encrypt an object by using server-side encryption with
// customer-provided encryption keys (SSE-C) when you store the object in Amazon
// S3, then when you GET the object, you must use the following headers:</p> <ul>
// <li> <p>x-amz-server-side-encryption-customer-algorithm</p> </li> <li>
// <p>x-amz-server-side-encryption-customer-key</p> </li> <li>
// <p>x-amz-server-side-encryption-customer-key-MD5</p> </li> </ul> <p>For more
// information about SSE-C, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerSideEncryptionCustomerKeys.html">Server-Side
// Encryption (Using Customer-Provided Encryption Keys)</a>.</p> <p>Assuming you
// have permission to read object tags (permission for the
// <code>s3:GetObjectVersionTagging</code> action), the response also returns the
// <code>x-amz-tagging-count</code> header that provides the count of number of
// tags associated with the object. You can use <a>GetObjectTagging</a> to retrieve
// the tag set associated with an object.</p> <p> <b>Permissions</b> </p> <p>You
// need the <code>s3:GetObject</code> permission for this operation. For more
// information, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html">Specifying
// Permissions in a Policy</a>. If the object you request does not exist, the error
// Amazon S3 returns depends on whether you also have the
// <code>s3:ListBucket</code> permission.</p> <ul> <li> <p>If you have the
// <code>s3:ListBucket</code> permission on the bucket, Amazon S3 will return an
// HTTP status code 404 ("no such key") error.</p> </li> <li> <p>If you don’t have
// the <code>s3:ListBucket</code> permission, Amazon S3 will return an HTTP status
// code 403 ("access denied") error.</p> </li> </ul> <p> <b>Versioning</b> </p>
// <p>By default, the GET operation returns the current version of an object. To
// return a different version, use the <code>versionId</code> subresource.</p>
// <note> <p>If the current version of the object is a delete marker, Amazon S3
// behaves as if the object was deleted and includes <code>x-amz-delete-marker:
// true</code> in the response.</p> </note> <p>For more information about
// versioning, see <a>PutBucketVersioning</a>. </p> <p> <b>Overriding Response
// Header Values</b> </p> <p>There are times when you want to override certain
// response header values in a GET response. For example, you might override the
// Content-Disposition response header value in your GET request.</p> <p>You can
// override values for a set of response headers using the following query
// parameters. These response header values are sent only on a successful request,
// that is, when status code 200 OK is returned. The set of headers you can
// override using these parameters is a subset of the headers that Amazon S3
// accepts when you create an object. The response headers that you can override
// for the GET response are <code>Content-Type</code>,
// <code>Content-Language</code>, <code>Expires</code>, <code>Cache-Control</code>,
// <code>Content-Disposition</code>, and <code>Content-Encoding</code>. To override
// these header values in the GET response, you use the following request
// parameters.</p> <note> <p>You must sign the request, either using an
// Authorization header or a presigned URL, when using these parameters. They
// cannot be used with an unsigned (anonymous) request.</p> </note> <ul> <li> <p>
// <code>response-content-type</code> </p> </li> <li> <p>
// <code>response-content-language</code> </p> </li> <li> <p>
// <code>response-expires</code> </p> </li> <li> <p>
// <code>response-cache-control</code> </p> </li> <li> <p>
// <code>response-content-disposition</code> </p> </li> <li> <p>
// <code>response-content-encoding</code> </p> </li> </ul> <p> <b>Additional
// Considerations about Request Headers</b> </p> <p>If both of the
// <code>If-Match</code> and <code>If-Unmodified-Since</code> headers are present
// in the request as follows: <code>If-Match</code> condition evaluates to
// <code>true</code>, and; <code>If-Unmodified-Since</code> condition evaluates to
// <code>false</code>; then, S3 returns 200 OK and the data requested. </p> <p>If
// both of the <code>If-None-Match</code> and <code>If-Modified-Since</code>
// headers are present in the request as follows:<code> If-None-Match</code>
// condition evaluates to <code>false</code>, and; <code>If-Modified-Since</code>
// condition evaluates to <code>true</code>; then, S3 returns 304 Not Modified
// response code.</p> <p>For more information about conditional requests, see <a
// href="https://tools.ietf.org/html/rfc7232">RFC 7232</a>.</p> <p>The following
// operations are related to <code>GetObject</code>:</p> <ul> <li> <p>
// <a>ListBuckets</a> </p> </li> <li> <p> <a>GetObjectAcl</a> </p> </li> </ul>
func (c *Client) GetObject(ctx context.Context, params *GetObjectInput, optFns ...func(*Options)) (*GetObjectOutput, error) {
	stack := middleware.NewStack("GetObject", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetObjectMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	addOpGetObjectValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetObject(options.Region), middleware.Before)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	addMetadataRetrieverMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "GetObject",
			Err:           err,
		}
	}
	out := result.(*GetObjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetObjectInput struct {

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer types.RequestPayer

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure that the
	// encryption key was transmitted without error.
	SSECustomerKeyMD5 *string

	// The bucket name containing the object. When using this API with an access point,
	// you must direct requests to the access point hostname. The access point hostname
	// takes the form AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com.
	// When using this operation using an access point through the AWS SDKs, you
	// provide the access point ARN in place of the bucket name. For more information
	// about access point ARNs, see Using Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) in
	// the Amazon Simple Storage Service Developer Guide.
	//
	// This member is required.
	Bucket *string

	// Return the object only if its entity tag (ETag) is different from the one
	// specified, otherwise return a 304 (not modified).
	IfNoneMatch *string

	// Return the object only if its entity tag (ETag) is the same as the one
	// specified, otherwise return a 412 (precondition failed).
	IfMatch *string

	// Specifies the customer-provided encryption key for Amazon S3 to use in
	// encrypting data. This value is used to store the object and then it is
	// discarded; Amazon S3 does not store the encryption key. The key must be
	// appropriate for use with the algorithm specified in the
	// x-amz-server-side-encryption-customer-algorithm header.
	SSECustomerKey *string

	// Downloads the specified range bytes of an object. For more information about the
	// HTTP Range header, see
	// https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35
	// (https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35). Amazon S3
	// doesn't support retrieving multiple ranges of data per GET request.
	Range *string

	// Sets the Content-Disposition header of the response
	ResponseContentDisposition *string

	// Sets the Content-Language header of the response.
	ResponseContentLanguage *string

	// VersionId used to reference a specific version of the object.
	VersionId *string

	// Sets the Content-Encoding header of the response.
	ResponseContentEncoding *string

	// Key of the object to get.
	//
	// This member is required.
	Key *string

	// Specifies the algorithm to use to when encrypting the object (for example,
	// AES256).
	SSECustomerAlgorithm *string

	// Return the object only if it has been modified since the specified time,
	// otherwise return a 304 (not modified).
	IfModifiedSince *time.Time

	// Sets the Cache-Control header of the response.
	ResponseCacheControl *string

	// Sets the Expires header of the response.
	ResponseExpires *time.Time

	// Sets the Content-Type header of the response.
	ResponseContentType *string

	// Return the object only if it has not been modified since the specified time,
	// otherwise return a 412 (precondition failed).
	IfUnmodifiedSince *time.Time

	// Part number of the object being read. This is a positive integer between 1 and
	// 10,000. Effectively performs a 'ranged' GET request for the part specified.
	// Useful for downloading just a part of an object.
	PartNumber *int32
}

type GetObjectOutput struct {

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm used.
	SSECustomerAlgorithm *string

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged

	// Provides information about object restoration operation and expiration time of
	// the restored object copy.
	Restore *string

	// Last modified date of the object
	LastModified *time.Time

	// The server-side encryption algorithm used when storing this object in Amazon S3
	// (for example, AES256, aws:kms).
	ServerSideEncryption types.ServerSideEncryption

	// The count of parts this object has.
	PartsCount *int32

	// The portion of the object returned in the response.
	ContentRange *string

	// Object data.
	Body io.ReadCloser

	// If present, specifies the ID of the AWS Key Management Service (AWS KMS)
	// symmetric customer managed customer master key (CMK) that was used for the
	// object.
	SSEKMSKeyId *string

	// The Object Lock mode currently in place for this object.
	ObjectLockMode types.ObjectLockMode

	// Indicates that a range of bytes was specified.
	AcceptRanges *string

	// Specifies what content encodings have been applied to the object and thus what
	// decoding mechanisms must be applied to obtain the media-type referenced by the
	// Content-Type header field.
	ContentEncoding *string

	// Provides storage class information of the object. Amazon S3 returns this header
	// for all objects except for S3 Standard storage class objects.
	StorageClass types.StorageClass

	// A map of metadata to store with the object in S3.
	Metadata map[string]*string

	// The number of tags, if any, on the object.
	TagCount *int32

	// Specifies whether the object retrieved was (true) or was not (false) a Delete
	// Marker. If false, this response header does not appear in the response.
	DeleteMarker *bool

	// This is set to the number of metadata entries not returned in x-amz-meta
	// headers. This can happen if you create metadata using an API like SOAP that
	// supports more flexible metadata than the REST API. For example, using SOAP, you
	// can create metadata whose values are not legal HTTP headers.
	MissingMeta *int32

	// The date and time when this object's Object Lock will expire.
	ObjectLockRetainUntilDate *time.Time

	// The language the content is in.
	ContentLanguage *string

	// The date and time at which the object is no longer cacheable.
	Expires *time.Time

	// Size of the body in bytes.
	ContentLength *int64

	// Specifies presentational information for the object.
	ContentDisposition *string

	// Version of the object.
	VersionId *string

	// Indicates whether this object has an active legal hold. This field is only
	// returned if you have permission to view an object's legal hold status.
	ObjectLockLegalHoldStatus types.ObjectLockLegalHoldStatus

	// A standard MIME type describing the format of the object data.
	ContentType *string

	// Specifies caching behavior along the request/reply chain.
	CacheControl *string

	// If the bucket is configured as a website, redirects requests for this object to
	// another object in the same bucket or to an external URL. Amazon S3 stores the
	// value of this header in the object metadata.
	WebsiteRedirectLocation *string

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round-trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string

	// An ETag is an opaque identifier assigned by a web server to a specific version
	// of a resource found at a URL.
	ETag *string

	// Amazon S3 can return this if your request involves a bucket that is either a
	// source or destination in a replication rule.
	ReplicationStatus types.ReplicationStatus

	// If the object expiration is configured (see PUT Bucket lifecycle), the response
	// includes this header. It includes the expiry-date and rule-id key-value pairs
	// providing object expiration information. The value of the rule-id is URL
	// encoded.
	Expiration *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetObjectMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetObject{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetObject{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetObject(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetObject",
	}
}
