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

// Retrieves the Savings Plans recommendations for your account. First use
// StartSavingsPlansPurchaseRecommendationGeneration to generate a new set of
// recommendations, and then use GetSavingsPlansPurchaseRecommendation to retrieve
// them.
func (c *Client) GetSavingsPlansPurchaseRecommendation(ctx context.Context, params *GetSavingsPlansPurchaseRecommendationInput, optFns ...func(*Options)) (*GetSavingsPlansPurchaseRecommendationOutput, error) {
	if params == nil {
		params = &GetSavingsPlansPurchaseRecommendationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSavingsPlansPurchaseRecommendation", params, optFns, c.addOperationGetSavingsPlansPurchaseRecommendationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSavingsPlansPurchaseRecommendationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSavingsPlansPurchaseRecommendationInput struct {

	// The lookback period that's used to generate the recommendation.
	//
	// This member is required.
	LookbackPeriodInDays types.LookbackPeriodInDays

	// The payment option that's used to generate these recommendations.
	//
	// This member is required.
	PaymentOption types.PaymentOption

	// The Savings Plans recommendation type that's requested.
	//
	// This member is required.
	SavingsPlansType types.SupportedSavingsPlansType

	// The savings plan recommendation term that's used to generate these
	// recommendations.
	//
	// This member is required.
	TermInYears types.TermInYears

	// The account scope that you want your recommendations for. Amazon Web Services
	// calculates recommendations including the management account and member accounts
	// if the value is set to PAYER . If the value is LINKED , recommendations are
	// calculated for individual member accounts only.
	AccountScope types.AccountScope

	// You can filter your recommendations by Account ID with the LINKED_ACCOUNT
	// dimension. To filter your recommendations by Account ID, specify Key as
	// LINKED_ACCOUNT and Value as the comma-separated Acount ID(s) that you want to
	// see Savings Plans purchase recommendations for.
	//
	// For GetSavingsPlansPurchaseRecommendation, the Filter doesn't include
	// CostCategories or Tags . It only includes Dimensions . With Dimensions , Key
	// must be LINKED_ACCOUNT and Value can be a single Account ID or multiple
	// comma-separated Account IDs that you want to see Savings Plans Purchase
	// Recommendations for. AND and OR operators are not supported.
	Filter *types.Expression

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextPageToken *string

	// The number of recommendations that you want returned in a single response
	// object.
	PageSize int32

	noSmithyDocumentSerde
}

type GetSavingsPlansPurchaseRecommendationOutput struct {

	// Information that regards this specific recommendation set.
	Metadata *types.SavingsPlansPurchaseRecommendationMetadata

	// The token for the next set of retrievable results. Amazon Web Services provides
	// the token when the response from a previous call has more results than the
	// maximum page size.
	NextPageToken *string

	// Contains your request parameters, Savings Plan Recommendations Summary, and
	// Details.
	SavingsPlansPurchaseRecommendation *types.SavingsPlansPurchaseRecommendation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSavingsPlansPurchaseRecommendationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSavingsPlansPurchaseRecommendation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSavingsPlansPurchaseRecommendation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetSavingsPlansPurchaseRecommendation"); err != nil {
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
	if err = addOpGetSavingsPlansPurchaseRecommendationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSavingsPlansPurchaseRecommendation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSavingsPlansPurchaseRecommendation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetSavingsPlansPurchaseRecommendation",
	}
}
