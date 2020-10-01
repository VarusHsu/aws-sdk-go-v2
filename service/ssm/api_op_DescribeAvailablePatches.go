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

// Lists all patches eligible to be included in a patch baseline.
func (c *Client) DescribeAvailablePatches(ctx context.Context, params *DescribeAvailablePatchesInput, optFns ...func(*Options)) (*DescribeAvailablePatchesOutput, error) {
	stack := middleware.NewStack("DescribeAvailablePatches", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeAvailablePatchesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAvailablePatches(options.Region), middleware.Before)
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
			OperationName: "DescribeAvailablePatches",
			Err:           err,
		}
	}
	out := result.(*DescribeAvailablePatchesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAvailablePatchesInput struct {

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	// The maximum number of patches to return (per page).
	MaxResults *int32

	// Filters used to scope down the returned patches.
	Filters []*types.PatchOrchestratorFilter
}

type DescribeAvailablePatchesOutput struct {

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// An array of patches. Each entry in the array is a patch structure.
	Patches []*types.Patch

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeAvailablePatchesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAvailablePatches{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAvailablePatches{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAvailablePatches(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeAvailablePatches",
	}
}
