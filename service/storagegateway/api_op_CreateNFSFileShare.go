// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a Network File System (NFS) file share on an existing file gateway. In
// Storage Gateway, a file share is a file system mount point backed by Amazon S3
// cloud storage. Storage Gateway exposes file shares using an NFS interface. This
// operation is only supported for file gateways.  <important> <p>File gateway
// requires AWS Security Token Service (AWS STS) to be activated to enable you to
// create a file share. Make sure AWS STS is activated in the AWS Region you are
// creating your file gateway in. If AWS STS is not activated in the AWS Region,
// activate it. For information about how to activate AWS STS, see <a
// href="https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html">Activating
// and deactivating AWS STS in an AWS Region</a> in the <i>AWS Identity and Access
// Management User Guide</i>.</p> <p>File gateway does not support creating hard or
// symbolic links on a file share.</p> </important>
func (c *Client) CreateNFSFileShare(ctx context.Context, params *CreateNFSFileShareInput, optFns ...func(*Options)) (*CreateNFSFileShareOutput, error) {
	stack := middleware.NewStack("CreateNFSFileShare", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateNFSFileShareMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpCreateNFSFileShareValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNFSFileShare(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

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
			OperationName: "CreateNFSFileShare",
			Err:           err,
		}
	}
	out := result.(*CreateNFSFileShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CreateNFSFileShareInput
type CreateNFSFileShareInput struct {

	// The default storage class for objects put into an Amazon S3 bucket by the file
	// gateway. The default value is S3_INTELLIGENT_TIERING. Optional.  <p>Valid
	// Values: <code>S3_STANDARD</code> | <code>S3_INTELLIGENT_TIERING</code> |
	// <code>S3_STANDARD_IA</code> | <code>S3_ONEZONE_IA</code> </p>
	DefaultStorageClass *string

	// Set to true to use Amazon S3 server-side encryption with your own AWS KMS key,
	// or false to use a key managed by Amazon S3. Optional.  <p>Valid Values:
	// <code>true</code> | <code>false</code> </p>
	KMSEncrypted *bool

	// The ARN of the AWS Identity and Access Management (IAM) role that a file gateway
	// assumes when it accesses the underlying storage.
	//
	// This member is required.
	Role *string

	// The list of clients that are allowed to access the file gateway. The list must
	// contain either valid IP addresses or valid CIDR blocks.
	ClientList []*string

	// The Amazon Resource Name (ARN) of a symmetric customer master key (CMK) used for
	// Amazon S3 server-side encryption. Storage Gateway does not support asymmetric
	// CMKs. This value can only be set when KMSEncrypted is true. Optional.
	KMSKey *string

	// A value that sets the write status of a file share. Set this value to true to
	// set the write status to read-only, otherwise set to false.  <p>Valid Values:
	// <code>true</code> | <code>false</code> </p>
	ReadOnly *bool

	// The Amazon Resource Name (ARN) of the file gateway on which you want to create a
	// file share.
	//
	// This member is required.
	GatewayARN *string

	// Refresh cache information.
	CacheAttributes *types.CacheAttributes

	// A list of up to 50 tags that can be assigned to the NFS file share. Each tag is
	// a key-value pair.  <note> <p>Valid characters for key and value are letters,
	// spaces, and numbers representable in UTF-8 format, and the following special
	// characters: + - = . _ : / @. The maximum length of a tag's key is 128
	// characters, and the maximum length for a tag's value is 256.</p> </note>
	Tags []*types.Tag

	// File share default values. Optional.
	NFSFileShareDefaults *types.NFSFileShareDefaults

	// A value that sets who pays the cost of the request and the cost associated with
	// data download from the S3 bucket. If this value is set to true, the requester
	// pays the costs; otherwise, the S3 bucket owner pays. However, the S3 bucket
	// owner always pays the cost of storing data.  <note> <p>
	// <code>RequesterPays</code> is a configuration for the S3 bucket that backs the
	// file share, so make sure that the configuration on the file share is the same as
	// the S3 bucket configuration.</p> </note> <p>Valid Values: <code>true</code> |
	// <code>false</code> </p>
	RequesterPays *bool

	// A unique string value that you supply that is used by file gateway to ensure
	// idempotent file share creation.
	//
	// This member is required.
	ClientToken *string

	// The ARN of the backend storage used for storing file data. A prefix name can be
	// added to the S3 bucket name. It must end with a "/".
	//
	// This member is required.
	LocationARN *string

	// A value that maps a user to anonymous user.  <p>Valid values are the
	// following:</p> <ul> <li> <p> <code>RootSquash</code>: Only root is mapped to
	// anonymous user.</p> </li> <li> <p> <code>NoSquash</code>: No one is mapped to
	// anonymous user.</p> </li> <li> <p> <code>AllSquash</code>: Everyone is mapped to
	// anonymous user.</p> </li> </ul>
	Squash *string

	// A value that enables guessing of the MIME type for uploaded objects based on
	// file extensions. Set this value to true to enable MIME type guessing, otherwise
	// set to false. The default value is true.  <p>Valid Values: <code>true</code> |
	// <code>false</code> </p>
	GuessMIMETypeEnabled *bool

	// The name of the file share. Optional.  <note> <p> <code>FileShareName</code>
	// must be set if an S3 prefix name is set in <code>LocationARN</code>.</p> </note>
	FileShareName *string

	// A value that sets the access control list (ACL) permission for objects in the S3
	// bucket that a file gateway puts objects into. The default value is private.
	ObjectACL types.ObjectACL
}

// CreateNFSFileShareOutput
type CreateNFSFileShareOutput struct {

	// The Amazon Resource Name (ARN) of the newly created file share.
	FileShareARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateNFSFileShareMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateNFSFileShare{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateNFSFileShare{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateNFSFileShare(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "CreateNFSFileShare",
	}
}
