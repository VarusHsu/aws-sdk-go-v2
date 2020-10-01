// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the metrics for the specified solution version.
func (c *Client) GetSolutionMetrics(ctx context.Context, params *GetSolutionMetricsInput, optFns ...func(*Options)) (*GetSolutionMetricsOutput, error) {
	stack := middleware.NewStack("GetSolutionMetrics", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetSolutionMetricsMiddlewares(stack)
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
	addOpGetSolutionMetricsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSolutionMetrics(options.Region), middleware.Before)
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
			OperationName: "GetSolutionMetrics",
			Err:           err,
		}
	}
	out := result.(*GetSolutionMetricsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSolutionMetricsInput struct {

	// The Amazon Resource Name (ARN) of the solution version for which to get metrics.
	//
	// This member is required.
	SolutionVersionArn *string
}

type GetSolutionMetricsOutput struct {

	// The metrics for the solution version.
	Metrics map[string]*float64

	// The same solution version ARN as specified in the request.
	SolutionVersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetSolutionMetricsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetSolutionMetrics{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSolutionMetrics{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSolutionMetrics(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "GetSolutionMetrics",
	}
}
