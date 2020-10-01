// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lets you enable Insights event logging by specifying the Insights selectors that
// you want to enable on an existing trail. You also use PutInsightSelectors to
// turn off Insights event logging, by passing an empty list of insight types. In
// this release, only ApiCallRateInsight is supported as an Insights selector.
func (c *Client) PutInsightSelectors(ctx context.Context, params *PutInsightSelectorsInput, optFns ...func(*Options)) (*PutInsightSelectorsOutput, error) {
	stack := middleware.NewStack("PutInsightSelectors", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutInsightSelectorsMiddlewares(stack)
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
	addOpPutInsightSelectorsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutInsightSelectors(options.Region), middleware.Before)
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
			OperationName: "PutInsightSelectors",
			Err:           err,
		}
	}
	out := result.(*PutInsightSelectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutInsightSelectorsInput struct {

	// The name of the CloudTrail trail for which you want to change or add Insights
	// selectors.
	//
	// This member is required.
	TrailName *string

	// A JSON string that contains the insight types you want to log on a trail. In
	// this release, only ApiCallRateInsight is supported as an insight type.
	//
	// This member is required.
	InsightSelectors []*types.InsightSelector
}

type PutInsightSelectorsOutput struct {

	// A JSON string that contains the insight types you want to log on a trail. In
	// this release, only ApiCallRateInsight is supported as an insight type.
	InsightSelectors []*types.InsightSelector

	// The Amazon Resource Name (ARN) of a trail for which you want to change or add
	// Insights selectors.
	TrailARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutInsightSelectorsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutInsightSelectors{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutInsightSelectors{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutInsightSelectors(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "PutInsightSelectors",
	}
}
