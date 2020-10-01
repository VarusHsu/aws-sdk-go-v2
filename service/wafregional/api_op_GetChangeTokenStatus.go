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
// global use. Returns the status of a ChangeToken that you got by calling
// GetChangeToken (). ChangeTokenStatus is one of the following values:
//
//     *
// PROVISIONED: You requested the change token by calling GetChangeToken, but you
// haven't used it yet in a call to create, update, or delete an AWS WAF object.
//
//
// * PENDING: AWS WAF is propagating the create, update, or delete request to all
// AWS WAF servers.
//
//     * INSYNC: Propagation is complete.
func (c *Client) GetChangeTokenStatus(ctx context.Context, params *GetChangeTokenStatusInput, optFns ...func(*Options)) (*GetChangeTokenStatusOutput, error) {
	stack := middleware.NewStack("GetChangeTokenStatus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetChangeTokenStatusMiddlewares(stack)
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
	addOpGetChangeTokenStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetChangeTokenStatus(options.Region), middleware.Before)
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
			OperationName: "GetChangeTokenStatus",
			Err:           err,
		}
	}
	out := result.(*GetChangeTokenStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetChangeTokenStatusInput struct {

	// The change token for which you want to get the status. This change token was
	// previously returned in the GetChangeToken response.
	//
	// This member is required.
	ChangeToken *string
}

type GetChangeTokenStatusOutput struct {

	// The status of the change token.
	ChangeTokenStatus types.ChangeTokenStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetChangeTokenStatusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetChangeTokenStatus{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetChangeTokenStatus{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetChangeTokenStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "GetChangeTokenStatus",
	}
}
