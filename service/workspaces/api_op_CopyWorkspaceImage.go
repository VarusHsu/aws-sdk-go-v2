// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Copies the specified image from the specified Region to the current Region.
func (c *Client) CopyWorkspaceImage(ctx context.Context, params *CopyWorkspaceImageInput, optFns ...func(*Options)) (*CopyWorkspaceImageOutput, error) {
	stack := middleware.NewStack("CopyWorkspaceImage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCopyWorkspaceImageMiddlewares(stack)
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
	addOpCopyWorkspaceImageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCopyWorkspaceImage(options.Region), middleware.Before)
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
			OperationName: "CopyWorkspaceImage",
			Err:           err,
		}
	}
	out := result.(*CopyWorkspaceImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopyWorkspaceImageInput struct {

	// The name of the image.
	//
	// This member is required.
	Name *string

	// The identifier of the source image.
	//
	// This member is required.
	SourceImageId *string

	// The tags for the image.
	Tags []*types.Tag

	// The identifier of the source Region.
	//
	// This member is required.
	SourceRegion *string

	// A description of the image.
	Description *string
}

type CopyWorkspaceImageOutput struct {

	// The identifier of the image.
	ImageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCopyWorkspaceImageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCopyWorkspaceImage{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCopyWorkspaceImage{}, middleware.After)
}

func newServiceMetadataMiddleware_opCopyWorkspaceImage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "CopyWorkspaceImage",
	}
}
