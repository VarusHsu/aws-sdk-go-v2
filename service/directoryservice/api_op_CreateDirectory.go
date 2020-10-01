// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a Simple AD directory. For more information, see Simple Active Directory
// (https://docs.aws.amazon.com/directoryservice/latest/admin-guide/directory_simple_ad.html)
// in the AWS Directory Service Admin Guide. Before you call CreateDirectory,
// ensure that all of the required permissions have been explicitly granted through
// a policy. For details about what permissions are required to run the
// CreateDirectory operation, see AWS Directory Service API Permissions: Actions,
// Resources, and Conditions Reference
// (http://docs.aws.amazon.com/directoryservice/latest/admin-guide/UsingWithDS_IAM_ResourcePermissions.html).
func (c *Client) CreateDirectory(ctx context.Context, params *CreateDirectoryInput, optFns ...func(*Options)) (*CreateDirectoryOutput, error) {
	stack := middleware.NewStack("CreateDirectory", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateDirectoryMiddlewares(stack)
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
	addOpCreateDirectoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDirectory(options.Region), middleware.Before)
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
			OperationName: "CreateDirectory",
			Err:           err,
		}
	}
	out := result.(*CreateDirectoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the CreateDirectory () operation.
type CreateDirectoryInput struct {

	// A description for the directory.
	Description *string

	// The size of the directory.
	//
	// This member is required.
	Size types.DirectorySize

	// The fully qualified name for the directory, such as corp.example.com.
	//
	// This member is required.
	Name *string

	// The password for the directory administrator. The directory creation process
	// creates a directory administrator account with the user name Administrator and
	// this password. If you need to change the password for the administrator account,
	// you can use the ResetUserPassword () API call.
	//
	// This member is required.
	Password *string

	// A DirectoryVpcSettings () object that contains additional information for the
	// operation.
	VpcSettings *types.DirectoryVpcSettings

	// The NetBIOS name of the directory, such as CORP.
	ShortName *string

	// The tags to be assigned to the Simple AD directory.
	Tags []*types.Tag
}

// Contains the results of the CreateDirectory () operation.
type CreateDirectoryOutput struct {

	// The identifier of the directory that was created.
	DirectoryId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateDirectoryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDirectory{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDirectory{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDirectory(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "CreateDirectory",
	}
}
