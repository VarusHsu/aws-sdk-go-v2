// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all your destinations. The results are ASCII-sorted by destination name.
func (c *Client) DescribeDestinations(ctx context.Context, params *DescribeDestinationsInput, optFns ...func(*Options)) (*DescribeDestinationsOutput, error) {
	stack := middleware.NewStack("DescribeDestinations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeDestinationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDestinations(options.Region), middleware.Before)
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
			OperationName: "DescribeDestinations",
			Err:           err,
		}
	}
	out := result.(*DescribeDestinationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDestinationsInput struct {

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	// The maximum number of items returned. If you don't specify a value, the default
	// is up to 50 items.
	Limit *int32

	// The prefix to match. If you don't specify a value, no prefix filter is applied.
	DestinationNamePrefix *string
}

type DescribeDestinationsOutput struct {

	// The destinations.
	Destinations []*types.Destination

	// The token for the next set of items to return. The token expires after 24 hours.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeDestinationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDestinations{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDestinations{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDestinations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "DescribeDestinations",
	}
}
