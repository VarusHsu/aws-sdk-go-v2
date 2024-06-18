// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Inserts or deletes Predicate objects in a Rule . Each Predicate object identifies a
// predicate, such as a ByteMatchSetor an IPSet, that specifies the web requests that you want to
// allow, block, or count. If you add more than one predicate to a Rule , a request
// must match all of the specifications to be allowed, blocked, or counted. For
// example, suppose that you add the following to a Rule :
//
//   - A ByteMatchSet that matches the value BadBot in the User-Agent header
//
//   - An IPSet that matches the IP address 192.0.2.44
//
// You then add the Rule to a WebACL and specify that you want to block requests
// that satisfy the Rule . For a request to be blocked, the User-Agent header in
// the request must contain the value BadBot and the request must originate from
// the IP address 192.0.2.44.
//
// To create and configure a Rule , perform the following steps:
//
//   - Create and update the predicates that you want to include in the Rule .
//
//   - Create the Rule . See CreateRule.
//
//   - Use GetChangeToken to get the change token that you provide in the
//     ChangeToken parameter of an UpdateRulerequest.
//
//   - Submit an UpdateRule request to add predicates to the Rule .
//
//   - Create and update a WebACL that contains the Rule . See CreateWebACL.
//
// If you want to replace one ByteMatchSet or IPSet with another, you delete the
// existing one and add the new one.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
func (c *Client) UpdateRule(ctx context.Context, params *UpdateRuleInput, optFns ...func(*Options)) (*UpdateRuleOutput, error) {
	if params == nil {
		params = &UpdateRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRule", params, optFns, c.addOperationUpdateRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRuleInput struct {

	// The value returned by the most recent call to GetChangeToken.
	//
	// This member is required.
	ChangeToken *string

	// The RuleId of the Rule that you want to update. RuleId is returned by CreateRule
	// and by ListRules.
	//
	// This member is required.
	RuleId *string

	// An array of RuleUpdate objects that you want to insert into or delete from a Rule.
	// For more information, see the applicable data types:
	//
	// RuleUpdate
	//   - : Contains Action and Predicate
	//
	// Predicate
	//   - : Contains DataId , Negated , and Type
	//
	// FieldToMatch
	//   - : Contains Data and Type
	//
	// This member is required.
	Updates []types.RuleUpdate

	noSmithyDocumentSerde
}

type UpdateRuleOutput struct {

	// The ChangeToken that you used to submit the UpdateRule request. You can also
	// use this value to query the status of the request. For more information, see GetChangeTokenStatus.
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateRule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateRule"); err != nil {
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
	if err = addOpUpdateRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateRule",
	}
}
