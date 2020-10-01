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
// global use. Returns an array of RegexMatchSetSummary () objects.
func (c *Client) ListRegexMatchSets(ctx context.Context, params *ListRegexMatchSetsInput, optFns ...func(*Options)) (*ListRegexMatchSetsOutput, error) {
	stack := middleware.NewStack("ListRegexMatchSets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListRegexMatchSetsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListRegexMatchSets(options.Region), middleware.Before)
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
			OperationName: "ListRegexMatchSets",
			Err:           err,
		}
	}
	out := result.(*ListRegexMatchSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRegexMatchSetsInput struct {

	// If you specify a value for Limit and you have more RegexMatchSet objects than
	// the value of Limit, AWS WAF returns a NextMarker value in the response that
	// allows you to list another group of ByteMatchSets. For the second and subsequent
	// ListRegexMatchSets requests, specify the value of NextMarker from the previous
	// response to get information about another batch of RegexMatchSet objects.
	NextMarker *string

	// Specifies the number of RegexMatchSet objects that you want AWS WAF to return
	// for this request. If you have more RegexMatchSet objects than the number you
	// specify for Limit, the response includes a NextMarker value that you can use to
	// get another batch of RegexMatchSet objects.
	Limit *int32
}

type ListRegexMatchSetsOutput struct {

	// If you have more RegexMatchSet objects than the number that you specified for
	// Limit in the request, the response includes a NextMarker value. To list more
	// RegexMatchSet objects, submit another ListRegexMatchSets request, and specify
	// the NextMarker value from the response in the NextMarker value in the next
	// request.
	NextMarker *string

	// An array of RegexMatchSetSummary () objects.
	RegexMatchSets []*types.RegexMatchSetSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListRegexMatchSetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListRegexMatchSets{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRegexMatchSets{}, middleware.After)
}

func newServiceMetadataMiddleware_opListRegexMatchSets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "ListRegexMatchSets",
	}
}
