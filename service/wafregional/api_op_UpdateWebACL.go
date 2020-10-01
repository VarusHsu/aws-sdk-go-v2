// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
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
// global use. Inserts or deletes ActivatedRule () objects in a WebACL. Each Rule
// identifies web requests that you want to allow, block, or count. When you update
// a WebACL, you specify the following values:
//
//     * A default action for the
// WebACL, either ALLOW or BLOCK. AWS WAF performs the default action if a request
// doesn't match the criteria in any of the Rules in a WebACL.
//
//     * The Rules
// that you want to add or delete. If you want to replace one Rule with another,
// you delete the existing Rule and add the new one.
//
//     * For each Rule, whether
// you want AWS WAF to allow requests, block requests, or count requests that match
// the conditions in the Rule.
//
//     * The order in which you want AWS WAF to
// evaluate the Rules in a WebACL. If you add more than one Rule to a WebACL, AWS
// WAF evaluates each request against the Rules in order based on the value of
// Priority. (The Rule that has the lowest value for Priority is evaluated first.)
// When a web request matches all the predicates (such as ByteMatchSets and IPSets)
// in a Rule, AWS WAF immediately takes the corresponding action, allow or block,
// and doesn't evaluate the request against the remaining Rules in the WebACL, if
// any.
//
//     <p>To create and configure a <code>WebACL</code>, perform the
// following steps:</p> <ol> <li> <p>Create and update the predicates that you want
// to include in <code>Rules</code>. For more information, see
// <a>CreateByteMatchSet</a>, <a>UpdateByteMatchSet</a>, <a>CreateIPSet</a>,
// <a>UpdateIPSet</a>, <a>CreateSqlInjectionMatchSet</a>, and
// <a>UpdateSqlInjectionMatchSet</a>.</p> </li> <li> <p>Create and update the
// <code>Rules</code> that you want to include in the <code>WebACL</code>. For more
// information, see <a>CreateRule</a> and <a>UpdateRule</a>.</p> </li> <li>
// <p>Create a <code>WebACL</code>. See <a>CreateWebACL</a>.</p> </li> <li> <p>Use
// <code>GetChangeToken</code> to get the change token that you provide in the
// <code>ChangeToken</code> parameter of an <a>UpdateWebACL</a> request.</p> </li>
// <li> <p>Submit an <code>UpdateWebACL</code> request to specify the
// <code>Rules</code> that you want to include in the <code>WebACL</code>, to
// specify the default action, and to associate the <code>WebACL</code> with a
// CloudFront distribution. </p> <p>The <code>ActivatedRule</code> can be a rule
// group. If you specify a rule group as your <code>ActivatedRule</code> , you can
// exclude specific rules from that rule group.</p> <p>If you already have a rule
// group associated with a web ACL and want to submit an <code>UpdateWebACL</code>
// request to exclude certain rules from that rule group, you must first remove the
// rule group from the web ACL, the re-insert it again, specifying the excluded
// rules. For details, see <a>ActivatedRule$ExcludedRules</a> . </p> </li> </ol>
// <p>Be aware that if you try to add a RATE_BASED rule to a web ACL without
// setting the rule type when first creating the rule, the <a>UpdateWebACL</a>
// request will fail because the request tries to add a REGULAR rule (the default
// rule type) with the specified ID, which does not exist. </p> <p>For more
// information about how to use the AWS WAF API to allow or block HTTP requests,
// see the <a href="https://docs.aws.amazon.com/waf/latest/developerguide/">AWS WAF
// Developer Guide</a>.</p>
func (c *Client) UpdateWebACL(ctx context.Context, params *UpdateWebACLInput, optFns ...func(*Options)) (*UpdateWebACLOutput, error) {
	stack := middleware.NewStack("UpdateWebACL", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateWebACLMiddlewares(stack)
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
	addOpUpdateWebACLValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWebACL(options.Region), middleware.Before)
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
			OperationName: "UpdateWebACL",
			Err:           err,
		}
	}
	out := result.(*UpdateWebACLOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateWebACLInput struct {

	// The WebACLId of the WebACL () that you want to update. WebACLId is returned by
	// CreateWebACL () and by ListWebACLs ().
	//
	// This member is required.
	WebACLId *string

	// A default action for the web ACL, either ALLOW or BLOCK. AWS WAF performs the
	// default action if a request doesn't match the criteria in any of the rules in a
	// web ACL.
	DefaultAction *types.WafAction

	// An array of updates to make to the WebACL (). An array of WebACLUpdate objects
	// that you want to insert into or delete from a WebACL (). For more information,
	// see the applicable data types:
	//
	//     * WebACLUpdate (): Contains Action and
	// ActivatedRule
	//
	//     * ActivatedRule (): Contains Action, OverrideAction,
	// Priority, RuleId, and Type. ActivatedRule|OverrideAction applies only when
	// updating or adding a RuleGroup to a WebACL. In this case, you do not use
	// ActivatedRule|Action. For all other update requests, ActivatedRule|Action is
	// used instead of ActivatedRule|OverrideAction.
	//
	//     * WafAction (): Contains Type
	Updates []*types.WebACLUpdate

	// The value returned by the most recent call to GetChangeToken ().
	//
	// This member is required.
	ChangeToken *string
}

type UpdateWebACLOutput struct {

	// The ChangeToken that you used to submit the UpdateWebACL request. You can also
	// use this value to query the status of the request. For more information, see
	// GetChangeTokenStatus ().
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateWebACLMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateWebACL{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateWebACL{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateWebACL(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "UpdateWebACL",
	}
}
