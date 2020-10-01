// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the contents of the specified folder.
func (c *Client) DeleteFolderContents(ctx context.Context, params *DeleteFolderContentsInput, optFns ...func(*Options)) (*DeleteFolderContentsOutput, error) {
	stack := middleware.NewStack("DeleteFolderContents", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteFolderContentsMiddlewares(stack)
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
	addOpDeleteFolderContentsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFolderContents(options.Region), middleware.Before)
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
			OperationName: "DeleteFolderContents",
			Err:           err,
		}
	}
	out := result.(*DeleteFolderContentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteFolderContentsInput struct {

	// The ID of the folder.
	//
	// This member is required.
	FolderId *string

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string
}

type DeleteFolderContentsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteFolderContentsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteFolderContents{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteFolderContents{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteFolderContents(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "DeleteFolderContents",
	}
}
