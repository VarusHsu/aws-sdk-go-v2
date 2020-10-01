// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the specified attributes of the specified folder. The user must have
// access to both the folder and its parent folder, if applicable.
func (c *Client) UpdateFolder(ctx context.Context, params *UpdateFolderInput, optFns ...func(*Options)) (*UpdateFolderOutput, error) {
	stack := middleware.NewStack("UpdateFolder", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateFolderMiddlewares(stack)
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
	addOpUpdateFolderValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFolder(options.Region), middleware.Before)
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
			OperationName: "UpdateFolder",
			Err:           err,
		}
	}
	out := result.(*UpdateFolderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFolderInput struct {

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string

	// The resource state of the folder. Only ACTIVE and RECYCLED are accepted values
	// from the API.
	ResourceState types.ResourceStateType

	// The name of the folder.
	Name *string

	// The ID of the parent folder.
	ParentFolderId *string

	// The ID of the folder.
	//
	// This member is required.
	FolderId *string
}

type UpdateFolderOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateFolderMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFolder{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFolder{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateFolder(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "UpdateFolder",
	}
}
