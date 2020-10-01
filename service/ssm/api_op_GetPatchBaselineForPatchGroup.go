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

// Retrieves the patch baseline that should be used for the specified patch group.
func (c *Client) GetPatchBaselineForPatchGroup(ctx context.Context, params *GetPatchBaselineForPatchGroupInput, optFns ...func(*Options)) (*GetPatchBaselineForPatchGroupOutput, error) {
	stack := middleware.NewStack("GetPatchBaselineForPatchGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetPatchBaselineForPatchGroupMiddlewares(stack)
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
	addOpGetPatchBaselineForPatchGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetPatchBaselineForPatchGroup(options.Region), middleware.Before)
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
			OperationName: "GetPatchBaselineForPatchGroup",
			Err:           err,
		}
	}
	out := result.(*GetPatchBaselineForPatchGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPatchBaselineForPatchGroupInput struct {

	// The name of the patch group whose patch baseline should be retrieved.
	//
	// This member is required.
	PatchGroup *string

	// Returns he operating system rule specified for patch groups using the patch
	// baseline.
	OperatingSystem types.OperatingSystem
}

type GetPatchBaselineForPatchGroupOutput struct {

	// The name of the patch group.
	PatchGroup *string

	// The ID of the patch baseline that should be used for the patch group.
	BaselineId *string

	// The operating system rule specified for patch groups using the patch baseline.
	OperatingSystem types.OperatingSystem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetPatchBaselineForPatchGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetPatchBaselineForPatchGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPatchBaselineForPatchGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetPatchBaselineForPatchGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetPatchBaselineForPatchGroup",
	}
}
