// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an AWS Directory Service directory. Before you call DeleteDirectory,
// ensure that all of the required permissions have been explicitly granted through
// a policy. For details about what permissions are required to run the
// DeleteDirectory operation, see AWS Directory Service API Permissions: Actions,
// Resources, and Conditions Reference
// (http://docs.aws.amazon.com/directoryservice/latest/admin-guide/UsingWithDS_IAM_ResourcePermissions.html).
func (c *Client) DeleteDirectory(ctx context.Context, params *DeleteDirectoryInput, optFns ...func(*Options)) (*DeleteDirectoryOutput, error) {
	stack := middleware.NewStack("DeleteDirectory", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteDirectoryMiddlewares(stack)
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
	addOpDeleteDirectoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDirectory(options.Region), middleware.Before)
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
			OperationName: "DeleteDirectory",
			Err:           err,
		}
	}
	out := result.(*DeleteDirectoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the DeleteDirectory () operation.
type DeleteDirectoryInput struct {

	// The identifier of the directory to delete.
	//
	// This member is required.
	DirectoryId *string
}

// Contains the results of the DeleteDirectory () operation.
type DeleteDirectoryOutput struct {

	// The directory identifier.
	DirectoryId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteDirectoryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteDirectory{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteDirectory{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteDirectory(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "DeleteDirectory",
	}
}
