// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves estimated usage records for hourly granularity or resource-level data
// at daily granularity.
func (c *Client) GetApproximateUsageRecords(ctx context.Context, params *GetApproximateUsageRecordsInput, optFns ...func(*Options)) (*GetApproximateUsageRecordsOutput, error) {
	if params == nil {
		params = &GetApproximateUsageRecordsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetApproximateUsageRecords", params, optFns, c.addOperationGetApproximateUsageRecordsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetApproximateUsageRecordsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetApproximateUsageRecordsInput struct {

	// The service to evaluate for the usage records. You can choose resource-level
	// data at daily granularity, or hourly granularity with or without resource-level
	// data.
	//
	// This member is required.
	ApproximationDimension types.ApproximationDimension

	// How granular you want the data to be. You can enable data at hourly or daily
	// granularity.
	//
	// This member is required.
	Granularity types.Granularity

	// The service metadata for the service or services you want to query. If not
	// specified, all elements are returned.
	Services []string

	noSmithyDocumentSerde
}

type GetApproximateUsageRecordsOutput struct {

	// The lookback period that's used for the estimation.
	LookbackPeriod *types.DateInterval

	// The service metadata for the service or services in the response.
	Services map[string]int64

	// The total number of usage records for all services in the services list.
	TotalRecords int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetApproximateUsageRecordsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetApproximateUsageRecords{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetApproximateUsageRecords{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetApproximateUsageRecords"); err != nil {
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
	if err = addOpGetApproximateUsageRecordsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetApproximateUsageRecords(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetApproximateUsageRecords(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetApproximateUsageRecords",
	}
}
