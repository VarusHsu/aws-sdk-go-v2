// Code generated by smithy-go-codegen DO NOT EDIT.

package forecastquery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecastquery/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a what-if forecast.
func (c *Client) QueryWhatIfForecast(ctx context.Context, params *QueryWhatIfForecastInput, optFns ...func(*Options)) (*QueryWhatIfForecastOutput, error) {
	if params == nil {
		params = &QueryWhatIfForecastInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "QueryWhatIfForecast", params, optFns, c.addOperationQueryWhatIfForecastMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*QueryWhatIfForecastOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type QueryWhatIfForecastInput struct {

	// The filtering criteria to apply when retrieving the forecast. For example, to
	// get the forecast for client_21 in the electricity usage dataset, specify the
	// following: {"item_id" : "client_21"} To get the full what-if forecast, use the
	// CreateForecastExportJob (https://docs.aws.amazon.com/en_us/forecast/latest/dg/API_CreateWhatIfForecastExport.html)
	// operation.
	//
	// This member is required.
	Filters map[string]string

	// The Amazon Resource Name (ARN) of the what-if forecast to query.
	//
	// This member is required.
	WhatIfForecastArn *string

	// The end date for the what-if forecast. Specify the date using this format:
	// yyyy-MM-dd'T'HH:mm:ss (ISO 8601 format). For example, 2015-01-01T20:00:00.
	EndDate *string

	// If the result of the previous request was truncated, the response includes a
	// NextToken . To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string

	// The start date for the what-if forecast. Specify the date using this format:
	// yyyy-MM-dd'T'HH:mm:ss (ISO 8601 format). For example, 2015-01-01T08:00:00.
	StartDate *string

	noSmithyDocumentSerde
}

type QueryWhatIfForecastOutput struct {

	// Provides information about a forecast. Returned as part of the QueryForecast
	// response.
	Forecast *types.Forecast

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationQueryWhatIfForecastMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpQueryWhatIfForecast{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpQueryWhatIfForecast{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "QueryWhatIfForecast"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpQueryWhatIfForecastValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opQueryWhatIfForecast(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opQueryWhatIfForecast(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "QueryWhatIfForecast",
	}
}
