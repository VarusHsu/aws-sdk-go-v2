// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves (queries) aggregated usage data for an account.
func (c *Client) GetUsageTotals(ctx context.Context, params *GetUsageTotalsInput, optFns ...func(*Options)) (*GetUsageTotalsOutput, error) {
	if params == nil {
		params = &GetUsageTotalsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUsageTotals", params, optFns, c.addOperationGetUsageTotalsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUsageTotalsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUsageTotalsInput struct {

	// The inclusive time period to retrieve the data for. Valid values are:
	// MONTH_TO_DATE, for the current calendar month to date; and, PAST_30_DAYS, for
	// the preceding 30 days. If you don't specify a value for this parameter, Amazon
	// Macie provides aggregated usage data for the preceding 30 days.
	TimeRange *string

	noSmithyDocumentSerde
}

type GetUsageTotalsOutput struct {

	// The inclusive time period that the usage data applies to. Possible values are:
	// MONTH_TO_DATE, for the current calendar month to date; and, PAST_30_DAYS, for
	// the preceding 30 days.
	TimeRange types.TimeRange

	// An array of objects that contains the results of the query. Each object
	// contains the data for a specific usage metric.
	UsageTotals []types.UsageTotal

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUsageTotalsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetUsageTotals{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUsageTotals{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetUsageTotals"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUsageTotals(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetUsageTotals(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetUsageTotals",
	}
}
