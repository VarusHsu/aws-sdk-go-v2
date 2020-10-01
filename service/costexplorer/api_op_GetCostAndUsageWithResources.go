// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves cost and usage metrics with resources for your account. You can
// specify which cost and usage-related metric, such as BlendedCosts or
// UsageQuantity, that you want the request to return. You can also filter and
// group your data by various dimensions, such as SERVICE or AZ, in a specific time
// range. For a complete list of valid dimensions, see the GetDimensionValues
// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_GetDimensionValues.html)
// operation. Master accounts in an organization in AWS Organizations have access
// to all member accounts. This API is currently available for the Amazon Elastic
// Compute Cloud – Compute service only. This is an opt-in only feature. You can
// enable this feature from the Cost Explorer Settings page. For information on how
// to access the Settings page, see Controlling Access for Cost Explorer
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/ce-access.html) in
// the AWS Billing and Cost Management User Guide.
func (c *Client) GetCostAndUsageWithResources(ctx context.Context, params *GetCostAndUsageWithResourcesInput, optFns ...func(*Options)) (*GetCostAndUsageWithResourcesOutput, error) {
	stack := middleware.NewStack("GetCostAndUsageWithResources", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetCostAndUsageWithResourcesMiddlewares(stack)
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
	addOpGetCostAndUsageWithResourcesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCostAndUsageWithResources(options.Region), middleware.Before)
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
			OperationName: "GetCostAndUsageWithResources",
			Err:           err,
		}
	}
	out := result.(*GetCostAndUsageWithResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCostAndUsageWithResourcesInput struct {

	// You can group Amazon Web Services costs using up to two different groups: either
	// dimensions, tag keys, or both.
	GroupBy []*types.GroupDefinition

	// Filters Amazon Web Services costs by different dimensions. For example, you can
	// specify SERVICE and LINKED_ACCOUNT and get the costs that are associated with
	// that account's usage of that service. You can nest Expression objects to define
	// any combination of dimension filters. For more information, see Expression
	// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html).
	// The GetCostAndUsageWithResources operation requires that you either group by or
	// filter by a ResourceId.
	Filter *types.Expression

	// Which metrics are returned in the query. For more information about blended and
	// unblended rates, see Why does the "blended" annotation appear on some line items
	// in my bill?
	// (http://aws.amazon.com/premiumsupport/knowledge-center/blended-rates-intro/).
	// Valid values are AmortizedCost, BlendedCost, NetAmortizedCost, NetUnblendedCost,
	// NormalizedUsageAmount, UnblendedCost, and UsageQuantity. If you return the
	// UsageQuantity metric, the service aggregates all usage numbers without taking
	// the units into account. For example, if you aggregate usageQuantity across all
	// of Amazon EC2, the results aren't meaningful because Amazon EC2 compute hours
	// and data transfer are measured in different units (for example, hours vs. GB).
	// To get more meaningful UsageQuantity metrics, filter by UsageType or
	// UsageTypeGroups. Metrics is required for GetCostAndUsageWithResources requests.
	Metrics []*string

	// Sets the start and end dates for retrieving Amazon Web Services costs. The range
	// must be within the last 14 days (the start date cannot be earlier than 14 days
	// ago). The start date is inclusive, but the end date is exclusive. For example,
	// if start is 2017-01-01 and end is 2017-05-01, then the cost and usage data is
	// retrieved from 2017-01-01 up to and including 2017-04-30 but not including
	// 2017-05-01.
	//
	// This member is required.
	TimePeriod *types.DateInterval

	// The token to retrieve the next set of results. AWS provides the token when the
	// response from a previous call has more results than the maximum page size.
	NextPageToken *string

	// Sets the AWS cost granularity to MONTHLY, DAILY, or HOURLY. If Granularity isn't
	// set, the response object doesn't include the Granularity, MONTHLY, DAILY, or
	// HOURLY.
	Granularity types.Granularity
}

type GetCostAndUsageWithResourcesOutput struct {

	// The groups that are specified by the Filter or GroupBy parameters in the
	// request.
	GroupDefinitions []*types.GroupDefinition

	// The token for the next set of retrievable results. AWS provides the token when
	// the response from a previous call has more results than the maximum page size.
	NextPageToken *string

	// The time period that is covered by the results in the response.
	ResultsByTime []*types.ResultByTime

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetCostAndUsageWithResourcesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetCostAndUsageWithResources{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCostAndUsageWithResources{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCostAndUsageWithResources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "GetCostAndUsageWithResources",
	}
}
