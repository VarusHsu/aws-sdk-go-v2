// Code generated by smithy-go-codegen DO NOT EDIT.

package forecastquery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/forecastquery/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a forecast for a single item, filtered by the supplied criteria.
//
// The criteria is a key-value pair. The key is either item_id (or the equivalent
// non-timestamp, non-target field) from the TARGET_TIME_SERIES dataset, or one of
// the forecast dimensions specified as part of the FeaturizationConfig object.
//
// By default, QueryForecast returns the complete date range for the filtered
// forecast. You can request a specific date range.
//
// To get the full forecast, use the [CreateForecastExportJob] operation.
//
// The forecasts generated by Amazon Forecast are in the same timezone as the
// dataset that was used to create the predictor.
//
// [CreateForecastExportJob]: https://docs.aws.amazon.com/en_us/forecast/latest/dg/API_CreateForecastExportJob.html
func (c *Client) QueryForecast(ctx context.Context, params *QueryForecastInput, optFns ...func(*Options)) (*QueryForecastOutput, error) {
	if params == nil {
		params = &QueryForecastInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "QueryForecast", params, optFns, c.addOperationQueryForecastMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*QueryForecastOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type QueryForecastInput struct {

	// The filtering criteria to apply when retrieving the forecast. For example, to
	// get the forecast for client_21 in the electricity usage dataset, specify the
	// following:
	//
	//     {"item_id" : "client_21"}
	//
	// To get the full forecast, use the [CreateForecastExportJob] operation.
	//
	// [CreateForecastExportJob]: https://docs.aws.amazon.com/en_us/forecast/latest/dg/API_CreateForecastExportJob.html
	//
	// This member is required.
	Filters map[string]string

	// The Amazon Resource Name (ARN) of the forecast to query.
	//
	// This member is required.
	ForecastArn *string

	// The end date for the forecast. Specify the date using this format:
	// yyyy-MM-dd'T'HH:mm:ss (ISO 8601 format). For example, 2015-01-01T20:00:00.
	EndDate *string

	// If the result of the previous request was truncated, the response includes a
	// NextToken . To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string

	// The start date for the forecast. Specify the date using this format:
	// yyyy-MM-dd'T'HH:mm:ss (ISO 8601 format). For example, 2015-01-01T08:00:00.
	StartDate *string

	noSmithyDocumentSerde
}

type QueryForecastOutput struct {

	// The forecast.
	Forecast *types.Forecast

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationQueryForecastMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpQueryForecast{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpQueryForecast{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "QueryForecast"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpQueryForecastValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opQueryForecast(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opQueryForecast(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "QueryForecast",
	}
}
