// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves an array of managed rule groups that are available for you to use.
// This list includes all Amazon Web Services Managed Rules rule groups and all of
// the Amazon Web Services Marketplace managed rule groups that you're subscribed
// to.
func (c *Client) ListAvailableManagedRuleGroups(ctx context.Context, params *ListAvailableManagedRuleGroupsInput, optFns ...func(*Options)) (*ListAvailableManagedRuleGroupsOutput, error) {
	if params == nil {
		params = &ListAvailableManagedRuleGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAvailableManagedRuleGroups", params, optFns, c.addOperationListAvailableManagedRuleGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAvailableManagedRuleGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAvailableManagedRuleGroupsInput struct {

	// Specifies whether this is for an Amazon CloudFront distribution or for a
	// regional application. A regional application can be an Application Load Balancer
	// (ALB), an Amazon API Gateway REST API, an AppSync GraphQL API, an Amazon Cognito
	// user pool, or an App Runner service. To work with CloudFront, you must also
	// specify the Region US East (N. Virginia) as follows:
	//   - CLI - Specify the Region when you use the CloudFront scope:
	//   --scope=CLOUDFRONT --region=us-east-1 .
	//   - API and SDKs - For all calls, use the Region endpoint us-east-1.
	//
	// This member is required.
	Scope types.Scope

	// The maximum number of objects that you want WAF to return for this request. If
	// more objects are available, in the response, WAF provides a NextMarker value
	// that you can use in a subsequent call to get the next batch of objects.
	Limit *int32

	// When you request a list of objects with a Limit setting, if the number of
	// objects that are still available for retrieval exceeds the limit, WAF returns a
	// NextMarker value in the response. To retrieve the next batch of objects, provide
	// the marker from the prior call in your next request.
	NextMarker *string

	noSmithyDocumentSerde
}

type ListAvailableManagedRuleGroupsOutput struct {

	// Array of managed rule groups that you can use. If you specified a Limit in your
	// request, this might not be the full list.
	ManagedRuleGroups []types.ManagedRuleGroupSummary

	// When you request a list of objects with a Limit setting, if the number of
	// objects that are still available for retrieval exceeds the limit, WAF returns a
	// NextMarker value in the response. To retrieve the next batch of objects, provide
	// the marker from the prior call in your next request.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAvailableManagedRuleGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAvailableManagedRuleGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAvailableManagedRuleGroups{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListAvailableManagedRuleGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAvailableManagedRuleGroups(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opListAvailableManagedRuleGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "ListAvailableManagedRuleGroups",
	}
}
