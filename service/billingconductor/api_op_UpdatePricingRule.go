// Code generated by smithy-go-codegen DO NOT EDIT.

package billingconductor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/billingconductor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing pricing rule.
func (c *Client) UpdatePricingRule(ctx context.Context, params *UpdatePricingRuleInput, optFns ...func(*Options)) (*UpdatePricingRuleOutput, error) {
	if params == nil {
		params = &UpdatePricingRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePricingRule", params, optFns, c.addOperationUpdatePricingRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePricingRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePricingRuleInput struct {

	//  The Amazon Resource Name (ARN) of the pricing rule to update.
	//
	// This member is required.
	Arn *string

	//  The new description for the pricing rule.
	Description *string

	//  The new modifier to show pricing plan rates as a percentage.
	ModifierPercentage *float64

	//  The new name of the pricing rule. The name must be unique to each pricing
	// rule.
	Name *string

	//  The set of tiering configurations for the pricing rule.
	Tiering *types.UpdateTieringInput

	//  The new pricing rule type.
	Type types.PricingRuleType

	noSmithyDocumentSerde
}

type UpdatePricingRuleOutput struct {

	//  The Amazon Resource Name (ARN) of the successfully updated pricing rule.
	Arn *string

	//  The pricing plans count that this pricing rule is associated with.
	AssociatedPricingPlanCount int64

	//  The seller of services provided by Amazon Web Services, their affiliates, or
	// third-party providers selling services via Amazon Web Services Marketplace.
	BillingEntity *string

	//  The new description for the pricing rule.
	Description *string

	//  The most recent time the pricing rule was modified.
	LastModifiedTime int64

	//  The new modifier to show pricing plan rates as a percentage.
	ModifierPercentage *float64

	//  The new name of the pricing rule. The name must be unique to each pricing
	// rule.
	Name *string

	// Operation refers to the specific Amazon Web Services covered by this line item.
	// This describes the specific usage of the line item.
	//
	// If the Scope attribute is set to SKU , this attribute indicates which operation
	// the PricingRule is modifying. For example, a value of RunInstances:0202
	// indicates the operation of running an Amazon EC2 instance.
	Operation *string

	//  The scope of pricing rule that indicates if it's globally applicable, or it's
	// service-specific.
	Scope types.PricingRuleScope

	//  If the Scope attribute is set to SERVICE , the attribute indicates which
	// service the PricingRule is applicable for.
	Service *string

	//  The set of tiering configurations for the pricing rule.
	Tiering *types.UpdateTieringInput

	//  The new pricing rule type.
	Type types.PricingRuleType

	// Usage type is the unit that each service uses to measure the usage of a
	// specific type of resource.
	//
	// If the Scope attribute is set to SKU , this attribute indicates which usage type
	// the PricingRule is modifying. For example, USW2-BoxUsage:m2.2xlarge describes
	// an M2 High Memory Double Extra Large instance in the US West (Oregon) Region.
	UsageType *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePricingRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdatePricingRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdatePricingRule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdatePricingRule"); err != nil {
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
	if err = addOpUpdatePricingRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePricingRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePricingRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdatePricingRule",
	}
}
