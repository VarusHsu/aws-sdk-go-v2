// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates settings for a specified resolver rule. ResolverRuleId is required, and
// all other parameters are optional. If you don't specify a parameter, it retains
// its current value.
func (c *Client) UpdateResolverRule(ctx context.Context, params *UpdateResolverRuleInput, optFns ...func(*Options)) (*UpdateResolverRuleOutput, error) {
	stack := middleware.NewStack("UpdateResolverRule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateResolverRuleMiddlewares(stack)
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
	addOpUpdateResolverRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateResolverRule(options.Region), middleware.Before)
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
			OperationName: "UpdateResolverRule",
			Err:           err,
		}
	}
	out := result.(*UpdateResolverRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateResolverRuleInput struct {

	// The new settings for the resolver rule.
	//
	// This member is required.
	Config *types.ResolverRuleConfig

	// The ID of the resolver rule that you want to update.
	//
	// This member is required.
	ResolverRuleId *string
}

type UpdateResolverRuleOutput struct {

	// The response to an UpdateResolverRule request.
	ResolverRule *types.ResolverRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateResolverRuleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateResolverRule{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateResolverRule{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateResolverRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "UpdateResolverRule",
	}
}
