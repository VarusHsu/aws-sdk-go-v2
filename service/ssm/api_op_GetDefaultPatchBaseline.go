// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the default patch baseline. Note that Systems Manager supports
// creating multiple default patch baselines. For example, you can create a default
// patch baseline for each operating system. If you do not specify an operating
// system value, the default patch baseline for Windows is returned.
func (c *Client) GetDefaultPatchBaseline(ctx context.Context, params *GetDefaultPatchBaselineInput, optFns ...func(*Options)) (*GetDefaultPatchBaselineOutput, error) {
	stack := middleware.NewStack("GetDefaultPatchBaseline", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetDefaultPatchBaselineMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDefaultPatchBaseline(options.Region), middleware.Before)
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
			OperationName: "GetDefaultPatchBaseline",
			Err:           err,
		}
	}
	out := result.(*GetDefaultPatchBaselineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDefaultPatchBaselineInput struct {

	// Returns the default patch baseline for the specified operating system.
	OperatingSystem types.OperatingSystem
}

type GetDefaultPatchBaselineOutput struct {

	// The operating system for the returned patch baseline.
	OperatingSystem types.OperatingSystem

	// The ID of the default patch baseline.
	BaselineId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetDefaultPatchBaselineMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetDefaultPatchBaseline{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDefaultPatchBaseline{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDefaultPatchBaseline(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetDefaultPatchBaseline",
	}
}
