// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the reservation utilization for your account. Management account in
// an organization have access to member accounts. You can filter data by
// dimensions in a time period. You can use GetDimensionValues to determine the
// possible dimension values. Currently, you can group only by SUBSCRIPTION_ID .
func (c *Client) GetReservationUtilization(ctx context.Context, params *GetReservationUtilizationInput, optFns ...func(*Options)) (*GetReservationUtilizationOutput, error) {
	if params == nil {
		params = &GetReservationUtilizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetReservationUtilization", params, optFns, c.addOperationGetReservationUtilizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetReservationUtilizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetReservationUtilizationInput struct {

	// Sets the start and end dates for retrieving Reserved Instance (RI) utilization.
	// The start date is inclusive, but the end date is exclusive. For example, if
	// start is 2017-01-01 and end is 2017-05-01 , then the cost and usage data is
	// retrieved from 2017-01-01 up to and including 2017-04-30 but not including
	// 2017-05-01 .
	//
	// This member is required.
	TimePeriod *types.DateInterval

	// Filters utilization data by dimensions. You can filter by the following
	// dimensions:
	//   - AZ
	//   - CACHE_ENGINE
	//   - DEPLOYMENT_OPTION
	//   - INSTANCE_TYPE
	//   - LINKED_ACCOUNT
	//   - OPERATING_SYSTEM
	//   - PLATFORM
	//   - REGION
	//   - SERVICE
	//   - SCOPE
	//   - TENANCY
	// GetReservationUtilization uses the same Expression (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html)
	// object as the other operations, but only AND is supported among each dimension,
	// and nesting is supported up to only one level deep. If there are multiple values
	// for a dimension, they are OR'd together.
	Filter *types.Expression

	// If GroupBy is set, Granularity can't be set. If Granularity isn't set, the
	// response object doesn't include Granularity , either MONTHLY or DAILY . If both
	// GroupBy and Granularity aren't set, GetReservationUtilization defaults to DAILY
	// . The GetReservationUtilization operation supports only DAILY and MONTHLY
	// granularities.
	Granularity types.Granularity

	// Groups only by SUBSCRIPTION_ID . Metadata is included.
	GroupBy []types.GroupDefinition

	// The maximum number of objects that you returned for this request. If more
	// objects are available, in the response, Amazon Web Services provides a
	// NextPageToken value that you can use in a subsequent call to get the next batch
	// of objects.
	MaxResults *int32

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextPageToken *string

	// The value that you want to sort the data by. The following values are supported
	// for Key :
	//   - UtilizationPercentage
	//   - UtilizationPercentageInUnits
	//   - PurchasedHours
	//   - PurchasedUnits
	//   - TotalActualHours
	//   - TotalActualUnits
	//   - UnusedHours
	//   - UnusedUnits
	//   - OnDemandCostOfRIHoursUsed
	//   - NetRISavings
	//   - TotalPotentialRISavings
	//   - AmortizedUpfrontFee
	//   - AmortizedRecurringFee
	//   - TotalAmortizedFee
	//   - RICostForUnusedHours
	//   - RealizedSavings
	//   - UnrealizedSavings
	// The supported values for SortOrder are ASCENDING and DESCENDING .
	SortBy *types.SortDefinition

	noSmithyDocumentSerde
}

type GetReservationUtilizationOutput struct {

	// The amount of time that you used your Reserved Instances (RIs).
	//
	// This member is required.
	UtilizationsByTime []types.UtilizationByTime

	// The token for the next set of retrievable results. Amazon Web Services provides
	// the token when the response from a previous call has more results than the
	// maximum page size.
	NextPageToken *string

	// The total amount of time that you used your Reserved Instances (RIs).
	Total *types.ReservationAggregates

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetReservationUtilizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetReservationUtilization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetReservationUtilization{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetReservationUtilization"); err != nil {
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
	if err = addOpGetReservationUtilizationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetReservationUtilization(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetReservationUtilization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetReservationUtilization",
	}
}
