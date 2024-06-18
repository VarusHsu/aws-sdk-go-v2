// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Generates a stickiness policy with sticky session lifetimes that follow that of
// an application-generated cookie. This policy can be associated only with
// HTTP/HTTPS listeners.
//
// This policy is similar to the policy created by CreateLBCookieStickinessPolicy, except that the lifetime of
// the special Elastic Load Balancing cookie, AWSELB , follows the lifetime of the
// application-generated cookie specified in the policy configuration. The load
// balancer only inserts a new stickiness cookie when the application response
// includes a new application cookie.
//
// If the application cookie is explicitly removed or expires, the session stops
// being sticky until a new application cookie is issued.
//
// For more information, see [Application-Controlled Session Stickiness] in the Classic Load Balancers Guide.
//
// [Application-Controlled Session Stickiness]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-sticky-sessions.html#enable-sticky-sessions-application
func (c *Client) CreateAppCookieStickinessPolicy(ctx context.Context, params *CreateAppCookieStickinessPolicyInput, optFns ...func(*Options)) (*CreateAppCookieStickinessPolicyOutput, error) {
	if params == nil {
		params = &CreateAppCookieStickinessPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAppCookieStickinessPolicy", params, optFns, c.addOperationCreateAppCookieStickinessPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAppCookieStickinessPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CreateAppCookieStickinessPolicy.
type CreateAppCookieStickinessPolicyInput struct {

	// The name of the application cookie used for stickiness.
	//
	// This member is required.
	CookieName *string

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string

	// The name of the policy being created. Policy names must consist of alphanumeric
	// characters and dashes (-). This name must be unique within the set of policies
	// for this load balancer.
	//
	// This member is required.
	PolicyName *string

	noSmithyDocumentSerde
}

// Contains the output for CreateAppCookieStickinessPolicy.
type CreateAppCookieStickinessPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAppCookieStickinessPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateAppCookieStickinessPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateAppCookieStickinessPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAppCookieStickinessPolicy"); err != nil {
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
	if err = addOpCreateAppCookieStickinessPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAppCookieStickinessPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAppCookieStickinessPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAppCookieStickinessPolicy",
	}
}
