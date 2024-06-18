// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the available versions for the specified managed rule group.
func (c *Client) ListAvailableManagedRuleGroupVersions(ctx context.Context, params *ListAvailableManagedRuleGroupVersionsInput, optFns ...func(*Options)) (*ListAvailableManagedRuleGroupVersionsOutput, error) {
	if params == nil {
		params = &ListAvailableManagedRuleGroupVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAvailableManagedRuleGroupVersions", params, optFns, c.addOperationListAvailableManagedRuleGroupVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAvailableManagedRuleGroupVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAvailableManagedRuleGroupVersionsInput struct {

	// The name of the managed rule group. You use this, along with the vendor name,
	// to identify the rule group.
	//
	// This member is required.
	Name *string

	// Specifies whether this is for an Amazon CloudFront distribution or for a
	// regional application. A regional application can be an Application Load Balancer
	// (ALB), an Amazon API Gateway REST API, an AppSync GraphQL API, an Amazon Cognito
	// user pool, an App Runner service, or an Amazon Web Services Verified Access
	// instance.
	//
	// To work with CloudFront, you must also specify the Region US East (N. Virginia)
	// as follows:
	//
	//   - CLI - Specify the Region when you use the CloudFront scope:
	//   --scope=CLOUDFRONT --region=us-east-1 .
	//
	//   - API and SDKs - For all calls, use the Region endpoint us-east-1.
	//
	// This member is required.
	Scope types.Scope

	// The name of the managed rule group vendor. You use this, along with the rule
	// group name, to identify a rule group.
	//
	// This member is required.
	VendorName *string

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

type ListAvailableManagedRuleGroupVersionsOutput struct {

	// The name of the version that's currently set as the default.
	CurrentDefaultVersion *string

	// When you request a list of objects with a Limit setting, if the number of
	// objects that are still available for retrieval exceeds the limit, WAF returns a
	// NextMarker value in the response. To retrieve the next batch of objects, provide
	// the marker from the prior call in your next request.
	NextMarker *string

	// The versions that are currently available for the specified managed rule group.
	// If you specified a Limit in your request, this might not be the full list.
	Versions []types.ManagedRuleGroupVersion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAvailableManagedRuleGroupVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAvailableManagedRuleGroupVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAvailableManagedRuleGroupVersions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAvailableManagedRuleGroupVersions"); err != nil {
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
	if err = addOpListAvailableManagedRuleGroupVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAvailableManagedRuleGroupVersions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListAvailableManagedRuleGroupVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAvailableManagedRuleGroupVersions",
	}
}
