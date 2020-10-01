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
// global use. Creates a RuleGroup. A rule group is a collection of predefined
// rules that you add to a web ACL. You use UpdateRuleGroup () to add rules to the
// rule group. Rule groups are subject to the following limits:
//
//     * Three rule
// groups per account. You can request an increase to this limit by contacting
// customer support.
//
//     * One rule group per web ACL.
//
//     * Ten rules per rule
// group.
//
// For more information about how to use the AWS WAF API to allow or block
// HTTP requests, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/).
func (c *Client) CreateRuleGroup(ctx context.Context, params *CreateRuleGroupInput, optFns ...func(*Options)) (*CreateRuleGroupOutput, error) {
	stack := middleware.NewStack("CreateRuleGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateRuleGroupMiddlewares(stack)
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
	addOpCreateRuleGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRuleGroup(options.Region), middleware.Before)
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
			OperationName: "CreateRuleGroup",
			Err:           err,
		}
	}
	out := result.(*CreateRuleGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRuleGroupInput struct {

	//
	Tags []*types.Tag

	// A friendly name or description for the metrics for this RuleGroup. The name can
	// contain only alphanumeric characters (A-Z, a-z, 0-9), with maximum length 128
	// and minimum length one. It can't contain whitespace or metric names reserved for
	// AWS WAF, including "All" and "Default_Action." You can't change the name of the
	// metric after you create the RuleGroup.
	//
	// This member is required.
	MetricName *string

	// A friendly name or description of the RuleGroup (). You can't change Name after
	// you create a RuleGroup.
	//
	// This member is required.
	Name *string

	// The value returned by the most recent call to GetChangeToken ().
	//
	// This member is required.
	ChangeToken *string
}

type CreateRuleGroupOutput struct {

	// The ChangeToken that you used to submit the CreateRuleGroup request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus ().
	ChangeToken *string

	// An empty RuleGroup ().
	RuleGroup *types.RuleGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateRuleGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRuleGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRuleGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateRuleGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "CreateRuleGroup",
	}
}
