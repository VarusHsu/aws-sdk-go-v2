// Code generated by smithy-go-codegen DO NOT EDIT.

package signer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/signer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information on a specific signing platform.
func (c *Client) GetSigningPlatform(ctx context.Context, params *GetSigningPlatformInput, optFns ...func(*Options)) (*GetSigningPlatformOutput, error) {
	stack := middleware.NewStack("GetSigningPlatform", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetSigningPlatformMiddlewares(stack)
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
	addOpGetSigningPlatformValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSigningPlatform(options.Region), middleware.Before)
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
			OperationName: "GetSigningPlatform",
			Err:           err,
		}
	}
	out := result.(*GetSigningPlatformOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSigningPlatformInput struct {

	// The ID of the target signing platform.
	//
	// This member is required.
	PlatformId *string
}

type GetSigningPlatformOutput struct {

	// A list of partner entities that use the target signing platform.
	Partner *string

	// The ID of the target signing platform.
	PlatformId *string

	// A list of configurations applied to the target platform at signing.
	SigningConfiguration *types.SigningConfiguration

	// The category type of the target signing platform.
	Category types.Category

	// The format of the target platform's signing image.
	SigningImageFormat *types.SigningImageFormat

	// The maximum size (in MB) of the payload that can be signed by the target
	// platform.
	MaxSizeInMB *int32

	// The validation template that is used by the target signing platform.
	Target *string

	// The display name of the target signing platform.
	DisplayName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetSigningPlatformMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetSigningPlatform{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSigningPlatform{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSigningPlatform(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "signer",
		OperationName: "GetSigningPlatform",
	}
}
