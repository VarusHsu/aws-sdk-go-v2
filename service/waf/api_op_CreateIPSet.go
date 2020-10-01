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
// global use. Creates an IPSet (), which you use to specify which web requests
// that you want to allow or block based on the IP addresses that the requests
// originate from. For example, if you're receiving a lot of requests from one or
// more individual IP addresses or one or more ranges of IP addresses and you want
// to block the requests, you can create an IPSet that contains those IP addresses
// and then configure AWS WAF to block the requests. To create and configure an
// IPSet, perform the following steps:
//
//     * Use GetChangeToken () to get the
// change token that you provide in the ChangeToken parameter of a CreateIPSet
// request.
//
//     * Submit a CreateIPSet request.
//
//     * Use GetChangeToken to get
// the change token that you provide in the ChangeToken parameter of an UpdateIPSet
// () request.
//
//     * Submit an UpdateIPSet request to specify the IP addresses
// that you want AWS WAF to watch for.
//
// For more information about how to use the
// AWS WAF API to allow or block HTTP requests, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/).
func (c *Client) CreateIPSet(ctx context.Context, params *CreateIPSetInput, optFns ...func(*Options)) (*CreateIPSetOutput, error) {
	stack := middleware.NewStack("CreateIPSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateIPSetMiddlewares(stack)
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
	addOpCreateIPSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIPSet(options.Region), middleware.Before)
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
			OperationName: "CreateIPSet",
			Err:           err,
		}
	}
	out := result.(*CreateIPSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIPSetInput struct {

	// A friendly name or description of the IPSet (). You can't change Name after you
	// create the IPSet.
	//
	// This member is required.
	Name *string

	// The value returned by the most recent call to GetChangeToken ().
	//
	// This member is required.
	ChangeToken *string
}

type CreateIPSetOutput struct {

	// The IPSet () returned in the CreateIPSet response.
	IPSet *types.IPSet

	// The ChangeToken that you used to submit the CreateIPSet request. You can also
	// use this value to query the status of the request. For more information, see
	// GetChangeTokenStatus ().
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateIPSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateIPSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateIPSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateIPSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "CreateIPSet",
	}
}
