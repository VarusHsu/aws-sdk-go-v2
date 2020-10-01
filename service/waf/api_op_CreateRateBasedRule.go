// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Creates a RateBasedRule (). The RateBasedRule contains a RateLimit,
// which specifies the maximum number of requests that AWS WAF allows from a
// specified IP address in a five-minute period. The RateBasedRule also contains
// the IPSet objects, ByteMatchSet objects, and other predicates that identify the
// requests that you want to count or block if these requests exceed the RateLimit.
// If you add more than one predicate to a RateBasedRule, a request not only must
// exceed the RateLimit, but it also must match all the conditions to be counted or
// blocked. For example, suppose you add the following to a RateBasedRule:
//
//     *
// An IPSet that matches the IP address 192.0.2.44/32
//
//     * A ByteMatchSet that
// matches BadBot in the User-Agent header
//
// Further, you specify a RateLimit of
// 1,000. You then add the RateBasedRule to a WebACL and specify that you want to
// block requests that meet the conditions in the rule. For a request to be
// blocked, it must come from the IP address 192.0.2.44 and the User-Agent header
// in the request must contain the value BadBot. Further, requests that match these
// two conditions must be received at a rate of more than 1,000 requests every five
// minutes. If both conditions are met and the rate is exceeded, AWS WAF blocks the
// requests. If the rate drops below 1,000 for a five-minute period, AWS WAF no
// longer blocks the requests.  <p>As a second example, suppose you want to limit
// requests to a particular page on your site. To do this, you could add the
// following to a <code>RateBasedRule</code>:</p> <ul> <li> <p>A
// <code>ByteMatchSet</code> with <code>FieldToMatch</code> of <code>URI</code>
// </p> </li> <li> <p>A <code>PositionalConstraint</code> of
// <code>STARTS_WITH</code> </p> </li> <li> <p>A <code>TargetString</code> of
// <code>login</code> </p> </li> </ul> <p>Further, you specify a
// <code>RateLimit</code> of 1,000.</p> <p>By adding this
// <code>RateBasedRule</code> to a <code>WebACL</code>, you could limit requests to
// your login page without affecting the rest of your site.</p> <p>To create and
// configure a <code>RateBasedRule</code>, perform the following steps:</p> <ol>
// <li> <p>Create and update the predicates that you want to include in the rule.
// For more information, see <a>CreateByteMatchSet</a>, <a>CreateIPSet</a>, and
// <a>CreateSqlInjectionMatchSet</a>.</p> </li> <li> <p>Use <a>GetChangeToken</a>
// to get the change token that you provide in the <code>ChangeToken</code>
// parameter of a <code>CreateRule</code> request.</p> </li> <li> <p>Submit a
// <code>CreateRateBasedRule</code> request.</p> </li> <li> <p>Use
// <code>GetChangeToken</code> to get the change token that you provide in the
// <code>ChangeToken</code> parameter of an <a>UpdateRule</a> request.</p> </li>
// <li> <p>Submit an <code>UpdateRateBasedRule</code> request to specify the
// predicates that you want to include in the rule.</p> </li> <li> <p>Create and
// update a <code>WebACL</code> that contains the <code>RateBasedRule</code>. For
// more information, see <a>CreateWebACL</a>.</p> </li> </ol> <p>For more
// information about how to use the AWS WAF API to allow or block HTTP requests,
// see the <a href="https://docs.aws.amazon.com/waf/latest/developerguide/">AWS WAF
// Developer Guide</a>.</p>
func (c *Client) CreateRateBasedRule(ctx context.Context, params *CreateRateBasedRuleInput, optFns ...func(*Options)) (*CreateRateBasedRuleOutput, error) {
	stack := middleware.NewStack("CreateRateBasedRule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateRateBasedRuleMiddlewares(stack)
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
	addOpCreateRateBasedRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRateBasedRule(options.Region), middleware.Before)
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
			OperationName: "CreateRateBasedRule",
			Err:           err,
		}
	}
	out := result.(*CreateRateBasedRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRateBasedRuleInput struct {

	// A friendly name or description of the RateBasedRule (). You can't change the
	// name of a RateBasedRule after you create it.
	//
	// This member is required.
	Name *string

	//
	Tags []*types.Tag

	// A friendly name or description for the metrics for this RateBasedRule. The name
	// can contain only alphanumeric characters (A-Z, a-z, 0-9), with maximum length
	// 128 and minimum length one. It can't contain whitespace or metric names reserved
	// for AWS WAF, including "All" and "Default_Action." You can't change the name of
	// the metric after you create the RateBasedRule.
	//
	// This member is required.
	MetricName *string

	// The maximum number of requests, which have an identical value in the field that
	// is specified by RateKey, allowed in a five-minute period. If the number of
	// requests exceeds the RateLimit and the other predicates specified in the rule
	// are also met, AWS WAF triggers the action that is specified for this rule.
	//
	// This member is required.
	RateLimit *int64

	// The field that AWS WAF uses to determine if requests are likely arriving from a
	// single source and thus subject to rate monitoring. The only valid value for
	// RateKey is IP. IP indicates that requests that arrive from the same IP address
	// are subject to the RateLimit that is specified in the RateBasedRule.
	//
	// This member is required.
	RateKey types.RateKey

	// The ChangeToken that you used to submit the CreateRateBasedRule request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus ().
	//
	// This member is required.
	ChangeToken *string
}

type CreateRateBasedRuleOutput struct {

	// The RateBasedRule () that is returned in the CreateRateBasedRule response.
	Rule *types.RateBasedRule

	// The ChangeToken that you used to submit the CreateRateBasedRule request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus ().
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateRateBasedRuleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRateBasedRule{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRateBasedRule{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateRateBasedRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "CreateRateBasedRule",
	}
}
