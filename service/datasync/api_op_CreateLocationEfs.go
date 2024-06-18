// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a transfer location for an Amazon EFS file system. DataSync can use
// this location as a source or destination for transferring data.
//
// Before you begin, make sure that you understand how DataSync [accesses Amazon EFS file systems].
//
// [accesses Amazon EFS file systems]: https://docs.aws.amazon.com/datasync/latest/userguide/create-efs-location.html#create-efs-location-access
func (c *Client) CreateLocationEfs(ctx context.Context, params *CreateLocationEfsInput, optFns ...func(*Options)) (*CreateLocationEfsOutput, error) {
	if params == nil {
		params = &CreateLocationEfsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLocationEfs", params, optFns, c.addOperationCreateLocationEfsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLocationEfsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CreateLocationEfsRequest
type CreateLocationEfsInput struct {

	// Specifies the subnet and security groups DataSync uses to access your Amazon
	// EFS file system.
	//
	// This member is required.
	Ec2Config *types.Ec2Config

	// Specifies the ARN for the Amazon EFS file system.
	//
	// This member is required.
	EfsFilesystemArn *string

	// Specifies the Amazon Resource Name (ARN) of the access point that DataSync uses
	// to access the Amazon EFS file system.
	AccessPointArn *string

	// Specifies an Identity and Access Management (IAM) role that DataSync assumes
	// when mounting the Amazon EFS file system.
	FileSystemAccessRoleArn *string

	// Specifies whether you want DataSync to use Transport Layer Security (TLS) 1.2
	// encryption when it copies data to or from the Amazon EFS file system.
	//
	// If you specify an access point using AccessPointArn or an IAM role using
	// FileSystemAccessRoleArn , you must set this parameter to TLS1_2 .
	InTransitEncryption types.EfsInTransitEncryption

	// Specifies a mount path for your Amazon EFS file system. This is where DataSync
	// reads or writes data (depending on if this is a source or destination location).
	// By default, DataSync uses the root directory, but you can also include
	// subdirectories.
	//
	// You must specify a value with forward slashes (for example, /path/to/folder ).
	Subdirectory *string

	// Specifies the key-value pair that represents a tag that you want to add to the
	// resource. The value can be an empty string. This value helps you manage, filter,
	// and search for your resources. We recommend that you create a name tag for your
	// location.
	Tags []types.TagListEntry

	noSmithyDocumentSerde
}

// CreateLocationEfs
type CreateLocationEfsOutput struct {

	// The Amazon Resource Name (ARN) of the Amazon EFS file system location that you
	// create.
	LocationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLocationEfsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLocationEfs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLocationEfs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLocationEfs"); err != nil {
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
	if err = addOpCreateLocationEfsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLocationEfs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLocationEfs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLocationEfs",
	}
}
