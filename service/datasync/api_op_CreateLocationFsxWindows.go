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

// Creates a transfer location for an Amazon FSx for Windows File Server file
// system. DataSync can use this location as a source or destination for
// transferring data.
//
// Before you begin, make sure that you understand how DataSync [accesses FSx for Windows File Server file systems].
//
// [accesses FSx for Windows File Server file systems]: https://docs.aws.amazon.com/datasync/latest/userguide/create-fsx-location.html#create-fsx-location-access
func (c *Client) CreateLocationFsxWindows(ctx context.Context, params *CreateLocationFsxWindowsInput, optFns ...func(*Options)) (*CreateLocationFsxWindowsOutput, error) {
	if params == nil {
		params = &CreateLocationFsxWindowsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLocationFsxWindows", params, optFns, c.addOperationCreateLocationFsxWindowsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLocationFsxWindowsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLocationFsxWindowsInput struct {

	// Specifies the Amazon Resource Name (ARN) for the FSx for Windows File Server
	// file system.
	//
	// This member is required.
	FsxFilesystemArn *string

	// Specifies the password of the user with the permissions to mount and access the
	// files, folders, and file metadata in your FSx for Windows File Server file
	// system.
	//
	// This member is required.
	Password *string

	// Specifies the ARNs of the security groups that provide access to your file
	// system's preferred subnet.
	//
	// If you choose a security group that doesn't allow connections from within
	// itself, do one of the following:
	//
	//   - Configure the security group to allow it to communicate within itself.
	//
	//   - Choose a different security group that can communicate with the mount
	//   target's security group.
	//
	// This member is required.
	SecurityGroupArns []string

	// Specifies the user with the permissions to mount and access the files, folders,
	// and file metadata in your FSx for Windows File Server file system.
	//
	// For information about choosing a user with the right level of access for your
	// transfer, see [required permissions]for FSx for Windows File Server locations.
	//
	// [required permissions]: https://docs.aws.amazon.com/datasync/latest/userguide/create-fsx-location.html#create-fsx-windows-location-permissions
	//
	// This member is required.
	User *string

	// Specifies the name of the Microsoft Active Directory domain that the FSx for
	// Windows File Server file system belongs to.
	//
	// If you have multiple Active Directory domains in your environment, configuring
	// this parameter makes sure that DataSync connects to the right file system.
	Domain *string

	// Specifies a mount path for your file system using forward slashes. This is
	// where DataSync reads or writes data (depending on if this is a source or
	// destination location).
	Subdirectory *string

	// Specifies labels that help you categorize, filter, and search for your Amazon
	// Web Services resources. We recommend creating at least a name tag for your
	// location.
	Tags []types.TagListEntry

	noSmithyDocumentSerde
}

type CreateLocationFsxWindowsOutput struct {

	// The ARN of the FSx for Windows File Server file system location you created.
	LocationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLocationFsxWindowsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLocationFsxWindows{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLocationFsxWindows{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLocationFsxWindows"); err != nil {
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
	if err = addOpCreateLocationFsxWindowsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLocationFsxWindows(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLocationFsxWindows(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLocationFsxWindows",
	}
}
