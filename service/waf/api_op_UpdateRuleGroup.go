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
// global use. Inserts or deletes ActivatedRule () objects in a RuleGroup. You can
// only insert REGULAR rules into a rule group. You can have a maximum of ten rules
// per rule group.  <p>To create and configure a <code>RuleGroup</code>, perform
// the following steps:</p> <ol> <li> <p>Create and update the <code>Rules</code>
// that you want to include in the <code>RuleGroup</code>. See
// <a>CreateRule</a>.</p> </li> <li> <p>Use <code>GetChangeToken</code> to get the
// change token that you provide in the <code>ChangeToken</code> parameter of an
// <a>UpdateRuleGroup</a> request.</p> </li> <li> <p>Submit an
// <code>UpdateRuleGroup</code> request to add <code>Rules</code> to the
// <code>RuleGroup</code>.</p> </li> <li> <p>Create and update a
// <code>WebACL</code> that contains the <code>RuleGroup</code>. See
// <a>CreateWebACL</a>.</p> </li> </ol> <p>If you want to replace one
// <code>Rule</code> with another, you delete the existing one and add the new
// one.</p> <p>For more information about how to use the AWS WAF API to allow or
// block HTTP requests, see the <a
// href="https://docs.aws.amazon.com/waf/latest/developerguide/">AWS WAF Developer
// Guide</a>.</p>
func (c *Client) UpdateRuleGroup(ctx context.Context, params *UpdateRuleGroupInput, optFns ...func(*Options)) (*UpdateRuleGroupOutput, error) {
	stack := middleware.NewStack("UpdateRuleGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateRuleGroupMiddlewares(stack)
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
	addOpUpdateRuleGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRuleGroup(options.Region), middleware.Before)
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
			OperationName: "UpdateRuleGroup",
			Err:           err,
		}
	}
	out := result.(*UpdateRuleGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRuleGroupInput struct {

	// An array of RuleGroupUpdate objects that you want to insert into or delete from
	// a RuleGroup (). You can only insert REGULAR rules into a rule group.
	// ActivatedRule|OverrideAction applies only when updating or adding a RuleGroup to
	// a WebACL. In this case you do not use ActivatedRule|Action. For all other update
	// requests, ActivatedRule|Action is used instead of ActivatedRule|OverrideAction.
	//
	// This member is required.
	Updates []*types.RuleGroupUpdate

	// The RuleGroupId of the RuleGroup () that you want to update. RuleGroupId is
	// returned by CreateRuleGroup () and by ListRuleGroups ().
	//
	// This member is required.
	RuleGroupId *string

	// The value returned by the most recent call to GetChangeToken ().
	//
	// This member is required.
	ChangeToken *string
}

type UpdateRuleGroupOutput struct {

	// The ChangeToken that you used to submit the UpdateRuleGroup request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus ().
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateRuleGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateRuleGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateRuleGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateRuleGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "UpdateRuleGroup",
	}
}
