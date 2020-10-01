// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of forecasts created using the CreateForecast () operation. For
// each forecast, this operation returns a summary of its properties, including its
// Amazon Resource Name (ARN). To retrieve the complete set of properties, specify
// the ARN with the DescribeForecast () operation. You can filter the list using an
// array of Filter () objects.
func (c *Client) ListForecasts(ctx context.Context, params *ListForecastsInput, optFns ...func(*Options)) (*ListForecastsOutput, error) {
	stack := middleware.NewStack("ListForecasts", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListForecastsMiddlewares(stack)
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
	addOpListForecastsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListForecasts(options.Region), middleware.Before)
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
			OperationName: "ListForecasts",
			Err:           err,
		}
	}
	out := result.(*ListForecastsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListForecastsInput struct {

	// An array of filters. For each filter, you provide a condition and a match
	// statement. The condition is either IS or IS_NOT, which specifies whether to
	// include or exclude the forecasts that match the statement from the list,
	// respectively. The match statement consists of a key and a value. Filter
	// properties
	//
	//     * Condition - The condition to apply. Valid values are IS and
	// IS_NOT. To include the forecasts that match the statement, specify IS. To
	// exclude matching forecasts, specify IS_NOT.
	//
	//     * Key - The name of the
	// parameter to filter on. Valid values are DatasetGroupArn, PredictorArn, and
	// Status.
	//
	//     * Value - The value to match.
	//
	// For example, to list all forecasts
	// whose status is not ACTIVE, you would specify: "Filters": [ { "Condition":
	// "IS_NOT", "Key": "Status", "Value": "ACTIVE" } ]
	Filters []*types.Filter

	// The number of items to return in the response.
	MaxResults *int32

	// If the result of the previous request was truncated, the response includes a
	// NextToken. To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string
}

type ListForecastsOutput struct {

	// If the response is truncated, Amazon Forecast returns this token. To retrieve
	// the next set of results, use the token in the next request.
	NextToken *string

	// An array of objects that summarize each forecast's properties.
	Forecasts []*types.ForecastSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListForecastsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListForecasts{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListForecasts{}, middleware.After)
}

func newServiceMetadataMiddleware_opListForecasts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "ListForecasts",
	}
}
