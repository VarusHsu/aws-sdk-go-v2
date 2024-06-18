// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Network File System (NFS) file share on an existing S3 File Gateway.
// In Storage Gateway, a file share is a file system mount point backed by Amazon
// S3 cloud storage. Storage Gateway exposes file shares using an NFS interface.
// This operation is only supported for S3 File Gateways.
//
// S3 File gateway requires Security Token Service (Amazon Web Services STS) to be
// activated to enable you to create a file share. Make sure Amazon Web Services
// STS is activated in the Amazon Web Services Region you are creating your S3 File
// Gateway in. If Amazon Web Services STS is not activated in the Amazon Web
// Services Region, activate it. For information about how to activate Amazon Web
// Services STS, see [Activating and deactivating Amazon Web Services STS in an Amazon Web Services Region]in the Identity and Access Management User Guide.
//
// S3 File Gateways do not support creating hard or symbolic links on a file share.
//
// [Activating and deactivating Amazon Web Services STS in an Amazon Web Services Region]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html
func (c *Client) CreateNFSFileShare(ctx context.Context, params *CreateNFSFileShareInput, optFns ...func(*Options)) (*CreateNFSFileShareOutput, error) {
	if params == nil {
		params = &CreateNFSFileShareInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNFSFileShare", params, optFns, c.addOperationCreateNFSFileShareMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNFSFileShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CreateNFSFileShareInput
type CreateNFSFileShareInput struct {

	// A unique string value that you supply that is used by S3 File Gateway to ensure
	// idempotent file share creation.
	//
	// This member is required.
	ClientToken *string

	// The Amazon Resource Name (ARN) of the S3 File Gateway on which you want to
	// create a file share.
	//
	// This member is required.
	GatewayARN *string

	// A custom ARN for the backend storage used for storing data for file shares. It
	// includes a resource ARN with an optional prefix concatenation. The prefix must
	// end with a forward slash (/).
	//
	// You can specify LocationARN as a bucket ARN, access point ARN or access point
	// alias, as shown in the following examples.
	//
	// Bucket ARN:
	//
	//     arn:aws:s3:::my-bucket/prefix/
	//
	// Access point ARN:
	//
	//     arn:aws:s3:region:account-id:accesspoint/access-point-name/prefix/
	//
	// If you specify an access point, the bucket policy must be configured to
	// delegate access control to the access point. For information, see [Delegating access control to access points]in the Amazon
	// S3 User Guide.
	//
	// Access point alias:
	//
	//     test-ap-ab123cdef4gehijklmn5opqrstuvuse1a-s3alias
	//
	// [Delegating access control to access points]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-policies.html#access-points-delegating-control
	//
	// This member is required.
	LocationARN *string

	// The ARN of the Identity and Access Management (IAM) role that an S3 File
	// Gateway assumes when it accesses the underlying storage.
	//
	// This member is required.
	Role *string

	// The Amazon Resource Name (ARN) of the storage used for audit logs.
	AuditDestinationARN *string

	// Specifies the Region of the S3 bucket where the NFS file share stores files.
	//
	// This parameter is required for NFS file shares that connect to Amazon S3
	// through a VPC endpoint, a VPC access point, or an access point alias that points
	// to a VPC access point.
	BucketRegion *string

	// Specifies refresh cache information for the file share.
	CacheAttributes *types.CacheAttributes

	// The list of clients that are allowed to access the S3 File Gateway. The list
	// must contain either valid IP addresses or valid CIDR blocks.
	ClientList []string

	// The default storage class for objects put into an Amazon S3 bucket by the S3
	// File Gateway. The default value is S3_STANDARD . Optional.
	//
	// Valid Values: S3_STANDARD | S3_INTELLIGENT_TIERING | S3_STANDARD_IA |
	// S3_ONEZONE_IA
	DefaultStorageClass *string

	// The name of the file share. Optional.
	//
	// FileShareName must be set if an S3 prefix name is set in LocationARN , or if an
	// access point or access point alias is used.
	FileShareName *string

	// A value that enables guessing of the MIME type for uploaded objects based on
	// file extensions. Set this value to true to enable MIME type guessing, otherwise
	// set to false . The default value is true .
	//
	// Valid Values: true | false
	GuessMIMETypeEnabled *bool

	// Set to true to use Amazon S3 server-side encryption with your own KMS key, or
	// false to use a key managed by Amazon S3. Optional.
	//
	// Valid Values: true | false
	KMSEncrypted *bool

	// The Amazon Resource Name (ARN) of a symmetric customer master key (CMK) used
	// for Amazon S3 server-side encryption. Storage Gateway does not support
	// asymmetric CMKs. This value can only be set when KMSEncrypted is true . Optional.
	KMSKey *string

	// File share default values. Optional.
	NFSFileShareDefaults *types.NFSFileShareDefaults

	// The notification policy of the file share. SettlingTimeInSeconds controls the
	// number of seconds to wait after the last point in time a client wrote to a file
	// before generating an ObjectUploaded notification. Because clients can make many
	// small writes to files, it's best to set this parameter for as long as possible
	// to avoid generating multiple notifications for the same file in a small time
	// period.
	//
	// SettlingTimeInSeconds has no effect on the timing of the object uploading to
	// Amazon S3, only the timing of the notification.
	//
	// The following example sets NotificationPolicy on with SettlingTimeInSeconds set
	// to 60.
	//
	//     {\"Upload\": {\"SettlingTimeInSeconds\": 60}}
	//
	// The following example sets NotificationPolicy off.
	//
	//     {}
	NotificationPolicy *string

	// A value that sets the access control list (ACL) permission for objects in the
	// S3 bucket that a S3 File Gateway puts objects into. The default value is private
	// .
	ObjectACL types.ObjectACL

	// A value that sets the write status of a file share. Set this value to true to
	// set the write status to read-only, otherwise set to false .
	//
	// Valid Values: true | false
	ReadOnly *bool

	// A value that sets who pays the cost of the request and the cost associated with
	// data download from the S3 bucket. If this value is set to true , the requester
	// pays the costs; otherwise, the S3 bucket owner pays. However, the S3 bucket
	// owner always pays the cost of storing data.
	//
	// RequesterPays is a configuration for the S3 bucket that backs the file share,
	// so make sure that the configuration on the file share is the same as the S3
	// bucket configuration.
	//
	// Valid Values: true | false
	RequesterPays *bool

	// A value that maps a user to anonymous user.
	//
	// Valid values are the following:
	//
	//   - RootSquash : Only root is mapped to anonymous user.
	//
	//   - NoSquash : No one is mapped to anonymous user.
	//
	//   - AllSquash : Everyone is mapped to anonymous user.
	Squash *string

	// A list of up to 50 tags that can be assigned to the NFS file share. Each tag is
	// a key-value pair.
	//
	// Valid characters for key and value are letters, spaces, and numbers
	// representable in UTF-8 format, and the following special characters: + - = . _ :
	// / @. The maximum length of a tag's key is 128 characters, and the maximum length
	// for a tag's value is 256.
	Tags []types.Tag

	// Specifies the DNS name for the VPC endpoint that the NFS file share uses to
	// connect to Amazon S3.
	//
	// This parameter is required for NFS file shares that connect to Amazon S3
	// through a VPC endpoint, a VPC access point, or an access point alias that points
	// to a VPC access point.
	VPCEndpointDNSName *string

	noSmithyDocumentSerde
}

// CreateNFSFileShareOutput
type CreateNFSFileShareOutput struct {

	// The Amazon Resource Name (ARN) of the newly created file share.
	FileShareARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateNFSFileShareMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateNFSFileShare{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateNFSFileShare{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateNFSFileShare"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpCreateNFSFileShareValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNFSFileShare(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateNFSFileShare(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateNFSFileShare",
	}
}
