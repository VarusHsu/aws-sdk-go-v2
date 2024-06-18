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

// Retrieves an array of Cost Category names and values incurred cost.
//
// If some Cost Category names and values are not associated with any cost, they
// will not be returned by this API.
func (c *Client) GetCostCategories(ctx context.Context, params *GetCostCategoriesInput, optFns ...func(*Options)) (*GetCostCategoriesOutput, error) {
	if params == nil {
		params = &GetCostCategoriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCostCategories", params, optFns, c.addOperationGetCostCategoriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCostCategoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCostCategoriesInput struct {

	// The time period of the request.
	//
	// This member is required.
	TimePeriod *types.DateInterval

	// The unique name of the Cost Category.
	CostCategoryName *string

	// Use Expression to filter in various Cost Explorer APIs.
	//
	// Not all Expression types are supported in each API. Refer to the documentation
	// for each specific API to see what is supported.
	//
	// There are two patterns:
	//
	//   - Simple dimension values.
	//
	//   - There are three types of simple dimension values: CostCategories , Tags ,
	//   and Dimensions .
	//
	//   - Specify the CostCategories field to define a filter that acts on Cost
	//   Categories.
	//
	//   - Specify the Tags field to define a filter that acts on Cost Allocation Tags.
	//
	//   - Specify the Dimensions field to define a filter that acts on the [DimensionValues]
	//   DimensionValues .
	//
	//   - For each filter type, you can set the dimension name and values for the
	//   filters that you plan to use.
	//
	//   - For example, you can filter for REGION==us-east-1 OR REGION==us-west-1 . For
	//   GetRightsizingRecommendation , the Region is a full name (for example,
	//   REGION==US East (N. Virginia) .
	//
	//   - The corresponding Expression for this example is as follows: {
	//   "Dimensions": { "Key": "REGION", "Values": [ "us-east-1", "us-west-1" ] } }
	//
	//   - As shown in the previous example, lists of dimension values are combined
	//   with OR when applying the filter.
	//
	//   - You can also set different match options to further control how the filter
	//   behaves. Not all APIs support match options. Refer to the documentation for each
	//   specific API to see what is supported.
	//
	//   - For example, you can filter for linked account names that start with "a".
	//
	//   - The corresponding Expression for this example is as follows: {
	//   "Dimensions": { "Key": "LINKED_ACCOUNT_NAME", "MatchOptions": [ "STARTS_WITH" ],
	//   "Values": [ "a" ] } }
	//
	//   - Compound Expression types with logical operations.
	//
	//   - You can use multiple Expression types and the logical operators AND/OR/NOT
	//   to create a list of one or more Expression objects. By doing this, you can
	//   filter by more advanced options.
	//
	//   - For example, you can filter by ((REGION == us-east-1 OR REGION ==
	//   us-west-1) OR (TAG.Type == Type1)) AND (USAGE_TYPE != DataTransfer) .
	//
	//   - The corresponding Expression for this example is as follows: { "And": [
	//   {"Or": [ {"Dimensions": { "Key": "REGION", "Values": [ "us-east-1", "us-west-1"
	//   ] }}, {"Tags": { "Key": "TagName", "Values": ["Value1"] } } ]}, {"Not":
	//   {"Dimensions": { "Key": "USAGE_TYPE", "Values": ["DataTransfer"] }}} ] }
	//
	// Because each Expression can have only one operator, the service returns an error
	//   if more than one is specified. The following example shows an Expression
	//   object that creates an error: { "And": [ ... ], "Dimensions": { "Key":
	//   "USAGE_TYPE", "Values": [ "DataTransfer" ] } }
	//
	// The following is an example of the corresponding error message: "Expression has
	//   more than one roots. Only one root operator is allowed for each expression: And,
	//   Or, Not, Dimensions, Tags, CostCategories"
	//
	// For the GetRightsizingRecommendation action, a combination of OR and NOT isn't
	// supported. OR isn't supported between different dimensions, or dimensions and
	// tags. NOT operators aren't supported. Dimensions are also limited to
	// LINKED_ACCOUNT , REGION , or RIGHTSIZING_TYPE .
	//
	// For the GetReservationPurchaseRecommendation action, only NOT is supported. AND
	// and OR aren't supported. Dimensions are limited to LINKED_ACCOUNT .
	//
	// [DimensionValues]: https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_DimensionValues.html
	Filter *types.Expression

	// This field is only used when the SortBy value is provided in the request.
	//
	// The maximum number of objects that are returned for this request. If MaxResults
	// isn't specified with the SortBy value, the request returns 1000 results as the
	// default value for this parameter.
	//
	// For GetCostCategories , MaxResults has an upper quota of 1000.
	MaxResults *int32

	// If the number of objects that are still available for retrieval exceeds the
	// quota, Amazon Web Services returns a NextPageToken value in the response. To
	// retrieve the next batch of objects, provide the NextPageToken from the previous
	// call in your next request.
	NextPageToken *string

	// The value that you want to search the filter values for.
	//
	// If you don't specify a CostCategoryName , SearchString is used to filter Cost
	// Category names that match the SearchString pattern. If you specify a
	// CostCategoryName , SearchString is used to filter Cost Category values that
	// match the SearchString pattern.
	SearchString *string

	// The value that you sort the data by.
	//
	// The key represents the cost and usage metrics. The following values are
	// supported:
	//
	//   - BlendedCost
	//
	//   - UnblendedCost
	//
	//   - AmortizedCost
	//
	//   - NetAmortizedCost
	//
	//   - NetUnblendedCost
	//
	//   - UsageQuantity
	//
	//   - NormalizedUsageAmount
	//
	// The supported key values for the SortOrder value are ASCENDING and DESCENDING .
	//
	// When you use the SortBy value, the NextPageToken and SearchString key values
	// aren't supported.
	SortBy []types.SortDefinition

	noSmithyDocumentSerde
}

type GetCostCategoriesOutput struct {

	// The number of objects that are returned.
	//
	// This member is required.
	ReturnSize *int32

	// The total number of objects.
	//
	// This member is required.
	TotalSize *int32

	// The names of the Cost Categories.
	CostCategoryNames []string

	// The Cost Category values.
	//
	// If the CostCategoryName key isn't specified in the request, the
	// CostCategoryValues fields aren't returned.
	CostCategoryValues []string

	// If the number of objects that are still available for retrieval exceeds the
	// quota, Amazon Web Services returns a NextPageToken value in the response. To
	// retrieve the next batch of objects, provide the marker from the prior call in
	// your next request.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCostCategoriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCostCategories{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCostCategories{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetCostCategories"); err != nil {
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
	if err = addOpGetCostCategoriesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCostCategories(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCostCategories(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetCostCategories",
	}
}
