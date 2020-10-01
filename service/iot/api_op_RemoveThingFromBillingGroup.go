// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the given thing from the billing group.
func (c *Client) RemoveThingFromBillingGroup(ctx context.Context, params *RemoveThingFromBillingGroupInput, optFns ...func(*Options)) (*RemoveThingFromBillingGroupOutput, error) {
	stack := middleware.NewStack("RemoveThingFromBillingGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpRemoveThingFromBillingGroupMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveThingFromBillingGroup(options.Region), middleware.Before)
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
			OperationName: "RemoveThingFromBillingGroup",
			Err:           err,
		}
	}
	out := result.(*RemoveThingFromBillingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveThingFromBillingGroupInput struct {

	// The ARN of the thing to be removed from the billing group.
	ThingArn *string

	// The name of the thing to be removed from the billing group.
	ThingName *string

	// The ARN of the billing group.
	BillingGroupArn *string

	// The name of the billing group.
	BillingGroupName *string
}

type RemoveThingFromBillingGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpRemoveThingFromBillingGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpRemoveThingFromBillingGroup{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpRemoveThingFromBillingGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opRemoveThingFromBillingGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "RemoveThingFromBillingGroup",
	}
}
