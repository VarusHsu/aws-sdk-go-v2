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

// View a summary of OpsItems based on specified filters and aggregators.
func (c *Client) GetOpsSummary(ctx context.Context, params *GetOpsSummaryInput, optFns ...func(*Options)) (*GetOpsSummaryOutput, error) {
	stack := middleware.NewStack("GetOpsSummary", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetOpsSummaryMiddlewares(stack)
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
	addOpGetOpsSummaryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetOpsSummary(options.Region), middleware.Before)
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
			OperationName: "GetOpsSummary",
			Err:           err,
		}
	}
	out := result.(*GetOpsSummaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOpsSummaryInput struct {

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// Optional aggregators that return counts of OpsItems based on one or more
	// expressions.
	Aggregators []*types.OpsAggregator

	// Optional filters used to scope down the returned OpsItems.
	Filters []*types.OpsFilter

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string

	// Specify the name of a resource data sync to get.
	SyncName *string

	// The OpsItem data type to return.
	ResultAttributes []*types.OpsResultAttribute
}

type GetOpsSummaryOutput struct {

	// The list of aggregated and filtered OpsItems.
	Entities []*types.OpsEntity

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetOpsSummaryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetOpsSummary{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetOpsSummary{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetOpsSummary(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetOpsSummary",
	}
}
